// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rcapis/rpc/motoko/webhook.proto

package webhook

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/rcrai/rcapis/rcapis/api/annotations"
	context "github.com/rcrai/rcapis/rcapis/rpc/motoko/context"
	intent "github.com/rcrai/rcapis/rcapis/rpc/motoko/intent"
	session "github.com/rcrai/rcapis/rcapis/rpc/motoko/session"
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

// The request message for a webhook call.
type WebhookRequest struct {
	// The unique identifier of detectIntent request session.
	// Can be used to identify end-user inside webhook implementation.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Session string `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	// The unique identifier of the response. Contains the same value as
	// `[Streaming]DetectIntentResponse.response_id`.
	ResponseId string `protobuf:"bytes,1,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
	// The result of the conversational query or event processing. Contains the
	// same value as `[Streaming]DetectIntentResponse.query_result`.
	QueryResult *session.QueryResult `protobuf:"bytes,2,opt,name=query_result,json=queryResult,proto3" json:"query_result,omitempty"`
	// Optional. The contents of the original request that was passed to
	// `[Streaming]DetectIntent` call.
	OriginalDetectIntentRequest *OriginalDetectIntentRequest `protobuf:"bytes,3,opt,name=original_detect_intent_request,json=originalDetectIntentRequest,proto3" json:"original_detect_intent_request,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                     `json:"-"`
	XXX_unrecognized            []byte                       `json:"-"`
	XXX_sizecache               int32                        `json:"-"`
}

func (m *WebhookRequest) Reset()         { *m = WebhookRequest{} }
func (m *WebhookRequest) String() string { return proto.CompactTextString(m) }
func (*WebhookRequest) ProtoMessage()    {}
func (*WebhookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de2160a3cdceef7a, []int{0}
}

func (m *WebhookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookRequest.Unmarshal(m, b)
}
func (m *WebhookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookRequest.Marshal(b, m, deterministic)
}
func (m *WebhookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookRequest.Merge(m, src)
}
func (m *WebhookRequest) XXX_Size() int {
	return xxx_messageInfo_WebhookRequest.Size(m)
}
func (m *WebhookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookRequest proto.InternalMessageInfo

func (m *WebhookRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *WebhookRequest) GetResponseId() string {
	if m != nil {
		return m.ResponseId
	}
	return ""
}

func (m *WebhookRequest) GetQueryResult() *session.QueryResult {
	if m != nil {
		return m.QueryResult
	}
	return nil
}

func (m *WebhookRequest) GetOriginalDetectIntentRequest() *OriginalDetectIntentRequest {
	if m != nil {
		return m.OriginalDetectIntentRequest
	}
	return nil
}

// The response message for a webhook call.
type WebhookResponse struct {
	// Optional. The text to be shown on the screen. This value is passed directly
	// to `QueryResult.fulfillment_text`.
	FulfillmentText string `protobuf:"bytes,1,opt,name=fulfillment_text,json=fulfillmentText,proto3" json:"fulfillment_text,omitempty"`
	// Optional. The collection of rich messages to present to the user. This
	// value is passed directly to `QueryResult.fulfillment_messages`.
	FulfillmentMessages []*intent.Intent_Message `protobuf:"bytes,2,rep,name=fulfillment_messages,json=fulfillmentMessages,proto3" json:"fulfillment_messages,omitempty"`
	// Optional. This value is passed directly to `QueryResult.webhook_source`.
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. This value is passed directly to `QueryResult.webhook_payload`.
	Payload *_struct.Struct `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	// Optional. The collection of output contexts. This value is passed directly
	// to `QueryResult.output_contexts`.
	OutputContexts []*context.Context `protobuf:"bytes,5,rep,name=output_contexts,json=outputContexts,proto3" json:"output_contexts,omitempty"`
	// Optional. Makes the platform immediately invoke another `DetectIntent` call
	// internally with the specified event as input.
	FollowupEventInput   *session.EventInput `protobuf:"bytes,6,opt,name=followup_event_input,json=followupEventInput,proto3" json:"followup_event_input,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *WebhookResponse) Reset()         { *m = WebhookResponse{} }
func (m *WebhookResponse) String() string { return proto.CompactTextString(m) }
func (*WebhookResponse) ProtoMessage()    {}
func (*WebhookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de2160a3cdceef7a, []int{1}
}

func (m *WebhookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookResponse.Unmarshal(m, b)
}
func (m *WebhookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookResponse.Marshal(b, m, deterministic)
}
func (m *WebhookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookResponse.Merge(m, src)
}
func (m *WebhookResponse) XXX_Size() int {
	return xxx_messageInfo_WebhookResponse.Size(m)
}
func (m *WebhookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookResponse proto.InternalMessageInfo

func (m *WebhookResponse) GetFulfillmentText() string {
	if m != nil {
		return m.FulfillmentText
	}
	return ""
}

func (m *WebhookResponse) GetFulfillmentMessages() []*intent.Intent_Message {
	if m != nil {
		return m.FulfillmentMessages
	}
	return nil
}

func (m *WebhookResponse) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *WebhookResponse) GetPayload() *_struct.Struct {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *WebhookResponse) GetOutputContexts() []*context.Context {
	if m != nil {
		return m.OutputContexts
	}
	return nil
}

func (m *WebhookResponse) GetFollowupEventInput() *session.EventInput {
	if m != nil {
		return m.FollowupEventInput
	}
	return nil
}

// Represents the contents of the original request that was passed to
// `[Streaming]DetectIntent call`.
type OriginalDetectIntentRequest struct {
	// The source of this request, e.g., Google, Facebook, Slack. It is set by
	// Dialogflow-owned servers.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. This field is set to the value of `QueryParameters.payload` field
	// passed in the request.
	Payload              *_struct.Struct `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OriginalDetectIntentRequest) Reset()         { *m = OriginalDetectIntentRequest{} }
func (m *OriginalDetectIntentRequest) String() string { return proto.CompactTextString(m) }
func (*OriginalDetectIntentRequest) ProtoMessage()    {}
func (*OriginalDetectIntentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de2160a3cdceef7a, []int{2}
}

func (m *OriginalDetectIntentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OriginalDetectIntentRequest.Unmarshal(m, b)
}
func (m *OriginalDetectIntentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OriginalDetectIntentRequest.Marshal(b, m, deterministic)
}
func (m *OriginalDetectIntentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalDetectIntentRequest.Merge(m, src)
}
func (m *OriginalDetectIntentRequest) XXX_Size() int {
	return xxx_messageInfo_OriginalDetectIntentRequest.Size(m)
}
func (m *OriginalDetectIntentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalDetectIntentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalDetectIntentRequest proto.InternalMessageInfo

func (m *OriginalDetectIntentRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *OriginalDetectIntentRequest) GetPayload() *_struct.Struct {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*WebhookRequest)(nil), "rcrai.rpc.motoko.WebhookRequest")
	proto.RegisterType((*WebhookResponse)(nil), "rcrai.rpc.motoko.WebhookResponse")
	proto.RegisterType((*OriginalDetectIntentRequest)(nil), "rcrai.rpc.motoko.OriginalDetectIntentRequest")
}

func init() { proto.RegisterFile("rcapis/rpc/motoko/webhook.proto", fileDescriptor_de2160a3cdceef7a) }

var fileDescriptor_de2160a3cdceef7a = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0x26, 0x90, 0xaa, 0x4e, 0xd5, 0x54, 0xa6, 0x82, 0xa5, 0x09, 0x6d, 0x94, 0x53, 0x38,
	0xe0, 0x15, 0xe1, 0x84, 0xb8, 0xa0, 0x00, 0x87, 0x1c, 0x80, 0xe2, 0x22, 0x21, 0x71, 0x59, 0x6d,
	0x1c, 0x27, 0xb1, 0xb2, 0xf1, 0x38, 0xfe, 0xa0, 0xed, 0x91, 0x7f, 0xcc, 0x91, 0x23, 0x5a, 0xdb,
	0x4b, 0x23, 0x42, 0x10, 0xa7, 0xd5, 0x8c, 0xdf, 0x1b, 0xbf, 0xe7, 0xb7, 0x83, 0x2e, 0x34, 0x2b,
	0x94, 0x30, 0x99, 0x56, 0x2c, 0x5b, 0x83, 0x85, 0x15, 0x64, 0xd7, 0x7c, 0xba, 0x04, 0x58, 0x11,
	0xa5, 0xc1, 0x02, 0x3e, 0xd1, 0x4c, 0x17, 0x82, 0x68, 0xc5, 0x48, 0x38, 0x3f, 0xeb, 0x45, 0x4a,
	0xa1, 0x44, 0x56, 0x48, 0x09, 0xb6, 0xb0, 0x02, 0xa4, 0x09, 0xf8, 0xb3, 0xbf, 0x0c, 0x64, 0x20,
	0x2d, 0xbf, 0xb1, 0x11, 0x70, 0xbe, 0x0b, 0x10, 0xd2, 0x72, 0x69, 0xf7, 0x0f, 0x30, 0xdc, 0x18,
	0x01, 0x32, 0x02, 0x7a, 0x0b, 0x80, 0x45, 0xc9, 0x33, 0x5f, 0x4d, 0xdd, 0x3c, 0x33, 0x56, 0x3b,
	0x16, 0xe9, 0x83, 0xef, 0x0d, 0x74, 0xfc, 0x25, 0x38, 0xa0, 0x7c, 0xe3, 0xb8, 0xb1, 0x38, 0x45,
	0x07, 0x71, 0x42, 0x7a, 0xaf, 0x9f, 0x0c, 0x0f, 0x69, 0x5d, 0xe2, 0x0b, 0xd4, 0xd6, 0xdc, 0x28,
	0x90, 0x86, 0xe7, 0x62, 0x96, 0x26, 0xfe, 0x14, 0xd5, 0xad, 0xc9, 0x0c, 0xbf, 0x46, 0x47, 0x1b,
	0xc7, 0xf5, 0x6d, 0xae, 0xb9, 0x71, 0xa5, 0x4d, 0x1b, 0xfd, 0x64, 0xd8, 0x1e, 0x3d, 0x21, 0x7f,
	0x3e, 0x0a, 0xf9, 0x54, 0xa1, 0xa8, 0x07, 0xd1, 0xf6, 0xe6, 0xae, 0xc0, 0x1a, 0x9d, 0x83, 0x16,
	0x0b, 0x21, 0x8b, 0x32, 0x9f, 0x71, 0xcb, 0x99, 0xcd, 0x83, 0xdd, 0x5c, 0x07, 0x79, 0x69, 0xd3,
	0xcf, 0x7c, 0xb6, 0x3b, 0xf3, 0x63, 0xe4, 0xbd, 0xf5, 0xb4, 0x89, 0x67, 0x45, 0x4f, 0xb4, 0x0b,
	0xfb, 0x0f, 0x07, 0x3f, 0x1a, 0xa8, 0xf3, 0xfb, 0x0d, 0x82, 0x17, 0xfc, 0x14, 0x9d, 0xcc, 0x5d,
	0x39, 0x17, 0x65, 0xb9, 0xae, 0x2e, 0xaf, 0x02, 0x89, 0x7e, 0x3b, 0x5b, 0xfd, 0xcf, 0xfc, 0xc6,
	0xe2, 0x2b, 0x74, 0xba, 0x0d, 0x5d, 0x73, 0x63, 0x8a, 0x05, 0x37, 0x69, 0xa3, 0xdf, 0x1c, 0xb6,
	0x47, 0xfd, 0x5d, 0xa1, 0xe1, 0x76, 0xf2, 0x3e, 0x00, 0xe9, 0x83, 0x2d, 0x76, 0xec, 0x19, 0xfc,
	0x10, 0xb5, 0x0c, 0x38, 0xcd, 0xb8, 0xf7, 0x7b, 0x48, 0x63, 0x85, 0x9f, 0xa3, 0x03, 0x55, 0xdc,
	0x96, 0x50, 0xcc, 0x7c, 0x38, 0xed, 0xd1, 0x23, 0x12, 0xf2, 0x25, 0x75, 0xbe, 0xe4, 0xca, 0xe7,
	0x4b, 0x6b, 0x1c, 0x1e, 0xa3, 0x0e, 0x38, 0xab, 0x9c, 0xcd, 0xe3, 0x9f, 0x65, 0xd2, 0xfb, 0x5e,
	0xda, 0xe3, 0x5d, 0x69, 0x6f, 0x02, 0x82, 0x1e, 0x07, 0x46, 0x2c, 0x0d, 0xfe, 0x80, 0x4e, 0xe7,
	0x50, 0x96, 0x70, 0xed, 0x54, 0xce, 0xbf, 0x55, 0x36, 0x85, 0x54, 0xce, 0xa6, 0x2d, 0xaf, 0xa1,
	0xb7, 0x3b, 0xe8, 0x5d, 0x05, 0x9a, 0x54, 0x18, 0x8a, 0x6b, 0xe6, 0x5d, 0x6f, 0xb0, 0x44, 0xdd,
	0x7f, 0xc4, 0xb5, 0xe5, 0x3e, 0xd9, 0xe7, 0xbe, 0xf9, 0x7f, 0xee, 0xc7, 0x2b, 0xd4, 0x65, 0xb0,
	0xae, 0x05, 0xfa, 0x5d, 0xd9, 0xd2, 0x39, 0x3e, 0x8a, 0xc1, 0x5f, 0x56, 0xfc, 0xcb, 0xe4, 0xeb,
	0xcb, 0x85, 0xb0, 0x4b, 0x37, 0x25, 0x0c, 0xd6, 0x99, 0xe7, 0x64, 0xf5, 0x7e, 0xed, 0x5b, 0xfc,
	0x57, 0xf1, 0xfb, 0x33, 0x49, 0xa6, 0x2d, 0x2f, 0xe3, 0xc5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2b, 0x40, 0x42, 0x22, 0x27, 0x04, 0x00, 0x00,
}
