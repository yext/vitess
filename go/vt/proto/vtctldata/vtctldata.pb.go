// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vtctldata.proto

package vtctldata

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	logutil "liquidata-inc/vitess/go/vt/proto/logutil"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ExecuteVtctlCommandRequest is the payload for ExecuteVtctlCommand.
// timeouts are in nanoseconds.
type ExecuteVtctlCommandRequest struct {
	Args                 []string `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	ActionTimeout        int64    `protobuf:"varint,2,opt,name=action_timeout,json=actionTimeout,proto3" json:"action_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteVtctlCommandRequest) Reset()         { *m = ExecuteVtctlCommandRequest{} }
func (m *ExecuteVtctlCommandRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteVtctlCommandRequest) ProtoMessage()    {}
func (*ExecuteVtctlCommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f41247b323a1ab2e, []int{0}
}

func (m *ExecuteVtctlCommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteVtctlCommandRequest.Unmarshal(m, b)
}
func (m *ExecuteVtctlCommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteVtctlCommandRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteVtctlCommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteVtctlCommandRequest.Merge(m, src)
}
func (m *ExecuteVtctlCommandRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteVtctlCommandRequest.Size(m)
}
func (m *ExecuteVtctlCommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteVtctlCommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteVtctlCommandRequest proto.InternalMessageInfo

func (m *ExecuteVtctlCommandRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ExecuteVtctlCommandRequest) GetActionTimeout() int64 {
	if m != nil {
		return m.ActionTimeout
	}
	return 0
}

// ExecuteVtctlCommandResponse is streamed back by ExecuteVtctlCommand.
type ExecuteVtctlCommandResponse struct {
	Event                *logutil.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExecuteVtctlCommandResponse) Reset()         { *m = ExecuteVtctlCommandResponse{} }
func (m *ExecuteVtctlCommandResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteVtctlCommandResponse) ProtoMessage()    {}
func (*ExecuteVtctlCommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f41247b323a1ab2e, []int{1}
}

func (m *ExecuteVtctlCommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteVtctlCommandResponse.Unmarshal(m, b)
}
func (m *ExecuteVtctlCommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteVtctlCommandResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteVtctlCommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteVtctlCommandResponse.Merge(m, src)
}
func (m *ExecuteVtctlCommandResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteVtctlCommandResponse.Size(m)
}
func (m *ExecuteVtctlCommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteVtctlCommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteVtctlCommandResponse proto.InternalMessageInfo

func (m *ExecuteVtctlCommandResponse) GetEvent() *logutil.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecuteVtctlCommandRequest)(nil), "vtctldata.ExecuteVtctlCommandRequest")
	proto.RegisterType((*ExecuteVtctlCommandResponse)(nil), "vtctldata.ExecuteVtctlCommandResponse")
}

func init() { proto.RegisterFile("vtctldata.proto", fileDescriptor_f41247b323a1ab2e) }

var fileDescriptor_f41247b323a1ab2e = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcf, 0xd1, 0x4a, 0x87, 0x30,
	0x14, 0x06, 0x70, 0xd6, 0xbf, 0x82, 0xff, 0x42, 0x83, 0x5d, 0x89, 0xdd, 0x88, 0x54, 0xec, 0xca,
	0x41, 0xbd, 0x41, 0xe2, 0x0b, 0x8c, 0x28, 0xe8, 0x26, 0x96, 0x1e, 0x64, 0xa0, 0x3b, 0xe6, 0xce,
	0x46, 0x8f, 0x1f, 0x3a, 0xf2, 0xaa, 0xbb, 0x8f, 0xdf, 0x37, 0xc6, 0x77, 0xf8, 0x6d, 0xa4, 0x9e,
	0xa6, 0xc1, 0x90, 0x69, 0x96, 0x15, 0x09, 0xc5, 0xf9, 0x80, 0x32, 0x9b, 0x70, 0x0c, 0x64, 0xa7,
	0xd4, 0xd4, 0xef, 0xbc, 0xec, 0x7e, 0xa0, 0x0f, 0x04, 0x6f, 0xdb, 0x93, 0x16, 0xe7, 0xd9, 0xb8,
	0x41, 0xc3, 0x77, 0x00, 0x4f, 0x42, 0xf0, 0x4b, 0xb3, 0x8e, 0xbe, 0x60, 0xd5, 0x49, 0x9e, 0xf5,
	0x9e, 0xc5, 0x03, 0xcf, 0x4d, 0x4f, 0x16, 0xdd, 0x27, 0xd9, 0x19, 0x30, 0x50, 0x71, 0x51, 0x31,
	0x79, 0xd2, 0x59, 0xd2, 0xd7, 0x84, 0x75, 0xcb, 0xef, 0xfe, 0xfd, 0xd8, 0x2f, 0xe8, 0x3c, 0x88,
	0x7b, 0x7e, 0x05, 0x11, 0x1c, 0x15, 0xac, 0x62, 0xf2, 0xe6, 0x29, 0x6f, 0xfe, 0x66, 0x75, 0x9b,
	0xea, 0x54, 0xbe, 0xc8, 0x8f, 0xc7, 0x68, 0x09, 0xbc, 0x6f, 0x2c, 0xaa, 0x94, 0xd4, 0x88, 0x2a,
	0x92, 0xda, 0xd7, 0xab, 0xe3, 0xac, 0xaf, 0xeb, 0x1d, 0x9e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xa9, 0xd5, 0x57, 0x1c, 0xfb, 0x00, 0x00, 0x00,
}
