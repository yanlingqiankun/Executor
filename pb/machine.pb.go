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

type StartMachineReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartMachineReq) Reset()         { *m = StartMachineReq{} }
func (m *StartMachineReq) String() string { return proto.CompactTextString(m) }
func (*StartMachineReq) ProtoMessage()    {}
func (*StartMachineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{7}
}

func (m *StartMachineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartMachineReq.Unmarshal(m, b)
}
func (m *StartMachineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartMachineReq.Marshal(b, m, deterministic)
}
func (m *StartMachineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartMachineReq.Merge(m, src)
}
func (m *StartMachineReq) XXX_Size() int {
	return xxx_messageInfo_StartMachineReq.Size(m)
}
func (m *StartMachineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StartMachineReq.DiscardUnknown(m)
}

var xxx_messageInfo_StartMachineReq proto.InternalMessageInfo

func (m *StartMachineReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MachineIdReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineIdReq) Reset()         { *m = MachineIdReq{} }
func (m *MachineIdReq) String() string { return proto.CompactTextString(m) }
func (*MachineIdReq) ProtoMessage()    {}
func (*MachineIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{8}
}

func (m *MachineIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MachineIdReq.Unmarshal(m, b)
}
func (m *MachineIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MachineIdReq.Marshal(b, m, deterministic)
}
func (m *MachineIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineIdReq.Merge(m, src)
}
func (m *MachineIdReq) XXX_Size() int {
	return xxx_messageInfo_MachineIdReq.Size(m)
}
func (m *MachineIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_MachineIdReq proto.InternalMessageInfo

func (m *MachineIdReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type KillMachineReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Signal               string   `protobuf:"bytes,2,opt,name=signal,proto3" json:"signal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KillMachineReq) Reset()         { *m = KillMachineReq{} }
func (m *KillMachineReq) String() string { return proto.CompactTextString(m) }
func (*KillMachineReq) ProtoMessage()    {}
func (*KillMachineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{9}
}

func (m *KillMachineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KillMachineReq.Unmarshal(m, b)
}
func (m *KillMachineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KillMachineReq.Marshal(b, m, deterministic)
}
func (m *KillMachineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KillMachineReq.Merge(m, src)
}
func (m *KillMachineReq) XXX_Size() int {
	return xxx_messageInfo_KillMachineReq.Size(m)
}
func (m *KillMachineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_KillMachineReq.DiscardUnknown(m)
}

var xxx_messageInfo_KillMachineReq proto.InternalMessageInfo

func (m *KillMachineReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *KillMachineReq) GetSignal() string {
	if m != nil {
		return m.Signal
	}
	return ""
}

type StopMachineReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timeout              int32    `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopMachineReq) Reset()         { *m = StopMachineReq{} }
func (m *StopMachineReq) String() string { return proto.CompactTextString(m) }
func (*StopMachineReq) ProtoMessage()    {}
func (*StopMachineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{10}
}

func (m *StopMachineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopMachineReq.Unmarshal(m, b)
}
func (m *StopMachineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopMachineReq.Marshal(b, m, deterministic)
}
func (m *StopMachineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopMachineReq.Merge(m, src)
}
func (m *StopMachineReq) XXX_Size() int {
	return xxx_messageInfo_StopMachineReq.Size(m)
}
func (m *StopMachineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StopMachineReq.DiscardUnknown(m)
}

var xxx_messageInfo_StopMachineReq proto.InternalMessageInfo

func (m *StopMachineReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StopMachineReq) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type RenameMachineReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameMachineReq) Reset()         { *m = RenameMachineReq{} }
func (m *RenameMachineReq) String() string { return proto.CompactTextString(m) }
func (*RenameMachineReq) ProtoMessage()    {}
func (*RenameMachineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{11}
}

func (m *RenameMachineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameMachineReq.Unmarshal(m, b)
}
func (m *RenameMachineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameMachineReq.Marshal(b, m, deterministic)
}
func (m *RenameMachineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameMachineReq.Merge(m, src)
}
func (m *RenameMachineReq) XXX_Size() int {
	return xxx_messageInfo_RenameMachineReq.Size(m)
}
func (m *RenameMachineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameMachineReq.DiscardUnknown(m)
}

var xxx_messageInfo_RenameMachineReq proto.InternalMessageInfo

func (m *RenameMachineReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RenameMachineReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RestartMachineReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timeout              int32    `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestartMachineReq) Reset()         { *m = RestartMachineReq{} }
func (m *RestartMachineReq) String() string { return proto.CompactTextString(m) }
func (*RestartMachineReq) ProtoMessage()    {}
func (*RestartMachineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{12}
}

func (m *RestartMachineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestartMachineReq.Unmarshal(m, b)
}
func (m *RestartMachineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestartMachineReq.Marshal(b, m, deterministic)
}
func (m *RestartMachineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestartMachineReq.Merge(m, src)
}
func (m *RestartMachineReq) XXX_Size() int {
	return xxx_messageInfo_RestartMachineReq.Size(m)
}
func (m *RestartMachineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RestartMachineReq.DiscardUnknown(m)
}

var xxx_messageInfo_RestartMachineReq proto.InternalMessageInfo

func (m *RestartMachineReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RestartMachineReq) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type AttachStreamIn struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachStreamIn) Reset()         { *m = AttachStreamIn{} }
func (m *AttachStreamIn) String() string { return proto.CompactTextString(m) }
func (*AttachStreamIn) ProtoMessage()    {}
func (*AttachStreamIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{13}
}

func (m *AttachStreamIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachStreamIn.Unmarshal(m, b)
}
func (m *AttachStreamIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachStreamIn.Marshal(b, m, deterministic)
}
func (m *AttachStreamIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachStreamIn.Merge(m, src)
}
func (m *AttachStreamIn) XXX_Size() int {
	return xxx_messageInfo_AttachStreamIn.Size(m)
}
func (m *AttachStreamIn) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachStreamIn.DiscardUnknown(m)
}

var xxx_messageInfo_AttachStreamIn proto.InternalMessageInfo

func (m *AttachStreamIn) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AttachStreamIn) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type AttachStreamOut struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachStreamOut) Reset()         { *m = AttachStreamOut{} }
func (m *AttachStreamOut) String() string { return proto.CompactTextString(m) }
func (*AttachStreamOut) ProtoMessage()    {}
func (*AttachStreamOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{14}
}

func (m *AttachStreamOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachStreamOut.Unmarshal(m, b)
}
func (m *AttachStreamOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachStreamOut.Marshal(b, m, deterministic)
}
func (m *AttachStreamOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachStreamOut.Merge(m, src)
}
func (m *AttachStreamOut) XXX_Size() int {
	return xxx_messageInfo_AttachStreamOut.Size(m)
}
func (m *AttachStreamOut) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachStreamOut.DiscardUnknown(m)
}

var xxx_messageInfo_AttachStreamOut proto.InternalMessageInfo

func (m *AttachStreamOut) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type ResizeTTYReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Width                uint32   `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResizeTTYReq) Reset()         { *m = ResizeTTYReq{} }
func (m *ResizeTTYReq) String() string { return proto.CompactTextString(m) }
func (*ResizeTTYReq) ProtoMessage()    {}
func (*ResizeTTYReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{15}
}

func (m *ResizeTTYReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResizeTTYReq.Unmarshal(m, b)
}
func (m *ResizeTTYReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResizeTTYReq.Marshal(b, m, deterministic)
}
func (m *ResizeTTYReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResizeTTYReq.Merge(m, src)
}
func (m *ResizeTTYReq) XXX_Size() int {
	return xxx_messageInfo_ResizeTTYReq.Size(m)
}
func (m *ResizeTTYReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ResizeTTYReq.DiscardUnknown(m)
}

var xxx_messageInfo_ResizeTTYReq proto.InternalMessageInfo

func (m *ResizeTTYReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResizeTTYReq) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ResizeTTYReq) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

type CanAttachJudgeResp struct {
	Tty                  bool     `protobuf:"varint,1,opt,name=tty,proto3" json:"tty,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	ImageType            string   `protobuf:"bytes,3,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CanAttachJudgeResp) Reset()         { *m = CanAttachJudgeResp{} }
func (m *CanAttachJudgeResp) String() string { return proto.CompactTextString(m) }
func (*CanAttachJudgeResp) ProtoMessage()    {}
func (*CanAttachJudgeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2daae29eac7cff8a, []int{16}
}

func (m *CanAttachJudgeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CanAttachJudgeResp.Unmarshal(m, b)
}
func (m *CanAttachJudgeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CanAttachJudgeResp.Marshal(b, m, deterministic)
}
func (m *CanAttachJudgeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanAttachJudgeResp.Merge(m, src)
}
func (m *CanAttachJudgeResp) XXX_Size() int {
	return xxx_messageInfo_CanAttachJudgeResp.Size(m)
}
func (m *CanAttachJudgeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CanAttachJudgeResp.DiscardUnknown(m)
}

var xxx_messageInfo_CanAttachJudgeResp proto.InternalMessageInfo

func (m *CanAttachJudgeResp) GetTty() bool {
	if m != nil {
		return m.Tty
	}
	return false
}

func (m *CanAttachJudgeResp) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CanAttachJudgeResp) GetImageType() string {
	if m != nil {
		return m.ImageType
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
	proto.RegisterType((*StartMachineReq)(nil), "pb.StartMachineReq")
	proto.RegisterType((*MachineIdReq)(nil), "pb.MachineIdReq")
	proto.RegisterType((*KillMachineReq)(nil), "pb.KillMachineReq")
	proto.RegisterType((*StopMachineReq)(nil), "pb.StopMachineReq")
	proto.RegisterType((*RenameMachineReq)(nil), "pb.RenameMachineReq")
	proto.RegisterType((*RestartMachineReq)(nil), "pb.RestartMachineReq")
	proto.RegisterType((*AttachStreamIn)(nil), "pb.AttachStreamIn")
	proto.RegisterType((*AttachStreamOut)(nil), "pb.AttachStreamOut")
	proto.RegisterType((*ResizeTTYReq)(nil), "pb.ResizeTTYReq")
	proto.RegisterType((*CanAttachJudgeResp)(nil), "pb.CanAttachJudgeResp")
}

func init() {
	proto.RegisterFile("pb/machine.proto", fileDescriptor_2daae29eac7cff8a)
}

var fileDescriptor_2daae29eac7cff8a = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4b, 0x4f, 0x1b, 0x3b,
	0x14, 0xd6, 0x24, 0x24, 0x21, 0x27, 0x93, 0x07, 0x16, 0xba, 0x77, 0x2e, 0xe8, 0xde, 0x9b, 0xce,
	0x02, 0x45, 0xaa, 0x14, 0x24, 0x8a, 0x2a, 0x84, 0x54, 0xa9, 0x2d, 0xb0, 0x48, 0xa1, 0x2f, 0x93,
	0x4d, 0xd5, 0xc5, 0x68, 0x32, 0x63, 0x12, 0x8b, 0x8c, 0xed, 0x7a, 0x1c, 0x68, 0xfa, 0xcb, 0xba,
	0xef, 0x1f, 0xab, 0xfc, 0x48, 0x99, 0x84, 0x90, 0x9d, 0xbf, 0xef, 0xbc, 0x8f, 0x3f, 0x1b, 0x3a,
	0x62, 0x74, 0x98, 0xc5, 0xc9, 0x84, 0x32, 0xd2, 0x17, 0x92, 0x2b, 0x8e, 0x4a, 0x62, 0xb4, 0xd7,
	0x16, 0xa3, 0xc3, 0x84, 0x67, 0x19, 0x67, 0x96, 0x0c, 0x7f, 0x96, 0xa0, 0x73, 0x26, 0x49, 0xac,
	0xc8, 0x7b, 0xeb, 0x8c, 0xc9, 0x37, 0xf4, 0x0f, 0x6c, 0xd3, 0x2c, 0x1e, 0x93, 0x88, 0xa6, 0x81,
	0xd7, 0xf5, 0x7a, 0x75, 0x5c, 0x33, 0x78, 0x90, 0x22, 0x04, 0x5b, 0x2c, 0xce, 0x48, 0x50, 0x32,
	0xb4, 0x39, 0xa3, 0x0e, 0x94, 0x09, 0xbb, 0x0b, 0xca, 0xdd, 0x72, 0xaf, 0x8e, 0xf5, 0x51, 0x33,
	0x4a, 0xcd, 0x83, 0xad, 0xae, 0xd7, 0xdb, 0xc6, 0xfa, 0xa8, 0x99, 0x24, 0x4b, 0x83, 0x8a, 0xf5,
	0x49, 0xb2, 0x14, 0xfd, 0x0f, 0x8d, 0x7b, 0x2e, 0x6f, 0x29, 0x1b, 0x47, 0x29, 0x95, 0x41, 0xd5,
	0x24, 0x04, 0x47, 0x9d, 0x53, 0x89, 0x2e, 0xa1, 0x49, 0xbe, 0x0b, 0x9e, 0x93, 0x34, 0x12, 0x5c,
	0xaa, 0x3c, 0xa8, 0x75, 0xcb, 0xbd, 0xc6, 0xd1, 0x41, 0x5f, 0x8c, 0xfa, 0xab, 0x2d, 0xf7, 0x2f,
	0xac, 0xe7, 0x27, 0xed, 0x78, 0xc1, 0x94, 0x9c, 0x63, 0x9f, 0x14, 0xa8, 0xbd, 0xcf, 0xb0, 0xf3,
	0xc8, 0x45, 0x37, 0x75, 0x4b, 0xe6, 0x6e, 0x44, 0x7d, 0x44, 0x07, 0x50, 0xb9, 0x8b, 0xa7, 0x33,
	0x3b, 0x5f, 0xe3, 0xa8, 0xa3, 0x6b, 0xe9, 0x80, 0xb7, 0x94, 0xa5, 0x94, 0x8d, 0x73, 0x6c, 0xcd,
	0xa7, 0xa5, 0x13, 0x2f, 0x3c, 0x83, 0x46, 0xc1, 0x84, 0xfe, 0x86, 0xda, 0x84, 0xe7, 0x2a, 0xa2,
	0xc2, 0x25, 0xac, 0x6a, 0x38, 0x10, 0x68, 0x1f, 0xea, 0xc6, 0xa0, 0x87, 0x70, 0x7b, 0xdb, 0xd6,
	0x84, 0x0e, 0x0e, 0xcf, 0xc1, 0x2f, 0xe6, 0x47, 0xc7, 0xd0, 0xd4, 0x7e, 0xd1, 0xc8, 0x11, 0x81,
	0x67, 0x86, 0x6e, 0xaf, 0x34, 0x82, 0x7d, 0x51, 0x88, 0x0a, 0x5f, 0xc3, 0xce, 0xca, 0x46, 0x72,
	0x81, 0x5a, 0x50, 0xfa, 0x73, 0x7f, 0x25, 0x9a, 0xa2, 0x7d, 0x28, 0x13, 0x29, 0xdd, 0x64, 0x75,
	0x9d, 0xf0, 0x42, 0x4a, 0x2e, 0xb1, 0x66, 0xc3, 0x10, 0x3a, 0xe7, 0x64, 0x4a, 0x96, 0x64, 0xb0,
	0x92, 0x20, 0x4c, 0xa1, 0x7d, 0x45, 0x73, 0x55, 0xac, 0x71, 0x0c, 0x4d, 0x27, 0xb2, 0x88, 0xb2,
	0x1b, 0xbe, 0xd4, 0xae, 0xf3, 0x1b, 0xb0, 0x1b, 0x8e, 0xfd, 0xec, 0x01, 0xe4, 0x9b, 0x3b, 0xf9,
	0xe5, 0x41, 0xa3, 0x10, 0xfa, 0x68, 0x8c, 0x75, 0x0a, 0xfc, 0x17, 0xc0, 0x0a, 0xd6, 0x58, 0xca,
	0xc6, 0x52, 0x37, 0xcc, 0x87, 0x25, 0xb3, 0x9a, 0x0b, 0x62, 0x54, 0xb9, 0x30, 0x0f, 0xe7, 0x82,
	0x68, 0x25, 0x26, 0x66, 0x7b, 0x91, 0xa2, 0x19, 0x09, 0x2a, 0x56, 0x89, 0x96, 0x1a, 0xd2, 0x8c,
	0xa0, 0xbf, 0xa0, 0x9a, 0xab, 0x58, 0xcd, 0x72, 0xa7, 0x52, 0x87, 0x96, 0xde, 0x49, 0x6d, 0xe9,
	0x9d, 0x84, 0xcf, 0xa0, 0x7d, 0xad, 0x62, 0xa9, 0x36, 0xac, 0xf3, 0x3f, 0xf0, 0x17, 0x73, 0xa6,
	0xeb, 0xec, 0x27, 0xd0, 0xba, 0xa4, 0xd3, 0xe9, 0xd3, 0x19, 0x4c, 0x5f, 0x74, 0xcc, 0xe2, 0xa9,
	0x5b, 0x86, 0x43, 0xe1, 0x29, 0xb4, 0xae, 0x15, 0x17, 0x1b, 0x22, 0x03, 0xa8, 0xe9, 0x59, 0xf9,
	0xcc, 0x2a, 0xb2, 0x82, 0x17, 0x30, 0x7c, 0x09, 0x1d, 0x4c, 0xf4, 0x1a, 0x37, 0x44, 0xaf, 0xb9,
	0x82, 0xf0, 0x15, 0xec, 0x60, 0x92, 0x6f, 0x1e, 0x79, 0x43, 0xd9, 0x53, 0x68, 0xbd, 0x51, 0x2a,
	0x4e, 0x26, 0xd7, 0x4a, 0x92, 0x38, 0x1b, 0xb0, 0x75, 0xb1, 0x09, 0x67, 0x8a, 0x30, 0x1b, 0xeb,
	0xe3, 0x05, 0x0c, 0x9f, 0x43, 0xbb, 0x18, 0xfb, 0x71, 0xa6, 0x8a, 0xce, 0xde, 0xb2, 0xf3, 0x15,
	0xf8, 0x98, 0xe4, 0xf4, 0x07, 0x19, 0x0e, 0xbf, 0x3c, 0xb1, 0xd3, 0x09, 0xa1, 0xe3, 0x89, 0xad,
	0xd2, 0xc4, 0x0e, 0xa1, 0x5d, 0xa8, 0xdc, 0xd3, 0x54, 0x4d, 0x8c, 0xba, 0x9a, 0xd8, 0x82, 0xf0,
	0x2b, 0xa0, 0xb3, 0x98, 0xd9, 0xea, 0xef, 0x66, 0xe9, 0xd8, 0xbe, 0x0a, 0xf7, 0xfd, 0x79, 0x0f,
	0xdf, 0xdf, 0x2e, 0x54, 0xb4, 0x66, 0x16, 0x2b, 0xb3, 0x60, 0x45, 0x97, 0xe5, 0x15, 0x5d, 0x8e,
	0xaa, 0xe6, 0x8b, 0x7e, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x76, 0xd1, 0xba, 0xcb, 0x05,
	0x00, 0x00,
}
