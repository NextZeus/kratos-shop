// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: api/cart/service/v1/cart.proto

package v1

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   int64 `protobuf:"varint,1,opt,name=itemId,proto3" json:"itemId,omitempty"`
	Quantity int64 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Item) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{1}
}

func (x *GetCartRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetCartReply) Reset() {
	*x = GetCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReply) ProtoMessage() {}

func (x *GetCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReply.ProtoReflect.Descriptor instead.
func (*GetCartReply) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{2}
}

func (x *GetCartReply) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DeleteCartRequest) Reset() {
	*x = DeleteCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartRequest) ProtoMessage() {}

func (x *DeleteCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartRequest.ProtoReflect.Descriptor instead.
func (*DeleteCartRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCartRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type DeleteCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCartReply) Reset() {
	*x = DeleteCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartReply) ProtoMessage() {}

func (x *DeleteCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartReply.ProtoReflect.Descriptor instead.
func (*DeleteCartReply) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{4}
}

type AddItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ItemId   int64 `protobuf:"varint,2,opt,name=itemId,proto3" json:"itemId,omitempty"`
	Quantity int64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *AddItemRequest) Reset() {
	*x = AddItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemRequest) ProtoMessage() {}

func (x *AddItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemRequest.ProtoReflect.Descriptor instead.
func (*AddItemRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{5}
}

func (x *AddItemRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AddItemRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *AddItemRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AddItemReply) Reset() {
	*x = AddItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemReply) ProtoMessage() {}

func (x *AddItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemReply.ProtoReflect.Descriptor instead.
func (*AddItemReply) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{6}
}

func (x *AddItemReply) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ItemId   int64 `protobuf:"varint,2,opt,name=itemId,proto3" json:"itemId,omitempty"`
	Quantity int64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UpdateItemRequest) Reset() {
	*x = UpdateItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemRequest) ProtoMessage() {}

func (x *UpdateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateItemRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateItemRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UpdateItemRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *UpdateItemRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UpdateItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *UpdateItemReply) Reset() {
	*x = UpdateItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemReply) ProtoMessage() {}

func (x *UpdateItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemReply.ProtoReflect.Descriptor instead.
func (*UpdateItemReply) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateItemReply) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ItemId int64 `protobuf:"varint,2,opt,name=itemId,proto3" json:"itemId,omitempty"`
}

func (x *DeleteItemRequest) Reset() {
	*x = DeleteItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemRequest) ProtoMessage() {}

func (x *DeleteItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteItemRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DeleteItemRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type DeleteItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *DeleteItemReply) Reset() {
	*x = DeleteItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_service_v1_cart_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemReply) ProtoMessage() {}

func (x *DeleteItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_service_v1_cart_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemReply.ProtoReflect.Descriptor instead.
func (*DeleteItemReply) Descriptor() ([]byte, []int) {
	return file_api_cart_service_v1_cart_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteItemReply) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_api_cart_service_v1_cart_proto protoreflect.FileDescriptor

var file_api_cart_service_v1_cart_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x3a, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0x22, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x25, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x11, 0x0a,
	0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x56, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3f, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x59, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x42, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xc0, 0x03, 0x0a, 0x04,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x51, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x51, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x23,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x5a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x4f,
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x78, 0x74, 0x7a, 0x65, 0x75, 0x73, 0x2f, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72,
	0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cart_service_v1_cart_proto_rawDescOnce sync.Once
	file_api_cart_service_v1_cart_proto_rawDescData = file_api_cart_service_v1_cart_proto_rawDesc
)

func file_api_cart_service_v1_cart_proto_rawDescGZIP() []byte {
	file_api_cart_service_v1_cart_proto_rawDescOnce.Do(func() {
		file_api_cart_service_v1_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cart_service_v1_cart_proto_rawDescData)
	})
	return file_api_cart_service_v1_cart_proto_rawDescData
}

var file_api_cart_service_v1_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_cart_service_v1_cart_proto_goTypes = []interface{}{
	(*Item)(nil),              // 0: api.cart.service.v1.Item
	(*GetCartRequest)(nil),    // 1: api.cart.service.v1.GetCartRequest
	(*GetCartReply)(nil),      // 2: api.cart.service.v1.GetCartReply
	(*DeleteCartRequest)(nil), // 3: api.cart.service.v1.DeleteCartRequest
	(*DeleteCartReply)(nil),   // 4: api.cart.service.v1.DeleteCartReply
	(*AddItemRequest)(nil),    // 5: api.cart.service.v1.AddItemRequest
	(*AddItemReply)(nil),      // 6: api.cart.service.v1.AddItemReply
	(*UpdateItemRequest)(nil), // 7: api.cart.service.v1.UpdateItemRequest
	(*UpdateItemReply)(nil),   // 8: api.cart.service.v1.UpdateItemReply
	(*DeleteItemRequest)(nil), // 9: api.cart.service.v1.DeleteItemRequest
	(*DeleteItemReply)(nil),   // 10: api.cart.service.v1.DeleteItemReply
}
var file_api_cart_service_v1_cart_proto_depIdxs = []int32{
	0,  // 0: api.cart.service.v1.GetCartReply.items:type_name -> api.cart.service.v1.Item
	0,  // 1: api.cart.service.v1.AddItemReply.items:type_name -> api.cart.service.v1.Item
	0,  // 2: api.cart.service.v1.UpdateItemReply.items:type_name -> api.cart.service.v1.Item
	0,  // 3: api.cart.service.v1.DeleteItemReply.items:type_name -> api.cart.service.v1.Item
	1,  // 4: api.cart.service.v1.Cart.GetCart:input_type -> api.cart.service.v1.GetCartRequest
	3,  // 5: api.cart.service.v1.Cart.DeleteCart:input_type -> api.cart.service.v1.DeleteCartRequest
	5,  // 6: api.cart.service.v1.Cart.AddItem:input_type -> api.cart.service.v1.AddItemRequest
	7,  // 7: api.cart.service.v1.Cart.UpdateItem:input_type -> api.cart.service.v1.UpdateItemRequest
	9,  // 8: api.cart.service.v1.Cart.DeleteItem:input_type -> api.cart.service.v1.DeleteItemRequest
	2,  // 9: api.cart.service.v1.Cart.GetCart:output_type -> api.cart.service.v1.GetCartReply
	4,  // 10: api.cart.service.v1.Cart.DeleteCart:output_type -> api.cart.service.v1.DeleteCartReply
	6,  // 11: api.cart.service.v1.Cart.AddItem:output_type -> api.cart.service.v1.AddItemReply
	8,  // 12: api.cart.service.v1.Cart.UpdateItem:output_type -> api.cart.service.v1.UpdateItemReply
	10, // 13: api.cart.service.v1.Cart.DeleteItem:output_type -> api.cart.service.v1.DeleteItemReply
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_cart_service_v1_cart_proto_init() }
func file_api_cart_service_v1_cart_proto_init() {
	if File_api_cart_service_v1_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_cart_service_v1_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_api_cart_service_v1_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartRequest); i {
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
		file_api_cart_service_v1_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartReply); i {
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
		file_api_cart_service_v1_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartRequest); i {
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
		file_api_cart_service_v1_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartReply); i {
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
		file_api_cart_service_v1_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemRequest); i {
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
		file_api_cart_service_v1_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemReply); i {
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
		file_api_cart_service_v1_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemRequest); i {
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
		file_api_cart_service_v1_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemReply); i {
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
		file_api_cart_service_v1_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemRequest); i {
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
		file_api_cart_service_v1_cart_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemReply); i {
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
			RawDescriptor: file_api_cart_service_v1_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_cart_service_v1_cart_proto_goTypes,
		DependencyIndexes: file_api_cart_service_v1_cart_proto_depIdxs,
		MessageInfos:      file_api_cart_service_v1_cart_proto_msgTypes,
	}.Build()
	File_api_cart_service_v1_cart_proto = out.File
	file_api_cart_service_v1_cart_proto_rawDesc = nil
	file_api_cart_service_v1_cart_proto_goTypes = nil
	file_api_cart_service_v1_cart_proto_depIdxs = nil
}
