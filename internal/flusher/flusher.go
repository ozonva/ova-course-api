package flusher

import (
	"github.com/ozonva/ova-course-api/internal/course"
	"github.com/ozonva/ova-course-api/internal/repo"
	"github.com/ozonva/ova-course-api/internal/utils"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(entities []course.Course) []course.Course
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(chunkSize int, courseRepo repo.CourseRepo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		courseRepo: courseRepo,
	}
}

type flusher struct {
	chunkSize  int
	courseRepo repo.CourseRepo
}

func (f *flusher) Flush(courses []course.Course) []course.Course {
	splitCourse := utils.SplitToBulks(courses, f.chunkSize)

	var unsavedCourses []course.Course

	for _, value := range splitCourse {
		if err := f.courseRepo.AddCourse(value); err != nil {
			unsavedCourses = append(unsavedCourses, value...)
		}
	}

	return unsavedCourses
}
