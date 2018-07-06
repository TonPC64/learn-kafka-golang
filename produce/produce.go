package main

import (
	"github.com/TonPC64/learn-kafka-golang/configs"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	produce("bello")
}

func produce(message string) {
	conn := configs.NewConn()
	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)

	conn.Close()
}
