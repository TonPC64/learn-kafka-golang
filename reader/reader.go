package main

import (
	"context"
	"fmt"
	"sync"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "test-kafka",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	// r.SetOffset(3)

	// for {
	// 	m, _ := r.ReadMessage(context.Background())
	// 	// if err != nil {
	// 	// 	break
	// 	// }
	// 	fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	// }

	// // r.Close()
	cont := context.Background()
	for {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(r *kafka.Reader, ctx context.Context, wg *sync.WaitGroup) {
			m, _ := r.FetchMessage(ctx)
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
			r.CommitMessages(ctx, m)
			wg.Done()
		}(r, cont, wg)
		wg.Wait()
	}
}
