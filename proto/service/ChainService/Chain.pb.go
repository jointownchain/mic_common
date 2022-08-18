// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: Chain.proto

package ChainService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

type GetHashContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerId int64  `protobuf:"varint,1,opt,name=managerId,proto3" json:"managerId"`
	Hash      string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash"`
	//目前支持的hash类型  COM CONFIG CONTRACT CONTRACT_MI FACTORING MEDICINE MI PAY SHIPMENT FACTORINGAGREEREF
	BusinessType string `protobuf:"bytes,3,opt,name=businessType,proto3" json:"businessType"`
}

func (x *GetHashContentReq) Reset() {
	*x = GetHashContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHashContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHashContentReq) ProtoMessage() {}

func (x *GetHashContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_Chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHashContentReq.ProtoReflect.Descriptor instead.
func (*GetHashContentReq) Descriptor() ([]byte, []int) {
	return file_Chain_proto_rawDescGZIP(), []int{0}
}

func (x *GetHashContentReq) GetManagerId() int64 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

func (x *GetHashContentReq) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *GetHashContentReq) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

type GetHashContentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//证书内容结构，json
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	//向上追溯用的hash
	LastHash         string `protobuf:"bytes,2,opt,name=lastHash,proto3" json:"lastHash"`
	UploadNode       string `protobuf:"bytes,3,opt,name=uploadNode,proto3" json:"uploadNode"`
	UploaderName     string `protobuf:"bytes,4,opt,name=uploaderName,proto3" json:"uploaderName"`
	UploaderRole     string `protobuf:"bytes,5,opt,name=uploaderRole,proto3" json:"uploaderRole"`
	UploaderRoleName string `protobuf:"bytes,6,opt,name=uploaderRoleName,proto3" json:"uploaderRoleName"`
	UploaderAccount  string `protobuf:"bytes,7,opt,name=uploaderAccount,proto3" json:"uploaderAccount"`
	CurrentHash      string `protobuf:"bytes,8,opt,name=currentHash,proto3" json:"currentHash"`
	// 上链成功时间
	UploadSuccessAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=uploadSuccessAt,proto3" json:"uploadSuccessAt"`
	// 证书描述
	Description string `protobuf:"bytes,10,opt,name=description,proto3" json:"description"`
}

func (x *GetHashContentResp) Reset() {
	*x = GetHashContentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHashContentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHashContentResp) ProtoMessage() {}

func (x *GetHashContentResp) ProtoReflect() protoreflect.Message {
	mi := &file_Chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHashContentResp.ProtoReflect.Descriptor instead.
func (*GetHashContentResp) Descriptor() ([]byte, []int) {
	return file_Chain_proto_rawDescGZIP(), []int{1}
}

func (x *GetHashContentResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *GetHashContentResp) GetLastHash() string {
	if x != nil {
		return x.LastHash
	}
	return ""
}

func (x *GetHashContentResp) GetUploadNode() string {
	if x != nil {
		return x.UploadNode
	}
	return ""
}

func (x *GetHashContentResp) GetUploaderName() string {
	if x != nil {
		return x.UploaderName
	}
	return ""
}

func (x *GetHashContentResp) GetUploaderRole() string {
	if x != nil {
		return x.UploaderRole
	}
	return ""
}

func (x *GetHashContentResp) GetUploaderRoleName() string {
	if x != nil {
		return x.UploaderRoleName
	}
	return ""
}

func (x *GetHashContentResp) GetUploaderAccount() string {
	if x != nil {
		return x.UploaderAccount
	}
	return ""
}

func (x *GetHashContentResp) GetCurrentHash() string {
	if x != nil {
		return x.CurrentHash
	}
	return ""
}

func (x *GetHashContentResp) GetUploadSuccessAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UploadSuccessAt
	}
	return nil
}

func (x *GetHashContentResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetDecryptContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerId     int64  `protobuf:"varint,1,opt,name=managerId,proto3" json:"managerId"`
	EcryptContent string `protobuf:"bytes,2,opt,name=ecryptContent,proto3" json:"ecryptContent"`
}

func (x *GetDecryptContentReq) Reset() {
	*x = GetDecryptContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDecryptContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDecryptContentReq) ProtoMessage() {}

func (x *GetDecryptContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_Chain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDecryptContentReq.ProtoReflect.Descriptor instead.
func (*GetDecryptContentReq) Descriptor() ([]byte, []int) {
	return file_Chain_proto_rawDescGZIP(), []int{2}
}

func (x *GetDecryptContentReq) GetManagerId() int64 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

func (x *GetDecryptContentReq) GetEcryptContent() string {
	if x != nil {
		return x.EcryptContent
	}
	return ""
}

type GetDecryptContentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
}

func (x *GetDecryptContentResp) Reset() {
	*x = GetDecryptContentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDecryptContentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDecryptContentResp) ProtoMessage() {}

func (x *GetDecryptContentResp) ProtoReflect() protoreflect.Message {
	mi := &file_Chain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDecryptContentResp.ProtoReflect.Descriptor instead.
func (*GetDecryptContentResp) Descriptor() ([]byte, []int) {
	return file_Chain_proto_rawDescGZIP(), []int{3}
}

func (x *GetDecryptContentResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_Chain_proto protoreflect.FileDescriptor

var file_Chain_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8c, 0x03, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x44,
	0x0a, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0xc1, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x23,
	0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x23, 0x5a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Chain_proto_rawDescOnce sync.Once
	file_Chain_proto_rawDescData = file_Chain_proto_rawDesc
)

func file_Chain_proto_rawDescGZIP() []byte {
	file_Chain_proto_rawDescOnce.Do(func() {
		file_Chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_Chain_proto_rawDescData)
	})
	return file_Chain_proto_rawDescData
}

var file_Chain_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Chain_proto_goTypes = []interface{}{
	(*GetHashContentReq)(nil),     // 0: ChainService.GetHashContentReq
	(*GetHashContentResp)(nil),    // 1: ChainService.GetHashContentResp
	(*GetDecryptContentReq)(nil),  // 2: ChainService.GetDecryptContentReq
	(*GetDecryptContentResp)(nil), // 3: ChainService.GetDecryptContentResp
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_Chain_proto_depIdxs = []int32{
	4, // 0: ChainService.GetHashContentResp.uploadSuccessAt:type_name -> google.protobuf.Timestamp
	0, // 1: ChainService.ChainService.GetHashContent:input_type -> ChainService.GetHashContentReq
	2, // 2: ChainService.ChainService.GetDecryptContent:input_type -> ChainService.GetDecryptContentReq
	1, // 3: ChainService.ChainService.GetHashContent:output_type -> ChainService.GetHashContentResp
	3, // 4: ChainService.ChainService.GetDecryptContent:output_type -> ChainService.GetDecryptContentResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Chain_proto_init() }
func file_Chain_proto_init() {
	if File_Chain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHashContentReq); i {
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
		file_Chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHashContentResp); i {
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
		file_Chain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDecryptContentReq); i {
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
		file_Chain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDecryptContentResp); i {
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
			RawDescriptor: file_Chain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Chain_proto_goTypes,
		DependencyIndexes: file_Chain_proto_depIdxs,
		MessageInfos:      file_Chain_proto_msgTypes,
	}.Build()
	File_Chain_proto = out.File
	file_Chain_proto_rawDesc = nil
	file_Chain_proto_goTypes = nil
	file_Chain_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChainServiceClient is the client API for ChainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChainServiceClient interface {
	//根据哈希查询证书内容
	GetHashContent(ctx context.Context, in *GetHashContentReq, opts ...grpc.CallOption) (*GetHashContentResp, error)
	GetDecryptContent(ctx context.Context, in *GetDecryptContentReq, opts ...grpc.CallOption) (*GetDecryptContentResp, error)
}

type chainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChainServiceClient(cc grpc.ClientConnInterface) ChainServiceClient {
	return &chainServiceClient{cc}
}

func (c *chainServiceClient) GetHashContent(ctx context.Context, in *GetHashContentReq, opts ...grpc.CallOption) (*GetHashContentResp, error) {
	out := new(GetHashContentResp)
	err := c.cc.Invoke(ctx, "/ChainService.ChainService/GetHashContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) GetDecryptContent(ctx context.Context, in *GetDecryptContentReq, opts ...grpc.CallOption) (*GetDecryptContentResp, error) {
	out := new(GetDecryptContentResp)
	err := c.cc.Invoke(ctx, "/ChainService.ChainService/GetDecryptContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainServiceServer is the server API for ChainService service.
type ChainServiceServer interface {
	//根据哈希查询证书内容
	GetHashContent(context.Context, *GetHashContentReq) (*GetHashContentResp, error)
	GetDecryptContent(context.Context, *GetDecryptContentReq) (*GetDecryptContentResp, error)
}

// UnimplementedChainServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChainServiceServer struct {
}

func (*UnimplementedChainServiceServer) GetHashContent(context.Context, *GetHashContentReq) (*GetHashContentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHashContent not implemented")
}
func (*UnimplementedChainServiceServer) GetDecryptContent(context.Context, *GetDecryptContentReq) (*GetDecryptContentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDecryptContent not implemented")
}

func RegisterChainServiceServer(s *grpc.Server, srv ChainServiceServer) {
	s.RegisterService(&_ChainService_serviceDesc, srv)
}

func _ChainService_GetHashContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHashContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).GetHashContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChainService.ChainService/GetHashContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).GetHashContent(ctx, req.(*GetHashContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_GetDecryptContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDecryptContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).GetDecryptContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChainService.ChainService/GetDecryptContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).GetDecryptContent(ctx, req.(*GetDecryptContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ChainService.ChainService",
	HandlerType: (*ChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHashContent",
			Handler:    _ChainService_GetHashContent_Handler,
		},
		{
			MethodName: "GetDecryptContent",
			Handler:    _ChainService_GetDecryptContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Chain.proto",
}
