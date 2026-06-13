package messaging

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(
	brokers []string,
	topic string,
) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:  kafka.TCP(brokers...),
			Topic: topic,
		},
	}
}

func (p *KafkaProducer) Publish(
	ctx context.Context,
	topic string,
	payload any,
) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(
		ctx,
		kafka.Message{
			Value: data,
		},
	)
}

// compiletime interface check
var _ Producer = (*KafkaProducer)(nil)
