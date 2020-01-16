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
	kafkaHome := []string{"localhost:9092"}
	topic := "my-zstd-topic"

	admin, err := NewClusterAdmin(kafkaHome, cfg)
	if err != nil {
		t.Fatal(err)
	}
	_ = admin.DeleteTopic(topic)
	err = admin.CreateTopic(topic, &TopicDetail{NumPartitions: 1, ReplicationFactor: 1}, false)
	if err.Error() != "kafka server: Topic with this name already exists. - Topic 'my-zstd-topic' already exists." {
		t.Fatal(err)
	}
	defer func() {
		if err := admin.Close(); err != nil {
			t.Error(err)
		}
	}()

	// producer
	producer, err := NewSyncProducer(kafkaHome, cfg)
	defer func() {
		if err := producer.Close(); err != nil {
			t.Error(err)
		}
	}()
	if err != nil {
		t.Fatalf("NewSyncProducer failed: %v", err)
	}
	_, _, err = producer.SendMessage(
		&ProducerMessage{
			Topic: topic,
			Value: StringEncoder("hello world!"),
		})
	if err != nil {
		t.Errorf("TEST:   sending message failed: %v\n", err)
	}
}
