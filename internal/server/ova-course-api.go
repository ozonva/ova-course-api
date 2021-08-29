package ova_course_api

import (
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-course-api/internal/repo"
	api "github.com/ozonva/ova-course-api/pkg/ova-course-api"
)

type GRPCServer struct {
	api.UnimplementedCourseServer
	repo repo.CourseRepo
}

func NewCourseServer(db *sqlx.DB) api.CourseServer {
	r := repo.NewRepo(db)
	return &GRPCServer{repo: r}
}
