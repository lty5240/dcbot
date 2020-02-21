// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/command.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
	context "golang.org/x/net/context"
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

type Command_CommandStatus int32

const (
	Command_ENABLE  Command_CommandStatus = 0
	Command_DISABLE Command_CommandStatus = 1
)

var Command_CommandStatus_name = map[int32]string{
	0: "ENABLE",
	1: "DISABLE",
}

var Command_CommandStatus_value = map[string]int32{
	"ENABLE":  0,
	"DISABLE": 1,
}

func (x Command_CommandStatus) String() string {
	return proto.EnumName(Command_CommandStatus_name, int32(x))
}

func (Command_CommandStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab3015d744d1176f, []int{0, 0}
}

type Command struct {
	Id                   int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Command              string                `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Service              string                `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	Status               Command_CommandStatus `protobuf:"varint,4,opt,name=status,proto3,enum=proto.Command_CommandStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab3015d744d1176f, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Command) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Command) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Command) GetStatus() Command_CommandStatus {
	if m != nil {
		return m.Status
	}
	return Command_ENABLE
}

type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab3015d744d1176f, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type CommandRequest struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandRequest) Reset()         { *m = CommandRequest{} }
func (m *CommandRequest) String() string { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()    {}
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab3015d744d1176f, []int{2}
}

func (m *CommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandRequest.Unmarshal(m, b)
}
func (m *CommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandRequest.Marshal(b, m, deterministic)
}
func (m *CommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandRequest.Merge(m, src)
}
func (m *CommandRequest) XXX_Size() int {
	return xxx_messageInfo_CommandRequest.Size(m)
}
func (m *CommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommandRequest proto.InternalMessageInfo

func (m *CommandRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

type ActionRequest struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Time                 string   `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab3015d744d1176f, []int{3}
}

func (m *ActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRequest.Unmarshal(m, b)
}
func (m *ActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRequest.Marshal(b, m, deterministic)
}
func (m *ActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRequest.Merge(m, src)
}
func (m *ActionRequest) XXX_Size() int {
	return xxx_messageInfo_ActionRequest.Size(m)
}
func (m *ActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRequest proto.InternalMessageInfo

func (m *ActionRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *ActionRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ActionRequest) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.Command_CommandStatus", Command_CommandStatus_name, Command_CommandStatus_value)
	proto.RegisterType((*Command)(nil), "proto.Command")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*CommandRequest)(nil), "proto.CommandRequest")
	proto.RegisterType((*ActionRequest)(nil), "proto.ActionRequest")
}

func init() { proto.RegisterFile("proto/command.proto", fileDescriptor_ab3015d744d1176f) }

var fileDescriptor_ab3015d744d1176f = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xee, 0xa6, 0xb1, 0xb1, 0x23, 0x8d, 0x32, 0x2a, 0x04, 0xf1, 0x10, 0xf6, 0x14, 0x3c, 0xb4,
	0x52, 0x3d, 0x0b, 0x55, 0x7b, 0x10, 0xd4, 0xc3, 0xe6, 0x09, 0x62, 0xb2, 0xca, 0x1e, 0xf2, 0x63,
	0x67, 0xeb, 0x5b, 0xf9, 0x8e, 0xd2, 0xcd, 0x44, 0xba, 0x20, 0x78, 0xda, 0xfd, 0xe6, 0x9b, 0xe1,
	0xfb, 0x81, 0xd3, 0x6e, 0xd3, 0xda, 0x76, 0x51, 0xb6, 0x75, 0x5d, 0x34, 0xd5, 0xdc, 0x21, 0x3c,
	0x70, 0x8f, 0xfc, 0x16, 0x10, 0x3d, 0xf4, 0x04, 0xc6, 0x10, 0x98, 0x2a, 0x11, 0xa9, 0xc8, 0xc6,
	0x2a, 0x30, 0x15, 0x26, 0x10, 0xf1, 0x4d, 0x12, 0xa4, 0x22, 0x9b, 0xaa, 0x01, 0xee, 0x18, 0xd2,
	0x9b, 0x2f, 0x53, 0xea, 0x64, 0xdc, 0x33, 0x0c, 0xf1, 0x16, 0x26, 0x64, 0x0b, 0xbb, 0xa5, 0x24,
	0x4c, 0x45, 0x16, 0x2f, 0x2f, 0x7b, 0xb9, 0x39, 0x6b, 0x0c, 0x6f, 0xee, 0x76, 0x14, 0xef, 0xca,
	0x0c, 0x66, 0x1e, 0x81, 0x00, 0x93, 0xf5, 0xeb, 0xea, 0xfe, 0x79, 0x7d, 0x32, 0xc2, 0x23, 0x88,
	0x1e, 0x9f, 0x72, 0x07, 0x84, 0xbc, 0x83, 0x43, 0xa5, 0xa9, 0x6b, 0x1b, 0xd2, 0x3b, 0x17, 0xb5,
	0x26, 0x2a, 0x3e, 0xb4, 0x33, 0x3d, 0x55, 0x03, 0xdc, 0xf7, 0x17, 0x78, 0xfe, 0xe4, 0x15, 0xc4,
	0xac, 0xa4, 0xf4, 0xe7, 0x56, 0x93, 0xdd, 0x4f, 0x29, 0xbc, 0x94, 0xf2, 0x05, 0x66, 0xab, 0xd2,
	0x9a, 0xb6, 0xf9, 0x77, 0xd5, 0x55, 0xd7, 0xb1, 0x56, 0x60, 0x3a, 0x44, 0x08, 0xad, 0xa9, 0x87,
	0x76, 0xdc, 0x7f, 0x49, 0xbf, 0xd2, 0x39, 0x97, 0x75, 0x0d, 0xe1, 0xbb, 0x69, 0x2a, 0x3c, 0xf7,
	0x4b, 0x62, 0xb9, 0x8b, 0x63, 0x1e, 0x0f, 0x81, 0xe5, 0x08, 0x17, 0x10, 0x92, 0x6e, 0x2a, 0x3c,
	0x63, 0xca, 0xf3, 0xf7, 0xc7, 0xc1, 0xdb, 0xc4, 0x4d, 0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb2, 0xf0, 0x2a, 0xce, 0x04, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CommandService service

type CommandServiceClient interface {
	Find(ctx context.Context, in *CommandRequest, opts ...client.CallOption) (*Response, error)
	Send(ctx context.Context, in *ActionRequest, opts ...client.CallOption) (*Response, error)
}

type commandServiceClient struct {
	c           client.Client
	serviceName string
}

func NewCommandServiceClient(serviceName string, c client.Client) CommandServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "proto"
	}
	return &commandServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *commandServiceClient) Find(ctx context.Context, in *CommandRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "CommandService.find", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) Send(ctx context.Context, in *ActionRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "CommandService.send", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CommandService service

type CommandServiceHandler interface {
	Find(context.Context, *CommandRequest, *Response) error
	Send(context.Context, *ActionRequest, *Response) error
}

func RegisterCommandServiceHandler(s server.Server, hdlr CommandServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&CommandService{hdlr}, opts...))
}

type CommandService struct {
	CommandServiceHandler
}

func (h *CommandService) Find(ctx context.Context, in *CommandRequest, out *Response) error {
	return h.CommandServiceHandler.Find(ctx, in, out)
}

func (h *CommandService) Send(ctx context.Context, in *ActionRequest, out *Response) error {
	return h.CommandServiceHandler.Send(ctx, in, out)
}
