package app

import (
	"context"
	"errors"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(config Config, producerConfig ProducerConfig) *Producer {
	return &Producer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:           config.Brokers,
			Topic:             config.Topic,
			Dialer:            config.Dialer,
			Balancer:          producerConfig.Balancer,
			MaxAttempts:       config.MaxAttempts,
			QueueCapacity:     config.QueueCapacity,
			BatchSize:         producerConfig.BatchSize,
			BatchBytes:        producerConfig.BatchBytes,
			BatchTimeout:      producerConfig.BatchTimeout,
			ReadTimeout:       producerConfig.ReadTimeout,
			WriteTimeout:      producerConfig.WriteTimeout,
			RebalanceInterval: producerConfig.RebalanceInterval,
			IdleConnTimeout:   producerConfig.IdleConnTimeout,
			RequiredAcks:      producerConfig.RequiredAcks,
			Async:             producerConfig.Async,
			CompressionCodec:  producerConfig.CompressionCodec.Codec(),
			Logger:            config.Logger,
			ErrorLogger:       config.ErrorLogger,
		}),
	}
}

func (p *Producer) Produce(ctx context.Context, record kafka.Message) error {

	if err := p.writer.WriteMessages(ctx, record); err != nil {
		return errors.New("record gönderilemedi." + err.Error())
	}
	return nil

}

func (p *Producer) Close() error {
	return p.writer.Close()
}
