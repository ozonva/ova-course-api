package kafka_client

import (
	"encoding/json"
	"github.com/ozonva/ova-course-api/internal/course"
	"github.com/segmentio/kafka-go"
	"strconv"
)

const (
	create = "course_create"
	update = "course_update"
	delete = "course_delete"
)

func (p *producer) CreateEvent(item course.Course) error {
	jsonCourse, err := json.Marshal(item)
	if err != nil {
		return err
	}

	err = p.Send([]byte(create), jsonCourse)
	return err
}

func (p *producer) UpdateEvent(item course.Course) error {
	jsonCourse, err := json.Marshal(item)
	if err != nil {
		return err
	}
	err = p.Send([]byte(update), jsonCourse)
	return err
}

func (p *producer) DeleteEvent(id uint64) error {
	strId := strconv.FormatUint(id, 10)
	_, err := p.conn.WriteMessages(
		kafka.Message{
			Key:   []byte(delete),
			Value: []byte(strId),
		},
	)
	err = p.Send([]byte(delete), []byte(strId))
	return err
}
