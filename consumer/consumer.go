package main

import (
	"context"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	topic := "test-kafka"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	for {
		b := make([]byte, 10e3) // 10KB max per message
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	batch.Close()
	conn.Close()
}
