// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rcapis/api/httpbody.proto

package httpbody

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Message that represents an arbitrary HTTP body. It should only be used for
// payload formats that can't be represented as JSON, such as raw binary or
// an HTML page.
//
//
// This message can be used both in streaming and non-streaming API methods in
// the request as well as the response.
//
// It can be used as a top-level request field, which is convenient if one
// wants to extract parameters from either the URL or HTTP template into the
// request fields and also want access to the raw HTTP body.
//
// Example:
//
//     message GetResourceRequest {
//       // A unique request id.
//       string request_id = 1;
//
//       // The raw HTTP body is bound to this field.
//       google.api.HttpBody http_body = 2;
//     }
//
//     service ResourceService {
//       rpc GetResource(GetResourceRequest) returns (google.api.HttpBody);
//       rpc UpdateResource(google.api.HttpBody) returns (google.protobuf.Empty);
//     }
//
// Example with streaming methods:
//
//     service CaldavService {
//       rpc GetCalendar(stream google.api.HttpBody)
//         returns (stream google.api.HttpBody);
//       rpc UpdateCalendar(stream google.api.HttpBody)
//         returns (stream google.api.HttpBody);
//     }
//
// Use of this type only changes how the request and response bodies are
// handled, all other features will continue to work unchanged.
type HttpBody struct {
	// The HTTP Content-Type string representing the content type of the body.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// HTTP body binary data.
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpBody) Reset()         { *m = HttpBody{} }
func (m *HttpBody) String() string { return proto.CompactTextString(m) }
func (*HttpBody) ProtoMessage()    {}
func (*HttpBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f5266f92101f56a, []int{0}
}

func (m *HttpBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpBody.Unmarshal(m, b)
}
func (m *HttpBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpBody.Marshal(b, m, deterministic)
}
func (m *HttpBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpBody.Merge(m, src)
}
func (m *HttpBody) XXX_Size() int {
	return xxx_messageInfo_HttpBody.Size(m)
}
func (m *HttpBody) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpBody.DiscardUnknown(m)
}

var xxx_messageInfo_HttpBody proto.InternalMessageInfo

func (m *HttpBody) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *HttpBody) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpBody)(nil), "rcrai.api.HttpBody")
}

func init() { proto.RegisterFile("rcapis/api/httpbody.proto", fileDescriptor_3f5266f92101f56a) }

var fileDescriptor_3f5266f92101f56a = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x4a, 0x4e, 0x2c,
	0xc8, 0x2c, 0xd6, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x28, 0x29, 0x29, 0x48, 0xca, 0x4f, 0xa9, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c, 0x4a, 0x2e, 0x4a, 0xcc, 0xd4, 0x4b, 0x2c, 0xc8,
	0x54, 0x72, 0xe4, 0xe2, 0xf0, 0x28, 0x29, 0x29, 0x70, 0xca, 0x4f, 0xa9, 0x14, 0x52, 0xe4, 0xe2,
	0x49, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x89, 0x2f, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0xe2, 0x86, 0x8a, 0x85, 0x54, 0x16, 0xa4, 0x0a, 0x09, 0x71, 0xb1, 0xa4, 0x24,
	0x96, 0x24, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x4e, 0x09, 0x5c, 0x22, 0xc9,
	0xf9, 0xb9, 0x7a, 0x10, 0x33, 0x21, 0x96, 0x82, 0x8c, 0x76, 0xe2, 0x85, 0x19, 0x1c, 0x00, 0xb2,
	0x34, 0x80, 0x31, 0xca, 0x24, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f,
	0xac, 0x5a, 0x1f, 0xea, 0x44, 0x2c, 0x2e, 0xb5, 0x86, 0x31, 0x7e, 0x30, 0x32, 0x26, 0xb1, 0x81,
	0x9d, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xb5, 0x2a, 0xb2, 0xd3, 0x00, 0x00, 0x00,
}
