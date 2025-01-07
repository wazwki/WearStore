package controllers

import (
	"github.com/IBM/sarama"
	"github.com/wazwki/WearStore/notification-service/internal/service"
)

type Consumer struct {
	service  service.ServiceInterface
	consumer sarama.Consumer
}

func NewConsumer(consumer sarama.Consumer, service service.ServiceInterface) *Consumer {
	return &Consumer{consumer: consumer, service: service}
}

func (c *Consumer) GetMessages(topic string, partition int32) error {
	partcon, err := c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		return err
	}
	defer partcon.Close()

	msgs := partcon.Messages()

	var messages []string
	for msg := range msgs {
		messages = append(messages, string(msg.Value))
	}

	err = c.service.SendNotification(messages)
	if err != nil {
		return err
	}

	return nil
}
