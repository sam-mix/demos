// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: todo-list.proto

package todolist

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskName string `protobuf:"bytes,1,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
}

func (x *AddTaskRequest) Reset() {
	*x = AddTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskRequest) ProtoMessage() {}

func (x *AddTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskRequest.ProtoReflect.Descriptor instead.
func (*AddTaskRequest) Descriptor() ([]byte, []int) {
	return file_todo_list_proto_rawDescGZIP(), []int{0}
}

func (x *AddTaskRequest) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

type AddTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *AddTaskResponse) Reset() {
	*x = AddTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskResponse) ProtoMessage() {}

func (x *AddTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskResponse.ProtoReflect.Descriptor instead.
func (*AddTaskResponse) Descriptor() ([]byte, []int) {
	return file_todo_list_proto_rawDescGZIP(), []int{1}
}

func (x *AddTaskResponse) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_proto_msgTypes[2]
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
	return file_todo_list_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTaskRequest) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_todo_list_proto_rawDescGZIP(), []int{3}
}

type GetTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTasksRequest) Reset() {
	*x = GetTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksRequest) ProtoMessage() {}

func (x *GetTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksRequest.ProtoReflect.Descriptor instead.
func (*GetTasksRequest) Descriptor() ([]byte, []int) {
	return file_todo_list_proto_rawDescGZIP(), []int{4}
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_proto_msgTypes[5]
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
	return file_todo_list_proto_rawDescGZIP(), []int{5}
}

func (x *Task) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetTasksResponse) Reset() {
	*x = GetTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksResponse) ProtoMessage() {}

func (x *GetTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksResponse.ProtoReflect.Descriptor instead.
func (*GetTasksResponse) Descriptor() ([]byte, []int) {
	return file_todo_list_proto_rawDescGZIP(), []int{6}
}

func (x *GetTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId   int32  `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TaskName string `protobuf:"bytes,2,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_proto_msgTypes[7]
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
	return file_todo_list_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTaskRequest) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *UpdateTaskRequest) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

type UpdateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTaskResponse) Reset() {
	*x = UpdateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_list_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponse) ProtoMessage() {}

func (x *UpdateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_list_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return file_todo_list_proto_rawDescGZIP(), []int{8}
}

var File_todo_list_proto protoreflect.FileDescriptor

var file_todo_list_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x22, 0x49, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x14,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa7, 0x02, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x1b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x19, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x1b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x3b, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_list_proto_rawDescOnce sync.Once
	file_todo_list_proto_rawDescData = file_todo_list_proto_rawDesc
)

func file_todo_list_proto_rawDescGZIP() []byte {
	file_todo_list_proto_rawDescOnce.Do(func() {
		file_todo_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_list_proto_rawDescData)
	})
	return file_todo_list_proto_rawDescData
}

var file_todo_list_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_todo_list_proto_goTypes = []interface{}{
	(*AddTaskRequest)(nil),     // 0: todolist.AddTaskRequest
	(*AddTaskResponse)(nil),    // 1: todolist.AddTaskResponse
	(*DeleteTaskRequest)(nil),  // 2: todolist.DeleteTaskRequest
	(*DeleteTaskResponse)(nil), // 3: todolist.DeleteTaskResponse
	(*GetTasksRequest)(nil),    // 4: todolist.GetTasksRequest
	(*Task)(nil),               // 5: todolist.Task
	(*GetTasksResponse)(nil),   // 6: todolist.GetTasksResponse
	(*UpdateTaskRequest)(nil),  // 7: todolist.UpdateTaskRequest
	(*UpdateTaskResponse)(nil), // 8: todolist.UpdateTaskResponse
}
var file_todo_list_proto_depIdxs = []int32{
	5, // 0: todolist.GetTasksResponse.tasks:type_name -> todolist.Task
	0, // 1: todolist.TodoList.AddTask:input_type -> todolist.AddTaskRequest
	2, // 2: todolist.TodoList.DeleteTask:input_type -> todolist.DeleteTaskRequest
	4, // 3: todolist.TodoList.GetTasks:input_type -> todolist.GetTasksRequest
	7, // 4: todolist.TodoList.UpdateTask:input_type -> todolist.UpdateTaskRequest
	1, // 5: todolist.TodoList.AddTask:output_type -> todolist.AddTaskResponse
	3, // 6: todolist.TodoList.DeleteTask:output_type -> todolist.DeleteTaskResponse
	6, // 7: todolist.TodoList.GetTasks:output_type -> todolist.GetTasksResponse
	8, // 8: todolist.TodoList.UpdateTask:output_type -> todolist.UpdateTaskResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_todo_list_proto_init() }
func file_todo_list_proto_init() {
	if File_todo_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskRequest); i {
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
		file_todo_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskResponse); i {
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
		file_todo_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_todo_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskResponse); i {
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
		file_todo_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTasksRequest); i {
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
		file_todo_list_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_todo_list_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTasksResponse); i {
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
		file_todo_list_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_todo_list_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskResponse); i {
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
			RawDescriptor: file_todo_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_list_proto_goTypes,
		DependencyIndexes: file_todo_list_proto_depIdxs,
		MessageInfos:      file_todo_list_proto_msgTypes,
	}.Build()
	File_todo_list_proto = out.File
	file_todo_list_proto_rawDesc = nil
	file_todo_list_proto_goTypes = nil
	file_todo_list_proto_depIdxs = nil
}
