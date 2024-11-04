package kafkax

import (
	"fmt"
	"net"
	"strconv"

	kafka "github.com/segmentio/kafka-go"
)

type TopicConfig struct {
	// TopicName topic name
	TopicName string

	// NumPartitions created. -1 indicates unset
	NumPartitions int

	// ReplicationFactor for the topic. -1 indicates unset
	ReplicationFactor int
}

func CreateTopics(brokerAddr string, topicConfigs ...TopicConfig) error {
	// to create topics when auto.create.topics.enable='false'
	conn, err := kafka.Dial("tcp", brokerAddr)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("failed to obtain controller: %v", err)
	}

	leaderConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return fmt.Errorf("failed to dial to controller: %v", err)
	}
	defer leaderConn.Close()

	nativeTopicConfigs := []kafka.TopicConfig{}
	for _, topicConfig := range topicConfigs {
		nativeTopicConfigs = append(nativeTopicConfigs, kafka.TopicConfig{
			Topic:             topicConfig.TopicName,
			NumPartitions:     topicConfig.NumPartitions,
			ReplicationFactor: topicConfig.ReplicationFactor,
		})
	}

	return leaderConn.CreateTopics(nativeTopicConfigs...)
}

func DeleteTopics(brokerAddr string, topics ...string) error {
	conn, err := kafka.Dial("tcp", brokerAddr)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("failed to obtain controller: %v", err)
	}

	leaderConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return fmt.Errorf("failed to dial to controller: %v", err)
	}
	defer leaderConn.Close()

	return leaderConn.DeleteTopics(topics...)
}

func ListTopics(brokerAddr string) ([]string, error) {
	conn, err := kafka.Dial("tcp", brokerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return nil, fmt.Errorf("failed to read partitions: %v", err)
	}

	topicMap := map[string]interface{}{}

	for _, partition := range partitions {
		topicMap[partition.Topic] = nil
	}

	topics := []string{}
	for key := range topicMap {
		topics = append(topics, key)
	}

	return topics, nil
}
