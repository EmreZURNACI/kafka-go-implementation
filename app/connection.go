package app

import (
	"errors"

	"github.com/segmentio/kafka-go"
)

type ConnectionConfig struct {
	Network string
	Address string
}

type Connection struct {
	Connect *kafka.Conn
}

func NewConnection(config ConnectionConfig) (*Connection, error) {

	if config.Network == "" {
		config.Network = "tcp"
	}

	if config.Address == "" {
		return nil, errors.New("host adresi boş bırakılamaz")
	}

	con, err := kafka.Dial(config.Network, config.Address)
	if err != nil {
		return nil, errors.New("host ile bağlantı sağlanamadı." + err.Error())
	}

	return &Connection{
		Connect: con,
	}, nil
}
func (c *Connection) Close() error {
	return c.Connect.Close()
}
