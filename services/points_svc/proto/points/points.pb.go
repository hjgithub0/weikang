// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: points.proto

package points

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

type CreatePointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Points int64 `protobuf:"varint,2,opt,name=Points,proto3" json:"Points,omitempty"`
}

func (x *CreatePointsRequest) Reset() {
	*x = CreatePointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePointsRequest) ProtoMessage() {}

func (x *CreatePointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePointsRequest.ProtoReflect.Descriptor instead.
func (*CreatePointsRequest) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePointsRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreatePointsRequest) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

type CreatePointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *CreatePointsResponse) Reset() {
	*x = CreatePointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePointsResponse) ProtoMessage() {}

func (x *CreatePointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePointsResponse.ProtoReflect.Descriptor instead.
func (*CreatePointsResponse) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePointsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PointsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Points int64 `protobuf:"varint,2,opt,name=Points,proto3" json:"Points,omitempty"`
}

func (x *PointsInfo) Reset() {
	*x = PointsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointsInfo) ProtoMessage() {}

func (x *PointsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointsInfo.ProtoReflect.Descriptor instead.
func (*PointsInfo) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{2}
}

func (x *PointsInfo) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *PointsInfo) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

type GetUserAllPointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserAllPointsRequest) Reset() {
	*x = GetUserAllPointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAllPointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAllPointsRequest) ProtoMessage() {}

func (x *GetUserAllPointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAllPointsRequest.ProtoReflect.Descriptor instead.
func (*GetUserAllPointsRequest) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{3}
}

type GetUserAllPointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PointsList []*PointsInfo `protobuf:"bytes,1,rep,name=PointsList,proto3" json:"PointsList,omitempty"`
}

func (x *GetUserAllPointsResponse) Reset() {
	*x = GetUserAllPointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAllPointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAllPointsResponse) ProtoMessage() {}

func (x *GetUserAllPointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAllPointsResponse.ProtoReflect.Descriptor instead.
func (*GetUserAllPointsResponse) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserAllPointsResponse) GetPointsList() []*PointsInfo {
	if x != nil {
		return x.PointsList
	}
	return nil
}

type UpdatePointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Points int64 `protobuf:"varint,2,opt,name=Points,proto3" json:"Points,omitempty"`
}

func (x *UpdatePointsRequest) Reset() {
	*x = UpdatePointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePointsRequest) ProtoMessage() {}

func (x *UpdatePointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePointsRequest.ProtoReflect.Descriptor instead.
func (*UpdatePointsRequest) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePointsRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UpdatePointsRequest) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

type UpdatePointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *UpdatePointsResponse) Reset() {
	*x = UpdatePointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePointsResponse) ProtoMessage() {}

func (x *UpdatePointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePointsResponse.ProtoReflect.Descriptor instead.
func (*UpdatePointsResponse) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePointsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeletePointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *DeletePointsRequest) Reset() {
	*x = DeletePointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePointsRequest) ProtoMessage() {}

func (x *DeletePointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePointsRequest.ProtoReflect.Descriptor instead.
func (*DeletePointsRequest) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePointsRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type DeletePointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *DeletePointsResponse) Reset() {
	*x = DeletePointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePointsResponse) ProtoMessage() {}

func (x *DeletePointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePointsResponse.ProtoReflect.Descriptor instead.
func (*DeletePointsResponse) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePointsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UpdatePointsCompensationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Points int64 `protobuf:"varint,2,opt,name=Points,proto3" json:"Points,omitempty"`
}

func (x *UpdatePointsCompensationRequest) Reset() {
	*x = UpdatePointsCompensationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePointsCompensationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePointsCompensationRequest) ProtoMessage() {}

func (x *UpdatePointsCompensationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePointsCompensationRequest.ProtoReflect.Descriptor instead.
func (*UpdatePointsCompensationRequest) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{9}
}

func (x *UpdatePointsCompensationRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UpdatePointsCompensationRequest) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

type UpdatePointsCompensationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *UpdatePointsCompensationResponse) Reset() {
	*x = UpdatePointsCompensationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_points_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePointsCompensationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePointsCompensationResponse) ProtoMessage() {}

func (x *UpdatePointsCompensationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePointsCompensationResponse.ProtoReflect.Descriptor instead.
func (*UpdatePointsCompensationResponse) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{10}
}

func (x *UpdatePointsCompensationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_points_proto protoreflect.FileDescriptor

var file_points_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x30, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x3c, 0x0a, 0x0a, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x19, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22,
	0x30, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x2d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x51, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x3c, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0xaf, 0x03, 0x0a, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x49,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1b,
	0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x2e,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c,
	0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x6c, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x49, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x65,
	0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_points_proto_rawDescOnce sync.Once
	file_points_proto_rawDescData = file_points_proto_rawDesc
)

func file_points_proto_rawDescGZIP() []byte {
	file_points_proto_rawDescOnce.Do(func() {
		file_points_proto_rawDescData = protoimpl.X.CompressGZIP(file_points_proto_rawDescData)
	})
	return file_points_proto_rawDescData
}

var file_points_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_points_proto_goTypes = []interface{}{
	(*CreatePointsRequest)(nil),              // 0: points.CreatePointsRequest
	(*CreatePointsResponse)(nil),             // 1: points.CreatePointsResponse
	(*PointsInfo)(nil),                       // 2: points.PointsInfo
	(*GetUserAllPointsRequest)(nil),          // 3: points.GetUserAllPointsRequest
	(*GetUserAllPointsResponse)(nil),         // 4: points.GetUserAllPointsResponse
	(*UpdatePointsRequest)(nil),              // 5: points.UpdatePointsRequest
	(*UpdatePointsResponse)(nil),             // 6: points.UpdatePointsResponse
	(*DeletePointsRequest)(nil),              // 7: points.DeletePointsRequest
	(*DeletePointsResponse)(nil),             // 8: points.DeletePointsResponse
	(*UpdatePointsCompensationRequest)(nil),  // 9: points.UpdatePointsCompensationRequest
	(*UpdatePointsCompensationResponse)(nil), // 10: points.UpdatePointsCompensationResponse
}
var file_points_proto_depIdxs = []int32{
	2,  // 0: points.GetUserAllPointsResponse.PointsList:type_name -> points.PointsInfo
	0,  // 1: points.Points.CreatePoints:input_type -> points.CreatePointsRequest
	3,  // 2: points.Points.GetUserAllPoints:input_type -> points.GetUserAllPointsRequest
	5,  // 3: points.Points.UpdatePoints:input_type -> points.UpdatePointsRequest
	9,  // 4: points.Points.UpdatePointsCompensation:input_type -> points.UpdatePointsCompensationRequest
	7,  // 5: points.Points.DeletePoints:input_type -> points.DeletePointsRequest
	1,  // 6: points.Points.CreatePoints:output_type -> points.CreatePointsResponse
	4,  // 7: points.Points.GetUserAllPoints:output_type -> points.GetUserAllPointsResponse
	6,  // 8: points.Points.UpdatePoints:output_type -> points.UpdatePointsResponse
	10, // 9: points.Points.UpdatePointsCompensation:output_type -> points.UpdatePointsCompensationResponse
	8,  // 10: points.Points.DeletePoints:output_type -> points.DeletePointsResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_points_proto_init() }
func file_points_proto_init() {
	if File_points_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_points_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePointsRequest); i {
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
		file_points_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePointsResponse); i {
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
		file_points_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointsInfo); i {
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
		file_points_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAllPointsRequest); i {
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
		file_points_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAllPointsResponse); i {
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
		file_points_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePointsRequest); i {
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
		file_points_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePointsResponse); i {
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
		file_points_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePointsRequest); i {
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
		file_points_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePointsResponse); i {
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
		file_points_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePointsCompensationRequest); i {
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
		file_points_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePointsCompensationResponse); i {
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
			RawDescriptor: file_points_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_points_proto_goTypes,
		DependencyIndexes: file_points_proto_depIdxs,
		MessageInfos:      file_points_proto_msgTypes,
	}.Build()
	File_points_proto = out.File
	file_points_proto_rawDesc = nil
	file_points_proto_goTypes = nil
	file_points_proto_depIdxs = nil
}
