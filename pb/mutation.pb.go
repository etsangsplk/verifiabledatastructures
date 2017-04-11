// Code generated by protoc-gen-go.
// source: mutation.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Mutation struct {
	// one of the following
	LogCreate   *LogCreateRequest   `protobuf:"bytes,1,opt,name=log_create,json=logCreate" json:"log_create,omitempty"`
	MapCreate   *MapCreateRequest   `protobuf:"bytes,2,opt,name=map_create,json=mapCreate" json:"map_create,omitempty"`
	LogDelete   *LogDeleteRequest   `protobuf:"bytes,3,opt,name=log_delete,json=logDelete" json:"log_delete,omitempty"`
	MapDelete   *MapDeleteRequest   `protobuf:"bytes,4,opt,name=map_delete,json=mapDelete" json:"map_delete,omitempty"`
	LogAddEntry *LogAddEntryRequest `protobuf:"bytes,5,opt,name=log_add_entry,json=logAddEntry" json:"log_add_entry,omitempty"`
	MapSetValue *MapSetValueRequest `protobuf:"bytes,6,opt,name=map_set_value,json=mapSetValue" json:"map_set_value,omitempty"`
}

func (m *Mutation) Reset()                    { *m = Mutation{} }
func (m *Mutation) String() string            { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()               {}
func (*Mutation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Mutation) GetLogCreate() *LogCreateRequest {
	if m != nil {
		return m.LogCreate
	}
	return nil
}

func (m *Mutation) GetMapCreate() *MapCreateRequest {
	if m != nil {
		return m.MapCreate
	}
	return nil
}

func (m *Mutation) GetLogDelete() *LogDeleteRequest {
	if m != nil {
		return m.LogDelete
	}
	return nil
}

func (m *Mutation) GetMapDelete() *MapDeleteRequest {
	if m != nil {
		return m.MapDelete
	}
	return nil
}

func (m *Mutation) GetLogAddEntry() *LogAddEntryRequest {
	if m != nil {
		return m.LogAddEntry
	}
	return nil
}

func (m *Mutation) GetMapSetValue() *MapSetValueRequest {
	if m != nil {
		return m.MapSetValue
	}
	return nil
}

type LeafNode struct {
	Mth []byte `protobuf:"bytes,1,opt,name=mth,proto3" json:"mth,omitempty"`
}

func (m *LeafNode) Reset()                    { *m = LeafNode{} }
func (m *LeafNode) String() string            { return proto.CompactTextString(m) }
func (*LeafNode) ProtoMessage()               {}
func (*LeafNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *LeafNode) GetMth() []byte {
	if m != nil {
		return m.Mth
	}
	return nil
}

type TreeNode struct {
	Mth []byte `protobuf:"bytes,1,opt,name=mth,proto3" json:"mth,omitempty"`
}

func (m *TreeNode) Reset()                    { *m = TreeNode{} }
func (m *TreeNode) String() string            { return proto.CompactTextString(m) }
func (*TreeNode) ProtoMessage()               {}
func (*TreeNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *TreeNode) GetMth() []byte {
	if m != nil {
		return m.Mth
	}
	return nil
}

type LogTreeHash struct {
	Mth []byte `protobuf:"bytes,1,opt,name=mth,proto3" json:"mth,omitempty"`
}

func (m *LogTreeHash) Reset()                    { *m = LogTreeHash{} }
func (m *LogTreeHash) String() string            { return proto.CompactTextString(m) }
func (*LogTreeHash) ProtoMessage()               {}
func (*LogTreeHash) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *LogTreeHash) GetMth() []byte {
	if m != nil {
		return m.Mth
	}
	return nil
}

type EntryIndex struct {
	Index int64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
}

func (m *EntryIndex) Reset()                    { *m = EntryIndex{} }
func (m *EntryIndex) String() string            { return proto.CompactTextString(m) }
func (*EntryIndex) ProtoMessage()               {}
func (*EntryIndex) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *EntryIndex) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*Mutation)(nil), "continusec.verifiabledatastructures.server.store.Mutation")
	proto.RegisterType((*LeafNode)(nil), "continusec.verifiabledatastructures.server.store.LeafNode")
	proto.RegisterType((*TreeNode)(nil), "continusec.verifiabledatastructures.server.store.TreeNode")
	proto.RegisterType((*LogTreeHash)(nil), "continusec.verifiabledatastructures.server.store.LogTreeHash")
	proto.RegisterType((*EntryIndex)(nil), "continusec.verifiabledatastructures.server.store.EntryIndex")
}

func init() { proto.RegisterFile("mutation.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0xe9, 0x27, 0x6d, 0xaa, 0x22, 0x8b, 0x87, 0x45, 0x04, 0xa5, 0x27, 0x4f, 0x8b, 0xe8,
	0x49, 0x3c, 0xf9, 0x05, 0x0a, 0xad, 0x87, 0x55, 0x44, 0xbc, 0x2c, 0xd3, 0xcd, 0xb4, 0x5d, 0xd8,
	0x6c, 0x62, 0x32, 0x29, 0xfa, 0xbf, 0xfc, 0x81, 0x92, 0xb4, 0xb1, 0x52, 0x14, 0xdc, 0x5b, 0xf6,
	0xcd, 0xbc, 0xcf, 0x3c, 0x0b, 0x61, 0x3b, 0xc2, 0x12, 0x50, 0x21, 0xab, 0x44, 0x69, 0x49, 0x32,
	0x3a, 0xc9, 0x65, 0x45, 0x45, 0x65, 0x0d, 0xe6, 0xc9, 0x02, 0x75, 0x31, 0x2d, 0x60, 0x52, 0x22,
	0x07, 0x02, 0x43, 0xda, 0xe6, 0x64, 0x35, 0x9a, 0xc4, 0xa0, 0x5e, 0xa0, 0x4e, 0x0c, 0x49, 0x8d,
	0xfb, 0x7d, 0x50, 0xc5, 0xb2, 0x3c, 0xfc, 0x6c, 0xb3, 0xde, 0x78, 0xc5, 0x8b, 0x5e, 0x18, 0x2b,
	0xe5, 0x2c, 0xcb, 0x35, 0x02, 0x61, 0xdc, 0x38, 0x6a, 0x1c, 0x0f, 0x4e, 0xcf, 0x93, 0xff, 0xe0,
	0x1d, 0x70, 0x24, 0x67, 0xd7, 0xbe, 0x99, 0xe2, 0x9b, 0x45, 0x43, 0x69, 0xbf, 0x0c, 0x89, 0x23,
	0x0b, 0x50, 0x81, 0xdc, 0xac, 0x49, 0x1e, 0x83, 0xda, 0x20, 0x8b, 0x90, 0x04, 0x67, 0x8e, 0x25,
	0x12, 0xc6, 0xad, 0xfa, 0xce, 0x37, 0xbe, 0xf9, 0xd3, 0x79, 0x99, 0x04, 0xe7, 0x15, 0xb9, 0x5d,
	0xdf, 0x79, 0x83, 0x2c, 0x42, 0x12, 0x65, 0x6c, 0xdb, 0x39, 0x03, 0xe7, 0x19, 0x56, 0xa4, 0x3f,
	0xe2, 0x8e, 0x87, 0x5f, 0xd4, 0xd1, 0xbe, 0xe4, 0xfc, 0xd6, 0x75, 0x03, 0x7e, 0x50, 0xae, 0x33,
	0xb7, 0xc0, 0xa9, 0x1b, 0xa4, 0x6c, 0x01, 0xa5, 0xc5, 0xb8, 0x5b, 0x73, 0xc1, 0x18, 0xd4, 0x23,
	0xd2, 0xb3, 0xeb, 0x7e, 0x2f, 0x10, 0xeb, 0x6c, 0x78, 0xc0, 0x7a, 0x23, 0x84, 0xe9, 0x83, 0xe4,
	0x18, 0xed, 0xb2, 0x96, 0xa0, 0xb9, 0x7f, 0x2e, 0x5b, 0xa9, 0x3b, 0xba, 0xdb, 0x27, 0x8d, 0xf8,
	0xc7, 0xed, 0x21, 0x1b, 0x8c, 0xe4, 0xcc, 0x0d, 0xdc, 0x81, 0x99, 0xff, 0x32, 0x30, 0x64, 0xcc,
	0xff, 0xc6, 0x7d, 0xc5, 0xf1, 0x3d, 0xda, 0x63, 0x9d, 0xc2, 0x1d, 0xfc, 0x44, 0x2b, 0x5d, 0x7e,
	0x5c, 0xb5, 0x5f, 0x9b, 0x6a, 0x32, 0xe9, 0xfa, 0x47, 0x7c, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff,
	0xdd, 0x63, 0x3c, 0x3b, 0x13, 0x03, 0x00, 0x00,
}
