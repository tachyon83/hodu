// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.6
// source: hodu.proto

package hodu

import (
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

type ROUTE_PROTO int32

const (
	ROUTE_PROTO_TCP  ROUTE_PROTO = 0
	ROUTE_PROTO_TCP4 ROUTE_PROTO = 1
	ROUTE_PROTO_TCP6 ROUTE_PROTO = 2
)

// Enum value maps for ROUTE_PROTO.
var (
	ROUTE_PROTO_name = map[int32]string{
		0: "TCP",
		1: "TCP4",
		2: "TCP6",
	}
	ROUTE_PROTO_value = map[string]int32{
		"TCP":  0,
		"TCP4": 1,
		"TCP6": 2,
	}
)

func (x ROUTE_PROTO) Enum() *ROUTE_PROTO {
	p := new(ROUTE_PROTO)
	*p = x
	return p
}

func (x ROUTE_PROTO) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ROUTE_PROTO) Descriptor() protoreflect.EnumDescriptor {
	return file_hodu_proto_enumTypes[0].Descriptor()
}

func (ROUTE_PROTO) Type() protoreflect.EnumType {
	return &file_hodu_proto_enumTypes[0]
}

func (x ROUTE_PROTO) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ROUTE_PROTO.Descriptor instead.
func (ROUTE_PROTO) EnumDescriptor() ([]byte, []int) {
	return file_hodu_proto_rawDescGZIP(), []int{0}
}

type PACKET_KIND int32

const (
	PACKET_KIND_ERROR         PACKET_KIND = 0 // generic error response
	PACKET_KIND_OK            PACKET_KIND = 1 // generic success response
	PACKET_KIND_ROUTE_START   PACKET_KIND = 2
	PACKET_KIND_ROUTE_STOP    PACKET_KIND = 3
	PACKET_KIND_ROUTE_STARTED PACKET_KIND = 4
	PACKET_KIND_ROUTE_STOPPED PACKET_KIND = 5
	PACKET_KIND_PEER_STARTED  PACKET_KIND = 6
	PACKET_KIND_PEER_STOPPED  PACKET_KIND = 7
	PACKET_KIND_PEER_ABORTED  PACKET_KIND = 8
	PACKET_KIND_PEER_EOF      PACKET_KIND = 9
	PACKET_KIND_PEER_DATA     PACKET_KIND = 10
)

// Enum value maps for PACKET_KIND.
var (
	PACKET_KIND_name = map[int32]string{
		0:  "ERROR",
		1:  "OK",
		2:  "ROUTE_START",
		3:  "ROUTE_STOP",
		4:  "ROUTE_STARTED",
		5:  "ROUTE_STOPPED",
		6:  "PEER_STARTED",
		7:  "PEER_STOPPED",
		8:  "PEER_ABORTED",
		9:  "PEER_EOF",
		10: "PEER_DATA",
	}
	PACKET_KIND_value = map[string]int32{
		"ERROR":         0,
		"OK":            1,
		"ROUTE_START":   2,
		"ROUTE_STOP":    3,
		"ROUTE_STARTED": 4,
		"ROUTE_STOPPED": 5,
		"PEER_STARTED":  6,
		"PEER_STOPPED":  7,
		"PEER_ABORTED":  8,
		"PEER_EOF":      9,
		"PEER_DATA":     10,
	}
)

func (x PACKET_KIND) Enum() *PACKET_KIND {
	p := new(PACKET_KIND)
	*p = x
	return p
}

func (x PACKET_KIND) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PACKET_KIND) Descriptor() protoreflect.EnumDescriptor {
	return file_hodu_proto_enumTypes[1].Descriptor()
}

func (PACKET_KIND) Type() protoreflect.EnumType {
	return &file_hodu_proto_enumTypes[1]
}

func (x PACKET_KIND) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PACKET_KIND.Descriptor instead.
func (PACKET_KIND) EnumDescriptor() ([]byte, []int) {
	return file_hodu_proto_rawDescGZIP(), []int{1}
}

type Seed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32 `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Flags   uint64 `protobuf:"varint,2,opt,name=Flags,proto3" json:"Flags,omitempty"`
}

func (x *Seed) Reset() {
	*x = Seed{}
	mi := &file_hodu_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Seed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seed) ProtoMessage() {}

func (x *Seed) ProtoReflect() protoreflect.Message {
	mi := &file_hodu_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seed.ProtoReflect.Descriptor instead.
func (*Seed) Descriptor() ([]byte, []int) {
	return file_hodu_proto_rawDescGZIP(), []int{0}
}

func (x *Seed) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Seed) GetFlags() uint64 {
	if x != nil {
		return x.Flags
	}
	return 0
}

type RouteDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId uint32      `protobuf:"varint,1,opt,name=RouteId,proto3" json:"RouteId,omitempty"`
	Proto   ROUTE_PROTO `protobuf:"varint,2,opt,name=Proto,proto3,enum=ROUTE_PROTO" json:"Proto,omitempty"`
	AddrStr string      `protobuf:"bytes,3,opt,name=AddrStr,proto3" json:"AddrStr,omitempty"`
}

func (x *RouteDesc) Reset() {
	*x = RouteDesc{}
	mi := &file_hodu_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteDesc) ProtoMessage() {}

func (x *RouteDesc) ProtoReflect() protoreflect.Message {
	mi := &file_hodu_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteDesc.ProtoReflect.Descriptor instead.
func (*RouteDesc) Descriptor() ([]byte, []int) {
	return file_hodu_proto_rawDescGZIP(), []int{1}
}

func (x *RouteDesc) GetRouteId() uint32 {
	if x != nil {
		return x.RouteId
	}
	return 0
}

func (x *RouteDesc) GetProto() ROUTE_PROTO {
	if x != nil {
		return x.Proto
	}
	return ROUTE_PROTO_TCP
}

func (x *RouteDesc) GetAddrStr() string {
	if x != nil {
		return x.AddrStr
	}
	return ""
}

type PeerDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId       uint32 `protobuf:"varint,1,opt,name=RouteId,proto3" json:"RouteId,omitempty"`
	PeerId        uint32 `protobuf:"varint,2,opt,name=PeerId,proto3" json:"PeerId,omitempty"`
	RemoteAddrStr string `protobuf:"bytes,3,opt,name=RemoteAddrStr,proto3" json:"RemoteAddrStr,omitempty"`
	LocalAddrStr  string `protobuf:"bytes,4,opt,name=LocalAddrStr,proto3" json:"LocalAddrStr,omitempty"`
}

func (x *PeerDesc) Reset() {
	*x = PeerDesc{}
	mi := &file_hodu_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PeerDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerDesc) ProtoMessage() {}

func (x *PeerDesc) ProtoReflect() protoreflect.Message {
	mi := &file_hodu_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerDesc.ProtoReflect.Descriptor instead.
func (*PeerDesc) Descriptor() ([]byte, []int) {
	return file_hodu_proto_rawDescGZIP(), []int{2}
}

func (x *PeerDesc) GetRouteId() uint32 {
	if x != nil {
		return x.RouteId
	}
	return 0
}

func (x *PeerDesc) GetPeerId() uint32 {
	if x != nil {
		return x.PeerId
	}
	return 0
}

func (x *PeerDesc) GetRemoteAddrStr() string {
	if x != nil {
		return x.RemoteAddrStr
	}
	return ""
}

func (x *PeerDesc) GetLocalAddrStr() string {
	if x != nil {
		return x.LocalAddrStr
	}
	return ""
}

type PeerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId uint32 `protobuf:"varint,1,opt,name=RouteId,proto3" json:"RouteId,omitempty"`
	PeerId  uint32 `protobuf:"varint,2,opt,name=PeerId,proto3" json:"PeerId,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *PeerData) Reset() {
	*x = PeerData{}
	mi := &file_hodu_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PeerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerData) ProtoMessage() {}

func (x *PeerData) ProtoReflect() protoreflect.Message {
	mi := &file_hodu_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerData.ProtoReflect.Descriptor instead.
func (*PeerData) Descriptor() ([]byte, []int) {
	return file_hodu_proto_rawDescGZIP(), []int{3}
}

func (x *PeerData) GetRouteId() uint32 {
	if x != nil {
		return x.RouteId
	}
	return 0
}

func (x *PeerData) GetPeerId() uint32 {
	if x != nil {
		return x.PeerId
	}
	return 0
}

func (x *PeerData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind PACKET_KIND `protobuf:"varint,1,opt,name=Kind,proto3,enum=PACKET_KIND" json:"Kind,omitempty"`
	// Types that are assignable to U:
	//
	//	*Packet_Route
	//	*Packet_Peer
	//	*Packet_Data
	U isPacket_U `protobuf_oneof:"U"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	mi := &file_hodu_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_hodu_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_hodu_proto_rawDescGZIP(), []int{4}
}

func (x *Packet) GetKind() PACKET_KIND {
	if x != nil {
		return x.Kind
	}
	return PACKET_KIND_ERROR
}

func (m *Packet) GetU() isPacket_U {
	if m != nil {
		return m.U
	}
	return nil
}

func (x *Packet) GetRoute() *RouteDesc {
	if x, ok := x.GetU().(*Packet_Route); ok {
		return x.Route
	}
	return nil
}

func (x *Packet) GetPeer() *PeerDesc {
	if x, ok := x.GetU().(*Packet_Peer); ok {
		return x.Peer
	}
	return nil
}

func (x *Packet) GetData() *PeerData {
	if x, ok := x.GetU().(*Packet_Data); ok {
		return x.Data
	}
	return nil
}

type isPacket_U interface {
	isPacket_U()
}

type Packet_Route struct {
	Route *RouteDesc `protobuf:"bytes,2,opt,name=Route,proto3,oneof"`
}

type Packet_Peer struct {
	Peer *PeerDesc `protobuf:"bytes,3,opt,name=Peer,proto3,oneof"`
}

type Packet_Data struct {
	Data *PeerData `protobuf:"bytes,4,opt,name=Data,proto3,oneof"`
}

func (*Packet_Route) isPacket_U() {}

func (*Packet_Peer) isPacket_U() {}

func (*Packet_Data) isPacket_U() {}

var File_hodu_proto protoreflect.FileDescriptor

var file_hodu_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x6f, 0x64, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x04,
	0x53, 0x65, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x46,
	0x6c, 0x61, 0x67, 0x73, 0x22, 0x63, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x52, 0x4f, 0x55,
	0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x50, 0x65,
	0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x72, 0x12, 0x22,
	0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x53,
	0x74, 0x72, 0x22, 0x50, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x65, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x95, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x20, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x52, 0x04, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x22, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x73, 0x63, 0x48, 0x00, 0x52, 0x05,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x48, 0x00,
	0x52, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x42, 0x03, 0x0a, 0x01, 0x55, 0x2a, 0x2a, 0x0a, 0x0b,
	0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x12, 0x07, 0x0a, 0x03, 0x54,
	0x43, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x43, 0x50, 0x34, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x54, 0x43, 0x50, 0x36, 0x10, 0x02, 0x2a, 0xba, 0x01, 0x0a, 0x0b, 0x50, 0x41, 0x43,
	0x4b, 0x45, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52,
	0x4f, 0x55, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a,
	0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d,
	0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44,
	0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x4f,
	0x50, 0x50, 0x45, 0x44, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x41,
	0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x45, 0x45, 0x52,
	0x5f, 0x45, 0x4f, 0x46, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x10, 0x0a, 0x32, 0x49, 0x0a, 0x04, 0x48, 0x6f, 0x64, 0x75, 0x12, 0x19, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x53, 0x65, 0x65, 0x64, 0x12, 0x05, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x1a,
	0x05, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x07, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x1a, 0x07, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x68, 0x6f, 0x64, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_hodu_proto_rawDescOnce sync.Once
	file_hodu_proto_rawDescData = file_hodu_proto_rawDesc
)

func file_hodu_proto_rawDescGZIP() []byte {
	file_hodu_proto_rawDescOnce.Do(func() {
		file_hodu_proto_rawDescData = protoimpl.X.CompressGZIP(file_hodu_proto_rawDescData)
	})
	return file_hodu_proto_rawDescData
}

var file_hodu_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_hodu_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hodu_proto_goTypes = []any{
	(ROUTE_PROTO)(0),  // 0: ROUTE_PROTO
	(PACKET_KIND)(0),  // 1: PACKET_KIND
	(*Seed)(nil),      // 2: Seed
	(*RouteDesc)(nil), // 3: RouteDesc
	(*PeerDesc)(nil),  // 4: PeerDesc
	(*PeerData)(nil),  // 5: PeerData
	(*Packet)(nil),    // 6: Packet
}
var file_hodu_proto_depIdxs = []int32{
	0, // 0: RouteDesc.Proto:type_name -> ROUTE_PROTO
	1, // 1: Packet.Kind:type_name -> PACKET_KIND
	3, // 2: Packet.Route:type_name -> RouteDesc
	4, // 3: Packet.Peer:type_name -> PeerDesc
	5, // 4: Packet.Data:type_name -> PeerData
	2, // 5: Hodu.GetSeed:input_type -> Seed
	6, // 6: Hodu.PacketStream:input_type -> Packet
	2, // 7: Hodu.GetSeed:output_type -> Seed
	6, // 8: Hodu.PacketStream:output_type -> Packet
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_hodu_proto_init() }
func file_hodu_proto_init() {
	if File_hodu_proto != nil {
		return
	}
	file_hodu_proto_msgTypes[4].OneofWrappers = []any{
		(*Packet_Route)(nil),
		(*Packet_Peer)(nil),
		(*Packet_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hodu_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hodu_proto_goTypes,
		DependencyIndexes: file_hodu_proto_depIdxs,
		EnumInfos:         file_hodu_proto_enumTypes,
		MessageInfos:      file_hodu_proto_msgTypes,
	}.Build()
	File_hodu_proto = out.File
	file_hodu_proto_rawDesc = nil
	file_hodu_proto_goTypes = nil
	file_hodu_proto_depIdxs = nil
}
