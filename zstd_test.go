package sarama

import (
	"testing"
)

func TestSaramaZSTD(t *testing.T) {
	cfg := NewConfig()
	cfg.ClientID = "sarama-zstd-test"
	cfg.Producer.Return.Errors = true
	cfg.Producer.Return.Successes = true
	cfg.Producer.Retry.Max = 0
	cfg.Producer.Compression = CompressionZSTD
	cfg.Version = V2_1_0_0

	producer, err := NewSyncProducer([]string{"localhost:9092"}, cfg)
	if err != nil {
		t.Fatalf("NewSyncProducer failed: %v", err)
	}
	_, _, err = producer.SendMessage(
		&ProducerMessage{
			Topic: "my-topic",
			Value: StringEncoder("hello world!"),
		})
	if err != nil {
		t.Errorf("TEST:   sending message failed: %v\n", err)
	}
	_ = producer.Close()
}
