package producer

import (
	kafkaConfig "helloworld/config"

	"log"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

func Produce(topic string, limit int) {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{kafkaConfig.CONST_HOST}, config)
	if err != nil {
		log.Fatal("failed to initialize NewSyncProducer, err:", err)
		return
	}
	defer producer.Close()

	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))
		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(str)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
		log.Printf("[producer] partition id: %d; offset:%d, value: %s\n", partition, offset, str)
	}
}
