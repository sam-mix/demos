package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	KafkaCluster "github.com/bsm/sarama-cluster"
)

func main() {
	KafkaConsumerCluster("")
}

func KafkaConsumerCluster(consumerId string) {

	brokers := []string{"127.0.0.1:9092"}
	topics := []string{"iris"}

	config := KafkaCluster.NewConfig()
	config.Consumer.Return.Errors = true
	//config.Consumer.Offsets.AutoCommit.Interval=1*time.Second
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Group.Return.Notifications = true

	//第二个参数是groupId
	consumer, err := KafkaCluster.NewConsumer(brokers, "consumer-group1", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 接收错误
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// 打印一些rebalance的信息
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// 消费消息
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s : %s/%d/%d\t%s\t%s\n", consumerId, msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // 提交offset
			}
		case <-signals:
			return
		}
	}
}
