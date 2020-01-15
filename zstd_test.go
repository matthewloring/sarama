package sarama

import (
	"testing"
)

func TestSaramaZSTD(t *testing.T) {
	kafkaVersion := V2_1_0_0
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
		Logger.Println("TEST:   sending compressed zstd message was successful")
	case err := <-producer.Errors():
		t.Errorf("TEST:   sending message failed: %v\n", err.Err)
	}

	producer.AsyncClose()
}
