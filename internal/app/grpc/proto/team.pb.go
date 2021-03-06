// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/app/grpc/proto/team.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Team struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_6441cd82d37bed08, []int{0}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Team)(nil), "proto.Team")
}

func init() { proto.RegisterFile("internal/app/grpc/proto/team.proto", fileDescriptor_6441cd82d37bed08) }

var fileDescriptor_6441cd82d37bed08 = []byte{
	// 106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0x2c, 0x28, 0xd0, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x28,
	0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x49, 0x4d, 0xcc, 0xd5, 0x03, 0x33, 0x85, 0x58, 0xc1, 0x94, 0x92,
	0x16, 0x17, 0x4b, 0x48, 0x6a, 0x62, 0xae, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x73, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x9d, 0xc4, 0x06, 0xd6, 0x62, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x2c, 0x98, 0x16, 0x0b, 0x5f, 0x00, 0x00, 0x00,
}
