package repo

import "github.com/ozonva/ova-course-api/internal/course"

//go:generate mockgen -source=repo.go -destination=repo_mock.go -package=repo

// CourseRepo - интерфейс хранилища для сущности Course
type CourseRepo interface {
	AddCourse(entities []course.Course) error
	ListCourses(limit, offset uint64) ([]course.Course, error)
	DescribeCourses(entityId uint64) (*course.Course, error)
}
