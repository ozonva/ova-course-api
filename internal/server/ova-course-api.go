package ova_course_api

import (
	api "github.com/ozonva/ova-course-api/pkg/ova-course-api"
)

type GRPCServer struct {
	api.UnimplementedCourseServer
}

func NewCourseServer() api.CourseServer {
	return &GRPCServer{}
}
