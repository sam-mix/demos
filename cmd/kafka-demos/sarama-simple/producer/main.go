package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func SyncProduce() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "kafka_demo"
	content := "this is a sync message"
	msg.Value = sarama.StringEncoder(content)

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9098"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()

	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}

	fmt.Printf("send msg partition/offset: %d/%d, value is: %s", partition, offset, content)

}

func AsyncProduce() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9098"}, config)
	if err != nil {
		fmt.Println("error is:", err.Error())
		return
	}
	defer client.AsyncClose()

	go func(p sarama.AsyncProducer) {
		for {
			select {
			case msg := <-p.Successes():
				value, _ := msg.Value.Encode()
				fmt.Printf("send msg partition/offset: %d/%d, value is: %s", msg.Partition, msg.Offset, string(value))
				return
			case fail := <-p.Errors():
				fmt.Println("err: ", fail.Err)
				return
			}
		}
	}(client)

	msg := &sarama.ProducerMessage{
		Topic: "kafka_demo",
		Value: sarama.ByteEncoder("this is a async message"),
	}
	client.Input() <- msg
	time.Sleep(time.Second * 1)
}

func main() {
	AsyncProduce()
	// fmt.Println()
	// SyncProduce()
}
