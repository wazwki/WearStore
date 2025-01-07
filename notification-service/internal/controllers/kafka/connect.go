package kafka

import "github.com/IBM/sarama"

func NewConsumer(address string) (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer([]string{address}, nil)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
