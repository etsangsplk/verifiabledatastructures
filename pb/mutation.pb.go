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
	Namespace []byte `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// one of the following
	LogAddEntry *LogAddEntryRequest `protobuf:"bytes,2,opt,name=log_add_entry,json=logAddEntry" json:"log_add_entry,omitempty"`
}

func (m *Mutation) Reset()                    { *m = Mutation{} }
func (m *Mutation) String() string            { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()               {}
func (*Mutation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Mutation) GetNamespace() []byte {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Mutation) GetLogAddEntry() *LogAddEntryRequest {
	if m != nil {
		return m.LogAddEntry
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

type MapNode struct {
	// for parent nodes only
	LeftNumber  int64 `protobuf:"varint,1,opt,name=left_number,json=leftNumber" json:"left_number,omitempty"`
	RightNumber int64 `protobuf:"varint,2,opt,name=right_number,json=rightNumber" json:"right_number,omitempty"`
	// set for parents only, if corresponding entry is non-zero
	LeftHash  []byte `protobuf:"bytes,3,opt,name=left_hash,json=leftHash,proto3" json:"left_hash,omitempty"`
	RightHash []byte `protobuf:"bytes,4,opt,name=right_hash,json=rightHash,proto3" json:"right_hash,omitempty"`
	// for leaf nodes only
	LeafHash []byte `protobuf:"bytes,6,opt,name=leaf_hash,json=leafHash,proto3" json:"leaf_hash,omitempty"`
	Path     []byte `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
}

func (m *MapNode) Reset()                    { *m = MapNode{} }
func (m *MapNode) String() string            { return proto.CompactTextString(m) }
func (*MapNode) ProtoMessage()               {}
func (*MapNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *MapNode) GetLeftNumber() int64 {
	if m != nil {
		return m.LeftNumber
	}
	return 0
}

func (m *MapNode) GetRightNumber() int64 {
	if m != nil {
		return m.RightNumber
	}
	return 0
}

func (m *MapNode) GetLeftHash() []byte {
	if m != nil {
		return m.LeftHash
	}
	return nil
}

func (m *MapNode) GetRightHash() []byte {
	if m != nil {
		return m.RightHash
	}
	return nil
}

func (m *MapNode) GetLeafHash() []byte {
	if m != nil {
		return m.LeafHash
	}
	return nil
}

func (m *MapNode) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*Mutation)(nil), "continusec.verifiabledatastructures.server.store.Mutation")
	proto.RegisterType((*LeafNode)(nil), "continusec.verifiabledatastructures.server.store.LeafNode")
	proto.RegisterType((*TreeNode)(nil), "continusec.verifiabledatastructures.server.store.TreeNode")
	proto.RegisterType((*LogTreeHash)(nil), "continusec.verifiabledatastructures.server.store.LogTreeHash")
	proto.RegisterType((*EntryIndex)(nil), "continusec.verifiabledatastructures.server.store.EntryIndex")
	proto.RegisterType((*MapNode)(nil), "continusec.verifiabledatastructures.server.store.MapNode")
}

func init() { proto.RegisterFile("mutation.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0xfb, 0x50,
	0x10, 0xc5, 0x49, 0xdb, 0x7f, 0x3f, 0x26, 0xfd, 0x8b, 0x04, 0x17, 0x41, 0x2b, 0xad, 0x59, 0x75,
	0x15, 0x44, 0x97, 0xae, 0x14, 0x04, 0x85, 0xb6, 0x8b, 0xe0, 0xca, 0x4d, 0x98, 0x24, 0x93, 0x0f,
	0x48, 0x72, 0xe3, 0xbd, 0x37, 0x45, 0x1f, 0xc1, 0x57, 0xf2, 0xe9, 0xe4, 0x4e, 0xa2, 0xdd, 0x28,
	0xb8, 0x9b, 0xfc, 0xce, 0x39, 0x73, 0x86, 0x5c, 0x38, 0xaa, 0x5a, 0x8d, 0xba, 0x10, 0xb5, 0xdf,
	0x48, 0xa1, 0x85, 0x73, 0x19, 0x8b, 0x5a, 0x17, 0x75, 0xab, 0x28, 0xf6, 0xf7, 0x24, 0x8b, 0xb4,
	0xc0, 0xa8, 0xa4, 0x04, 0x35, 0x2a, 0x2d, 0xdb, 0x58, 0xb7, 0x92, 0x94, 0xaf, 0x48, 0xee, 0x49,
	0xfa, 0x4a, 0x0b, 0x49, 0xa7, 0x33, 0x6c, 0x8a, 0x2e, 0xec, 0xbd, 0x5b, 0x30, 0xdd, 0xf6, 0xfb,
	0x9c, 0x05, 0xcc, 0x6a, 0xac, 0x48, 0x35, 0x18, 0x93, 0x6b, 0xad, 0xac, 0xf5, 0x3c, 0x38, 0x00,
	0x27, 0x84, 0xff, 0xa5, 0xc8, 0x42, 0x4c, 0x92, 0x90, 0x6a, 0x2d, 0xdf, 0xdc, 0xc1, 0xca, 0x5a,
	0xdb, 0x57, 0x37, 0xfe, 0x5f, 0xfa, 0x4d, 0xe3, 0x46, 0x64, 0xb7, 0x49, 0x72, 0x6f, 0xb2, 0x01,
	0xbd, 0xb4, 0xa4, 0x74, 0x60, 0x97, 0x07, 0xe6, 0x2d, 0x60, 0xba, 0x21, 0x4c, 0x77, 0x22, 0x21,
	0xe7, 0x18, 0x86, 0x95, 0xce, 0xfb, 0x23, 0xcc, 0x68, 0xd4, 0x27, 0x49, 0xf4, 0x8b, 0xba, 0x04,
	0x7b, 0x23, 0x32, 0x63, 0x78, 0x40, 0x95, 0xff, 0x60, 0xf0, 0x00, 0xb8, 0xe5, 0xb1, 0x4e, 0xe8,
	0xd5, 0x39, 0x81, 0x7f, 0x85, 0x19, 0xd8, 0x31, 0x0c, 0xba, 0x0f, 0xef, 0xc3, 0x82, 0xc9, 0x16,
	0x1b, 0xae, 0x58, 0x82, 0x5d, 0x52, 0xaa, 0xc3, 0xba, 0xad, 0x22, 0x92, 0xbd, 0x0f, 0x0c, 0xda,
	0x31, 0x71, 0x2e, 0x60, 0x2e, 0x8b, 0x2c, 0xff, 0x76, 0x0c, 0xd8, 0x61, 0x33, 0xeb, 0x2d, 0x67,
	0x30, 0xe3, 0x1d, 0x39, 0xaa, 0xdc, 0x1d, 0xf2, 0x2d, 0x53, 0x03, 0xf8, 0xc4, 0x73, 0x80, 0x2e,
	0xcf, 0xea, 0xa8, 0xfb, 0xdb, 0x4c, 0x58, 0xe6, 0x2c, 0xa6, 0x9d, 0x3a, 0xfe, 0xca, 0x62, 0xca,
	0xa2, 0x03, 0xa3, 0x06, 0x75, 0xee, 0x4e, 0x98, 0xf3, 0x7c, 0x37, 0x7a, 0x1e, 0x34, 0x51, 0x34,
	0xe6, 0x67, 0xbd, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x55, 0x17, 0x31, 0x9a, 0x25, 0x02, 0x00,
	0x00,
}
