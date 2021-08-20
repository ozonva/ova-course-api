package course

import (
	"fmt"
	"time"
)

const (
	dateFormat = "02.01.2006 15:04:05"
)

type Course struct {
	id          uint64
	userId      uint64
	name        string
	description string
	dateStart   time.Time
	dateFinish  time.Time
	dateCreated time.Time
}

func New(id uint64, userId uint64, Name string, dateStart time.Time, dateFinish time.Time) Course {
	return Course{
		id,
		userId,
		Name,
		"",
		dateStart,
		dateFinish,
		time.Now(),
	}
}

func (c Course) String() string {
	return fmt.Sprintf(
		"Course(%v):\nUserID: %v,\nName: %s,\nDescription: %s,\nDate start: %s,\nDate finish: %s,\nDate created: %s",
		c.id,
		c.userId,
		c.name,
		c.description,
		c.dateStart.Format(dateFormat),
		c.dateFinish.Format(dateFormat),
		c.dateCreated.Format(dateFormat),
	)
}

func (c Course) IsActive() bool {
	return (c.dateStart.Before(time.Now()) || c.dateStart.Equal(time.Now())) && c.dateFinish.After(time.Now())
}

func (c Course) Id() uint64 {
	return c.id
}

func (c Course) UserId() uint64 {
	return c.userId
}

func (c Course) Name() string {
	return c.name
}

func (c *Course) SetName(name string) {
	c.name = name
}

func (c Course) Description() string {
	return c.description
}

func (c *Course) SetDescription(description string) {
	c.description = description
}

func (c Course) DateStart() time.Time {
	return c.dateStart
}

func (c *Course) SetDateStart(dateStart time.Time) {
	c.dateStart = dateStart
}

func (c Course) DateFinish() time.Time {
	return c.dateFinish
}

func (c *Course) SetDateFinish(dateFinish time.Time) {
	c.dateFinish = dateFinish
}

func (c Course) DateCreated() time.Time {
	return c.dateCreated
}
