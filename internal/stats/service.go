package stats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	
	"github.com/ClickHouse/clickhouse-go/v2"
    "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/IBM/sarama"

	"github.com/Q0un/architect/proto/stats"
)

type StatsService struct {
	logger        *log.Logger
	db            driver.Conn
	kafkaConsumer sarama.Consumer
}

func NewStatsService(logger *log.Logger, config *Config) (*StatsService, error) {
	options := &clickhouse.Options{
		Addr: []string{config.Db.Host + ":" + config.Db.Port},
		Auth: clickhouse.Auth{
			Database: config.Db.Name,
			Username: config.Db.User,
			Password: config.Db.Password,
		},
	}
	db, err := clickhouse.Open(options)
	if err != nil {
		return nil, err
	}

	kafkaConsumer, err := setupKafkaConsumer(config)
	if err != nil {
		return nil, err
	}

	return &StatsService{
		logger:        logger,
		db:            db,
		kafkaConsumer: kafkaConsumer,
	}, nil
}

func setupKafkaConsumer(config *Config) (sarama.Consumer, error) {
	brokers := []string{config.Kafka.Host}

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer(brokers, kafkaConfig)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func (service *StatsService) runKafkaConsumer(ctx context.Context) error {
	partitionList, err := service.kafkaConsumer.Partitions("stats")
	if err != nil {
		return err
	}

	for _, partition := range partitionList {
		err = service.consumePartition(partition)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *StatsService) consumePartition(partition int32) error {
	pc, err := service.kafkaConsumer.ConsumePartition("stats", partition, sarama.OffsetNewest)
	if err != nil {
		return fmt.Errorf("Failed to consume partition %d: %v", partition, err)
	}

	for message := range pc.Messages() {
		var event stats.StatsEvent
		err := json.Unmarshal(message.Value, &event)
		if err != nil {
			service.logger.Println("Error unmarshalling JSON:", err)
			continue
		}

		err = service.addToDb(&event)
		if err != nil {
			service.logger.Println("Error sending to clickhouse:", err)
			continue
		}
	}
	return nil
}

func (service *StatsService) addToDb(event *stats.StatsEvent) error {
	batch, err := service.db.PrepareBatch(context.Background(), "INSERT INTO stats (user_id, ticket_id, type) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("Failed to prepare batch: %v", err)
	}

	if err := batch.Append(event.GetUserId(), event.GetTicketId(), event.GetType()); err != nil {
		return fmt.Errorf("Failed to append row: %v", err)
	}

	if err := batch.Send(); err != nil {
		return fmt.Errorf("Failed to send batch: %v", err)
	}

	return nil
}


func (service *StatsService) TicketStats(ticket_id uint64) (uint64, uint64, error) {
	query := fmt.Sprintf("SELECT uniqExactIf(user_id, type = 'view' AND ticket_id = %d) FROM stats", ticket_id)
	row, err := service.db.Query(context.Background(), query)
	if err != nil {
		return 0, 0, err
	}

	var views uint64
	row.Next()
	err = row.Scan(&views)
	if err != nil {
		return 0, 0, err
	}
	
	query = fmt.Sprintf("SELECT uniqExactIf(user_id, type = 'like' AND ticket_id = %d) FROM stats", ticket_id)
	row, err = service.db.Query(context.Background(), query)
	if err != nil {
		return 0, 0, err
	}

	var likes uint64
	row.Next()
	err = row.Scan(&likes)
	if err != nil {
		return 0, 0, err
	}

	return views, likes, nil
}

func (service *StatsService) TopTickets(evType string) ([]*stats.TicketInfo, error) {
	query := fmt.Sprintf("SELECT ticket_id, uniqExactIf(user_id, type = '%s') AS stat FROM stats GROUP BY ticket_id ORDER BY stat DESC LIMIT 5", evType)
	rows, err := service.db.Query(context.Background(), query)
	if err != nil {
		return []*stats.TicketInfo{}, err
	}

	var top []*stats.TicketInfo
	for rows.Next() {
		var ticketId uint64
		var count uint64
		err = rows.Scan(&ticketId, &count)
		if err != nil {
			return []*stats.TicketInfo{}, err
		}
		top = append(top, &stats.TicketInfo{
			TicketId: ticketId,
			Count: count,
		})
	}
	return top, nil
}

func (service *StatsService) TopUsers() ([]*stats.UserInfo, error) {
	query := "SELECT user_id, uniqExactIf(ticket_id, type = 'like') AS likes FROM stats GROUP BY user_id ORDER BY likes DESC LIMIT 3"
	rows, err := service.db.Query(context.Background(), query)
	if err != nil {
		return []*stats.UserInfo{}, err
	}

	var top []*stats.UserInfo
	for rows.Next() {
		var userId uint64
		var likes uint64
		err = rows.Scan(&userId, &likes)
		if err != nil {
			return []*stats.UserInfo{}, err
		}
		top = append(top, &stats.UserInfo{
			UserId: userId,
			Likes: likes,
		})
	}
	return top, nil
}
