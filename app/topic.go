package app

import (
	"errors"

	"github.com/segmentio/kafka-go"
)

type TopicConfig struct {
	Name              string
	Partitions        int16
	ReplicationFactor int16
}

func (c *Connection) CreateTopic(config TopicConfig) error {

	err := c.Connect.CreateTopics(kafka.TopicConfig{
		Topic:             config.Name,
		NumPartitions:     int(config.Partitions),
		ReplicationFactor: int(config.ReplicationFactor),
	})

	if err != nil {
		return errors.New("topic oluşturulamadı." + err.Error())
	}

	if err := c.Close(); err != nil {
		return errors.New("bağlantı kapatılamadı." + err.Error())
	}
	return nil
}

func (c *Connection) DeleteTopic(topic []string) error {
	if err := c.Connect.DeleteTopics(topic...); err != nil {
		return errors.New("topic silinemedi." + err.Error())
	}
	if err := c.Close(); err != nil {
		return errors.New("bağlantı kapatılamadı." + err.Error())
	}
	return nil
}
