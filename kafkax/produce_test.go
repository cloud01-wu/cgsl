package kafkax

import (
	"context"
	"testing"
	"time"

	"github.com/cloud01-wu/cgsl/utils"
	kafka "github.com/segmentio/kafka-go"
)

func TestProduce(t *testing.T) {
	// make a writer that produces to the given topic, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP("192.168.99.251:9091", "192.168.99.251:9092", "192.168.99.251:9093"),
		Topic:    "my-topic-1",
		Balancer: &kafka.LeastBytes{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// batch produce messages
	err := w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(utils.RandomUUIDString()),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte(utils.RandomUUIDString()),
			Value: []byte("Two!"),
		},
		kafka.Message{
			Key:   []byte(utils.RandomUUIDString()),
			Value: []byte("Three!"),
		},
	)
	defer w.Close()

	if err != nil {
		t.Fatal("failed to write messages:", err)
	}

	t.Log("messages produced successfully")
}
