package app

import (
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

type Config struct {
	Brokers       []string
	Topic         string
	Dialer        *kafka.Dialer
	MaxAttempts   int
	QueueCapacity int
	Logger        kafka.Logger
	ErrorLogger   kafka.Logger
}

type ProducerConfig struct {
	Balancer          kafka.Balancer
	BatchSize         int
	BatchBytes        int
	BatchTimeout      time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	RebalanceInterval time.Duration
	IdleConnTimeout   time.Duration
	RequiredAcks      int
	Async             bool
	CompressionCodec  compress.Compression
}

type ConsumerConfig struct {
	GroupID                string
	GroupTopics            []string
	Partition              int
	MinBytes               int
	MaxBytes               int
	MaxWait                time.Duration
	ReadBatchTimeout       time.Duration
	ReadLagInterval        time.Duration
	GroupBalancers         []kafka.GroupBalancer
	HeartbeatInterval      time.Duration
	CommitInterval         time.Duration
	PartitionWatchInterval time.Duration
	WatchPartitionChanges  bool
	SessionTimeout         time.Duration
	RebalanceTimeout       time.Duration
	JoinGroupBackoff       time.Duration
	RetentionTime          time.Duration
	StartOffset            int64
	ReadBackoffMin         time.Duration
	ReadBackoffMax         time.Duration
	IsolationLevel         kafka.IsolationLevel
	OffsetOutOfRangeError  bool
}
