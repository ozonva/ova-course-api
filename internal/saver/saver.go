package saver

import (
	"log"
	"sync"
	"time"

	"github.com/ozonva/ova-course-api/internal/course"
	"github.com/ozonva/ova-course-api/internal/flusher"
)

type Saver interface {
	Save(entity course.Course)
	Close()
}

type saver struct {
	interval   time.Duration
	mu         sync.Locker
	once       sync.Once
	ticketChan *time.Ticker
	killChan   chan struct{}
	capacity   uint
	flusher    flusher.Flusher
	courses    []course.Course
	closed     bool
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(capacity uint, flusher flusher.Flusher, interval time.Duration) Saver {
	s := &saver{
		capacity: capacity,
		flusher:  flusher,
		courses:  make([]course.Course, 0, capacity),
		interval: interval,
	}

	s.init()

	return s
}

func (s *saver) init() {
	s.once.Do(func() {
		s.ticketChan = time.NewTicker(s.interval)
		s.killChan = make(chan struct{}, 1)

		go func() {
			for {
				select {
				case <-s.ticketChan.C:
					s.flush()
				case <-s.killChan:
					s.ticketChan.Stop()
					s.flush()
					return
				}
			}
		}()
	})
}

func (s *saver) Save(entity course.Course) {
	if s.closed {
		log.Println("error: unable to save the object is closed")
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.courses = append(s.courses, entity)

}

func (s *saver) flush() {
	s.mu.Lock()
	defer s.mu.Unlock()

	unsavedCourses := s.flusher.Flush(s.courses)
	s.courses = make([]course.Course, 0, cap(s.courses))
	if unsavedCourses != nil {
		log.Printf("warning: some courses were not saved: \n%v\n", unsavedCourses)
		s.courses = append(s.courses, unsavedCourses...)
	}
}

func (s *saver) Close() {
	if s.closed {
		log.Println("warning: unable to closed the object is closed")
		return
	}

	s.closed = true
	s.killChan <- struct{}{}
}
