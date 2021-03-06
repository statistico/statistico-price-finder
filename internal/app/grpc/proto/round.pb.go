// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/app/grpc/proto/round.proto

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

type Round struct {
	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SeasonId int64  `protobuf:"varint,3,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	// RFC3339 formatted string i.e "2006-01-02T15:04:05Z07:00"
	StartDate string `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// RFC3339 formatted string i.e "2006-01-02T15:04:05Z07:00"
	EndDate              string   `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Round) Reset()         { *m = Round{} }
func (m *Round) String() string { return proto.CompactTextString(m) }
func (*Round) ProtoMessage()    {}
func (*Round) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa3b616c4946ca71, []int{0}
}

func (m *Round) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Round.Unmarshal(m, b)
}
func (m *Round) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Round.Marshal(b, m, deterministic)
}
func (m *Round) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Round.Merge(m, src)
}
func (m *Round) XXX_Size() int {
	return xxx_messageInfo_Round.Size(m)
}
func (m *Round) XXX_DiscardUnknown() {
	xxx_messageInfo_Round.DiscardUnknown(m)
}

var xxx_messageInfo_Round proto.InternalMessageInfo

func (m *Round) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Round) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Round) GetSeasonId() int64 {
	if m != nil {
		return m.SeasonId
	}
	return 0
}

func (m *Round) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *Round) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func init() {
	proto.RegisterType((*Round)(nil), "proto.Round")
}

func init() {
	proto.RegisterFile("internal/app/grpc/proto/round.proto", fileDescriptor_aa3b616c4946ca71)
}

var fileDescriptor_aa3b616c4946ca71 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0x8e, 0x3d, 0xae, 0xc2, 0x30,
	0x0c, 0x80, 0x95, 0xfe, 0xbc, 0xd7, 0x7a, 0x78, 0x43, 0xa6, 0x3c, 0x21, 0xa4, 0x0a, 0x96, 0x4e,
	0x64, 0xe0, 0x0a, 0x2c, 0xac, 0xbd, 0x40, 0x65, 0xb0, 0x85, 0x22, 0x81, 0x13, 0xa5, 0xe1, 0x02,
	0x9c, 0x1c, 0xd5, 0x9d, 0xfc, 0xe9, 0xfb, 0x6c, 0xc9, 0x70, 0x0c, 0x52, 0x38, 0x0b, 0x3e, 0x3d,
	0xa6, 0xe4, 0x1f, 0x39, 0xdd, 0x7d, 0xca, 0xb1, 0x44, 0x9f, 0xe3, 0x5b, 0xe8, 0xa4, 0x6c, 0x5b,
	0x1d, 0x87, 0x8f, 0x81, 0x76, 0x5a, 0xb5, 0xfd, 0x83, 0x2a, 0x90, 0x33, 0x83, 0x19, 0xeb, 0xa9,
	0x0a, 0x64, 0x2d, 0x34, 0x82, 0x2f, 0x76, 0xd5, 0x60, 0xc6, 0x7e, 0x52, 0xb6, 0x3b, 0xe8, 0x17,
	0xc6, 0x25, 0xca, 0x1c, 0xc8, 0xd5, 0xba, 0xda, 0x6d, 0xe2, 0x4a, 0x76, 0x0f, 0xb0, 0x14, 0xcc,
	0x65, 0x26, 0x2c, 0xec, 0x1a, 0x3d, 0xeb, 0xd5, 0x5c, 0xb0, 0xb0, 0xfd, 0x87, 0x8e, 0x85, 0xb6,
	0xd8, 0x6a, 0xfc, 0x65, 0xa1, 0x35, 0xdd, 0x7e, 0xf4, 0x97, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x33, 0x57, 0xcb, 0x76, 0xb9, 0x00, 0x00, 0x00,
}
