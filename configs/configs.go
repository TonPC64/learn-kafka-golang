package configs

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

func NewConn() *kafka.Conn {
	topic := "test-kafka"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	return conn
}
