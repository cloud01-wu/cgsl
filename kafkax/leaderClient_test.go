package kafkax

import "testing"

const (
	brokerAddr = "192.168.99.251:9092"
)

func TestCreateTopics(t *testing.T) {
	configs := []TopicConfig{
		{
			TopicName:         "my-topic-1",
			NumPartitions:     2,
			ReplicationFactor: 1,
		},
		{
			TopicName:         "my-topic-2",
			NumPartitions:     2,
			ReplicationFactor: 1,
		},
	}

	err := CreateTopics(brokerAddr, configs...)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTopics(t *testing.T) {
	err := DeleteTopics(brokerAddr, []string{
		"my-topic-1",
		"my-topic-2",
	}...)

	if err != nil {
		t.Error(err)
	}
}

func TestListTopics(t *testing.T) {
	topics, err := ListTopics(brokerAddr)
	if err != nil {
		t.Error(err)
	}

	for _, topic := range topics {
		t.Log(topic)
	}
}
