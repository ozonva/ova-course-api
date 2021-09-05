package kafka_client

import (
	"context"
	"github.com/ozonva/ova-course-api/internal/course"
	"github.com/segmentio/kafka-go"
)

type Producer interface {
	CreateEvent(item course.Course) error
	UpdateEvent(item course.Course) error
	DeleteEvent(id uint64) error
}

type producer struct {
	conn *kafka.Conn
}

func New(cnx context.Context, network string, address string, topic string, partition int) (Producer, error) {
	conn, err := kafka.DialLeader(cnx, network, address, topic, partition)
	if err != nil {
		return nil, err
	}

	return &producer{conn: conn}, nil
}

func (p *producer) Send(key []byte, value []byte) error {
	_, err := p.conn.WriteMessages(
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	return err
}
