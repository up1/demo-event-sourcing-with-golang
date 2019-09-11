package main

import (
	"demo/consumer"
	"demo/producer"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	act := flag.String("act", "producer", "Either: producer or consumer")
	partition := flag.String("partition", "0",
		"Partition which the consumer program will be subscribing")

	flag.Parse()

	fmt.Printf("Welcome to Banku service: %s\n\n", *act)

	switch *act {
	case "producer":
		producer.MainProducer()
	case "consumer":
		if part32int, err := strconv.ParseInt(*partition, 10, 32); err == nil {
			consumer.MainConsumer(int32(part32int))
		}
	}
}
