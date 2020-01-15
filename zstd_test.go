package sarama

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSaramaZSTD(t *testing.T) {
	kafkaVersion := V2_1_0_0
	Logger = log.New(os.Stdout, "", 0)

	Logger.Println(">>> kafka version", kafkaVersion)
	cfg := NewConfig()

	cfg.ClientID = "sarama-zstd-test"

	cfg.Producer.Return.Errors = true
	cfg.Producer.Return.Successes = true
	cfg.Producer.Retry.Max = 0

	cfg.Version = kafkaVersion

	if err := cfg.Validate(); err != nil {
		t.Fatalf("configuration validation failed: %v", err)
	}

	cfg.Producer.Compression = CompressionZSTD

	producer, err := NewAsyncProducer([]string{"localhost:9092"}, cfg)
	if err != nil {
		t.Fatalf("NewAsyncProducer failed: %v", err)
	}

	producer.Input() <- &ProducerMessage{
		Topic: "sarama-test",
		Value: StringEncoder("hello world!"),
	}

	select {
	case <-producer.Successes():
		fmt.Println("TEST:   sending message was successful")
	case err := <-producer.Errors():
		t.Errorf("TEST:   sending message failed: %v\n", err.Err)
	}

	if err := producer.Close(); err != nil {
		t.Errorf("producer.Close failed: %v", err)
	}
}
