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
	LeftNumber  int64  `protobuf:"varint,1,opt,name=left_number,json=leftNumber" json:"left_number,omitempty"`
	RightNumber int64  `protobuf:"varint,2,opt,name=right_number,json=rightNumber" json:"right_number,omitempty"`
	LeftHash    []byte `protobuf:"bytes,3,opt,name=left_hash,json=leftHash,proto3" json:"left_hash,omitempty"`
	RightHash   []byte `protobuf:"bytes,4,opt,name=right_hash,json=rightHash,proto3" json:"right_hash,omitempty"`
	// for leaf nodes only
	LeafHash      []byte `protobuf:"bytes,6,opt,name=leaf_hash,json=leafHash,proto3" json:"leaf_hash,omitempty"`
	RemainingPath []byte `protobuf:"bytes,7,opt,name=remaining_path,json=remainingPath,proto3" json:"remaining_path,omitempty"`
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

func (m *MapNode) GetRemainingPath() []byte {
	if m != nil {
		return m.RemainingPath
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
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0xcb, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x5b, 0xfb, 0x67, 0xd2, 0x16, 0x09, 0x1e, 0x82, 0x56, 0x5a, 0x03, 0x42, 0x4f,
	0x41, 0xf4, 0xe8, 0x49, 0x41, 0x50, 0x68, 0x8b, 0x04, 0x4f, 0x5e, 0xc2, 0x24, 0x99, 0x24, 0x0b,
	0xc9, 0x6e, 0xdc, 0xdd, 0x14, 0xfd, 0x08, 0x7e, 0x38, 0xbf, 0x93, 0xec, 0x24, 0xb6, 0x17, 0x5f,
	0x78, 0x6f, 0x9b, 0xdf, 0xf3, 0x7b, 0x76, 0x86, 0x2c, 0x6c, 0xdb, 0xde, 0xa2, 0x15, 0x4a, 0xc6,
	0x9d, 0x56, 0x56, 0x05, 0x6f, 0x72, 0x25, 0xad, 0x90, 0xbd, 0xa1, 0x3c, 0xbe, 0x92, 0x16, 0xa5,
	0xc0, 0xac, 0xa1, 0x02, 0x2d, 0x1a, 0xab, 0xfb, 0xdc, 0xf6, 0x9a, 0x4c, 0x6c, 0x48, 0x5f, 0x49,
	0xc7, 0xc6, 0x2a, 0x4d, 0xcf, 0x57, 0xd8, 0x89, 0xa1, 0x1c, 0xfd, 0xf6, 0x60, 0x79, 0x1e, 0xef,
	0x0b, 0x76, 0xb0, 0x92, 0xd8, 0x92, 0xe9, 0x30, 0xa7, 0xd0, 0x3b, 0x78, 0xc7, 0x75, 0x72, 0x07,
	0x41, 0x0a, 0x9b, 0x46, 0x55, 0x29, 0x16, 0x45, 0x4a, 0xd2, 0xea, 0x5f, 0xe1, 0xe4, 0xe0, 0x1d,
	0xfd, 0xb7, 0xef, 0xe3, 0xc7, 0xcc, 0x77, 0x13, 0x4f, 0xaa, 0xfa, 0x50, 0x14, 0x9f, 0x5c, 0x37,
	0xa1, 0x1f, 0x3d, 0x19, 0x9b, 0xf8, 0xcd, 0x9d, 0x45, 0x3b, 0x58, 0x9e, 0x08, 0xcb, 0x8b, 0x2a,
	0x28, 0x78, 0x0a, 0xd3, 0xd6, 0xd6, 0xe3, 0x12, 0xee, 0xe8, 0xd2, 0x6f, 0x9a, 0xe8, 0x81, 0x74,
	0x0f, 0xfe, 0x49, 0x55, 0x4e, 0xf8, 0x8c, 0xa6, 0xfe, 0x8f, 0x10, 0x01, 0xf0, 0x94, 0x2f, 0xb2,
	0xa0, 0x9f, 0xc1, 0x33, 0x78, 0x22, 0xdc, 0x81, 0x8d, 0x69, 0x32, 0x7c, 0x44, 0x7f, 0x3c, 0x58,
	0x9c, 0xb1, 0xe3, 0x11, 0x7b, 0xf0, 0x1b, 0x2a, 0x6d, 0x2a, 0xfb, 0x36, 0x23, 0x3d, 0x7a, 0xe0,
	0xd0, 0x85, 0x49, 0xf0, 0x0a, 0xd6, 0x5a, 0x54, 0xf5, 0xcd, 0x98, 0xb0, 0xe1, 0x33, 0x1b, 0x95,
	0x17, 0xb0, 0xe2, 0x3b, 0x6a, 0x34, 0x75, 0x38, 0xe5, 0x5d, 0x96, 0x0e, 0xf0, 0x8a, 0x2f, 0x01,
	0x86, 0x3e, 0xa7, 0xb3, 0xe1, 0x6f, 0x33, 0xe1, 0x98, 0xbb, 0x58, 0x0e, 0xe9, 0xfc, 0x5f, 0x17,
	0x4b, 0x0e, 0x5f, 0xc3, 0x56, 0x53, 0x8b, 0x42, 0x0a, 0x59, 0xa5, 0x1d, 0xda, 0x3a, 0x5c, 0xb0,
	0xb1, 0xb9, 0xd1, 0xaf, 0x68, 0xeb, 0x8f, 0xb3, 0xef, 0x93, 0x2e, 0xcb, 0xe6, 0xfc, 0xd2, 0xef,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x80, 0x49, 0xd6, 0x38, 0x02, 0x00, 0x00,
}
