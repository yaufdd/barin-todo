// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/todo/v1/todo.proto

package todov1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_UNKNOWN Status = 0
	Status_OPEN    Status = 1
	Status_DONE    Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "OPEN",
		2: "DONE",
	}
	Status_value = map[string]int32{
		"UNKNOWN": 0,
		"OPEN":    1,
		"DONE":    2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_todo_v1_todo_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_proto_todo_v1_todo_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{0}
}

type SortBy int32

const (
	SortBy_CREATED_AT SortBy = 0
	SortBy_DEADLINE   SortBy = 1
	SortBy_TITLE      SortBy = 2
	SortBy_STATUS     SortBy = 3
)

// Enum value maps for SortBy.
var (
	SortBy_name = map[int32]string{
		0: "CREATED_AT",
		1: "DEADLINE",
		2: "TITLE",
		3: "STATUS",
	}
	SortBy_value = map[string]int32{
		"CREATED_AT": 0,
		"DEADLINE":   1,
		"TITLE":      2,
		"STATUS":     3,
	}
)

func (x SortBy) Enum() *SortBy {
	p := new(SortBy)
	*p = x
	return p
}

func (x SortBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortBy) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_todo_v1_todo_proto_enumTypes[1].Descriptor()
}

func (SortBy) Type() protoreflect.EnumType {
	return &file_proto_todo_v1_todo_proto_enumTypes[1]
}

func (x SortBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortBy.Descriptor instead.
func (SortBy) EnumDescriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{1}
}

type SortOrder int32

const (
	SortOrder_ASC  SortOrder = 0
	SortOrder_DESC SortOrder = 1
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	SortOrder_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_todo_v1_todo_proto_enumTypes[2].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_proto_todo_v1_todo_proto_enumTypes[2]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{2}
}

type Todo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status        Status                 `protobuf:"varint,4,opt,name=status,proto3,enum=todo.v1.Status" json:"status,omitempty"`
	Deadline      string                 `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Deleted       bool                   `protobuf:"varint,8,opt,name=deleted,proto3" json:"deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Todo) Reset() {
	*x = Todo{}
	mi := &file_proto_todo_v1_todo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_v1_todo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Todo) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

func (x *Todo) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *Todo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Todo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Todo) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type CreateTodoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Deadline      string                 `protobuf:"bytes,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	mi := &file_proto_todo_v1_todo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_v1_todo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTodoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTodoRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTodoRequest) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

type CreateTodoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Todo          *Todo                  `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTodoResponse) Reset() {
	*x = CreateTodoResponse{}
	mi := &file_proto_todo_v1_todo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoResponse) ProtoMessage() {}

func (x *CreateTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_v1_todo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoResponse.ProtoReflect.Descriptor instead.
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTodoResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type UpdateTodoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status        Status                 `protobuf:"varint,4,opt,name=status,proto3,enum=todo.v1.Status" json:"status,omitempty"`
	Deadline      string                 `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTodoRequest) Reset() {
	*x = UpdateTodoRequest{}
	mi := &file_proto_todo_v1_todo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoRequest) ProtoMessage() {}

func (x *UpdateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_v1_todo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoRequest.ProtoReflect.Descriptor instead.
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTodoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTodoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTodoRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateTodoRequest) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

func (x *UpdateTodoRequest) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

type UpdateTodoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTodoResponse) Reset() {
	*x = UpdateTodoResponse{}
	mi := &file_proto_todo_v1_todo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoResponse) ProtoMessage() {}

func (x *UpdateTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_v1_todo_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoResponse.ProtoReflect.Descriptor instead.
func (*UpdateTodoResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{4}
}

type DeleteTodoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTodoRequest) Reset() {
	*x = DeleteTodoRequest{}
	mi := &file_proto_todo_v1_todo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoRequest) ProtoMessage() {}

func (x *DeleteTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_v1_todo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoRequest.ProtoReflect.Descriptor instead.
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTodoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTodoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTodoResponse) Reset() {
	*x = DeleteTodoResponse{}
	mi := &file_proto_todo_v1_todo_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoResponse) ProtoMessage() {}

func (x *DeleteTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_v1_todo_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoResponse.ProtoReflect.Descriptor instead.
func (*DeleteTodoResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{6}
}

type ListTodosRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusFilter  Status                 `protobuf:"varint,1,opt,name=status_filter,json=statusFilter,proto3,enum=todo.v1.Status" json:"status_filter,omitempty"`
	SortBy        SortBy                 `protobuf:"varint,2,opt,name=sort_by,json=sortBy,proto3,enum=todo.v1.SortBy" json:"sort_by,omitempty"`
	SortOrder     SortOrder              `protobuf:"varint,3,opt,name=sort_order,json=sortOrder,proto3,enum=todo.v1.SortOrder" json:"sort_order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTodosRequest) Reset() {
	*x = ListTodosRequest{}
	mi := &file_proto_todo_v1_todo_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodosRequest) ProtoMessage() {}

func (x *ListTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_v1_todo_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodosRequest.ProtoReflect.Descriptor instead.
func (*ListTodosRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{7}
}

func (x *ListTodosRequest) GetStatusFilter() Status {
	if x != nil {
		return x.StatusFilter
	}
	return Status_UNKNOWN
}

func (x *ListTodosRequest) GetSortBy() SortBy {
	if x != nil {
		return x.SortBy
	}
	return SortBy_CREATED_AT
}

func (x *ListTodosRequest) GetSortOrder() SortOrder {
	if x != nil {
		return x.SortOrder
	}
	return SortOrder_ASC
}

type ListTodosResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Todos         []*Todo                `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTodosResponse) Reset() {
	*x = ListTodosResponse{}
	mi := &file_proto_todo_v1_todo_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodosResponse) ProtoMessage() {}

func (x *ListTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_v1_todo_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodosResponse.ProtoReflect.Descriptor instead.
func (*ListTodosResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_v1_todo_proto_rawDescGZIP(), []int{8}
}

func (x *ListTodosResponse) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

var File_proto_todo_v1_todo_proto protoreflect.FileDescriptor

const file_proto_todo_v1_todo_proto_rawDesc = "" +
	"\n" +
	"\x18proto/todo/v1/todo.proto\x12\atodo.v1\"\xeb\x01\n" +
	"\x04Todo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12'\n" +
	"\x06status\x18\x04 \x01(\x0e2\x0f.todo.v1.StatusR\x06status\x12\x1a\n" +
	"\bdeadline\x18\x05 \x01(\tR\bdeadline\x12\x1d\n" +
	"\n" +
	"created_at\x18\x06 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\a \x01(\tR\tupdatedAt\x12\x18\n" +
	"\adeleted\x18\b \x01(\bR\adeleted\"g\n" +
	"\x11CreateTodoRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x1a\n" +
	"\bdeadline\x18\x03 \x01(\tR\bdeadline\"7\n" +
	"\x12CreateTodoResponse\x12!\n" +
	"\x04todo\x18\x01 \x01(\v2\r.todo.v1.TodoR\x04todo\"\xa0\x01\n" +
	"\x11UpdateTodoRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12'\n" +
	"\x06status\x18\x04 \x01(\x0e2\x0f.todo.v1.StatusR\x06status\x12\x1a\n" +
	"\bdeadline\x18\x05 \x01(\tR\bdeadline\"\x14\n" +
	"\x12UpdateTodoResponse\"#\n" +
	"\x11DeleteTodoRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"\x14\n" +
	"\x12DeleteTodoResponse\"\xa5\x01\n" +
	"\x10ListTodosRequest\x124\n" +
	"\rstatus_filter\x18\x01 \x01(\x0e2\x0f.todo.v1.StatusR\fstatusFilter\x12(\n" +
	"\asort_by\x18\x02 \x01(\x0e2\x0f.todo.v1.SortByR\x06sortBy\x121\n" +
	"\n" +
	"sort_order\x18\x03 \x01(\x0e2\x12.todo.v1.SortOrderR\tsortOrder\"8\n" +
	"\x11ListTodosResponse\x12#\n" +
	"\x05todos\x18\x01 \x03(\v2\r.todo.v1.TodoR\x05todos*)\n" +
	"\x06Status\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\b\n" +
	"\x04OPEN\x10\x01\x12\b\n" +
	"\x04DONE\x10\x02*=\n" +
	"\x06SortBy\x12\x0e\n" +
	"\n" +
	"CREATED_AT\x10\x00\x12\f\n" +
	"\bDEADLINE\x10\x01\x12\t\n" +
	"\x05TITLE\x10\x02\x12\n" +
	"\n" +
	"\x06STATUS\x10\x03*\x1e\n" +
	"\tSortOrder\x12\a\n" +
	"\x03ASC\x10\x00\x12\b\n" +
	"\x04DESC\x10\x012\xa6\x02\n" +
	"\vTodoService\x12E\n" +
	"\n" +
	"CreateTodo\x12\x1a.todo.v1.CreateTodoRequest\x1a\x1b.todo.v1.CreateTodoResponse\x12E\n" +
	"\n" +
	"UpdateTodo\x12\x1a.todo.v1.UpdateTodoRequest\x1a\x1b.todo.v1.UpdateTodoResponse\x12E\n" +
	"\n" +
	"DeleteTodo\x12\x1a.todo.v1.DeleteTodoRequest\x1a\x1b.todo.v1.DeleteTodoResponse\x12B\n" +
	"\tListTodos\x12\x19.todo.v1.ListTodosRequest\x1a\x1a.todo.v1.ListTodosResponseB:Z8github.com/ericbamba/barin-todo/gen/proto/todo/v1;todov1b\x06proto3"

var (
	file_proto_todo_v1_todo_proto_rawDescOnce sync.Once
	file_proto_todo_v1_todo_proto_rawDescData []byte
)

func file_proto_todo_v1_todo_proto_rawDescGZIP() []byte {
	file_proto_todo_v1_todo_proto_rawDescOnce.Do(func() {
		file_proto_todo_v1_todo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_todo_v1_todo_proto_rawDesc), len(file_proto_todo_v1_todo_proto_rawDesc)))
	})
	return file_proto_todo_v1_todo_proto_rawDescData
}

var file_proto_todo_v1_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_todo_v1_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_todo_v1_todo_proto_goTypes = []any{
	(Status)(0),                // 0: todo.v1.Status
	(SortBy)(0),                // 1: todo.v1.SortBy
	(SortOrder)(0),             // 2: todo.v1.SortOrder
	(*Todo)(nil),               // 3: todo.v1.Todo
	(*CreateTodoRequest)(nil),  // 4: todo.v1.CreateTodoRequest
	(*CreateTodoResponse)(nil), // 5: todo.v1.CreateTodoResponse
	(*UpdateTodoRequest)(nil),  // 6: todo.v1.UpdateTodoRequest
	(*UpdateTodoResponse)(nil), // 7: todo.v1.UpdateTodoResponse
	(*DeleteTodoRequest)(nil),  // 8: todo.v1.DeleteTodoRequest
	(*DeleteTodoResponse)(nil), // 9: todo.v1.DeleteTodoResponse
	(*ListTodosRequest)(nil),   // 10: todo.v1.ListTodosRequest
	(*ListTodosResponse)(nil),  // 11: todo.v1.ListTodosResponse
}
var file_proto_todo_v1_todo_proto_depIdxs = []int32{
	0,  // 0: todo.v1.Todo.status:type_name -> todo.v1.Status
	3,  // 1: todo.v1.CreateTodoResponse.todo:type_name -> todo.v1.Todo
	0,  // 2: todo.v1.UpdateTodoRequest.status:type_name -> todo.v1.Status
	0,  // 3: todo.v1.ListTodosRequest.status_filter:type_name -> todo.v1.Status
	1,  // 4: todo.v1.ListTodosRequest.sort_by:type_name -> todo.v1.SortBy
	2,  // 5: todo.v1.ListTodosRequest.sort_order:type_name -> todo.v1.SortOrder
	3,  // 6: todo.v1.ListTodosResponse.todos:type_name -> todo.v1.Todo
	4,  // 7: todo.v1.TodoService.CreateTodo:input_type -> todo.v1.CreateTodoRequest
	6,  // 8: todo.v1.TodoService.UpdateTodo:input_type -> todo.v1.UpdateTodoRequest
	8,  // 9: todo.v1.TodoService.DeleteTodo:input_type -> todo.v1.DeleteTodoRequest
	10, // 10: todo.v1.TodoService.ListTodos:input_type -> todo.v1.ListTodosRequest
	5,  // 11: todo.v1.TodoService.CreateTodo:output_type -> todo.v1.CreateTodoResponse
	7,  // 12: todo.v1.TodoService.UpdateTodo:output_type -> todo.v1.UpdateTodoResponse
	9,  // 13: todo.v1.TodoService.DeleteTodo:output_type -> todo.v1.DeleteTodoResponse
	11, // 14: todo.v1.TodoService.ListTodos:output_type -> todo.v1.ListTodosResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_todo_v1_todo_proto_init() }
func file_proto_todo_v1_todo_proto_init() {
	if File_proto_todo_v1_todo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_todo_v1_todo_proto_rawDesc), len(file_proto_todo_v1_todo_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todo_v1_todo_proto_goTypes,
		DependencyIndexes: file_proto_todo_v1_todo_proto_depIdxs,
		EnumInfos:         file_proto_todo_v1_todo_proto_enumTypes,
		MessageInfos:      file_proto_todo_v1_todo_proto_msgTypes,
	}.Build()
	File_proto_todo_v1_todo_proto = out.File
	file_proto_todo_v1_todo_proto_goTypes = nil
	file_proto_todo_v1_todo_proto_depIdxs = nil
}
