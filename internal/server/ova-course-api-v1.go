package ova_course_api

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	opentracingLog "github.com/opentracing/opentracing-go/log"
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
func (g *GRPCServer) CreateCourseV1(ctx context.Context, req *api.CreateCourseV1Request) (*api.SuccessV1, error) {
	log.Info().Msg("Request: CreateCourseV1")

	log.Debug().
		Uint64("userId", req.UserId).
		Str("name", req.Name).
		Time("dateStart", req.DateStart.AsTime()).
		Time("dateFinish", req.DateFinish.AsTime()).
		Msg("request createcourseV1")

	span, ctx := opentracing.StartSpanFromContext(ctx, "create_course_v1")
	defer span.Finish()
	span.LogFields(opentracingLog.Int("count of courses", 1))

	item, err := convertItemToCourse(req)
	courseId, err := g.repo.AddOneCourse(item)

	if err != nil {
		return &api.SuccessV1{Success: false}, err
	}

	item.SetId(uint64(courseId))

	// TODO Нужно делать в транзакции на случай если сообщение не отошлется, откатить insert
	err = g.kafkaProducer.CreateEvent(item)
	if err != nil {
		return &api.SuccessV1{Success: false}, err
	}

	g.metrics.CreateSuccessInc()
	return &api.SuccessV1{Success: true}, err
}

func (g *GRPCServer) MultiCreateCourseV1(ctx context.Context, req *api.MultiCreateCourseV1Request) (*api.MultiCreateCourseV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "multi_create_course_v1")
	defer span.Finish()
	span.LogFields(opentracingLog.Int("count of courses", len(req.Data)))

	length := len(req.GetData())
	flush := flusher.NewFlusher(length, g.repo)

	srv := saver.NewSaver(ctx, span, uint(length), flush, addPeriod)
	defer srv.Close()
	result := make([]*api.SuccessErrorV1, 0, length)
	var errorGeneral error = nil
	for _, item := range req.GetData() {
		success := true
		var errMsg string
		itemCourse, err := convertItemToCourse(item)
		if err != nil {
			success = false
			errMsg = err.Error()
			if errorGeneral == nil {
				errorGeneral = errors.New("an error occurred while saving objects")
			}
		}
		srv.Save(itemCourse)
		result = append(result, &api.SuccessErrorV1{Success: success, Error: errMsg})
	}

	return &api.MultiCreateCourseV1Response{Data: result}, errorGeneral
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
	g.metrics.UpdateSuccessInc()
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
func (g *GRPCServer) RemoveCourseV1(ctx context.Context, req *api.RemoveCourseV1Request) (*api.SuccessV1, error) {
	log.Info().Msg("Request: RemoveCourseV1")
	if req.GetId() == 0 {
		err := errors.New("id can't be zero")
		log.Error().Err(err).Msg("")
		return nil, err
	}
	err := g.repo.RemoveCourse(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("")
		return &api.SuccessV1{Success: false}, err
	}

	// TODO Нужно делать в транзакции на случай если сообщение не отошлется, откатить delete
	err = g.kafkaProducer.DeleteEvent(req.GetId())
	if err != nil {
		return &api.SuccessV1{Success: false}, err
	}

	g.metrics.RemoveSuccessInc()
	return &api.SuccessV1{Success: true}, nil
}

// UpdateCourseV1  Обновление курса
func (g *GRPCServer) UpdateCourseV1(ctx context.Context, req *api.UpdateCourseV1Request) (*api.SuccessV1, error) {
	if req.GetId() == 0 {
		err := errors.New("id can't be zero")
		log.Error().Err(err).Msg("")
		return nil, err
	}

	course := course.New(req.GetUserId(), req.GetName(), req.GetDateStart().AsTime(), req.GetDateFinish().AsTime())
	course.SetDescription(req.GetDescription())
	course.SetId(req.GetId())
	err := g.repo.UpdateCourse(course)

	if err != nil {
		log.Error().Err(err).Msg("")
		return &api.SuccessV1{Success: false}, err
	}

	// TODO Нужно делать в транзакции на случай если сообщение не отошлется, откатить update
	err = g.kafkaProducer.UpdateEvent(course)
	if err != nil {
		return &api.SuccessV1{Success: false}, err
	}
	return &api.SuccessV1{Success: true}, nil
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

func convertItemToCourse(req *api.CreateCourseV1Request) (course.Course, error) {
	if err := req.DateStart.CheckValid(); err != nil {
		log.Error().Err(err).Msg("")
		return course.Course{}, err
	}

	if err := req.DateFinish.CheckValid(); err != nil {
		log.Error().Err(err).Msg("")
		return course.Course{}, err
	}

	obj := course.New(req.GetUserId(), req.GetName(), req.GetDateStart().AsTime(), req.GetDateFinish().AsTime())

	return obj, nil
}
