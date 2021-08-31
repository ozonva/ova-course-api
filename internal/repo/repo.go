package repo

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-course-api/internal/course"
)

//go:generate mockgen -source=repo.go -destination=repo_mock.go -package=repo

// CourseRepo - интерфейс хранилища для сущности Course
type CourseRepo interface {
	AddCourse(entities []course.Course) error
	ListCourses(limit, offset uint64) ([]course.Course, error)
	DescribeCourses(entityId uint64) (*course.Course, error)
	RemoveCourse(courseId uint64) error
}

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) CourseRepo {
	return &repo{db: db}
}

func (r repo) AddCourse(entities []course.Course) error {

	bld := squirrel.
		Insert("courses").
		Columns("user_id", "name", "description", "date_start", "date_finish", "date_created")
	for _, entity := range entities {
		bld = bld.Values(entity.UserId(), entity.Name(), entity.Description(), entity.DateStart(), entity.DateFinish(), entity.DateCreated())
	}
	query, args, err := bld.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r repo) ListCourses(limit, offset uint64) ([]course.Course, error) {

	result := make([]course.Course, 0, limit)

	query, args, err := squirrel.
		Select("id", "user_id", "name", "description", "date_start", "date_finish", "date_created").
		From("courses").
		OrderBy("id desc").
		Limit(limit).
		Offset(offset).
		ToSql()
	if err != nil {
		return nil, err
	}
	err = r.db.Select(&result, query, args...)
	return result, err
}

func (r repo) DescribeCourses(courseId uint64) (*course.Course, error) {

	var result *course.Course
	query, args, err := squirrel.
		Select("id", "user_id", "name", "description", "date_start", "date_finish", "date_created").
		From("courses").
		Where(squirrel.Eq{"id": courseId}).
		ToSql()
	if err != nil {
		return nil, err
	}
	err = r.db.Get(&result, query, args)
	return result, err
}

func (r repo) RemoveCourse(courseId uint64) error {
	query := fmt.Sprintf("DELETE FROM courses id = $1")
	query, args, err := squirrel.
		Delete("courses").
		Where(squirrel.Eq{"id": courseId}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args)
	return err
}
