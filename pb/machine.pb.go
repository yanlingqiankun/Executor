// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/machine.proto

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

type CreateMachineReq struct {
	ImageId              string                   `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Name                 string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Env                  []string                 `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	Tty                  bool                     `protobuf:"varint,4,opt,name=tty,proto3" json:"tty,omitempty"`
	Cmd                  []string                 `protobuf:"bytes,5,rep,name=cmd,proto3" json:"cmd,omitempty"`
	WorkingDir           string                   `protobuf:"bytes,6,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
	ExposedPorts         map[string]*PortBindings `protobuf:"bytes,7,rep,name=exposed_ports,json=exposedPorts,proto3" json:"exposed_ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CreateMachineReq) Reset()         { *m = CreateMachineReq{} }
func (m *CreateMachineReq) String() string { return proto.CompactTextString(m) }
func (*CreateMachineReq) ProtoMessage()    {}
func (*CreateMachineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{0}
}

func (m *CreateMachineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMachineReq.Unmarshal(m, b)
}
func (m *CreateMachineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMachineReq.Marshal(b, m, deterministic)
}
func (m *CreateMachineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMachineReq.Merge(m, src)
}
func (m *CreateMachineReq) XXX_Size() int {
	return xxx_messageInfo_CreateMachineReq.Size(m)
}
func (m *CreateMachineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMachineReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMachineReq proto.InternalMessageInfo

func (m *CreateMachineReq) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *CreateMachineReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateMachineReq) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *CreateMachineReq) GetTty() bool {
	if m != nil {
		return m.Tty
	}
	return false
}

func (m *CreateMachineReq) GetCmd() []string {
	if m != nil {
		return m.Cmd
	}
	return nil
}

func (m *CreateMachineReq) GetWorkingDir() string {
	if m != nil {
		return m.WorkingDir
	}
	return ""
}

func (m *CreateMachineReq) GetExposedPorts() map[string]*PortBindings {
	if m != nil {
		return m.ExposedPorts
	}
	return nil
}

type PortBinding struct {
	HostIp               string   `protobuf:"bytes,1,opt,name=host_ip,json=hostIp,proto3" json:"host_ip,omitempty"`
	HostPort             string   `protobuf:"bytes,2,opt,name=host_port,json=hostPort,proto3" json:"host_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortBinding) Reset()         { *m = PortBinding{} }
func (m *PortBinding) String() string { return proto.CompactTextString(m) }
func (*PortBinding) ProtoMessage()    {}
func (*PortBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{1}
}

func (m *PortBinding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortBinding.Unmarshal(m, b)
}
func (m *PortBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortBinding.Marshal(b, m, deterministic)
}
func (m *PortBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortBinding.Merge(m, src)
}
func (m *PortBinding) XXX_Size() int {
	return xxx_messageInfo_PortBinding.Size(m)
}
func (m *PortBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_PortBinding.DiscardUnknown(m)
}

var xxx_messageInfo_PortBinding proto.InternalMessageInfo

func (m *PortBinding) GetHostIp() string {
	if m != nil {
		return m.HostIp
	}
	return ""
}

func (m *PortBinding) GetHostPort() string {
	if m != nil {
		return m.HostPort
	}
	return ""
}

type PortBindings struct {
	PortBindings         []*PortBinding `protobuf:"bytes,1,rep,name=port_bindings,json=portBindings,proto3" json:"port_bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PortBindings) Reset()         { *m = PortBindings{} }
func (m *PortBindings) String() string { return proto.CompactTextString(m) }
func (*PortBindings) ProtoMessage()    {}
func (*PortBindings) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{2}
}

func (m *PortBindings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortBindings.Unmarshal(m, b)
}
func (m *PortBindings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortBindings.Marshal(b, m, deterministic)
}
func (m *PortBindings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortBindings.Merge(m, src)
}
func (m *PortBindings) XXX_Size() int {
	return xxx_messageInfo_PortBindings.Size(m)
}
func (m *PortBindings) XXX_DiscardUnknown() {
	xxx_messageInfo_PortBindings.DiscardUnknown(m)
}

var xxx_messageInfo_PortBindings proto.InternalMessageInfo

func (m *PortBindings) GetPortBindings() []*PortBinding {
	if m != nil {
		return m.PortBindings
	}
	return nil
}

type CreateMachineResp struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Err                  *Error   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMachineResp) Reset()         { *m = CreateMachineResp{} }
func (m *CreateMachineResp) String() string { return proto.CompactTextString(m) }
func (*CreateMachineResp) ProtoMessage()    {}
func (*CreateMachineResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{3}
}

func (m *CreateMachineResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMachineResp.Unmarshal(m, b)
}
func (m *CreateMachineResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMachineResp.Marshal(b, m, deterministic)
}
func (m *CreateMachineResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMachineResp.Merge(m, src)
}
func (m *CreateMachineResp) XXX_Size() int {
	return xxx_messageInfo_CreateMachineResp.Size(m)
}
func (m *CreateMachineResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMachineResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMachineResp proto.InternalMessageInfo

func (m *CreateMachineResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateMachineResp) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type DeleteMachineReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMachineReq) Reset()         { *m = DeleteMachineReq{} }
func (m *DeleteMachineReq) String() string { return proto.CompactTextString(m) }
func (*DeleteMachineReq) ProtoMessage()    {}
func (*DeleteMachineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{4}
}

func (m *DeleteMachineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMachineReq.Unmarshal(m, b)
}
func (m *DeleteMachineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMachineReq.Marshal(b, m, deterministic)
}
func (m *DeleteMachineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMachineReq.Merge(m, src)
}
func (m *DeleteMachineReq) XXX_Size() int {
	return xxx_messageInfo_DeleteMachineReq.Size(m)
}
func (m *DeleteMachineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMachineReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMachineReq proto.InternalMessageInfo

func (m *DeleteMachineReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListMachineResp struct {
	MachineInfos         []*MachineInfo `protobuf:"bytes,1,rep,name=machine_infos,json=machineInfos,proto3" json:"machine_infos,omitempty"`
	Err                  *Error         `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListMachineResp) Reset()         { *m = ListMachineResp{} }
func (m *ListMachineResp) String() string { return proto.CompactTextString(m) }
func (*ListMachineResp) ProtoMessage()    {}
func (*ListMachineResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{5}
}

func (m *ListMachineResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMachineResp.Unmarshal(m, b)
}
func (m *ListMachineResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMachineResp.Marshal(b, m, deterministic)
}
func (m *ListMachineResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMachineResp.Merge(m, src)
}
func (m *ListMachineResp) XXX_Size() int {
	return xxx_messageInfo_ListMachineResp.Size(m)
}
func (m *ListMachineResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMachineResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListMachineResp proto.InternalMessageInfo

func (m *ListMachineResp) GetMachineInfos() []*MachineInfo {
	if m != nil {
		return m.MachineInfos
	}
	return nil
}

func (m *ListMachineResp) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type MachineInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageName            string   `protobuf:"bytes,3,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	ImageType            string   `protobuf:"bytes,4,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
	CreateTime           string   `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	ImageId              string   `protobuf:"bytes,7,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineInfo) Reset()         { *m = MachineInfo{} }
func (m *MachineInfo) String() string { return proto.CompactTextString(m) }
func (*MachineInfo) ProtoMessage()    {}
func (*MachineInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{6}
}

func (m *MachineInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MachineInfo.Unmarshal(m, b)
}
func (m *MachineInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MachineInfo.Marshal(b, m, deterministic)
}
func (m *MachineInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineInfo.Merge(m, src)
}
func (m *MachineInfo) XXX_Size() int {
	return xxx_messageInfo_MachineInfo.Size(m)
}
func (m *MachineInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MachineInfo proto.InternalMessageInfo

func (m *MachineInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MachineInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MachineInfo) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *MachineInfo) GetImageType() string {
	if m != nil {
		return m.ImageType
	}
	return ""
}

func (m *MachineInfo) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *MachineInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *MachineInfo) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateMachineReq)(nil), "pb.CreateMachineReq")
	proto.RegisterMapType((map[string]*PortBindings)(nil), "pb.CreateMachineReq.ExposedPortsEntry")
	proto.RegisterType((*PortBinding)(nil), "pb.PortBinding")
	proto.RegisterType((*PortBindings)(nil), "pb.PortBindings")
	proto.RegisterType((*CreateMachineResp)(nil), "pb.CreateMachineResp")
	proto.RegisterType((*DeleteMachineReq)(nil), "pb.DeleteMachineReq")
	proto.RegisterType((*ListMachineResp)(nil), "pb.ListMachineResp")
	proto.RegisterType((*MachineInfo)(nil), "pb.MachineInfo")
}

func init() {
	proto.RegisterFile("pb/machine.proto", fileDescriptor_2daae29eac7cff8a)
}

var fileDescriptor_2daae29eac7cff8a = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xed, 0x36, 0x89, 0xc7, 0x09, 0x71, 0xf7, 0x00, 0xa6, 0x15, 0x22, 0xf2, 0xa1, 0xca,
	0x29, 0x95, 0x4a, 0x0f, 0x88, 0x13, 0xa2, 0xc9, 0x21, 0xe2, 0x43, 0x60, 0xf5, 0x6e, 0xf9, 0x63,
	0x9a, 0xae, 0xda, 0xfd, 0x60, 0xbd, 0x2d, 0xf8, 0x9f, 0x71, 0xe7, 0x8f, 0xa1, 0x5d, 0x2f, 0x60,
	0xa7, 0x88, 0xdb, 0xcc, 0x7b, 0x6f, 0xc6, 0x33, 0xeb, 0x37, 0x10, 0xcb, 0xf2, 0x8c, 0x15, 0xd5,
	0x0d, 0xe5, 0xb8, 0x92, 0x4a, 0x68, 0x41, 0x7c, 0x59, 0x1e, 0xcf, 0x65, 0x79, 0x56, 0x09, 0xc6,
	0x04, 0xef, 0xc0, 0xf4, 0x87, 0x0f, 0xf1, 0xa5, 0xc2, 0x42, 0xe3, 0xc7, 0x4e, 0x9c, 0xe1, 0x57,
	0xf2, 0x1c, 0x26, 0x94, 0x15, 0x3b, 0xcc, 0x69, 0x9d, 0x78, 0x0b, 0x6f, 0x19, 0x66, 0x63, 0x9b,
	0x6f, 0x6b, 0x42, 0xe0, 0x80, 0x17, 0x0c, 0x13, 0xdf, 0xc2, 0x36, 0x26, 0x31, 0x04, 0xc8, 0x1f,
	0x92, 0x60, 0x11, 0x2c, 0xc3, 0xcc, 0x84, 0x06, 0xd1, 0xba, 0x4d, 0x0e, 0x16, 0xde, 0x72, 0x92,
	0x99, 0xd0, 0x20, 0x15, 0xab, 0x93, 0xc3, 0x4e, 0x53, 0xb1, 0x9a, 0xbc, 0x84, 0xe8, 0x9b, 0x50,
	0xb7, 0x94, 0xef, 0xf2, 0x9a, 0xaa, 0x64, 0x64, 0x1b, 0x82, 0x83, 0xd6, 0x54, 0x91, 0xf7, 0x30,
	0xc3, 0xef, 0x52, 0x34, 0x58, 0xe7, 0x52, 0x28, 0xdd, 0x24, 0xe3, 0x45, 0xb0, 0x8c, 0xce, 0x4f,
	0x57, 0xb2, 0x5c, 0xed, 0x8f, 0xbc, 0xda, 0x74, 0xca, 0xcf, 0x46, 0xb8, 0xe1, 0x5a, 0xb5, 0xd9,
	0x14, 0x7b, 0xd0, 0xf1, 0x17, 0x38, 0x7a, 0x24, 0x31, 0x43, 0xdd, 0x62, 0xeb, 0x56, 0x34, 0x21,
	0x39, 0x85, 0xc3, 0x87, 0xe2, 0xee, 0xbe, 0xdb, 0x2f, 0x3a, 0x8f, 0xcd, 0xb7, 0x4c, 0xc1, 0x3b,
	0xca, 0x6b, 0xca, 0x77, 0x4d, 0xd6, 0xd1, 0x6f, 0xfc, 0xd7, 0x5e, 0x7a, 0x09, 0x51, 0x8f, 0x22,
	0xcf, 0x60, 0x7c, 0x23, 0x1a, 0x9d, 0x53, 0xe9, 0x1a, 0x8e, 0x4c, 0xba, 0x95, 0xe4, 0x04, 0x42,
	0x4b, 0x98, 0x25, 0xdc, 0xbb, 0x4d, 0x0c, 0x60, 0x8a, 0xd3, 0x35, 0x4c, 0xfb, 0xfd, 0xc9, 0x05,
	0xcc, 0x8c, 0x2e, 0x2f, 0x1d, 0x90, 0x78, 0x76, 0xe9, 0xf9, 0xde, 0x20, 0xd9, 0x54, 0xf6, 0xaa,
	0xd2, 0xb7, 0x70, 0xb4, 0xf7, 0x22, 0x8d, 0x24, 0x4f, 0xc0, 0xff, 0xf3, 0xff, 0x7c, 0x5a, 0x93,
	0x13, 0x08, 0x50, 0x29, 0xb7, 0x59, 0x68, 0x1a, 0x6e, 0x94, 0x12, 0x2a, 0x33, 0x68, 0x9a, 0x42,
	0xbc, 0xc6, 0x3b, 0x1c, 0xd8, 0x60, 0xaf, 0x41, 0x5a, 0xc3, 0xfc, 0x03, 0x6d, 0x74, 0xff, 0x1b,
	0x17, 0x30, 0x73, 0x26, 0xcb, 0x29, 0xbf, 0x16, 0x83, 0x71, 0x9d, 0x6e, 0xcb, 0xaf, 0x45, 0x36,
	0x65, 0x7f, 0x93, 0xe6, 0xff, 0x93, 0xfc, 0xf4, 0x20, 0xea, 0x95, 0x3e, 0x5a, 0xe3, 0x5f, 0x0e,
	0x7c, 0x01, 0xd0, 0x19, 0xd6, 0x32, 0x81, 0x65, 0x42, 0x8b, 0x7c, 0x1a, 0xd0, 0xba, 0x95, 0x68,
	0x5d, 0xf9, 0x9b, 0xbe, 0x6a, 0x25, 0x1a, 0x27, 0x56, 0xf6, 0xf5, 0x72, 0x4d, 0x19, 0x26, 0x87,
	0x9d, 0x13, 0x3b, 0xe8, 0x8a, 0x32, 0x24, 0x4f, 0x61, 0xd4, 0xe8, 0x42, 0xdf, 0x37, 0xce, 0xa5,
	0x2e, 0x1b, 0xdc, 0xc9, 0x78, 0x70, 0x27, 0xe5, 0xc8, 0x9e, 0xd7, 0xab, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xd1, 0x67, 0xf8, 0xf1, 0x87, 0x03, 0x00, 0x00,
}
