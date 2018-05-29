// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snapshot.proto

package snapshot

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Cmd int32

const (
	Cmd_Begin    Cmd = 0
	Cmd_Clear    Cmd = 1
	Cmd_Snapshot Cmd = 2
	Cmd_Restore  Cmd = 3
	Cmd_End      Cmd = 4
)

var Cmd_name = map[int32]string{
	0: "Begin",
	1: "Clear",
	2: "Snapshot",
	3: "Restore",
	4: "End",
}
var Cmd_value = map[string]int32{
	"Begin":    0,
	"Clear":    1,
	"Snapshot": 2,
	"Restore":  3,
	"End":      4,
}

func (x Cmd) String() string {
	return proto.EnumName(Cmd_name, int32(x))
}
func (Cmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_snapshot_c4331fe13f41e491, []int{0}
}

type Resp int32

const (
	Resp_BeginAck    Resp = 0
	Resp_ClearAck    Resp = 1
	Resp_SnapshotAck Resp = 2
	Resp_RestoreAck  Resp = 3
	Resp_EndAck      Resp = 4
)

var Resp_name = map[int32]string{
	0: "BeginAck",
	1: "ClearAck",
	2: "SnapshotAck",
	3: "RestoreAck",
	4: "EndAck",
}
var Resp_value = map[string]int32{
	"BeginAck":    0,
	"ClearAck":    1,
	"SnapshotAck": 2,
	"RestoreAck":  3,
	"EndAck":      4,
}

func (x Resp) String() string {
	return proto.EnumName(Resp_name, int32(x))
}
func (Resp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_snapshot_c4331fe13f41e491, []int{1}
}

type SnapshotReq struct {
	Cmd                  Cmd      `protobuf:"varint,1,opt,name=cmd,enum=Cmd" json:"cmd,omitempty"`
	StartHeight          uint64   `protobuf:"varint,2,opt,name=start_height,json=startHeight" json:"start_height,omitempty"`
	EndHeight            uint64   `protobuf:"varint,3,opt,name=end_height,json=endHeight" json:"end_height,omitempty"`
	File                 string   `protobuf:"bytes,4,opt,name=file" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotReq) Reset()         { *m = SnapshotReq{} }
func (m *SnapshotReq) String() string { return proto.CompactTextString(m) }
func (*SnapshotReq) ProtoMessage()    {}
func (*SnapshotReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_snapshot_c4331fe13f41e491, []int{0}
}
func (m *SnapshotReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotReq.Unmarshal(m, b)
}
func (m *SnapshotReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotReq.Marshal(b, m, deterministic)
}
func (dst *SnapshotReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotReq.Merge(dst, src)
}
func (m *SnapshotReq) XXX_Size() int {
	return xxx_messageInfo_SnapshotReq.Size(m)
}
func (m *SnapshotReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotReq.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotReq proto.InternalMessageInfo

func (m *SnapshotReq) GetCmd() Cmd {
	if m != nil {
		return m.Cmd
	}
	return Cmd_Begin
}

func (m *SnapshotReq) GetStartHeight() uint64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *SnapshotReq) GetEndHeight() uint64 {
	if m != nil {
		return m.EndHeight
	}
	return 0
}

func (m *SnapshotReq) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

type SnapshotResp struct {
	Resp                 Resp     `protobuf:"varint,1,opt,name=resp,enum=Resp" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotResp) Reset()         { *m = SnapshotResp{} }
func (m *SnapshotResp) String() string { return proto.CompactTextString(m) }
func (*SnapshotResp) ProtoMessage()    {}
func (*SnapshotResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_snapshot_c4331fe13f41e491, []int{1}
}
func (m *SnapshotResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotResp.Unmarshal(m, b)
}
func (m *SnapshotResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotResp.Marshal(b, m, deterministic)
}
func (dst *SnapshotResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotResp.Merge(dst, src)
}
func (m *SnapshotResp) XXX_Size() int {
	return xxx_messageInfo_SnapshotResp.Size(m)
}
func (m *SnapshotResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotResp.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotResp proto.InternalMessageInfo

func (m *SnapshotResp) GetResp() Resp {
	if m != nil {
		return m.Resp
	}
	return Resp_BeginAck
}

func init() {
	proto.RegisterType((*SnapshotReq)(nil), "SnapshotReq")
	proto.RegisterType((*SnapshotResp)(nil), "SnapshotResp")
	proto.RegisterEnum("Cmd", Cmd_name, Cmd_value)
	proto.RegisterEnum("Resp", Resp_name, Resp_value)
}

func init() { proto.RegisterFile("snapshot.proto", fileDescriptor_snapshot_c4331fe13f41e491) }

var fileDescriptor_snapshot_c4331fe13f41e491 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x85, 0xdd, 0xb6, 0xf2, 0x43, 0x70, 0x33, 0x07, 0x83, 0x07, 0x13, 0xed, 0x49, 0x7b,
	0xe0, 0xa0, 0x0f, 0x60, 0xb4, 0x31, 0xf1, 0x66, 0xb2, 0x3e, 0x80, 0xc1, 0xb2, 0x16, 0x62, 0x0b,
	0xb8, 0xec, 0xcd, 0x97, 0x77, 0x76, 0x43, 0xc3, 0x89, 0x99, 0x6f, 0x7e, 0xbe, 0xd9, 0x0c, 0xf2,
	0xb1, 0xab, 0x86, 0xb1, 0xe9, 0x5d, 0x39, 0xd8, 0xde, 0xf5, 0xeb, 0x3f, 0xa4, 0x1f, 0x13, 0xd1,
	0xe6, 0x97, 0x2e, 0x21, 0x76, 0xc7, 0xba, 0x88, 0x6e, 0xa2, 0xbb, 0xfc, 0x41, 0x96, 0xdb, 0x63,
	0xad, 0x3d, 0xa0, 0x5b, 0x64, 0xa3, 0xab, 0xac, 0xfb, 0x6c, 0x4c, 0xbb, 0x6f, 0x5c, 0x11, 0x73,
	0x40, 0xea, 0x34, 0xb0, 0xb7, 0x80, 0xe8, 0x1a, 0x30, 0x5d, 0x7d, 0x0a, 0x88, 0x10, 0x48, 0x98,
	0x4c, 0x63, 0x82, 0xfc, 0x6e, 0x0f, 0xa6, 0x90, 0x3c, 0x48, 0x74, 0xa8, 0xd7, 0xf7, 0xc8, 0xe6,
	0xe5, 0xe3, 0x40, 0x57, 0x90, 0x96, 0xbf, 0xd3, 0xfa, 0x45, 0xe9, 0xa1, 0x0e, 0x68, 0xf3, 0x04,
	0xc1, 0x8f, 0xa1, 0x04, 0x8b, 0x17, 0xb3, 0x6f, 0x3b, 0x75, 0xe6, 0xcb, 0xed, 0xc1, 0x54, 0x56,
	0x45, 0x94, 0xe1, 0xfc, 0xe4, 0x51, 0x31, 0xa5, 0x58, 0xf1, 0x8f, 0xae, 0xb7, 0x46, 0x09, 0x5a,
	0x41, 0xbc, 0x76, 0xb5, 0x92, 0x9b, 0x77, 0xc8, 0xb0, 0x83, 0xb3, 0xc1, 0xf0, 0xbc, 0xfb, 0x61,
	0x09, 0x77, 0x41, 0xe2, 0xbb, 0x88, 0x2e, 0xe6, 0x63, 0x78, 0x10, 0x53, 0x0e, 0x4c, 0x2a, 0xdf,
	0x0b, 0x02, 0x96, 0x6c, 0xf3, 0xb5, 0xfc, 0x5a, 0x86, 0x03, 0x3e, 0xfe, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x4d, 0xaf, 0xd1, 0x68, 0x52, 0x01, 0x00, 0x00,
}