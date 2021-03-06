// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/proto3/block.proto

package proto3

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PartSetHeader struct {
	Total                int32    `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartSetHeader) Reset()         { *m = PartSetHeader{} }
func (m *PartSetHeader) String() string { return proto.CompactTextString(m) }
func (*PartSetHeader) ProtoMessage()    {}
func (*PartSetHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_760f4d5ceb2a11f0, []int{0}
}
func (m *PartSetHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartSetHeader.Unmarshal(m, b)
}
func (m *PartSetHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartSetHeader.Marshal(b, m, deterministic)
}
func (m *PartSetHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartSetHeader.Merge(m, src)
}
func (m *PartSetHeader) XXX_Size() int {
	return xxx_messageInfo_PartSetHeader.Size(m)
}
func (m *PartSetHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_PartSetHeader.DiscardUnknown(m)
}

var xxx_messageInfo_PartSetHeader proto.InternalMessageInfo

func (m *PartSetHeader) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PartSetHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type BlockID struct {
	Hash                 []byte         `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	PartsHeader          *PartSetHeader `protobuf:"bytes,2,opt,name=PartsHeader,proto3" json:"PartsHeader,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BlockID) Reset()         { *m = BlockID{} }
func (m *BlockID) String() string { return proto.CompactTextString(m) }
func (*BlockID) ProtoMessage()    {}
func (*BlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_760f4d5ceb2a11f0, []int{1}
}
func (m *BlockID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockID.Unmarshal(m, b)
}
func (m *BlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockID.Marshal(b, m, deterministic)
}
func (m *BlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockID.Merge(m, src)
}
func (m *BlockID) XXX_Size() int {
	return xxx_messageInfo_BlockID.Size(m)
}
func (m *BlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockID.DiscardUnknown(m)
}

var xxx_messageInfo_BlockID proto.InternalMessageInfo

func (m *BlockID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockID) GetPartsHeader() *PartSetHeader {
	if m != nil {
		return m.PartsHeader
	}
	return nil
}

type Header struct {
	// basic block info
	Version *Version   `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	ChainID string     `protobuf:"bytes,2,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Height  int64      `protobuf:"varint,3,opt,name=Height,proto3" json:"Height,omitempty"`
	Time    *Timestamp `protobuf:"bytes,4,opt,name=Time,proto3" json:"Time,omitempty"`
	// prev block info
	LastBlockID *BlockID `protobuf:"bytes,5,opt,name=LastBlockID,proto3" json:"LastBlockID,omitempty"`
	// hashes of block data
	LastCommitHash []byte `protobuf:"bytes,6,opt,name=LastCommitHash,proto3" json:"LastCommitHash,omitempty"`
	DataHash       []byte `protobuf:"bytes,7,opt,name=DataHash,proto3" json:"DataHash,omitempty"`
	// hashes from the app output from the prev block
	ValidatorsHash     []byte `protobuf:"bytes,8,opt,name=ValidatorsHash,proto3" json:"ValidatorsHash,omitempty"`
	NextValidatorsHash []byte `protobuf:"bytes,9,opt,name=NextValidatorsHash,proto3" json:"NextValidatorsHash,omitempty"`
	ConsensusHash      []byte `protobuf:"bytes,10,opt,name=ConsensusHash,proto3" json:"ConsensusHash,omitempty"`
	AppHash            []byte `protobuf:"bytes,11,opt,name=AppHash,proto3" json:"AppHash,omitempty"`
	LastResultsHash    []byte `protobuf:"bytes,12,opt,name=LastResultsHash,proto3" json:"LastResultsHash,omitempty"`
	// consensus info
	EvidenceHash         []byte   `protobuf:"bytes,13,opt,name=EvidenceHash,proto3" json:"EvidenceHash,omitempty"`
	ProposerAddress      []byte   `protobuf:"bytes,14,opt,name=ProposerAddress,proto3" json:"ProposerAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_760f4d5ceb2a11f0, []int{2}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Header) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetTime() *Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Header) GetLastBlockID() *BlockID {
	if m != nil {
		return m.LastBlockID
	}
	return nil
}

func (m *Header) GetLastCommitHash() []byte {
	if m != nil {
		return m.LastCommitHash
	}
	return nil
}

func (m *Header) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

func (m *Header) GetValidatorsHash() []byte {
	if m != nil {
		return m.ValidatorsHash
	}
	return nil
}

func (m *Header) GetNextValidatorsHash() []byte {
	if m != nil {
		return m.NextValidatorsHash
	}
	return nil
}

func (m *Header) GetConsensusHash() []byte {
	if m != nil {
		return m.ConsensusHash
	}
	return nil
}

func (m *Header) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *Header) GetLastResultsHash() []byte {
	if m != nil {
		return m.LastResultsHash
	}
	return nil
}

func (m *Header) GetEvidenceHash() []byte {
	if m != nil {
		return m.EvidenceHash
	}
	return nil
}

func (m *Header) GetProposerAddress() []byte {
	if m != nil {
		return m.ProposerAddress
	}
	return nil
}

type Version struct {
	Block                uint64   `protobuf:"varint,1,opt,name=Block,proto3" json:"Block,omitempty"`
	App                  uint64   `protobuf:"varint,2,opt,name=App,proto3" json:"App,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_760f4d5ceb2a11f0, []int{3}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Version) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

// Timestamp wraps how amino encodes time.
// This is the protobuf well-known type protobuf/timestamp.proto
// See:
// https://github.com/google/protobuf/blob/d2980062c859649523d5fd51d6b55ab310e47482/src/google/protobuf/timestamp.proto#L123-L135
// NOTE/XXX: nanos do not get skipped if they are zero in amino.
type Timestamp struct {
	Seconds              int64    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos                int32    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_760f4d5ceb2a11f0, []int{4}
}
func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterType((*PartSetHeader)(nil), "proto3.PartSetHeader")
	proto.RegisterType((*BlockID)(nil), "proto3.BlockID")
	proto.RegisterType((*Header)(nil), "proto3.Header")
	proto.RegisterType((*Version)(nil), "proto3.Version")
	proto.RegisterType((*Timestamp)(nil), "proto3.Timestamp")
}

func init() { proto.RegisterFile("types/proto3/block.proto", fileDescriptor_760f4d5ceb2a11f0) }

var fileDescriptor_760f4d5ceb2a11f0 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x41, 0x6f, 0x13, 0x31,
	0x10, 0x85, 0xb5, 0x64, 0x93, 0x34, 0xb3, 0x49, 0x0b, 0x23, 0x40, 0x16, 0xa7, 0x68, 0x05, 0x28,
	0x5c, 0x52, 0xb5, 0x3d, 0x20, 0xc4, 0x29, 0x24, 0x48, 0xad, 0x84, 0x50, 0x65, 0xaa, 0xdc, 0xdd,
	0xac, 0x45, 0x56, 0x24, 0xb6, 0xe5, 0x71, 0x11, 0xfc, 0x0e, 0xfe, 0x30, 0xf2, 0x78, 0x37, 0x74,
	0x23, 0x6e, 0xfb, 0xde, 0xbc, 0xf9, 0x6c, 0x8f, 0xbd, 0x20, 0xc2, 0x6f, 0xa7, 0xe9, 0xdc, 0x79,
	0x1b, 0xec, 0xd5, 0xf9, 0xfd, 0xce, 0x6e, 0x7e, 0xcc, 0x59, 0xe0, 0x20, 0x79, 0xe5, 0x07, 0x98,
	0xdc, 0x2a, 0x1f, 0xbe, 0xe9, 0x70, 0xad, 0x55, 0xa5, 0x3d, 0x3e, 0x87, 0xfe, 0x9d, 0x0d, 0x6a,
	0x27, 0xb2, 0x69, 0x36, 0xeb, 0xcb, 0x24, 0x10, 0x21, 0xbf, 0x56, 0xb4, 0x15, 0x4f, 0xa6, 0xd9,
	0x6c, 0x2c, 0xf9, 0xbb, 0x5c, 0xc3, 0xf0, 0x53, 0x24, 0xde, 0xac, 0x0e, 0xe5, 0xec, 0x5f, 0x19,
	0xdf, 0x43, 0x11, 0xc9, 0x94, 0xb8, 0xdc, 0x59, 0x5c, 0xbe, 0x48, 0xcb, 0x5f, 0xcd, 0x3b, 0x8b,
	0xca, 0xc7, 0xc9, 0xf2, 0x4f, 0x0e, 0x83, 0x66, 0x33, 0xef, 0x60, 0xb8, 0xd6, 0x9e, 0x6a, 0x6b,
	0x18, 0x5d, 0x5c, 0x9e, 0xb5, 0xfd, 0x8d, 0x2d, 0xdb, 0x3a, 0x0a, 0x18, 0x2e, 0xb7, 0xaa, 0x36,
	0x37, 0x2b, 0x5e, 0x6a, 0x24, 0x5b, 0x89, 0x2f, 0x23, 0xae, 0xfe, 0xbe, 0x0d, 0xa2, 0x37, 0xcd,
	0x66, 0x3d, 0xd9, 0x28, 0x7c, 0x03, 0xf9, 0x5d, 0xbd, 0xd7, 0x22, 0x67, 0xf2, 0xb3, 0x96, 0x1c,
	0x3d, 0x0a, 0x6a, 0xef, 0x24, 0x97, 0xf1, 0x02, 0x8a, 0x2f, 0x8a, 0x42, 0x73, 0x54, 0xd1, 0xef,
	0xee, 0xa3, 0xb1, 0xe5, 0xe3, 0x0c, 0xbe, 0x85, 0xd3, 0x28, 0x97, 0x76, 0xbf, 0xaf, 0x03, 0x0f,
	0x66, 0xc0, 0x83, 0x39, 0x72, 0xf1, 0x15, 0x9c, 0xac, 0x54, 0x50, 0x9c, 0x18, 0x72, 0xe2, 0xa0,
	0x23, 0x63, 0xad, 0x76, 0x75, 0xa5, 0x82, 0xf5, 0xc4, 0x89, 0x93, 0xc4, 0xe8, 0xba, 0x38, 0x07,
	0xfc, 0xaa, 0x7f, 0x85, 0xa3, 0xec, 0x88, 0xb3, 0xff, 0xa9, 0xe0, 0x6b, 0x98, 0x2c, 0xad, 0x21,
	0x6d, 0xe8, 0x21, 0x45, 0x81, 0xa3, 0x5d, 0x33, 0x4e, 0x73, 0xe1, 0x1c, 0xd7, 0x0b, 0xae, 0xb7,
	0x12, 0x67, 0x70, 0x16, 0x4f, 0x21, 0x35, 0x3d, 0xec, 0x42, 0x22, 0x8c, 0x39, 0x71, 0x6c, 0x63,
	0x09, 0xe3, 0xcf, 0x3f, 0xeb, 0x4a, 0x9b, 0x8d, 0xe6, 0xd8, 0x84, 0x63, 0x1d, 0x2f, 0xd2, 0x6e,
	0xbd, 0x75, 0x96, 0xb4, 0x5f, 0x54, 0x95, 0xd7, 0x44, 0xe2, 0x34, 0xd1, 0x8e, 0xec, 0xf2, 0xe2,
	0xf0, 0x14, 0xe2, 0x13, 0xe5, 0x49, 0xf3, 0x9b, 0xc8, 0x65, 0x12, 0xf8, 0x14, 0x7a, 0x0b, 0xe7,
	0xf8, 0xf2, 0x73, 0x19, 0x3f, 0xcb, 0x8f, 0x30, 0x3a, 0x5c, 0x66, 0x3c, 0x11, 0xe9, 0x8d, 0x35,
	0x15, 0x71, 0x5b, 0x4f, 0xb6, 0x32, 0xe2, 0x8c, 0x32, 0x96, 0xb8, 0xb5, 0x2f, 0x93, 0xb8, 0x6f,
	0x7e, 0x90, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x17, 0xe4, 0x46, 0x43, 0x03, 0x00, 0x00,
}
