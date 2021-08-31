package ova_course_api

import (
	"context"
	"errors"
	"github.com/ozonva/ova-course-api/internal/course"
	"github.com/ozonva/ova-course-api/internal/flusher"
	"github.com/ozonva/ova-course-api/internal/saver"
	api "github.com/ozonva/ova-course-api/pkg/ova-course-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	addCap    = 3
	addPeriod = time.Second
)

// CreateCourseV1  Создание курса
func (g *GRPCServer) CreateCourseV1(ctx context.Context, req *api.CreateCourseV1Request) (*api.CreateCourseV1Response, error) {
	log.Info().Msg("Request: CreateCourseV1")

	log.Debug().
		Uint64("userId", req.UserId).
		Str("name", req.Name).
		Time("dateStart", req.DateStart.AsTime()).
		Time("dateFinish", req.DateFinish.AsTime()).
		Msg("request createcourseV1")

	flush := flusher.NewFlusher(addCap, g.repo)

	s := saver.NewSaver(addCap, flush, addPeriod)
	defer s.Close()

	if err := req.DateStart.CheckValid(); err != nil {
		log.Error().Err(err).Msg("")
		return &api.CreateCourseV1Response{Success: false}, err
	}

	if err := req.DateFinish.CheckValid(); err != nil {
		log.Error().Err(err).Msg("")
		return &api.CreateCourseV1Response{Success: false}, err
	}

	obj := course.New(req.GetUserId(), req.GetName(), req.GetDateStart().AsTime(), req.GetDateFinish().AsTime())
	s.Save(obj)

	return &api.CreateCourseV1Response{Success: true}, nil
}

// DescribeCourseV1  Описание курса
func (g *GRPCServer) DescribeCourseV1(ctx context.Context, req *api.DescribeCourseV1Request) (*api.DescribeCourseV1Response, error) {
	log.Info().Msg("Request: DescribeCourseV1")

	log.Debug().
		Uint64("userId", req.GetId()).
		Msg("request describeCourseV1")

	if req.GetId() == 0 {
		err := errors.New("id can't be 0")
		log.Error().Err(err).Msg("")
		return nil, err
	}

	course, err := g.repo.DescribeCourses(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	return convertCourseToResponse(course), nil
}

// ListCourseV1  Список всех доступных курсов
func (g *GRPCServer) ListCourseV1(ctx context.Context, req *api.ListCourseV1Request) (*api.ListCourseV1Response, error) {
	log.Info().Msg("Request: ListCourseV1")

	log.Debug().
		Uint64("limit", req.GetLimit()).
		Uint64("offset", req.GetOffset()).
		Msg("request listCourseV1")

	list, err := g.repo.ListCourses(req.GetLimit(), req.GetOffset())
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}
	result := make([]*api.DescribeCourseV1Response, 0, req.Limit)
	for _, course := range list {
		result = append(result, convertCourseToResponse(&course))
	}
	return &api.ListCourseV1Response{Data: result}, nil
}

// RemoveCourseV1  Удаление курса
func (g *GRPCServer) RemoveCourseV1(ctx context.Context, req *api.RemoveCourseV1Request) (*api.RemoveCourseV1Response, error) {
	log.Info().Msg("Request: RemoveCourseV1")
	if req.GetId() == 0 {
		err := errors.New("id can't be zero")
		log.Error().Err(err).Msg("")
		return nil, err
	}
	err := g.repo.RemoveCourse(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("")
		return &api.RemoveCourseV1Response{Success: false}, err
	}
	return &api.RemoveCourseV1Response{Success: true}, nil
}

func convertCourseToResponse(origin *course.Course) *api.DescribeCourseV1Response {
	return &api.DescribeCourseV1Response{
		Id:          origin.Id(),
		UserId:      origin.UserId(),
		Name:        origin.Name(),
		Description: origin.Description(),
		DateStart:   timestamppb.New(origin.DateStart()),
		DateFinish:  timestamppb.New(origin.DateFinish()),
		DateCreated: timestamppb.New(origin.DateCreated()),
	}
}
