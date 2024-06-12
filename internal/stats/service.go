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
