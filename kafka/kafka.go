package kafka

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

var (
	Brokers = []string{"localhost:29092"}
	Topic   = "bank-transactions"
)

func newKafkaConfiguration() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_1_0
	return conf
}

func NewKafkaSyncProducer() sarama.SyncProducer {
	kafka, err := sarama.NewSyncProducer(Brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	return kafka
}

func NewKafkaConsumer() sarama.Consumer {
	config := sarama.NewConfig()
    config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(Brokers, config)

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	return consumer
}

func SendMsg(kafka sarama.SyncProducer, event interface{}) error {
	json, err := json.Marshal(event)

	if err != nil {
		return err
	}

	msgLog := &sarama.ProducerMessage{
		Topic: Topic,
		Value: sarama.StringEncoder(string(json)),
	}

	partition, offset, err := kafka.SendMessage(msgLog)
	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
	}

	fmt.Printf("Message: %+v\n", event)
	fmt.Printf("Message is stored in partition %d, offset %d\n",
		partition, offset)

	return nil
}
