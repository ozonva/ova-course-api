package ova_course_api

import (
	"github.com/jmoiron/sqlx"
	kafkaClient "github.com/ozonva/ova-course-api/internal/kafka-client"
	"github.com/ozonva/ova-course-api/internal/metrics"
	"github.com/ozonva/ova-course-api/internal/repo"
	api "github.com/ozonva/ova-course-api/pkg/ova-course-api"
)

type GRPCServer struct {
	api.UnimplementedCourseServer
	repo          repo.CourseRepo
	kafkaProducer kafkaClient.Producer
	metrics       metrics.Metrics
}

func NewCourseServer(db *sqlx.DB, kafkaProducer kafkaClient.Producer, metrics metrics.Metrics) api.CourseServer {
	r := repo.NewRepo(db)
	return &GRPCServer{
		repo:          r,
		kafkaProducer: kafkaProducer,
		metrics:       metrics,
	}
}
