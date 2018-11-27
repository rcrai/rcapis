// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rcapis/rpc/motoko/context.proto

package context

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/rcrai/rcapis/rcapis/api/annotations"
	context "golang.org/x/net/context"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

// Represents a context.
type Context struct {
	// Required. The unique identifier of the context. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/contexts/<Context ID>`.
	// Note: The Context ID is always converted to lowercase.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The number of conversational query requests after which the
	// context expires. If set to `0` (the default) the context expires
	// immediately. Contexts expire automatically after 10 minutes even if there
	// are no matching queries.
	LifespanCount int32 `protobuf:"varint,2,opt,name=lifespan_count,json=lifespanCount,proto3" json:"lifespan_count,omitempty"`
	// Optional. The collection of parameters associated with this context.
	// Refer to [this doc](https://dialogflow.com/docs/actions-and-parameters) for
	// syntax.
	Parameters           *_struct.Struct `protobuf:"bytes,3,opt,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2d3b67c2eec1e9, []int{0}
}

func (m *Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Context.Unmarshal(m, b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Context.Marshal(b, m, deterministic)
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return xxx_messageInfo_Context.Size(m)
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func (m *Context) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Context) GetLifespanCount() int32 {
	if m != nil {
		return m.LifespanCount
	}
	return 0
}

func (m *Context) GetParameters() *_struct.Struct {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// The request message for [Contexts.ListContexts].
type ListContextsRequest struct {
	// Required. The session to list all contexts from.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of items to return in a single page. By
	// default 100 and at most 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListContextsRequest) Reset()         { *m = ListContextsRequest{} }
func (m *ListContextsRequest) String() string { return proto.CompactTextString(m) }
func (*ListContextsRequest) ProtoMessage()    {}
func (*ListContextsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2d3b67c2eec1e9, []int{1}
}

func (m *ListContextsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListContextsRequest.Unmarshal(m, b)
}
func (m *ListContextsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListContextsRequest.Marshal(b, m, deterministic)
}
func (m *ListContextsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContextsRequest.Merge(m, src)
}
func (m *ListContextsRequest) XXX_Size() int {
	return xxx_messageInfo_ListContextsRequest.Size(m)
}
func (m *ListContextsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContextsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListContextsRequest proto.InternalMessageInfo

func (m *ListContextsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListContextsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListContextsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for [Contexts.ListContexts].
type ListContextsResponse struct {
	// The list of contexts. There will be a maximum number of items
	// returned based on the page_size field in the request.
	Contexts []*Context `protobuf:"bytes,1,rep,name=contexts,proto3" json:"contexts,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListContextsResponse) Reset()         { *m = ListContextsResponse{} }
func (m *ListContextsResponse) String() string { return proto.CompactTextString(m) }
func (*ListContextsResponse) ProtoMessage()    {}
func (*ListContextsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2d3b67c2eec1e9, []int{2}
}

func (m *ListContextsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListContextsResponse.Unmarshal(m, b)
}
func (m *ListContextsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListContextsResponse.Marshal(b, m, deterministic)
}
func (m *ListContextsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContextsResponse.Merge(m, src)
}
func (m *ListContextsResponse) XXX_Size() int {
	return xxx_messageInfo_ListContextsResponse.Size(m)
}
func (m *ListContextsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContextsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListContextsResponse proto.InternalMessageInfo

func (m *ListContextsResponse) GetContexts() []*Context {
	if m != nil {
		return m.Contexts
	}
	return nil
}

func (m *ListContextsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for [Contexts.GetContext].
type GetContextRequest struct {
	// Required. The name of the context. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/contexts/<Context ID>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContextRequest) Reset()         { *m = GetContextRequest{} }
func (m *GetContextRequest) String() string { return proto.CompactTextString(m) }
func (*GetContextRequest) ProtoMessage()    {}
func (*GetContextRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2d3b67c2eec1e9, []int{3}
}

func (m *GetContextRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContextRequest.Unmarshal(m, b)
}
func (m *GetContextRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContextRequest.Marshal(b, m, deterministic)
}
func (m *GetContextRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContextRequest.Merge(m, src)
}
func (m *GetContextRequest) XXX_Size() int {
	return xxx_messageInfo_GetContextRequest.Size(m)
}
func (m *GetContextRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContextRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetContextRequest proto.InternalMessageInfo

func (m *GetContextRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Contexts.CreateContext].
type CreateContextRequest struct {
	// Required. The session to create a context for.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The context to create.
	Context              *Context `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateContextRequest) Reset()         { *m = CreateContextRequest{} }
func (m *CreateContextRequest) String() string { return proto.CompactTextString(m) }
func (*CreateContextRequest) ProtoMessage()    {}
func (*CreateContextRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2d3b67c2eec1e9, []int{4}
}

func (m *CreateContextRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateContextRequest.Unmarshal(m, b)
}
func (m *CreateContextRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateContextRequest.Marshal(b, m, deterministic)
}
func (m *CreateContextRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateContextRequest.Merge(m, src)
}
func (m *CreateContextRequest) XXX_Size() int {
	return xxx_messageInfo_CreateContextRequest.Size(m)
}
func (m *CreateContextRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateContextRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateContextRequest proto.InternalMessageInfo

func (m *CreateContextRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateContextRequest) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

// The request message for [Contexts.UpdateContext].
type UpdateContextRequest struct {
	// Required. The context to update. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/contexts/<Context ID>`.
	Context *Context `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	// Optional. The mask to control which fields get updated.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateContextRequest) Reset()         { *m = UpdateContextRequest{} }
func (m *UpdateContextRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateContextRequest) ProtoMessage()    {}
func (*UpdateContextRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2d3b67c2eec1e9, []int{5}
}

func (m *UpdateContextRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContextRequest.Unmarshal(m, b)
}
func (m *UpdateContextRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContextRequest.Marshal(b, m, deterministic)
}
func (m *UpdateContextRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContextRequest.Merge(m, src)
}
func (m *UpdateContextRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateContextRequest.Size(m)
}
func (m *UpdateContextRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContextRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContextRequest proto.InternalMessageInfo

func (m *UpdateContextRequest) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *UpdateContextRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The request message for [Contexts.DeleteContext].
type DeleteContextRequest struct {
	// Required. The name of the context to delete. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/contexts/<Context ID>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteContextRequest) Reset()         { *m = DeleteContextRequest{} }
func (m *DeleteContextRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteContextRequest) ProtoMessage()    {}
func (*DeleteContextRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2d3b67c2eec1e9, []int{6}
}

func (m *DeleteContextRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteContextRequest.Unmarshal(m, b)
}
func (m *DeleteContextRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteContextRequest.Marshal(b, m, deterministic)
}
func (m *DeleteContextRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteContextRequest.Merge(m, src)
}
func (m *DeleteContextRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteContextRequest.Size(m)
}
func (m *DeleteContextRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteContextRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteContextRequest proto.InternalMessageInfo

func (m *DeleteContextRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Contexts.DeleteAllContexts].
type DeleteAllContextsRequest struct {
	// Required. The name of the session to delete all contexts from. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>`.
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAllContextsRequest) Reset()         { *m = DeleteAllContextsRequest{} }
func (m *DeleteAllContextsRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAllContextsRequest) ProtoMessage()    {}
func (*DeleteAllContextsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec2d3b67c2eec1e9, []int{7}
}

func (m *DeleteAllContextsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAllContextsRequest.Unmarshal(m, b)
}
func (m *DeleteAllContextsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAllContextsRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAllContextsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAllContextsRequest.Merge(m, src)
}
func (m *DeleteAllContextsRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAllContextsRequest.Size(m)
}
func (m *DeleteAllContextsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAllContextsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAllContextsRequest proto.InternalMessageInfo

func (m *DeleteAllContextsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func init() {
	proto.RegisterType((*Context)(nil), "rcrai.rpc.motoko.Context")
	proto.RegisterType((*ListContextsRequest)(nil), "rcrai.rpc.motoko.ListContextsRequest")
	proto.RegisterType((*ListContextsResponse)(nil), "rcrai.rpc.motoko.ListContextsResponse")
	proto.RegisterType((*GetContextRequest)(nil), "rcrai.rpc.motoko.GetContextRequest")
	proto.RegisterType((*CreateContextRequest)(nil), "rcrai.rpc.motoko.CreateContextRequest")
	proto.RegisterType((*UpdateContextRequest)(nil), "rcrai.rpc.motoko.UpdateContextRequest")
	proto.RegisterType((*DeleteContextRequest)(nil), "rcrai.rpc.motoko.DeleteContextRequest")
	proto.RegisterType((*DeleteAllContextsRequest)(nil), "rcrai.rpc.motoko.DeleteAllContextsRequest")
}

func init() { proto.RegisterFile("rcapis/rpc/motoko/context.proto", fileDescriptor_ec2d3b67c2eec1e9) }

var fileDescriptor_ec2d3b67c2eec1e9 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0xc7, 0xe5, 0xed, 0xf7, 0xdb, 0x9f, 0x67, 0x2b, 0x30, 0x53, 0x8d, 0xd2, 0x0d, 0x51, 0x05,
	0x6d, 0x54, 0x3d, 0xc4, 0x5a, 0x27, 0x84, 0xc6, 0xe0, 0xc0, 0x36, 0x40, 0x42, 0x20, 0x4d, 0x19,
	0x5c, 0xb8, 0x54, 0x5e, 0xf6, 0xac, 0x84, 0x36, 0x71, 0x88, 0x1d, 0x34, 0x86, 0x76, 0xe1, 0x02,
	0x82, 0x1b, 0x48, 0x88, 0x13, 0x6f, 0x8a, 0xb7, 0xc0, 0x8b, 0xe0, 0x88, 0xe2, 0x38, 0x5d, 0xdb,
	0xa4, 0x65, 0xe3, 0xd4, 0xc6, 0xf9, 0xda, 0xdf, 0x8f, 0x9f, 0xe7, 0xeb, 0x18, 0xae, 0x47, 0x2e,
	0x0f, 0x3d, 0xc9, 0xa2, 0xd0, 0x65, 0xbe, 0x50, 0xa2, 0x23, 0x98, 0x2b, 0x02, 0x85, 0x47, 0xca,
	0x0e, 0x23, 0xa1, 0x04, 0xbd, 0x14, 0xb9, 0x11, 0xf7, 0xec, 0x28, 0x74, 0xed, 0xf4, 0x7d, 0x75,
	0xd9, 0x4c, 0xe1, 0xa1, 0xc7, 0x78, 0x10, 0x08, 0xc5, 0x95, 0x27, 0x02, 0x99, 0xea, 0xab, 0x4b,
	0x6d, 0x21, 0xda, 0x5d, 0x64, 0xfa, 0x69, 0x3f, 0x3e, 0x64, 0xe8, 0x87, 0xea, 0xad, 0x79, 0x59,
	0x1b, 0x7e, 0x79, 0xe8, 0x61, 0xf7, 0xa0, 0xe5, 0x73, 0xd9, 0x31, 0x8a, 0xe5, 0x61, 0x85, 0x54,
	0x51, 0xec, 0x1a, 0x18, 0xeb, 0x04, 0xa6, 0xb7, 0x53, 0x3a, 0x4a, 0xe1, 0xbf, 0x80, 0xfb, 0x58,
	0x21, 0x35, 0x52, 0x9f, 0x75, 0xf4, 0x7f, 0xba, 0x02, 0x17, 0xba, 0xde, 0x21, 0xca, 0x90, 0x07,
	0x2d, 0x57, 0xc4, 0x81, 0xaa, 0x4c, 0xd4, 0x48, 0xfd, 0x7f, 0xa7, 0x94, 0x8d, 0x6e, 0x27, 0x83,
	0xf4, 0x36, 0x40, 0xc8, 0x23, 0xee, 0xa3, 0xc2, 0x48, 0x56, 0x26, 0x6b, 0xa4, 0x3e, 0xd7, 0xbc,
	0x62, 0xa7, 0xc6, 0x76, 0x66, 0x6c, 0xef, 0x69, 0x63, 0xa7, 0x4f, 0x6a, 0x79, 0x70, 0xf9, 0x89,
	0x27, 0x95, 0x41, 0x90, 0x0e, 0xbe, 0x8e, 0x51, 0x2a, 0xba, 0x08, 0x53, 0x21, 0x8f, 0x30, 0x50,
	0x06, 0xc6, 0x3c, 0xd1, 0x25, 0x98, 0x0d, 0x79, 0x1b, 0x5b, 0xd2, 0x3b, 0x46, 0x43, 0x32, 0x93,
	0x0c, 0xec, 0x79, 0xc7, 0x48, 0xaf, 0x25, 0x10, 0x6d, 0x6c, 0x29, 0xd1, 0xc1, 0x40, 0x43, 0xcc,
	0x3a, 0x5a, 0xfe, 0x2c, 0x19, 0xb0, 0x62, 0x28, 0x0f, 0x5a, 0xc9, 0x50, 0x04, 0x12, 0xe9, 0x2d,
	0x98, 0x31, 0xfd, 0x91, 0x15, 0x52, 0x9b, 0xac, 0xcf, 0x35, 0xaf, 0xda, 0xc3, 0x1d, 0xb2, 0xcd,
	0x2c, 0xa7, 0x27, 0xa5, 0xab, 0x70, 0x31, 0xc0, 0x23, 0xd5, 0xea, 0xb3, 0x9c, 0xd0, 0x96, 0xa5,
	0x64, 0x78, 0xb7, 0x67, 0x7b, 0x13, 0x16, 0x1e, 0x61, 0xe6, 0x9a, 0xed, 0xaf, 0xa0, 0xd4, 0x96,
	0x0b, 0xe5, 0xed, 0x08, 0xb9, 0xc2, 0x21, 0xed, 0xa8, 0x5a, 0xac, 0xc3, 0xb4, 0x81, 0xd1, 0xc6,
	0x63, 0xb1, 0x33, 0xa5, 0xf5, 0x91, 0x40, 0xf9, 0x79, 0x78, 0x90, 0x77, 0xe9, 0x5b, 0x8d, 0x9c,
	0x75, 0x35, 0xba, 0x09, 0x73, 0xb1, 0x5e, 0x4c, 0xe7, 0xcd, 0x60, 0x54, 0x73, 0x7d, 0x7f, 0x98,
	0x44, 0xf2, 0x29, 0x97, 0x1d, 0x07, 0x52, 0x79, 0xf2, 0xdf, 0x6a, 0x40, 0x79, 0x07, 0xbb, 0x98,
	0x23, 0x29, 0xaa, 0x4d, 0x13, 0x2a, 0xa9, 0xf6, 0x7e, 0xb7, 0x7b, 0xc6, 0xac, 0x34, 0x3f, 0x4f,
	0xc3, 0x4c, 0xa6, 0xa5, 0xdf, 0x09, 0xcc, 0xf7, 0x77, 0x9f, 0xae, 0xe4, 0xb7, 0x57, 0x10, 0xc4,
	0xea, 0xea, 0xdf, 0x64, 0x69, 0x88, 0xac, 0x8d, 0xf7, 0x3f, 0x7f, 0x7d, 0x9d, 0x58, 0xa7, 0x6b,
	0xec, 0xcd, 0x1a, 0x7b, 0x97, 0x12, 0xdc, 0x0b, 0x23, 0xf1, 0x0a, 0x5d, 0x25, 0x59, 0x83, 0xf1,
	0x36, 0x06, 0x8a, 0x49, 0x94, 0x32, 0x39, 0xd8, 0xac, 0x71, 0xc2, 0x7a, 0x41, 0xfa, 0x40, 0x00,
	0x4e, 0x13, 0x42, 0x6f, 0xe4, 0x1d, 0x73, 0xf9, 0xa9, 0x8e, 0x6e, 0xce, 0x10, 0x49, 0x52, 0xbd,
	0x71, 0x1c, 0x3d, 0x0c, 0xd6, 0x38, 0xa1, 0xdf, 0x08, 0x94, 0x06, 0x22, 0x48, 0x0b, 0xb6, 0x5f,
	0x94, 0xd1, 0x71, 0x3c, 0x5b, 0x9a, 0xe7, 0xae, 0x75, 0xfe, 0xca, 0xdc, 0xe9, 0xe5, 0xec, 0x07,
	0x81, 0xd2, 0x40, 0x6a, 0x8b, 0xc0, 0x8a, 0x62, 0x3d, 0x0e, 0xec, 0xb1, 0x06, 0xdb, 0x69, 0x6e,
	0x68, 0xb0, 0xec, 0x13, 0x7d, 0x9e, 0x82, 0x9d, 0x02, 0x7e, 0x22, 0x50, 0x1a, 0x08, 0x73, 0x11,
	0x60, 0x51, 0xda, 0xab, 0x8b, 0xb9, 0xd3, 0xf2, 0x20, 0xf9, 0xba, 0x67, 0x6d, 0x6c, 0xfc, 0x43,
	0x1b, 0xbf, 0x10, 0x58, 0xc8, 0x9d, 0x16, 0xda, 0x18, 0x05, 0x94, 0x3f, 0x52, 0x67, 0x84, 0x3a,
	0x4f, 0x2f, 0xb7, 0x3a, 0xb0, 0xe4, 0x0a, 0x3f, 0x63, 0xd0, 0xd7, 0x5d, 0x1f, 0xca, 0xd6, 0xbc,
	0x41, 0xd8, 0x4d, 0x0c, 0x77, 0xc9, 0x8b, 0x8d, 0xb6, 0xa7, 0x5e, 0xc6, 0xfb, 0xb6, 0x2b, 0x7c,
	0xa6, 0xe7, 0xb0, 0xec, 0x56, 0x1d, 0x75, 0xb9, 0x6e, 0x9a, 0xdf, 0xdf, 0x84, 0xec, 0x4f, 0x69,
	0xee, 0xf5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x54, 0x46, 0xde, 0xbd, 0x8b, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContextsClient is the client API for Contexts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContextsClient interface {
	// Returns the list of all contexts in the specified session.
	ListContexts(ctx context.Context, in *ListContextsRequest, opts ...grpc.CallOption) (*ListContextsResponse, error)
	// Retrieves the specified context.
	GetContext(ctx context.Context, in *GetContextRequest, opts ...grpc.CallOption) (*Context, error)
	// Creates a context.
	CreateContext(ctx context.Context, in *CreateContextRequest, opts ...grpc.CallOption) (*Context, error)
	// Updates the specified context.
	UpdateContext(ctx context.Context, in *UpdateContextRequest, opts ...grpc.CallOption) (*Context, error)
	// Deletes the specified context.
	DeleteContext(ctx context.Context, in *DeleteContextRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Deletes all active contexts in the specified session.
	DeleteAllContexts(ctx context.Context, in *DeleteAllContextsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type contextsClient struct {
	cc *grpc.ClientConn
}

func NewContextsClient(cc *grpc.ClientConn) ContextsClient {
	return &contextsClient{cc}
}

func (c *contextsClient) ListContexts(ctx context.Context, in *ListContextsRequest, opts ...grpc.CallOption) (*ListContextsResponse, error) {
	out := new(ListContextsResponse)
	err := c.cc.Invoke(ctx, "/rcrai.rpc.motoko.Contexts/ListContexts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextsClient) GetContext(ctx context.Context, in *GetContextRequest, opts ...grpc.CallOption) (*Context, error) {
	out := new(Context)
	err := c.cc.Invoke(ctx, "/rcrai.rpc.motoko.Contexts/GetContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextsClient) CreateContext(ctx context.Context, in *CreateContextRequest, opts ...grpc.CallOption) (*Context, error) {
	out := new(Context)
	err := c.cc.Invoke(ctx, "/rcrai.rpc.motoko.Contexts/CreateContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextsClient) UpdateContext(ctx context.Context, in *UpdateContextRequest, opts ...grpc.CallOption) (*Context, error) {
	out := new(Context)
	err := c.cc.Invoke(ctx, "/rcrai.rpc.motoko.Contexts/UpdateContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextsClient) DeleteContext(ctx context.Context, in *DeleteContextRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rcrai.rpc.motoko.Contexts/DeleteContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextsClient) DeleteAllContexts(ctx context.Context, in *DeleteAllContextsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rcrai.rpc.motoko.Contexts/DeleteAllContexts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContextsServer is the server API for Contexts service.
type ContextsServer interface {
	// Returns the list of all contexts in the specified session.
	ListContexts(context.Context, *ListContextsRequest) (*ListContextsResponse, error)
	// Retrieves the specified context.
	GetContext(context.Context, *GetContextRequest) (*Context, error)
	// Creates a context.
	CreateContext(context.Context, *CreateContextRequest) (*Context, error)
	// Updates the specified context.
	UpdateContext(context.Context, *UpdateContextRequest) (*Context, error)
	// Deletes the specified context.
	DeleteContext(context.Context, *DeleteContextRequest) (*empty.Empty, error)
	// Deletes all active contexts in the specified session.
	DeleteAllContexts(context.Context, *DeleteAllContextsRequest) (*empty.Empty, error)
}

func RegisterContextsServer(s *grpc.Server, srv ContextsServer) {
	s.RegisterService(&_Contexts_serviceDesc, srv)
}

func _Contexts_ListContexts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContextsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextsServer).ListContexts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcrai.rpc.motoko.Contexts/ListContexts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextsServer).ListContexts(ctx, req.(*ListContextsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contexts_GetContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextsServer).GetContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcrai.rpc.motoko.Contexts/GetContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextsServer).GetContext(ctx, req.(*GetContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contexts_CreateContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextsServer).CreateContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcrai.rpc.motoko.Contexts/CreateContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextsServer).CreateContext(ctx, req.(*CreateContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contexts_UpdateContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextsServer).UpdateContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcrai.rpc.motoko.Contexts/UpdateContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextsServer).UpdateContext(ctx, req.(*UpdateContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contexts_DeleteContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextsServer).DeleteContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcrai.rpc.motoko.Contexts/DeleteContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextsServer).DeleteContext(ctx, req.(*DeleteContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contexts_DeleteAllContexts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllContextsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextsServer).DeleteAllContexts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcrai.rpc.motoko.Contexts/DeleteAllContexts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextsServer).DeleteAllContexts(ctx, req.(*DeleteAllContextsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contexts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rcrai.rpc.motoko.Contexts",
	HandlerType: (*ContextsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListContexts",
			Handler:    _Contexts_ListContexts_Handler,
		},
		{
			MethodName: "GetContext",
			Handler:    _Contexts_GetContext_Handler,
		},
		{
			MethodName: "CreateContext",
			Handler:    _Contexts_CreateContext_Handler,
		},
		{
			MethodName: "UpdateContext",
			Handler:    _Contexts_UpdateContext_Handler,
		},
		{
			MethodName: "DeleteContext",
			Handler:    _Contexts_DeleteContext_Handler,
		},
		{
			MethodName: "DeleteAllContexts",
			Handler:    _Contexts_DeleteAllContexts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rcapis/rpc/motoko/context.proto",
}
