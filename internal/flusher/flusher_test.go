package flusher_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-course-api/internal/course"
	"github.com/ozonva/ova-course-api/internal/repo"
	"time"

	"github.com/ozonva/ova-course-api/internal/flusher"
)

var _ = Describe("Flusher", func() {
	const chunkSize = 2
	var (
		mockCtrl    *gomock.Controller
		mockRepo    *repo.MockCourseRepo
		testFlusher flusher.Flusher

		courses = []course.Course{
			course.New(1, "", time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, time.April, 14, 0, 0, 0, 0, time.UTC)),
			course.New(2, "", time.Date(2021, time.May, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, time.May, 14, 0, 0, 0, 0, time.UTC)),
			course.New(4, "", time.Date(2021, time.June, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, time.June, 14, 0, 0, 0, 0, time.UTC)),
		}
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockCourseRepo(mockCtrl)
		testFlusher = flusher.NewFlusher(chunkSize, mockRepo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Save data", func() {
		When("save all data", func() {
			AssertReturnNil := func(courses []course.Course) {
				Expect(testFlusher.Flush(courses)).To(BeNil())
			}
			Context("count of courses less than chunkSize", func() {
				oneCourse := courses[:1]
				BeforeEach(func() {
					mockRepo.EXPECT().AddCourse(oneCourse).Return(nil).Times(1)
				})
				It("Flush should return nil", func() {
					AssertReturnNil(oneCourse)
				})
			})
			Context("count of courses more than chunkSize", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddCourse(courses[:2]).Return(nil).Times(1),
						mockRepo.EXPECT().AddCourse(courses[2:]).Return(nil).Times(1),
					)
				})
				It("Flush should return nil", func() {
					AssertReturnNil(courses)
				})
			})
		})
		When("save with error", func() {
			err := fmt.Errorf("can't save data")
			AssertFlushReturnCourses := func(CoursesIn []course.Course, CoursesOut []course.Course) {
				Expect(testFlusher.Flush(CoursesIn)).To(Equal(CoursesOut))
			}
			Context("all courses", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddCourse(courses[:2]).Return(err).Times(1),
						mockRepo.EXPECT().AddCourse(courses[2:]).Return(err).Times(1),
					)
				})
				It("Flush should return all courses", func() {
					AssertFlushReturnCourses(courses, courses)
				})
			})
			Context("part of courses", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddCourse(courses[:2]).Return(nil).Times(1),
						mockRepo.EXPECT().AddCourse(courses[2:]).Return(err).Times(1),
					)
				})
				It("Flush should return nil", func() {
					AssertFlushReturnCourses(courses, courses[2:])
				})
			})
		})
	})
})
