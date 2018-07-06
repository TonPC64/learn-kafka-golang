package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	topic := "test-kafka"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
}
