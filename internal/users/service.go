package users

import (
	"context"
	"crypto/md5"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Q0un/architect/proto/api"
)

type UsersService struct {
	logger     *log.Logger
	config     *Config
	db         *sqlx.DB
	jwtPublic  *rsa.PublicKey
	jwtPrivate *rsa.PrivateKey
}

func NewUsersService(logger *log.Logger, config *Config) (*UsersService, error) {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			config.DbUser,
			config.DbPassword,
			config.DbName,
			config.DbHost,
			config.DbPort,
		),
	)
	if err != nil {
		return nil, err
	}

	private, err := os.ReadFile(config.JwtPrivateFile)
	if err != nil {
		return nil, err
	}
	public, err := os.ReadFile(config.JwtPublicFile)
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

	return &UsersService{
		logger:     logger,
		config:     config,
		db:         db,
		jwtPublic:  jwtPublic,
		jwtPrivate: jwtPrivate,
	}, nil
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

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"login": req.GetLogin(),
	})
	token, err := t.SignedString(service.jwtPrivate)
	if err != nil {
		return "", err
	}

	tx := service.db.MustBegin()
	tx.MustExec(
		"INSERT INTO users (login, password) VALUES ($1, $2)",
		req.GetLogin(),
		md5Password(req.GetLogin(), req.GetPassword()),
	)
	tx.Commit()

	return token, nil
}

func (service *UsersService) SignIn(ctx context.Context, req *api.SignInRequest) (string, error) {
	user := User{}
	err := service.db.Get(&user, "SELECT * FROM users WHERE login=$1", req.GetLogin())
	if err != nil {
		return "", err
	}

	if md5Password(req.GetLogin(), req.GetPassword()) != user.Password {
		return "", fmt.Errorf("Wrong password")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"login": req.GetLogin(),
	})
	token, err := t.SignedString(service.jwtPrivate)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UsersService) EditInfo(ctx context.Context, req *api.EditInfoRequest, tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return service.jwtPublic, nil
	})
	if err != nil || !token.Valid {
		return fmt.Errorf("Bad jwt-token header")
	}

	login := token.Claims.(jwt.MapClaims)["login"].(string)

	user := User{}
	err = service.db.Get(&user, "SELECT * FROM users WHERE login=$1", login)
	if err != nil {
		return fmt.Errorf("Bad jwt-token header")
	}

	if req.Name != nil {
		user.Name.Scan(req.GetName())
	}
	if req.Surname != nil {
		user.Surname.Scan(req.GetSurname())
	}
	if req.Birthday != nil {
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
