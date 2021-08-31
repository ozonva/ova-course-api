package ova_course_api

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-course-api/internal/course"
	api "github.com/ozonva/ova-course-api/pkg/ova-course-api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"github.com/ozonva/ova-course-api/internal/repo"
)

var _ = Describe("OvaCourseApiV1", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *repo.MockCourseRepo
		ctx      context.Context

		courses = []course.Course{
			course.New(1, "", time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, time.April, 14, 0, 0, 0, 0, time.UTC)),
			course.New(2, "", time.Date(2021, time.May, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, time.May, 14, 0, 0, 0, 0, time.UTC)),
			course.New(4, "", time.Date(2021, time.June, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, time.June, 14, 0, 0, 0, 0, time.UTC)),
		}
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockCourseRepo(mockCtrl)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("CreateCourseV1", func() {
		It("no errors", func() {
			oneCourse := courses[:1]
			mockRepo.EXPECT().AddCourse(oneCourse).Return(nil).Times(1)

			srv := &GRPCServer{repo: mockRepo}
			_, err := srv.CreateCourseV1(ctx, &api.CreateCourseV1Request{
				UserId:     courses[0].UserId(),
				Name:       courses[0].Name(),
				DateStart:  timestamppb.New(courses[0].DateStart()),
				DateFinish: timestamppb.New(courses[0].DateFinish()),
			})
			Expect(err).To(BeNil())
		})
	})

	Context("RemoveCourseV1", func() {
		It("no errors", func() {

			id := uint64(1)
			mockRepo.EXPECT().RemoveCourse(id).Return(nil).Times(1)

			srv := &GRPCServer{repo: mockRepo}
			_, err := srv.RemoveCourseV1(ctx, &api.RemoveCourseV1Request{Id: id})

			Expect(err).To(BeNil())

		})
	})

	Context("ListCourseV1", func() {
		It("no errors", func() {

			limit := uint64(10)
			offset := uint64(0)
			mockRepo.EXPECT().ListCourses(limit, offset).Return(courses, nil).Times(1)

			srv := &GRPCServer{repo: mockRepo}
			_, err := srv.ListCourseV1(ctx, &api.ListCourseV1Request{Limit: limit, Offset: offset})

			Expect(err).To(BeNil())

		})
	})

	Context("DescribeCourseV1", func() {
		It("no errors", func() {

			id := uint64(1)
			mockRepo.EXPECT().DescribeCourses(id).Return(courses[0], nil).Times(1)

			srv := &GRPCServer{repo: mockRepo}
			_, err := srv.DescribeCourseV1(ctx, &api.DescribeCourseV1Request{Id: id})

			Expect(err).To(BeNil())

		})
	})

})
