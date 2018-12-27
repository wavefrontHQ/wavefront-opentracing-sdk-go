// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wire.proto

package wire

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

type TracerState struct {
	TraceId              string            `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId               string            `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	Sampled              bool              `protobuf:"varint,3,opt,name=sampled,proto3" json:"sampled,omitempty"`
	BaggageItems         map[string]string `protobuf:"bytes,4,rep,name=baggage_items,json=baggageItems,proto3" json:"baggage_items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TracerState) Reset()         { *m = TracerState{} }
func (m *TracerState) String() string { return proto.CompactTextString(m) }
func (*TracerState) ProtoMessage()    {}
func (*TracerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2dcdddcdf68d8e0, []int{0}
}

func (m *TracerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TracerState.Unmarshal(m, b)
}
func (m *TracerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TracerState.Marshal(b, m, deterministic)
}
func (m *TracerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracerState.Merge(m, src)
}
func (m *TracerState) XXX_Size() int {
	return xxx_messageInfo_TracerState.Size(m)
}
func (m *TracerState) XXX_DiscardUnknown() {
	xxx_messageInfo_TracerState.DiscardUnknown(m)
}

var xxx_messageInfo_TracerState proto.InternalMessageInfo

func (m *TracerState) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *TracerState) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *TracerState) GetSampled() bool {
	if m != nil {
		return m.Sampled
	}
	return false
}

func (m *TracerState) GetBaggageItems() map[string]string {
	if m != nil {
		return m.BaggageItems
	}
	return nil
}

func init() {
	proto.RegisterType((*TracerState)(nil), "wavefront_go.wire.TracerState")
	proto.RegisterMapType((map[string]string)(nil), "wavefront_go.wire.TracerState.BaggageItemsEntry")
}

func init() { proto.RegisterFile("wire.proto", fileDescriptor_f2dcdddcdf68d8e0) }

var fileDescriptor_f2dcdddcdf68d8e0 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xcf, 0x2c, 0x4a,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x4f, 0x2c, 0x4b, 0x4d, 0x2b, 0xca, 0xcf,
	0x2b, 0x89, 0x4f, 0xcf, 0xd7, 0x03, 0x49, 0x28, 0x7d, 0x66, 0xe4, 0xe2, 0x0e, 0x29, 0x4a, 0x4c,
	0x4e, 0x2d, 0x0a, 0x2e, 0x49, 0x2c, 0x49, 0x15, 0x92, 0xe4, 0xe2, 0x28, 0x01, 0x71, 0xe3, 0x33,
	0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xd8, 0xc1, 0x7c, 0xcf, 0x14, 0x21, 0x71, 0x2e,
	0xf6, 0xe2, 0x82, 0xc4, 0x3c, 0x90, 0x0c, 0x13, 0x58, 0x86, 0x0d, 0xc4, 0xf5, 0x4c, 0x11, 0x92,
	0xe0, 0x62, 0x2f, 0x4e, 0xcc, 0x2d, 0xc8, 0x49, 0x4d, 0x91, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x08,
	0x82, 0x71, 0x85, 0x42, 0xb9, 0x78, 0x93, 0x12, 0xd3, 0xd3, 0x13, 0xd3, 0x53, 0xe3, 0x33, 0x4b,
	0x52, 0x73, 0x8b, 0x25, 0x58, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0x0c, 0xf4, 0x30, 0x1c, 0xa2, 0x87,
	0xe4, 0x08, 0x3d, 0x27, 0x88, 0x1e, 0x4f, 0x90, 0x16, 0xd7, 0xbc, 0x92, 0xa2, 0xca, 0x20, 0x9e,
	0x24, 0x24, 0x21, 0x29, 0x7b, 0x2e, 0x41, 0x0c, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95,
	0x50, 0x47, 0x83, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x50, 0xe7, 0x42,
	0x38, 0x56, 0x4c, 0x16, 0x8c, 0x4e, 0x6c, 0x51, 0x2c, 0x20, 0x4b, 0x93, 0xd8, 0xc0, 0xe1, 0x62,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x65, 0x5d, 0x44, 0x25, 0x01, 0x00, 0x00,
}
