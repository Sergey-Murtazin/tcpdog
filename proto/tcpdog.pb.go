// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: tcpdog.proto

package tcpdog

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FieldsPBS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields *_struct.Struct `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *FieldsPBS) Reset() {
	*x = FieldsPBS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcpdog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldsPBS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldsPBS) ProtoMessage() {}

func (x *FieldsPBS) ProtoReflect() protoreflect.Message {
	mi := &file_tcpdog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldsPBS.ProtoReflect.Descriptor instead.
func (*FieldsPBS) Descriptor() ([]byte, []int) {
	return file_tcpdog_proto_rawDescGZIP(), []int{0}
}

func (x *FieldsPBS) GetFields() *_struct.Struct {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Fields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TCPHeaderLen  int32  `protobuf:"varint,1,opt,name=TCPHeaderLen,proto3" json:"TCPHeaderLen,omitempty"`
	SRTT          int32  `protobuf:"varint,2,opt,name=SRTT,proto3" json:"SRTT,omitempty"`
	MDev          int32  `protobuf:"varint,3,opt,name=MDev,proto3" json:"MDev,omitempty"`
	TotalRetrans  int32  `protobuf:"varint,4,opt,name=TotalRetrans,proto3" json:"TotalRetrans,omitempty"`
	AdvMSS        int32  `protobuf:"varint,5,opt,name=AdvMSS,proto3" json:"AdvMSS,omitempty"`
	SAddr         string `protobuf:"bytes,6,opt,name=SAddr,proto3" json:"SAddr,omitempty"`
	DAddr         string `protobuf:"bytes,7,opt,name=DAddr,proto3" json:"DAddr,omitempty"`
	DPort         int32  `protobuf:"varint,8,opt,name=DPort,proto3" json:"DPort,omitempty"`
	LPort         int32  `protobuf:"varint,9,opt,name=LPort,proto3" json:"LPort,omitempty"`
	BytesReceived int64  `protobuf:"varint,10,opt,name=BytesReceived,proto3" json:"BytesReceived,omitempty"`
	BytesSent     int64  `protobuf:"varint,11,opt,name=BytesSent,proto3" json:"BytesSent,omitempty"`
	BytesAcked    int64  `protobuf:"varint,12,opt,name=BytesAcked,proto3" json:"BytesAcked,omitempty"`
	NumSAcks      int32  `protobuf:"varint,13,opt,name=NumSAcks,proto3" json:"NumSAcks,omitempty"`
	UserMSS       int32  `protobuf:"varint,14,opt,name=UserMSS,proto3" json:"UserMSS,omitempty"`
	RTT           int32  `protobuf:"varint,15,opt,name=RTT,proto3" json:"RTT,omitempty"`
	MSSClamp      int32  `protobuf:"varint,16,opt,name=MSSClamp,proto3" json:"MSSClamp,omitempty"`
	Task          string `protobuf:"bytes,17,opt,name=Task,proto3" json:"Task,omitempty"`
	PID           int32  `protobuf:"varint,18,opt,name=PID,proto3" json:"PID,omitempty"`
	SegsIn        int32  `protobuf:"varint,19,opt,name=SegsIn,proto3" json:"SegsIn,omitempty"`
	SegsOut       int32  `protobuf:"varint,20,opt,name=SegsOut,proto3" json:"SegsOut,omitempty"`
	DsackDups     int32  `protobuf:"varint,21,opt,name=DsackDups,proto3" json:"DsackDups,omitempty"`
	RateDelivered int32  `protobuf:"varint,22,opt,name=RateDelivered,proto3" json:"RateDelivered,omitempty"`
	RateInterval  int32  `protobuf:"varint,23,opt,name=RateInterval,proto3" json:"RateInterval,omitempty"`
}

func (x *Fields) Reset() {
	*x = Fields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcpdog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fields) ProtoMessage() {}

func (x *Fields) ProtoReflect() protoreflect.Message {
	mi := &file_tcpdog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fields.ProtoReflect.Descriptor instead.
func (*Fields) Descriptor() ([]byte, []int) {
	return file_tcpdog_proto_rawDescGZIP(), []int{1}
}

func (x *Fields) GetTCPHeaderLen() int32 {
	if x != nil {
		return x.TCPHeaderLen
	}
	return 0
}

func (x *Fields) GetSRTT() int32 {
	if x != nil {
		return x.SRTT
	}
	return 0
}

func (x *Fields) GetMDev() int32 {
	if x != nil {
		return x.MDev
	}
	return 0
}

func (x *Fields) GetTotalRetrans() int32 {
	if x != nil {
		return x.TotalRetrans
	}
	return 0
}

func (x *Fields) GetAdvMSS() int32 {
	if x != nil {
		return x.AdvMSS
	}
	return 0
}

func (x *Fields) GetSAddr() string {
	if x != nil {
		return x.SAddr
	}
	return ""
}

func (x *Fields) GetDAddr() string {
	if x != nil {
		return x.DAddr
	}
	return ""
}

func (x *Fields) GetDPort() int32 {
	if x != nil {
		return x.DPort
	}
	return 0
}

func (x *Fields) GetLPort() int32 {
	if x != nil {
		return x.LPort
	}
	return 0
}

func (x *Fields) GetBytesReceived() int64 {
	if x != nil {
		return x.BytesReceived
	}
	return 0
}

func (x *Fields) GetBytesSent() int64 {
	if x != nil {
		return x.BytesSent
	}
	return 0
}

func (x *Fields) GetBytesAcked() int64 {
	if x != nil {
		return x.BytesAcked
	}
	return 0
}

func (x *Fields) GetNumSAcks() int32 {
	if x != nil {
		return x.NumSAcks
	}
	return 0
}

func (x *Fields) GetUserMSS() int32 {
	if x != nil {
		return x.UserMSS
	}
	return 0
}

func (x *Fields) GetRTT() int32 {
	if x != nil {
		return x.RTT
	}
	return 0
}

func (x *Fields) GetMSSClamp() int32 {
	if x != nil {
		return x.MSSClamp
	}
	return 0
}

func (x *Fields) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *Fields) GetPID() int32 {
	if x != nil {
		return x.PID
	}
	return 0
}

func (x *Fields) GetSegsIn() int32 {
	if x != nil {
		return x.SegsIn
	}
	return 0
}

func (x *Fields) GetSegsOut() int32 {
	if x != nil {
		return x.SegsOut
	}
	return 0
}

func (x *Fields) GetDsackDups() int32 {
	if x != nil {
		return x.DsackDups
	}
	return 0
}

func (x *Fields) GetRateDelivered() int32 {
	if x != nil {
		return x.RateDelivered
	}
	return 0
}

func (x *Fields) GetRateInterval() int32 {
	if x != nil {
		return x.RateInterval
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcpdog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_tcpdog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_tcpdog_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_tcpdog_proto protoreflect.FileDescriptor

var file_tcpdog_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x63, 0x70, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x74, 0x63, 0x70, 0x64, 0x6f, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x50, 0x42,
	0x53, 0x12, 0x2f, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x22, 0xf0, 0x04, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x54, 0x43, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x54, 0x43, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x65,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x52, 0x54, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x53, 0x52, 0x54, 0x54, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x44, 0x65, 0x76, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x4d, 0x44, 0x65, 0x76, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x52, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x64, 0x76, 0x4d, 0x53, 0x53, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41,
	0x64, 0x76, 0x4d, 0x53, 0x53, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x41, 0x64, 0x64, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x44,
	0x41, 0x64, 0x64, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x44, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x44, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x50, 0x6f, 0x72, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x79, 0x74, 0x65, 0x73, 0x41, 0x63, 0x6b, 0x65, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x42, 0x79, 0x74, 0x65, 0x73, 0x41, 0x63, 0x6b, 0x65,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x75, 0x6d, 0x53, 0x41, 0x63, 0x6b, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x4e, 0x75, 0x6d, 0x53, 0x41, 0x63, 0x6b, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x53, 0x53, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x53, 0x53, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x54, 0x54, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x52, 0x54, 0x54, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x53, 0x53,
	0x43, 0x6c, 0x61, 0x6d, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4d, 0x53, 0x53,
	0x43, 0x6c, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x49, 0x44,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x65, 0x67, 0x73, 0x49, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x65, 0x67,
	0x73, 0x49, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x67, 0x73, 0x4f, 0x75, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x65, 0x67, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x44, 0x73, 0x61, 0x63, 0x6b, 0x44, 0x75, 0x70, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x44, 0x73, 0x61, 0x63, 0x6b, 0x44, 0x75, 0x70, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x52,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x52, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x76, 0x0a, 0x06, 0x54, 0x43, 0x50, 0x44, 0x6f, 0x67, 0x12,
	0x32, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x63, 0x65, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x2e,
	0x74, 0x63, 0x70, 0x64, 0x6f, 0x67, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x10, 0x2e,
	0x74, 0x63, 0x70, 0x64, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x38, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x65, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x50, 0x42, 0x53, 0x12, 0x11, 0x2e, 0x74, 0x63, 0x70, 0x64, 0x6f, 0x67, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x50, 0x42, 0x53, 0x1a, 0x10, 0x2e, 0x74, 0x63, 0x70, 0x64, 0x6f, 0x67,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tcpdog_proto_rawDescOnce sync.Once
	file_tcpdog_proto_rawDescData = file_tcpdog_proto_rawDesc
)

func file_tcpdog_proto_rawDescGZIP() []byte {
	file_tcpdog_proto_rawDescOnce.Do(func() {
		file_tcpdog_proto_rawDescData = protoimpl.X.CompressGZIP(file_tcpdog_proto_rawDescData)
	})
	return file_tcpdog_proto_rawDescData
}

var file_tcpdog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tcpdog_proto_goTypes = []interface{}{
	(*FieldsPBS)(nil),      // 0: tcpdog.FieldsPBS
	(*Fields)(nil),         // 1: tcpdog.Fields
	(*Response)(nil),       // 2: tcpdog.Response
	(*_struct.Struct)(nil), // 3: google.protobuf.Struct
}
var file_tcpdog_proto_depIdxs = []int32{
	3, // 0: tcpdog.FieldsPBS.fields:type_name -> google.protobuf.Struct
	1, // 1: tcpdog.TCPDog.Tracepoint:input_type -> tcpdog.Fields
	0, // 2: tcpdog.TCPDog.TracepointPBS:input_type -> tcpdog.FieldsPBS
	2, // 3: tcpdog.TCPDog.Tracepoint:output_type -> tcpdog.Response
	2, // 4: tcpdog.TCPDog.TracepointPBS:output_type -> tcpdog.Response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tcpdog_proto_init() }
func file_tcpdog_proto_init() {
	if File_tcpdog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tcpdog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldsPBS); i {
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
		file_tcpdog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fields); i {
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
		file_tcpdog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_tcpdog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tcpdog_proto_goTypes,
		DependencyIndexes: file_tcpdog_proto_depIdxs,
		MessageInfos:      file_tcpdog_proto_msgTypes,
	}.Build()
	File_tcpdog_proto = out.File
	file_tcpdog_proto_rawDesc = nil
	file_tcpdog_proto_goTypes = nil
	file_tcpdog_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TCPDogClient is the client API for TCPDog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TCPDogClient interface {
	Tracepoint(ctx context.Context, opts ...grpc.CallOption) (TCPDog_TracepointClient, error)
	TracepointPBS(ctx context.Context, opts ...grpc.CallOption) (TCPDog_TracepointPBSClient, error)
}

type tCPDogClient struct {
	cc grpc.ClientConnInterface
}

func NewTCPDogClient(cc grpc.ClientConnInterface) TCPDogClient {
	return &tCPDogClient{cc}
}

func (c *tCPDogClient) Tracepoint(ctx context.Context, opts ...grpc.CallOption) (TCPDog_TracepointClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TCPDog_serviceDesc.Streams[0], "/tcpdog.TCPDog/Tracepoint", opts...)
	if err != nil {
		return nil, err
	}
	x := &tCPDogTracepointClient{stream}
	return x, nil
}

type TCPDog_TracepointClient interface {
	Send(*Fields) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type tCPDogTracepointClient struct {
	grpc.ClientStream
}

func (x *tCPDogTracepointClient) Send(m *Fields) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tCPDogTracepointClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tCPDogClient) TracepointPBS(ctx context.Context, opts ...grpc.CallOption) (TCPDog_TracepointPBSClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TCPDog_serviceDesc.Streams[1], "/tcpdog.TCPDog/TracepointPBS", opts...)
	if err != nil {
		return nil, err
	}
	x := &tCPDogTracepointPBSClient{stream}
	return x, nil
}

type TCPDog_TracepointPBSClient interface {
	Send(*FieldsPBS) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type tCPDogTracepointPBSClient struct {
	grpc.ClientStream
}

func (x *tCPDogTracepointPBSClient) Send(m *FieldsPBS) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tCPDogTracepointPBSClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TCPDogServer is the server API for TCPDog service.
type TCPDogServer interface {
	Tracepoint(TCPDog_TracepointServer) error
	TracepointPBS(TCPDog_TracepointPBSServer) error
}

// UnimplementedTCPDogServer can be embedded to have forward compatible implementations.
type UnimplementedTCPDogServer struct {
}

func (*UnimplementedTCPDogServer) Tracepoint(TCPDog_TracepointServer) error {
	return status.Errorf(codes.Unimplemented, "method Tracepoint not implemented")
}
func (*UnimplementedTCPDogServer) TracepointPBS(TCPDog_TracepointPBSServer) error {
	return status.Errorf(codes.Unimplemented, "method TracepointPBS not implemented")
}

func RegisterTCPDogServer(s *grpc.Server, srv TCPDogServer) {
	s.RegisterService(&_TCPDog_serviceDesc, srv)
}

func _TCPDog_Tracepoint_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TCPDogServer).Tracepoint(&tCPDogTracepointServer{stream})
}

type TCPDog_TracepointServer interface {
	SendAndClose(*Response) error
	Recv() (*Fields, error)
	grpc.ServerStream
}

type tCPDogTracepointServer struct {
	grpc.ServerStream
}

func (x *tCPDogTracepointServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tCPDogTracepointServer) Recv() (*Fields, error) {
	m := new(Fields)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TCPDog_TracepointPBS_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TCPDogServer).TracepointPBS(&tCPDogTracepointPBSServer{stream})
}

type TCPDog_TracepointPBSServer interface {
	SendAndClose(*Response) error
	Recv() (*FieldsPBS, error)
	grpc.ServerStream
}

type tCPDogTracepointPBSServer struct {
	grpc.ServerStream
}

func (x *tCPDogTracepointPBSServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tCPDogTracepointPBSServer) Recv() (*FieldsPBS, error) {
	m := new(FieldsPBS)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TCPDog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tcpdog.TCPDog",
	HandlerType: (*TCPDogServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tracepoint",
			Handler:       _TCPDog_Tracepoint_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TracepointPBS",
			Handler:       _TCPDog_TracepointPBS_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "tcpdog.proto",
}
