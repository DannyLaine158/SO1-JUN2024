package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"serverGRPC/model"
	"time"
)

func Produce(value model.Data) {
	topic := "mytopic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		panic(err)
	}

	valueBytes, err := json.Marshal(&value)
	if err != nil {
		panic(err)
	}
	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return
	}
	_, err = conn.WriteMessages(
		kafka.Message{
			Value: valueBytes,
		})

	if err != nil {
		panic(err)
	}

	if err := conn.Close(); err != nil {
		panic(err)
	}

	log.Printf("Produced message to topic %s at partition %d", topic, partition)
}
