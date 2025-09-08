package app

import (
	"context"
	"errors"
	"time"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(config Config, consumerConfig ConsumerConfig) (*Consumer, error) {

	if consumerConfig.GroupID != "" && consumerConfig.Partition > 0 {
		return nil, errors.New("either specify a consumer group ID or a partition, but not both")
	}

	if !((consumerConfig.GroupID != "" && len(consumerConfig.GroupTopics) != 0) || (consumerConfig.GroupID != "" && config.Topic != "")) {
		return nil, errors.New("if a consumer consumer group id is specified, you must also define either a groups topic or a topic")
	}

	return &Consumer{
		reader: kafka.NewReader(
			kafka.ReaderConfig{
				Brokers:                config.Brokers,
				GroupID:                consumerConfig.GroupID,
				GroupTopics:            consumerConfig.GroupTopics,
				Topic:                  config.Topic,
				Partition:              consumerConfig.Partition,
				Dialer:                 config.Dialer,
				QueueCapacity:          config.QueueCapacity,
				MinBytes:               consumerConfig.MinBytes,
				MaxBytes:               consumerConfig.MinBytes,
				MaxWait:                consumerConfig.MaxWait,
				ReadBatchTimeout:       consumerConfig.ReadBatchTimeout,
				ReadLagInterval:        consumerConfig.ReadLagInterval,
				GroupBalancers:         consumerConfig.GroupBalancers,
				HeartbeatInterval:      consumerConfig.HeartbeatInterval,
				CommitInterval:         consumerConfig.CommitInterval,
				PartitionWatchInterval: consumerConfig.PartitionWatchInterval,
				WatchPartitionChanges:  consumerConfig.WatchPartitionChanges,
				SessionTimeout:         consumerConfig.SessionTimeout,
				RebalanceTimeout:       consumerConfig.RebalanceTimeout,
				JoinGroupBackoff:       consumerConfig.JoinGroupBackoff,
				RetentionTime:          consumerConfig.RetentionTime,
				StartOffset:            consumerConfig.StartOffset,
				ReadBackoffMin:         consumerConfig.ReadBackoffMin,
				ReadBackoffMax:         consumerConfig.ReadBackoffMax,
				Logger:                 config.Logger,
				ErrorLogger:            config.ErrorLogger,
				IsolationLevel:         consumerConfig.IsolationLevel,
				MaxAttempts:            config.MaxAttempts,
				OffsetOutOfRangeError:  consumerConfig.OffsetOutOfRangeError,
			},
		),
	}, nil
}

type Message struct {
	Headers   []kafka.Header
	Key       string
	Value     string
	Timestamp time.Time
	Topic     string
	Partition int
	Offset    int64
}

func (c *Consumer) Consume(ctx context.Context) (*Message, error) {
	message, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return nil, errors.New("veri okunamadı." + err.Error())
	}

	return &Message{
		Headers:   message.Headers,
		Value:     string(message.Value),
		Key:       string(message.Key),
		Timestamp: message.Time,
		Topic:     message.Topic,
		Partition: message.Partition,
		Offset:    message.Offset,
	}, nil
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
