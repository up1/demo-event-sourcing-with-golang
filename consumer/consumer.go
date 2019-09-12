package consumer

import (
	"bufio"
	"demo/bank"
	"demo/event"
	k "demo/kafka"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

func MainConsumer(partition int32) {
	kafka := k.NewKafkaConsumer()
	defer kafka.Close()

	consumer, err := kafka.ConsumePartition(k.Topic, partition, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	go consumeEvents(consumer)

	fmt.Println("Press [enter] to exit consumer\n")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Terminating...")
}

func consumeEvents(consumer sarama.PartitionConsumer) {
	var msgVal []byte
	var log interface{}
	var logMap map[string]interface{}
	var bankAccount *bank.BankAccount
	var err error

	for {
		select {
		case err := <-consumer.Errors():
			fmt.Printf("Kafka error: %s\n", err)
		case msg := <-consumer.Messages():
			msgVal = msg.Value

			if err = json.Unmarshal(msgVal, &log); err != nil {
				fmt.Printf("Failed parsing: %s", err)
			} else {
				logMap = log.(map[string]interface{})
				logType := logMap["Type"]
				fmt.Printf("Processing %s:\n%s\n", logMap["Type"], string(msgVal))

				switch logType {
				case "CreateEvent":
					_event := new(event.CreateEvent)
					if err = json.Unmarshal(msgVal, &_event); err == nil {
						bankAccount, err = _event.Process()
					}
				case "BalanceEvent":
					_event := new(event.BalanceEvent)
					if err = json.Unmarshal(msgVal, &_event); err == nil {
						bankAccount, err = _event.Process()
					}
				case "DepositEvent":
					_event := new(event.DepositEvent)
					if err = json.Unmarshal(msgVal, &_event); err == nil {
						bankAccount, err = _event.Process()
					}
				case "WithdrawEvent":
					_event := new(event.WithdrawEvent)
					if err = json.Unmarshal(msgVal, &_event); err == nil {
						bankAccount, err = _event.Process()
					}
				case "TransferEvent":
					_event := new(event.TransferEvent)
					if err = json.Unmarshal(msgVal, &_event); err == nil {
						if bankAccount, err = _event.Process(); err == nil {
							if targetAcc, err := bank.FetchAccount(_event.TargetId); err == nil {
								fmt.Printf("%+v\n", *targetAcc)
							}
						}
					}
				default:
					fmt.Println("Unknown command: ", logType)
				}

				if err != nil {
					fmt.Printf("Error processing: %s\n", err)
				} else {
					fmt.Printf("%+v\n\n", *bankAccount)
				}
			}
		}
	}
}
