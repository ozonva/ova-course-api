// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/ova-course-api/ova-course-api.proto

package ova_course_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateCourseV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint64                 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DateStart  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateFinish *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_finish,json=dateFinish,proto3" json:"date_finish,omitempty"`
}

func (x *CreateCourseV1Request) Reset() {
	*x = CreateCourseV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseV1Request) ProtoMessage() {}

func (x *CreateCourseV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseV1Request.ProtoReflect.Descriptor instead.
func (*CreateCourseV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_course_api_ova_course_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCourseV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCourseV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCourseV1Request) GetDateStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DateStart
	}
	return nil
}

func (x *CreateCourseV1Request) GetDateFinish() *timestamppb.Timestamp {
	if x != nil {
		return x.DateFinish
	}
	return nil
}

type CreateCourseV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateCourseV1Response) Reset() {
	*x = CreateCourseV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseV1Response) ProtoMessage() {}

func (x *CreateCourseV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseV1Response.ProtoReflect.Descriptor instead.
func (*CreateCourseV1Response) Descriptor() ([]byte, []int) {
	return file_api_ova_course_api_ova_course_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCourseV1Response) GetId() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DescribeCourseV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeCourseV1Request) Reset() {
	*x = DescribeCourseV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCourseV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCourseV1Request) ProtoMessage() {}

func (x *DescribeCourseV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCourseV1Request.ProtoReflect.Descriptor instead.
func (*DescribeCourseV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_course_api_ova_course_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeCourseV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeCourseV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      uint64                 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DateStart   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateFinish  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_finish,json=dateFinish,proto3" json:"date_finish,omitempty"`
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
}

func (x *DescribeCourseV1Response) Reset() {
	*x = DescribeCourseV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCourseV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCourseV1Response) ProtoMessage() {}

func (x *DescribeCourseV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCourseV1Response.ProtoReflect.Descriptor instead.
func (*DescribeCourseV1Response) Descriptor() ([]byte, []int) {
	return file_api_ova_course_api_ova_course_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeCourseV1Response) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DescribeCourseV1Response) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DescribeCourseV1Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeCourseV1Response) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DescribeCourseV1Response) GetDateStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DateStart
	}
	return nil
}

func (x *DescribeCourseV1Response) GetDateFinish() *timestamppb.Timestamp {
	if x != nil {
		return x.DateFinish
	}
	return nil
}

func (x *DescribeCourseV1Response) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

type ListCourseV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListCourseV1Request) Reset() {
	*x = ListCourseV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCourseV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCourseV1Request) ProtoMessage() {}

func (x *ListCourseV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCourseV1Request.ProtoReflect.Descriptor instead.
func (*ListCourseV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_course_api_ova_course_api_proto_rawDescGZIP(), []int{0}
}

func (x *ListCourseV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListCourseV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListCourseV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*DescribeCourseV1Response `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListCourseV1Response) Reset() {
	*x = ListCourseV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCourseV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCourseV1Response) ProtoMessage() {}

func (x *ListCourseV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCourseV1Response.ProtoReflect.Descriptor instead.
func (*ListCourseV1Response) Descriptor() ([]byte, []int) {
	return file_api_ova_course_api_ova_course_api_proto_rawDescGZIP(), []int{4}
}

func (x *ListCourseV1Response) GetData() []*DescribeCourseV1Response {
	if x != nil {
		return x.Data
	}
	return nil
}

type RemoveCourseV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveCourseV1Request) Reset() {
	*x = RemoveCourseV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCourseV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCourseV1Request) ProtoMessage() {}

func (x *RemoveCourseV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCourseV1Request.ProtoReflect.Descriptor instead.
func (*RemoveCourseV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_course_api_ova_course_api_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveCourseV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveCourseV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveCourseV1Response) Reset() {
	*x = RemoveCourseV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCourseV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCourseV1Response) ProtoMessage() {}

func (x *RemoveCourseV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_course_api_ova_course_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCourseV1Response.ProtoReflect.Descriptor instead.
func (*RemoveCourseV1Response) Descriptor() ([]byte, []int) {
	return file_api_ova_course_api_ova_course_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveCourseV1Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_ova_course_api_ova_course_api_proto protoreflect.FileDescriptor

var file_api_ova_course_api_ova_course_api_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x76, 0x61, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x29, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xaf, 0x02, 0x0a, 0x18, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x3d, 0x0a,
	0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x54, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x27, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32,
	0x87, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x12, 0x25, 0x2e, 0x6f,
	0x76, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a,
	0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56,
	0x31, 0x12, 0x27, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6f, 0x76, 0x61,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x24,
	0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x12, 0x25, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2f, 0x6f,
	0x76, 0x61, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x69,
	0x3b, 0x6f, 0x76, 0x61, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ova_course_api_ova_course_api_proto_rawDescOnce sync.Once
	file_api_ova_course_api_ova_course_api_proto_rawDescData = file_api_ova_course_api_ova_course_api_proto_rawDesc
)

func file_api_ova_course_api_ova_course_api_proto_rawDescGZIP() []byte {
	file_api_ova_course_api_ova_course_api_proto_rawDescOnce.Do(func() {
		file_api_ova_course_api_ova_course_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ova_course_api_ova_course_api_proto_rawDescData)
	})
	return file_api_ova_course_api_ova_course_api_proto_rawDescData
}

var file_api_ova_course_api_ova_course_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_ova_course_api_ova_course_api_proto_goTypes = []interface{}{
	(*CreateCourseV1Request)(nil),    // 0: ova.course.api.CreateCourseV1Request
	(*CreateCourseV1Response)(nil),   // 1: ova.course.api.CreateCourseV1Response
	(*DescribeCourseV1Request)(nil),  // 2: ova.course.api.DescribeCourseV1Request
	(*DescribeCourseV1Response)(nil), // 3: ova.course.api.DescribeCourseV1Response
	(*ListCourseV1Response)(nil),     // 4: ova.course.api.ListCourseV1Response
	(*RemoveCourseV1Request)(nil),    // 5: ova.course.api.RemoveCourseV1Request
	(*RemoveCourseV1Response)(nil),   // 6: ova.course.api.RemoveCourseV1Response
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_api_ova_course_api_ova_course_api_proto_depIdxs = []int32{
	7,  // 0: ova.course.api.CreateCourseV1Request.date_start:type_name -> google.protobuf.Timestamp
	7,  // 1: ova.course.api.CreateCourseV1Request.date_finish:type_name -> google.protobuf.Timestamp
	7,  // 2: ova.course.api.DescribeCourseV1Response.date_start:type_name -> google.protobuf.Timestamp
	7,  // 3: ova.course.api.DescribeCourseV1Response.date_finish:type_name -> google.protobuf.Timestamp
	7,  // 4: ova.course.api.DescribeCourseV1Response.date_created:type_name -> google.protobuf.Timestamp
	3,  // 5: ova.course.api.ListCourseV1Response.data:type_name -> ova.course.api.DescribeCourseV1Response
	0,  // 6: ova.course.api.Course.CreateCourseV1:input_type -> ova.course.api.CreateCourseV1Request
	2,  // 7: ova.course.api.Course.DescribeCourseV1:input_type -> ova.course.api.DescribeCourseV1Request
	8,  // 8: ova.course.api.Course.ListCourseV1:input_type -> google.protobuf.Empty
	5,  // 9: ova.course.api.Course.RemoveCourseV1:input_type -> ova.course.api.RemoveCourseV1Request
	1,  // 10: ova.course.api.Course.CreateCourseV1:output_type -> ova.course.api.CreateCourseV1Response
	3,  // 11: ova.course.api.Course.DescribeCourseV1:output_type -> ova.course.api.DescribeCourseV1Response
	4,  // 12: ova.course.api.Course.ListCourseV1:output_type -> ova.course.api.ListCourseV1Response
	6,  // 13: ova.course.api.Course.RemoveCourseV1:output_type -> ova.course.api.RemoveCourseV1Response
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_ova_course_api_ova_course_api_proto_init() }
func file_api_ova_course_api_ova_course_api_proto_init() {
	if File_api_ova_course_api_ova_course_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ova_course_api_ova_course_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCourseV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ova_course_api_ova_course_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCourseV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ova_course_api_ova_course_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCourseV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ova_course_api_ova_course_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCourseV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ova_course_api_ova_course_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCourseV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ova_course_api_ova_course_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCourseV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ova_course_api_ova_course_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCourseV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_ova_course_api_ova_course_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ova_course_api_ova_course_api_proto_goTypes,
		DependencyIndexes: file_api_ova_course_api_ova_course_api_proto_depIdxs,
		MessageInfos:      file_api_ova_course_api_ova_course_api_proto_msgTypes,
	}.Build()
	File_api_ova_course_api_ova_course_api_proto = out.File
	file_api_ova_course_api_ova_course_api_proto_rawDesc = nil
	file_api_ova_course_api_ova_course_api_proto_goTypes = nil
	file_api_ova_course_api_ova_course_api_proto_depIdxs = nil
}
