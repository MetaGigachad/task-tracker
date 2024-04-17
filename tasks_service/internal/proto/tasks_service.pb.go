// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: api/tasks_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type TaskStatus int32

const (
	TaskStatus_Open          TaskStatus = 0
	TaskStatus_InProgress    TaskStatus = 1
	TaskStatus_PendingReview TaskStatus = 2
	TaskStatus_Testing       TaskStatus = 3
	TaskStatus_Completed     TaskStatus = 4
	TaskStatus_Cancelled     TaskStatus = 5
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "Open",
		1: "InProgress",
		2: "PendingReview",
		3: "Testing",
		4: "Completed",
		5: "Cancelled",
	}
	TaskStatus_value = map[string]int32{
		"Open":          0,
		"InProgress":    1,
		"PendingReview": 2,
		"Testing":       3,
		"Completed":     4,
		"Cancelled":     5,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_tasks_service_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_api_tasks_service_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{0}
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tasks_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tasks_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskId string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tasks_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tasks_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTaskRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskId         string  `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	NewTitle       *string `protobuf:"bytes,3,opt,name=new_title,json=newTitle,proto3,oneof" json:"new_title,omitempty"`
	NewDescription *string `protobuf:"bytes,4,opt,name=new_description,json=newDescription,proto3,oneof" json:"new_description,omitempty"`
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tasks_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tasks_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTaskRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UpdateTaskRequest) GetNewTitle() string {
	if x != nil && x.NewTitle != nil {
		return *x.NewTitle
	}
	return ""
}

func (x *UpdateTaskRequest) GetNewDescription() string {
	if x != nil && x.NewDescription != nil {
		return *x.NewDescription
	}
	return ""
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskId string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tasks_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tasks_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTaskRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type GetTaskPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StartId  int32  `protobuf:"varint,2,opt,name=start_id,json=startId,proto3" json:"start_id,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetTaskPageRequest) Reset() {
	*x = GetTaskPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tasks_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskPageRequest) ProtoMessage() {}

func (x *GetTaskPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tasks_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskPageRequest.ProtoReflect.Descriptor instead.
func (*GetTaskPageRequest) Descriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTaskPageRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetTaskPageRequest) GetStartId() int32 {
	if x != nil {
		return x.StartId
	}
	return 0
}

func (x *GetTaskPageRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*TaskResponse_Task
	//	*TaskResponse_Error
	Response isTaskResponse_Response `protobuf_oneof:"response"`
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tasks_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tasks_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{5}
}

func (m *TaskResponse) GetResponse() isTaskResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *TaskResponse) GetTask() *Task {
	if x, ok := x.GetResponse().(*TaskResponse_Task); ok {
		return x.Task
	}
	return nil
}

func (x *TaskResponse) GetError() *Error {
	if x, ok := x.GetResponse().(*TaskResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isTaskResponse_Response interface {
	isTaskResponse_Response()
}

type TaskResponse_Task struct {
	Task *Task `protobuf:"bytes,1,opt,name=task,proto3,oneof"`
}

type TaskResponse_Error struct {
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*TaskResponse_Task) isTaskResponse_Response() {}

func (*TaskResponse_Error) isTaskResponse_Response() {}

type TaskPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*TaskPageResponse_TaskPage
	//	*TaskPageResponse_Error
	Response isTaskPageResponse_Response `protobuf_oneof:"response"`
}

func (x *TaskPageResponse) Reset() {
	*x = TaskPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tasks_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskPageResponse) ProtoMessage() {}

func (x *TaskPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tasks_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskPageResponse.ProtoReflect.Descriptor instead.
func (*TaskPageResponse) Descriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{6}
}

func (m *TaskPageResponse) GetResponse() isTaskPageResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *TaskPageResponse) GetTaskPage() *TaskPage {
	if x, ok := x.GetResponse().(*TaskPageResponse_TaskPage); ok {
		return x.TaskPage
	}
	return nil
}

func (x *TaskPageResponse) GetError() *Error {
	if x, ok := x.GetResponse().(*TaskPageResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isTaskPageResponse_Response interface {
	isTaskPageResponse_Response()
}

type TaskPageResponse_TaskPage struct {
	TaskPage *TaskPage `protobuf:"bytes,1,opt,name=task_page,json=taskPage,proto3,oneof"`
}

type TaskPageResponse_Error struct {
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*TaskPageResponse_TaskPage) isTaskPageResponse_Response() {}

func (*TaskPageResponse_Error) isTaskPageResponse_Response() {}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Title       string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status      TaskStatus             `protobuf:"varint,5,opt,name=status,proto3,enum=tasks_service.TaskStatus" json:"status,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tasks_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_api_tasks_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{7}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_Open
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tasks_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_api_tasks_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{8}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TaskPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TaskPage) Reset() {
	*x = TaskPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tasks_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskPage) ProtoMessage() {}

func (x *TaskPage) ProtoReflect() protoreflect.Message {
	mi := &file_api_tasks_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskPage.ProtoReflect.Descriptor instead.
func (*TaskPage) Descriptor() ([]byte, []int) {
	return file_api_tasks_service_proto_rawDescGZIP(), []int{9}
}

func (x *TaskPage) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_api_tasks_service_proto protoreflect.FileDescriptor

var file_api_tasks_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x6e,
	0x65, 0x77, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x6e, 0x65, 0x77, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a,
	0x0f, 0x6e, 0x65, 0x77, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x6e, 0x65, 0x77, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6e, 0x65,
	0x77, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x73, 0x0a, 0x0c, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x48, 0x00,
	0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x84, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x35, 0x0a,
	0x08, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2a, 0x64, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x05, 0x32, 0x8f, 0x03, 0x0a, 0x0c, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x20, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x1d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x20, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x20, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10,
	0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_tasks_service_proto_rawDescOnce sync.Once
	file_api_tasks_service_proto_rawDescData = file_api_tasks_service_proto_rawDesc
)

func file_api_tasks_service_proto_rawDescGZIP() []byte {
	file_api_tasks_service_proto_rawDescOnce.Do(func() {
		file_api_tasks_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_tasks_service_proto_rawDescData)
	})
	return file_api_tasks_service_proto_rawDescData
}

var file_api_tasks_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_tasks_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_tasks_service_proto_goTypes = []interface{}{
	(TaskStatus)(0),               // 0: tasks_service.TaskStatus
	(*CreateTaskRequest)(nil),     // 1: tasks_service.CreateTaskRequest
	(*GetTaskRequest)(nil),        // 2: tasks_service.GetTaskRequest
	(*UpdateTaskRequest)(nil),     // 3: tasks_service.UpdateTaskRequest
	(*DeleteTaskRequest)(nil),     // 4: tasks_service.DeleteTaskRequest
	(*GetTaskPageRequest)(nil),    // 5: tasks_service.GetTaskPageRequest
	(*TaskResponse)(nil),          // 6: tasks_service.TaskResponse
	(*TaskPageResponse)(nil),      // 7: tasks_service.TaskPageResponse
	(*Task)(nil),                  // 8: tasks_service.Task
	(*Error)(nil),                 // 9: tasks_service.Error
	(*TaskPage)(nil),              // 10: tasks_service.TaskPage
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_api_tasks_service_proto_depIdxs = []int32{
	8,  // 0: tasks_service.TaskResponse.task:type_name -> tasks_service.Task
	9,  // 1: tasks_service.TaskResponse.error:type_name -> tasks_service.Error
	10, // 2: tasks_service.TaskPageResponse.task_page:type_name -> tasks_service.TaskPage
	9,  // 3: tasks_service.TaskPageResponse.error:type_name -> tasks_service.Error
	11, // 4: tasks_service.Task.created_at:type_name -> google.protobuf.Timestamp
	0,  // 5: tasks_service.Task.status:type_name -> tasks_service.TaskStatus
	8,  // 6: tasks_service.TaskPage.tasks:type_name -> tasks_service.Task
	1,  // 7: tasks_service.TasksService.CreateTask:input_type -> tasks_service.CreateTaskRequest
	2,  // 8: tasks_service.TasksService.GetTask:input_type -> tasks_service.GetTaskRequest
	3,  // 9: tasks_service.TasksService.UpdateTask:input_type -> tasks_service.UpdateTaskRequest
	4,  // 10: tasks_service.TasksService.DeleteTask:input_type -> tasks_service.DeleteTaskRequest
	5,  // 11: tasks_service.TasksService.GetTaskPage:input_type -> tasks_service.GetTaskPageRequest
	6,  // 12: tasks_service.TasksService.CreateTask:output_type -> tasks_service.TaskResponse
	6,  // 13: tasks_service.TasksService.GetTask:output_type -> tasks_service.TaskResponse
	6,  // 14: tasks_service.TasksService.UpdateTask:output_type -> tasks_service.TaskResponse
	6,  // 15: tasks_service.TasksService.DeleteTask:output_type -> tasks_service.TaskResponse
	7,  // 16: tasks_service.TasksService.GetTaskPage:output_type -> tasks_service.TaskPageResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_tasks_service_proto_init() }
func file_api_tasks_service_proto_init() {
	if File_api_tasks_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_tasks_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_api_tasks_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskRequest); i {
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
		file_api_tasks_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskRequest); i {
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
		file_api_tasks_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskRequest); i {
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
		file_api_tasks_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskPageRequest); i {
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
		file_api_tasks_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResponse); i {
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
		file_api_tasks_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskPageResponse); i {
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
		file_api_tasks_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_api_tasks_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_api_tasks_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskPage); i {
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
	file_api_tasks_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_api_tasks_service_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*TaskResponse_Task)(nil),
		(*TaskResponse_Error)(nil),
	}
	file_api_tasks_service_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*TaskPageResponse_TaskPage)(nil),
		(*TaskPageResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_tasks_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_tasks_service_proto_goTypes,
		DependencyIndexes: file_api_tasks_service_proto_depIdxs,
		EnumInfos:         file_api_tasks_service_proto_enumTypes,
		MessageInfos:      file_api_tasks_service_proto_msgTypes,
	}.Build()
	File_api_tasks_service_proto = out.File
	file_api_tasks_service_proto_rawDesc = nil
	file_api_tasks_service_proto_goTypes = nil
	file_api_tasks_service_proto_depIdxs = nil
}
