package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-course-api/internal/course"
)

//go:generate mockgen -source=repo.go -destination=repo_mock.go -package=repo

// CourseRepo - интерфейс хранилища для сущности Course
type CourseRepo interface {
	AddCourse(entities []course.Course) error
	ListCourses(limit, offset uint64) ([]course.Course, error)
	DescribeCourses(entityId uint64) (*course.Course, error)
}

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) CourseRepo {
	return &repo{db: db}
}

func (r repo) AddCourse(entity course.Course) error {
	query := fmt.Sprintf("INSERT INTO courses (user_id, name, description, date_start, date_finish, date_created) values ($1, $2, $3, $4, $5, $6) RETURNING id")
	r.db.QueryRow(query, entities)
}

func (r repo) ListCourses(limit, offset uint64) ([]course.Course, error) {

}

func (r repo) DescribeCourses(entityId uint64) (*course.Course, error) {

}
