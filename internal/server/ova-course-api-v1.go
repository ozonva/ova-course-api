package ova_course_api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/ozonva/ova-course-api/pkg/ova-course-api"
	"github.com/rs/zerolog/log"
)

// CreateCourseV1  Создание курса
func (g *GRPCServer) CreateCourseV1(ctx context.Context, request *api.CreateCourseV1Request) (*api.CreateCourseV1Response, error) {
	log.Info().Msg("Request: CreateCourseV1")
	return &api.CreateCourseV1Response{Success: true}, nil
}

// DescribeCourseV1  Описание курса
func (g *GRPCServer) DescribeCourseV1(ctx context.Context, request *api.DescribeCourseV1Request) (*api.DescribeCourseV1Response, error) {
	log.Info().Msg("Request: DescribeCourseV1")
	return &api.DescribeCourseV1Response{Id: 1}, nil
}

// ListCourseV1  Список всех доступных курсов
func (g *GRPCServer) ListCourseV1(ctx context.Context, request *empty.Empty) (*api.ListCourseV1Response, error) {
	log.Info().Msg("Request: ListCourseV1")
	return &api.ListCourseV1Response{}, nil
}

// RemoveCourseV1  Удаление курса
func (g *GRPCServer) RemoveCourseV1(ctx context.Context, request *api.RemoveCourseV1Request) (*api.RemoveCourseV1Response, error) {
	log.Info().Msg("Request: RemoveCourseV1")
	return &api.RemoveCourseV1Response{Success: true}, nil
}
