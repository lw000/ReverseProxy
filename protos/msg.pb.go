// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package proxy

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

//请求注册服务
type ReqRegService struct {
	SrvId                int32    `protobuf:"varint,1,opt,name=srvId,proto3" json:"srvId,omitempty"`
	SrvType              int32    `protobuf:"varint,2,opt,name=srvType,proto3" json:"srvType,omitempty"`
	SrvVersion           uint32   `protobuf:"varint,3,opt,name=srvVersion,proto3" json:"srvVersion,omitempty"`
	License              string   `protobuf:"bytes,4,opt,name=license,proto3" json:"license,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRegService) Reset()         { *m = ReqRegService{} }
func (m *ReqRegService) String() string { return proto.CompactTextString(m) }
func (*ReqRegService) ProtoMessage()    {}
func (*ReqRegService) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *ReqRegService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRegService.Unmarshal(m, b)
}
func (m *ReqRegService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRegService.Marshal(b, m, deterministic)
}
func (m *ReqRegService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRegService.Merge(m, src)
}
func (m *ReqRegService) XXX_Size() int {
	return xxx_messageInfo_ReqRegService.Size(m)
}
func (m *ReqRegService) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRegService.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRegService proto.InternalMessageInfo

func (m *ReqRegService) GetSrvId() int32 {
	if m != nil {
		return m.SrvId
	}
	return 0
}

func (m *ReqRegService) GetSrvType() int32 {
	if m != nil {
		return m.SrvType
	}
	return 0
}

func (m *ReqRegService) GetSrvVersion() uint32 {
	if m != nil {
		return m.SrvVersion
	}
	return 0
}

func (m *ReqRegService) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

//响应注册服务
type AckRegService struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckRegService) Reset()         { *m = AckRegService{} }
func (m *AckRegService) String() string { return proto.CompactTextString(m) }
func (*AckRegService) ProtoMessage()    {}
func (*AckRegService) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *AckRegService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckRegService.Unmarshal(m, b)
}
func (m *AckRegService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckRegService.Marshal(b, m, deterministic)
}
func (m *AckRegService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckRegService.Merge(m, src)
}
func (m *AckRegService) XXX_Size() int {
	return xxx_messageInfo_AckRegService.Size(m)
}
func (m *AckRegService) XXX_DiscardUnknown() {
	xxx_messageInfo_AckRegService.DiscardUnknown(m)
}

var xxx_messageInfo_AckRegService proto.InternalMessageInfo

func (m *AckRegService) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AckRegService) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

//请求更新服务配置
type ReqUpdateConfig struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Priority             uint32   `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUpdateConfig) Reset()         { *m = ReqUpdateConfig{} }
func (m *ReqUpdateConfig) String() string { return proto.CompactTextString(m) }
func (*ReqUpdateConfig) ProtoMessage()    {}
func (*ReqUpdateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *ReqUpdateConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUpdateConfig.Unmarshal(m, b)
}
func (m *ReqUpdateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUpdateConfig.Marshal(b, m, deterministic)
}
func (m *ReqUpdateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUpdateConfig.Merge(m, src)
}
func (m *ReqUpdateConfig) XXX_Size() int {
	return xxx_messageInfo_ReqUpdateConfig.Size(m)
}
func (m *ReqUpdateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUpdateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUpdateConfig proto.InternalMessageInfo

func (m *ReqUpdateConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ReqUpdateConfig) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

//回应更新服务配置
type AckUpdateConfig struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckUpdateConfig) Reset()         { *m = AckUpdateConfig{} }
func (m *AckUpdateConfig) String() string { return proto.CompactTextString(m) }
func (*AckUpdateConfig) ProtoMessage()    {}
func (*AckUpdateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *AckUpdateConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckUpdateConfig.Unmarshal(m, b)
}
func (m *AckUpdateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckUpdateConfig.Marshal(b, m, deterministic)
}
func (m *AckUpdateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckUpdateConfig.Merge(m, src)
}
func (m *AckUpdateConfig) XXX_Size() int {
	return xxx_messageInfo_AckUpdateConfig.Size(m)
}
func (m *AckUpdateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AckUpdateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AckUpdateConfig proto.InternalMessageInfo

func (m *AckUpdateConfig) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AckUpdateConfig) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqRegService)(nil), "proxy.ReqRegService")
	proto.RegisterType((*AckRegService)(nil), "proxy.AckRegService")
	proto.RegisterType((*ReqUpdateConfig)(nil), "proxy.ReqUpdateConfig")
	proto.RegisterType((*AckUpdateConfig)(nil), "proxy.AckUpdateConfig")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x89, 0x77, 0x51, 0x33, 0x10, 0x4e, 0x16, 0x8b, 0xc5, 0x42, 0x42, 0xaa, 0x54, 0x36,
	0x22, 0xd6, 0xc1, 0xca, 0x76, 0xfc, 0xd1, 0xeb, 0x66, 0x8c, 0xcb, 0x79, 0x99, 0xbd, 0x99, 0x10,
	0xdc, 0xff, 0x5e, 0xb2, 0x7a, 0x12, 0x6c, 0xec, 0xde, 0xf7, 0xe0, 0xcd, 0x07, 0x03, 0xc5, 0x4e,
	0xfb, 0xab, 0x20, 0x3c, 0xb2, 0xc9, 0x83, 0xf0, 0x67, 0xac, 0x23, 0x94, 0x48, 0x7b, 0xa4, 0xfe,
	0x81, 0x64, 0xf2, 0x8e, 0xcc, 0x39, 0xe4, 0x2a, 0xd3, 0x7d, 0x67, 0xb3, 0x2a, 0x6b, 0x72, 0xfc,
	0x06, 0x63, 0xe1, 0x44, 0x65, 0x7a, 0x8c, 0x81, 0xec, 0x51, 0xea, 0x0f, 0x68, 0x2e, 0x01, 0x54,
	0xa6, 0x67, 0x12, 0xf5, 0x3c, 0xd8, 0x55, 0x95, 0x35, 0x25, 0x2e, 0x9a, 0x79, 0xf9, 0xe1, 0x1d,
	0x0d, 0x4a, 0x76, 0x5d, 0x65, 0x4d, 0x81, 0x07, 0xac, 0x6f, 0xa0, 0x6c, 0xdd, 0x76, 0xa1, 0x36,
	0xb0, 0x76, 0xdc, 0xd1, 0x8f, 0x39, 0x65, 0x73, 0x06, 0xab, 0x9d, 0xf6, 0x49, 0x5a, 0xe0, 0x1c,
	0xeb, 0x16, 0x36, 0x48, 0xfb, 0xa7, 0xd0, 0xbd, 0x8c, 0x74, 0xc7, 0xc3, 0x9b, 0xef, 0xe7, 0xe1,
	0x3b, 0xeb, 0x98, 0x86, 0x05, 0xa6, 0x6c, 0x2e, 0xe0, 0x34, 0x88, 0x67, 0xf1, 0x63, 0x4c, 0xe2,
	0x12, 0x7f, 0xb9, 0xbe, 0x85, 0x4d, 0xeb, 0xb6, 0x7f, 0x4f, 0xfc, 0xef, 0x7e, 0x3d, 0x4e, 0xbf,
	0xbb, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x39, 0xa9, 0x90, 0xa9, 0x48, 0x01, 0x00, 0x00,
}
