package kafkax

import (
	"context"
	"testing"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func TestConsume(t *testing.T) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"192.168.99.251:9091", "192.168.99.251:9092", "192.168.99.251:9093"},
		Topic:     "my-topic-1",
		Partition: 0,               // default is 0 as well
		MinBytes:  10e3,            // 10KB
		MaxBytes:  10e6,            // 10MB
		MaxWait:   time.Second * 3, // each 3 seconds
	})
	reader.SetOffset(0)

	// close the reader after 10s
	time.AfterFunc(time.Second*10, func() {
		defer reader.Close()
	})

	defer reader.Close()
	for {
		m, err := reader.ReadMessage(context.Background())
		// r.FetchMessage() && r.CommitMessages()
		if err != nil {
			break
		}
		t.Logf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

}

func TestConsumeGroup(t *testing.T) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"192.168.99.251:9091", "192.168.99.251:9092", "192.168.99.251:9093"},
		Topic:    "my-topic-1",
		GroupID:  "consumer-group-id1",
		MinBytes: 10e3,            // 10KB
		MaxBytes: 10e6,            // 10MB
		MaxWait:  time.Second * 3, // each 3 seconds
	})

	// close the reader after 10s
	time.AfterFunc(time.Second*10, func() {
		defer reader.Close()
	})

	defer reader.Close()
	for {
		m, err := reader.ReadMessage(context.Background())
		// r.FetchMessage() && r.CommitMessages()
		if err != nil {
			break
		}
		t.Logf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
