package users

import (
	"context"
	"crypto/md5"
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/IBM/sarama"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Q0un/architect/proto/api"
	"github.com/Q0un/architect/proto/stats"
)

type UsersService struct {
	logger        *log.Logger
	config        *Config
	db            *sqlx.DB
	jwtPublic     *rsa.PublicKey
	jwtPrivate    *rsa.PrivateKey
	tickenator    *TickenatorClient
	kafkaProducer sarama.SyncProducer
}

func NewUsersService(logger *log.Logger, config *Config) (*UsersService, error) {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			config.Db.User,
			config.Db.Password,
			config.Db.Name,
			config.Db.Host,
			config.Db.Port,
		),
	)
	if err != nil {
		return nil, err
	}

	private, err := os.ReadFile(config.Jwt.PrivateFile)
	if err != nil {
		return nil, err
	}
	public, err := os.ReadFile(config.Jwt.PublicFile)
	if err != nil {
		return nil, err
	}
	jwtPrivate, err := jwt.ParseRSAPrivateKeyFromPEM(private)
	if err != nil {
		return nil, err
	}
	jwtPublic, err := jwt.ParseRSAPublicKeyFromPEM(public)
	if err != nil {
		return nil, err
	}

	tickenator, err := NewTickenatorClient(logger, config)
	if err != nil {
		return nil, err
	}

	kafkaProducer, err := setupKafkaProducer(config)
	if err != nil {
		return nil, err
	}

	return &UsersService{
		logger:        logger,
		config:        config,
		db:            db,
		jwtPublic:     jwtPublic,
		jwtPrivate:    jwtPrivate,
		tickenator:    tickenator,
		kafkaProducer: kafkaProducer,
	}, nil
}

func setupKafkaProducer(config *Config) (sarama.SyncProducer, error) {
	brokers := []string{config.Kafka.Host}

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, kafkaConfig)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func md5Password(login string, password string) string {
	b := []byte(login)
	s := []byte(password)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func (service *UsersService) SignUp(ctx context.Context, req *api.SignUpRequest) (string, error) {
	user := User{}
	err := service.db.Get(&user, "SELECT * FROM users WHERE login=$1", req.GetLogin())
	if err == nil {
		return "", fmt.Errorf("User with such login already exists")
	}

	tx := service.db.MustBegin()
	var id uint64
	err = tx.QueryRowx(
		"INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id",
		req.GetLogin(),
		md5Password(req.GetLogin(), req.GetPassword()),
	).Scan(&id)
	if err != nil {
		tx.Rollback()
		return "", fmt.Errorf("Cannot add user to database")
	}
	tx.Commit()
	fmt.Println("signup:", id)

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id": strconv.FormatUint(id, 10),
	})
	token, err := t.SignedString(service.jwtPrivate)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UsersService) SignIn(ctx context.Context, req *api.SignInRequest) (string, error) {
	user := User{}
	err := service.db.Get(&user, "SELECT * FROM users WHERE login=$1", req.GetLogin())
	if err != nil {
		return "", fmt.Errorf("Wrong login or password")
	}

	if md5Password(req.GetLogin(), req.GetPassword()) != user.Password {
		return "", fmt.Errorf("Wrong login or password")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id": strconv.FormatUint(user.Id, 10),
	})
	token, err := t.SignedString(service.jwtPrivate)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UsersService) EditInfo(ctx context.Context, req *api.EditInfoRequest, id uint64) error {
	user := User{}
	err := service.db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("Bad auth header")
	}

	if req.Name != nil {
		user.Name.Scan(req.GetName())
	}
	if req.Surname != nil {
		user.Surname.Scan(req.GetSurname())
	}
	if req.Birthday != nil {
		_, err := time.Parse(time.DateOnly, req.GetBirthday())
		if err != nil {
			return fmt.Errorf("Bad birthday format. Valid one is YYYY-MM-DD")
		}
		user.Birthday.Scan(req.GetBirthday())
	}
	if req.Mail != nil {
		user.Mail.Scan(req.GetMail())
	}
	if req.Phone != nil {
		user.Phone.Scan(req.GetPhone())
	}

	_, err = service.db.NamedExec(
		"UPDATE users SET name=:name, surname=:surname, birthday=:birthday, mail=:mail, phone=:phone WHERE id = :id",
		user,
	)
	if err != nil {
		return err
	}

	return nil
}

func (service *UsersService) CheckUser(id uint64) bool {
	user := User{}
	fmt.Println("check:", id)
	err := service.db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	fmt.Println(err)
	return err == nil
}

func (service *UsersService) SendKafkaEvent(ctx context.Context, ticketId uint64, userId uint64, evType string) error {
	ev := stats.StatsEvent{
		TicketId: ticketId,
		UserId: userId,
		Type: evType,
	}

	jsonEvent, err := json.Marshal(ev)
    if err != nil {
        return fmt.Errorf("Failed to serialize message to JSON: %v", err)
    }

	message := &sarama.ProducerMessage{
        Topic: "stats",
        Key:   sarama.StringEncoder(evType),
        Value: sarama.ByteEncoder(jsonEvent),
    }

    _, _, err = service.kafkaProducer.SendMessage(message)
    if err != nil {
        return fmt.Errorf("Failed to send message to kafka: %v", err)
    }
	return nil
}
