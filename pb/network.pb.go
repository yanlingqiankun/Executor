// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/network.proto

package pb

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

type NetworkName struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkName) Reset()         { *m = NetworkName{} }
func (m *NetworkName) String() string { return proto.CompactTextString(m) }
func (*NetworkName) ProtoMessage()    {}
func (*NetworkName) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ca4385605dfbfc, []int{0}
}

func (m *NetworkName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkName.Unmarshal(m, b)
}
func (m *NetworkName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkName.Marshal(b, m, deterministic)
}
func (m *NetworkName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkName.Merge(m, src)
}
func (m *NetworkName) XXX_Size() int {
	return xxx_messageInfo_NetworkName.Size(m)
}
func (m *NetworkName) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkName.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkName proto.InternalMessageInfo

func (m *NetworkName) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type NetworkCreateReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Subnet               string   `protobuf:"bytes,2,opt,name=subnet,proto3" json:"subnet,omitempty"`
	Gateway              string   `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Isolated             bool     `protobuf:"varint,4,opt,name=isolated,proto3" json:"isolated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkCreateReq) Reset()         { *m = NetworkCreateReq{} }
func (m *NetworkCreateReq) String() string { return proto.CompactTextString(m) }
func (*NetworkCreateReq) ProtoMessage()    {}
func (*NetworkCreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ca4385605dfbfc, []int{1}
}

func (m *NetworkCreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkCreateReq.Unmarshal(m, b)
}
func (m *NetworkCreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkCreateReq.Marshal(b, m, deterministic)
}
func (m *NetworkCreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkCreateReq.Merge(m, src)
}
func (m *NetworkCreateReq) XXX_Size() int {
	return xxx_messageInfo_NetworkCreateReq.Size(m)
}
func (m *NetworkCreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkCreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkCreateReq proto.InternalMessageInfo

func (m *NetworkCreateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkCreateReq) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *NetworkCreateReq) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *NetworkCreateReq) GetIsolated() bool {
	if m != nil {
		return m.Isolated
	}
	return false
}

type NetworkCreateResp struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkCreateResp) Reset()         { *m = NetworkCreateResp{} }
func (m *NetworkCreateResp) String() string { return proto.CompactTextString(m) }
func (*NetworkCreateResp) ProtoMessage()    {}
func (*NetworkCreateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ca4385605dfbfc, []int{2}
}

func (m *NetworkCreateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkCreateResp.Unmarshal(m, b)
}
func (m *NetworkCreateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkCreateResp.Marshal(b, m, deterministic)
}
func (m *NetworkCreateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkCreateResp.Merge(m, src)
}
func (m *NetworkCreateResp) XXX_Size() int {
	return xxx_messageInfo_NetworkCreateResp.Size(m)
}
func (m *NetworkCreateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkCreateResp.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkCreateResp proto.InternalMessageInfo

func (m *NetworkCreateResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NetworkCreateResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type NetworkDeleteReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Force                bool     `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkDeleteReq) Reset()         { *m = NetworkDeleteReq{} }
func (m *NetworkDeleteReq) String() string { return proto.CompactTextString(m) }
func (*NetworkDeleteReq) ProtoMessage()    {}
func (*NetworkDeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ca4385605dfbfc, []int{3}
}

func (m *NetworkDeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkDeleteReq.Unmarshal(m, b)
}
func (m *NetworkDeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkDeleteReq.Marshal(b, m, deterministic)
}
func (m *NetworkDeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkDeleteReq.Merge(m, src)
}
func (m *NetworkDeleteReq) XXX_Size() int {
	return xxx_messageInfo_NetworkDeleteReq.Size(m)
}
func (m *NetworkDeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkDeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkDeleteReq proto.InternalMessageInfo

func (m *NetworkDeleteReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkDeleteReq) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

type NetworkDeleteResp struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkDeleteResp) Reset()         { *m = NetworkDeleteResp{} }
func (m *NetworkDeleteResp) String() string { return proto.CompactTextString(m) }
func (*NetworkDeleteResp) ProtoMessage()    {}
func (*NetworkDeleteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ca4385605dfbfc, []int{4}
}

func (m *NetworkDeleteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkDeleteResp.Unmarshal(m, b)
}
func (m *NetworkDeleteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkDeleteResp.Marshal(b, m, deterministic)
}
func (m *NetworkDeleteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkDeleteResp.Merge(m, src)
}
func (m *NetworkDeleteResp) XXX_Size() int {
	return xxx_messageInfo_NetworkDeleteResp.Size(m)
}
func (m *NetworkDeleteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkDeleteResp.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkDeleteResp proto.InternalMessageInfo

func (m *NetworkDeleteResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type NetworkInspectReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInspectReq) Reset()         { *m = NetworkInspectReq{} }
func (m *NetworkInspectReq) String() string { return proto.CompactTextString(m) }
func (*NetworkInspectReq) ProtoMessage()    {}
func (*NetworkInspectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ca4385605dfbfc, []int{5}
}

func (m *NetworkInspectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInspectReq.Unmarshal(m, b)
}
func (m *NetworkInspectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInspectReq.Marshal(b, m, deterministic)
}
func (m *NetworkInspectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInspectReq.Merge(m, src)
}
func (m *NetworkInspectReq) XXX_Size() int {
	return xxx_messageInfo_NetworkInspectReq.Size(m)
}
func (m *NetworkInspectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInspectReq.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInspectReq proto.InternalMessageInfo

func (m *NetworkInspectReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NetworkInspectResp struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	NetInfo              string   `protobuf:"bytes,2,opt,name=net_info,json=netInfo,proto3" json:"net_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInspectResp) Reset()         { *m = NetworkInspectResp{} }
func (m *NetworkInspectResp) String() string { return proto.CompactTextString(m) }
func (*NetworkInspectResp) ProtoMessage()    {}
func (*NetworkInspectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ca4385605dfbfc, []int{6}
}

func (m *NetworkInspectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInspectResp.Unmarshal(m, b)
}
func (m *NetworkInspectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInspectResp.Marshal(b, m, deterministic)
}
func (m *NetworkInspectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInspectResp.Merge(m, src)
}
func (m *NetworkInspectResp) XXX_Size() int {
	return xxx_messageInfo_NetworkInspectResp.Size(m)
}
func (m *NetworkInspectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInspectResp.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInspectResp proto.InternalMessageInfo

func (m *NetworkInspectResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *NetworkInspectResp) GetNetInfo() string {
	if m != nil {
		return m.NetInfo
	}
	return ""
}

type NetworkInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           string   `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Gateway              string   `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Subnet               string   `protobuf:"bytes,4,opt,name=subnet,proto3" json:"subnet,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInfo) Reset()         { *m = NetworkInfo{} }
func (m *NetworkInfo) String() string { return proto.CompactTextString(m) }
func (*NetworkInfo) ProtoMessage()    {}
func (*NetworkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ca4385605dfbfc, []int{7}
}

func (m *NetworkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInfo.Unmarshal(m, b)
}
func (m *NetworkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInfo.Marshal(b, m, deterministic)
}
func (m *NetworkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInfo.Merge(m, src)
}
func (m *NetworkInfo) XXX_Size() int {
	return xxx_messageInfo_NetworkInfo.Size(m)
}
func (m *NetworkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInfo proto.InternalMessageInfo

func (m *NetworkInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkInfo) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *NetworkInfo) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *NetworkInfo) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *NetworkInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type NetworkListResp struct {
	Networks             []*NetworkInfo `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
	Error                *Error         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NetworkListResp) Reset()         { *m = NetworkListResp{} }
func (m *NetworkListResp) String() string { return proto.CompactTextString(m) }
func (*NetworkListResp) ProtoMessage()    {}
func (*NetworkListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ca4385605dfbfc, []int{8}
}

func (m *NetworkListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkListResp.Unmarshal(m, b)
}
func (m *NetworkListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkListResp.Marshal(b, m, deterministic)
}
func (m *NetworkListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkListResp.Merge(m, src)
}
func (m *NetworkListResp) XXX_Size() int {
	return xxx_messageInfo_NetworkListResp.Size(m)
}
func (m *NetworkListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkListResp.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkListResp proto.InternalMessageInfo

func (m *NetworkListResp) GetNetworks() []*NetworkInfo {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *NetworkListResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkName)(nil), "pb.NetworkName")
	proto.RegisterType((*NetworkCreateReq)(nil), "pb.NetworkCreateReq")
	proto.RegisterType((*NetworkCreateResp)(nil), "pb.NetworkCreateResp")
	proto.RegisterType((*NetworkDeleteReq)(nil), "pb.NetworkDeleteReq")
	proto.RegisterType((*NetworkDeleteResp)(nil), "pb.NetworkDeleteResp")
	proto.RegisterType((*NetworkInspectReq)(nil), "pb.NetworkInspectReq")
	proto.RegisterType((*NetworkInspectResp)(nil), "pb.NetworkInspectResp")
	proto.RegisterType((*NetworkInfo)(nil), "pb.NetworkInfo")
	proto.RegisterType((*NetworkListResp)(nil), "pb.NetworkListResp")
}

func init() {
	proto.RegisterFile("pb/network.proto", fileDescriptor_18ca4385605dfbfc)
}

var fileDescriptor_18ca4385605dfbfc = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x51, 0x6b, 0xf2, 0x30,
	0x14, 0xa5, 0xb5, 0x6a, 0xbd, 0x85, 0x4f, 0xbf, 0xf0, 0xf1, 0xd1, 0x09, 0x43, 0xc9, 0xcb, 0x84,
	0x81, 0x82, 0xdb, 0xe3, 0xde, 0xe6, 0x1e, 0x84, 0x21, 0xa3, 0xec, 0x5d, 0xda, 0x7a, 0x1d, 0x61,
	0x36, 0xc9, 0x92, 0x0c, 0xf1, 0x1f, 0xec, 0x67, 0x8f, 0x26, 0xb5, 0x6e, 0x32, 0xb7, 0xbd, 0xe5,
	0x9c, 0xdc, 0xdc, 0x73, 0xee, 0x3d, 0x81, 0x9e, 0xcc, 0x26, 0x1c, 0xcd, 0x56, 0xa8, 0xe7, 0xb1,
	0x54, 0xc2, 0x08, 0xe2, 0xcb, 0xac, 0xdf, 0x95, 0xd9, 0x24, 0x17, 0x45, 0x21, 0xb8, 0x23, 0xe9,
	0x39, 0x44, 0x0b, 0x57, 0xb5, 0x48, 0x0b, 0x24, 0x7f, 0xc0, 0x67, 0xab, 0xd8, 0x1b, 0x7a, 0xa3,
	0x4e, 0xe2, 0xb3, 0x15, 0x35, 0xd0, 0xab, 0xae, 0x6f, 0x15, 0xa6, 0x06, 0x13, 0x7c, 0x21, 0x04,
	0x02, 0x9e, 0x16, 0x58, 0x55, 0xd9, 0x33, 0xf9, 0x0f, 0x2d, 0xfd, 0x9a, 0x71, 0x34, 0xb1, 0x6f,
	0xd9, 0x0a, 0x91, 0x18, 0xda, 0x4f, 0xa9, 0xc1, 0x6d, 0xba, 0x8b, 0x1b, 0xf6, 0x62, 0x0f, 0x49,
	0x1f, 0x42, 0xa6, 0xc5, 0x26, 0x35, 0xb8, 0x8a, 0x83, 0xa1, 0x37, 0x0a, 0x93, 0x1a, 0xd3, 0x19,
	0xfc, 0x3d, 0x52, 0xd5, 0xf2, 0xd8, 0x1a, 0x19, 0x40, 0x13, 0x95, 0x12, 0xca, 0x2a, 0x46, 0xd3,
	0xce, 0x58, 0x66, 0xe3, 0xbb, 0x92, 0x48, 0x1c, 0x4f, 0x6f, 0x6a, 0xef, 0x33, 0xdc, 0xe0, 0x69,
	0xef, 0xff, 0xa0, 0xb9, 0x16, 0x2a, 0x47, 0xdb, 0x28, 0x4c, 0x1c, 0xa0, 0xd7, 0xb5, 0x87, 0xfd,
	0x6b, 0x2d, 0x0f, 0x9a, 0xde, 0x09, 0xcd, 0x8b, 0xfa, 0xd5, 0x9c, 0x6b, 0x89, 0xb9, 0x39, 0x21,
	0x4a, 0x1f, 0x80, 0x1c, 0x17, 0xfe, 0xa2, 0x3f, 0x39, 0x83, 0x90, 0xa3, 0x59, 0x32, 0xbe, 0x16,
	0xd5, 0xa6, 0xdb, 0x1c, 0xcd, 0x9c, 0xaf, 0x05, 0x7d, 0xf3, 0xea, 0x28, 0x4b, 0xfc, 0xe5, 0xa8,
	0x03, 0x88, 0x72, 0xbb, 0xd1, 0xa5, 0x61, 0x05, 0x56, 0x1d, 0xc0, 0x51, 0x8f, 0xac, 0xc0, 0x6f,
	0xf2, 0x3a, 0x24, 0x1c, 0x7c, 0x4a, 0x98, 0x40, 0x60, 0x76, 0x12, 0xe3, 0xa6, 0x93, 0x29, 0xcf,
	0x74, 0x09, 0xdd, 0xca, 0xc9, 0x3d, 0xd3, 0x6e, 0xb2, 0x4b, 0x6b, 0xbc, 0xa4, 0x74, 0xec, 0x0d,
	0x1b, 0xa3, 0x68, 0xda, 0x2d, 0x87, 0xfb, 0x60, 0x38, 0xa9, 0x0b, 0x7e, 0x8c, 0x36, 0x6b, 0xd9,
	0xcf, 0x7b, 0xf5, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x27, 0xe9, 0xe2, 0xed, 0xe5, 0x02, 0x00, 0x00,
}
