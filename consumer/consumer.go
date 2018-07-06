package main

import (
	"fmt"
	"time"

	"github.com/TonPC64/learn-kafka-golang/configs"
)

func main() {

	conn := configs.NewConn()

	conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	for {
		message, err := batch.ReadMessage()
		if err != nil {
			break
		}

		// blob, _ := json.Marshal(message)
		fmt.Println(string(message.Value))
		// fmt.Println(string(blob))

	}

	batch.Close()
	conn.Close()
}
