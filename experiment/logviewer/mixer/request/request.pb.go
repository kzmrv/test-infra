// Code generated by protoc-gen-go. DO NOT EDIT.
// source: experiment/logviewer/mixer/request/request.proto

package request

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	work "k8s.io/test-infra/experiment/logviewer/worker/work"
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

type MixerRequest struct {
	BuildNumber          int32                `protobuf:"varint,1,opt,name=buildNumber,proto3" json:"buildNumber,omitempty"`
	TargetSubstring      string               `protobuf:"bytes,2,opt,name=targetSubstring,proto3" json:"targetSubstring,omitempty"`
	FilePrefix           string               `protobuf:"bytes,3,opt,name=filePrefix,proto3" json:"filePrefix,omitempty"`
	Since                *timestamp.Timestamp `protobuf:"bytes,4,opt,name=since,proto3" json:"since,omitempty"`
	Until                *timestamp.Timestamp `protobuf:"bytes,5,opt,name=until,proto3" json:"until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MixerRequest) Reset()         { *m = MixerRequest{} }
func (m *MixerRequest) String() string { return proto.CompactTextString(m) }
func (*MixerRequest) ProtoMessage()    {}
func (*MixerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6237ed2f89e79ffe, []int{0}
}

func (m *MixerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MixerRequest.Unmarshal(m, b)
}
func (m *MixerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MixerRequest.Marshal(b, m, deterministic)
}
func (m *MixerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MixerRequest.Merge(m, src)
}
func (m *MixerRequest) XXX_Size() int {
	return xxx_messageInfo_MixerRequest.Size(m)
}
func (m *MixerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MixerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MixerRequest proto.InternalMessageInfo

func (m *MixerRequest) GetBuildNumber() int32 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *MixerRequest) GetTargetSubstring() string {
	if m != nil {
		return m.TargetSubstring
	}
	return ""
}

func (m *MixerRequest) GetFilePrefix() string {
	if m != nil {
		return m.FilePrefix
	}
	return ""
}

func (m *MixerRequest) GetSince() *timestamp.Timestamp {
	if m != nil {
		return m.Since
	}
	return nil
}

func (m *MixerRequest) GetUntil() *timestamp.Timestamp {
	if m != nil {
		return m.Until
	}
	return nil
}

type MixerResult struct {
	LogLines             []*work.LogLine `protobuf:"bytes,1,rep,name=logLines,proto3" json:"logLines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MixerResult) Reset()         { *m = MixerResult{} }
func (m *MixerResult) String() string { return proto.CompactTextString(m) }
func (*MixerResult) ProtoMessage()    {}
func (*MixerResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_6237ed2f89e79ffe, []int{1}
}

func (m *MixerResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MixerResult.Unmarshal(m, b)
}
func (m *MixerResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MixerResult.Marshal(b, m, deterministic)
}
func (m *MixerResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MixerResult.Merge(m, src)
}
func (m *MixerResult) XXX_Size() int {
	return xxx_messageInfo_MixerResult.Size(m)
}
func (m *MixerResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MixerResult.DiscardUnknown(m)
}

var xxx_messageInfo_MixerResult proto.InternalMessageInfo

func (m *MixerResult) GetLogLines() []*work.LogLine {
	if m != nil {
		return m.LogLines
	}
	return nil
}

func init() {
	proto.RegisterType((*MixerRequest)(nil), "request.MixerRequest")
	proto.RegisterType((*MixerResult)(nil), "request.MixerResult")
}

func init() {
	proto.RegisterFile("experiment/logviewer/mixer/request/request.proto", fileDescriptor_6237ed2f89e79ffe)
}

var fileDescriptor_6237ed2f89e79ffe = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdf, 0x7d, 0x6b, 0xab, 0x6e, 0x14, 0x61, 0x51, 0x08, 0x39, 0x68, 0xe8, 0x29, 0x22,
	0xcd, 0x96, 0x7a, 0xd1, 0xab, 0x78, 0x11, 0xaa, 0x48, 0x2a, 0x08, 0xde, 0x9a, 0x3a, 0x09, 0x43,
	0x37, 0xd9, 0x3a, 0xbb, 0xdb, 0xf6, 0xd3, 0xfa, 0x59, 0xa4, 0xf9, 0x23, 0xa5, 0x08, 0xbd, 0x64,
	0xc8, 0x6f, 0x9f, 0x99, 0xe1, 0x79, 0x86, 0x0f, 0x61, 0xbd, 0x00, 0xc2, 0x02, 0x4a, 0x2b, 0x95,
	0xce, 0x97, 0x08, 0x2b, 0x20, 0x59, 0xe0, 0x1a, 0x48, 0x12, 0x7c, 0x39, 0x30, 0xb6, 0xad, 0xf1,
	0x82, 0xb4, 0xd5, 0xe2, 0xb0, 0xf9, 0x0d, 0xae, 0x72, 0xad, 0x73, 0x05, 0xb2, 0xc2, 0xa9, 0xcb,
	0xa4, 0xc5, 0x02, 0x8c, 0x9d, 0x16, 0x8b, 0x5a, 0x19, 0xdc, 0xfc, 0x39, 0x7b, 0xa5, 0x69, 0xde,
	0x94, 0xea, 0x53, 0x8b, 0xfb, 0xdf, 0x8c, 0x9f, 0x3c, 0x6f, 0xd6, 0x26, 0xf5, 0x78, 0x11, 0x72,
	0x2f, 0x75, 0xa8, 0x3e, 0x5f, 0x5c, 0x91, 0x02, 0xf9, 0x2c, 0x64, 0x51, 0x37, 0xd9, 0x46, 0x22,
	0xe2, 0x67, 0x76, 0x4a, 0x39, 0xd8, 0x89, 0x4b, 0x8d, 0x25, 0x2c, 0x73, 0xff, 0x7f, 0xc8, 0xa2,
	0xe3, 0x64, 0x17, 0x8b, 0x4b, 0xce, 0x33, 0x54, 0xf0, 0x4a, 0x90, 0xe1, 0xda, 0xef, 0x54, 0xa2,
	0x2d, 0x22, 0x86, 0xbc, 0x6b, 0xb0, 0x9c, 0x81, 0x7f, 0x10, 0xb2, 0xc8, 0x1b, 0x05, 0x71, 0x6d,
	0x2d, 0x6e, 0xad, 0xc5, 0x6f, 0xad, 0xb5, 0xa4, 0x16, 0x6e, 0x3a, 0x5c, 0x69, 0x51, 0xf9, 0xdd,
	0xfd, 0x1d, 0x95, 0xb0, 0x7f, 0xc7, 0xbd, 0xc6, 0x9f, 0x71, 0xca, 0x8a, 0x6b, 0x7e, 0xa4, 0x74,
	0x3e, 0xc6, 0x12, 0x8c, 0xcf, 0xc2, 0x4e, 0xe4, 0x8d, 0x4e, 0xe3, 0x2a, 0x8e, 0x71, 0x4d, 0x93,
	0xdf, 0xe7, 0xd1, 0x53, 0x93, 0xcc, 0x04, 0x68, 0x89, 0x33, 0x10, 0xf7, 0xbc, 0xf7, 0xa8, 0xdf,
	0x35, 0xcd, 0xc5, 0x45, 0xdc, 0xde, 0x66, 0x3b, 0xba, 0xe0, 0x7c, 0x17, 0x6f, 0x36, 0xf6, 0xff,
	0x0d, 0xd9, 0x83, 0xfc, 0x18, 0x58, 0x30, 0x76, 0x80, 0x65, 0x46, 0x53, 0xb9, 0xff, 0xf6, 0x69,
	0xaf, 0x32, 0x74, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xb5, 0x17, 0xd8, 0x28, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MixerServiceClient is the client API for MixerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MixerServiceClient interface {
	DoWork(ctx context.Context, in *MixerRequest, opts ...grpc.CallOption) (MixerService_DoWorkClient, error)
}

type mixerServiceClient struct {
	cc *grpc.ClientConn
}

func NewMixerServiceClient(cc *grpc.ClientConn) MixerServiceClient {
	return &mixerServiceClient{cc}
}

func (c *mixerServiceClient) DoWork(ctx context.Context, in *MixerRequest, opts ...grpc.CallOption) (MixerService_DoWorkClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MixerService_serviceDesc.Streams[0], "/request.MixerService/DoWork", opts...)
	if err != nil {
		return nil, err
	}
	x := &mixerServiceDoWorkClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MixerService_DoWorkClient interface {
	Recv() (*MixerResult, error)
	grpc.ClientStream
}

type mixerServiceDoWorkClient struct {
	grpc.ClientStream
}

func (x *mixerServiceDoWorkClient) Recv() (*MixerResult, error) {
	m := new(MixerResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MixerServiceServer is the server API for MixerService service.
type MixerServiceServer interface {
	DoWork(*MixerRequest, MixerService_DoWorkServer) error
}

// UnimplementedMixerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMixerServiceServer struct {
}

func (*UnimplementedMixerServiceServer) DoWork(req *MixerRequest, srv MixerService_DoWorkServer) error {
	return status.Errorf(codes.Unimplemented, "method DoWork not implemented")
}

func RegisterMixerServiceServer(s *grpc.Server, srv MixerServiceServer) {
	s.RegisterService(&_MixerService_serviceDesc, srv)
}

func _MixerService_DoWork_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MixerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MixerServiceServer).DoWork(m, &mixerServiceDoWorkServer{stream})
}

type MixerService_DoWorkServer interface {
	Send(*MixerResult) error
	grpc.ServerStream
}

type mixerServiceDoWorkServer struct {
	grpc.ServerStream
}

func (x *mixerServiceDoWorkServer) Send(m *MixerResult) error {
	return x.ServerStream.SendMsg(m)
}

var _MixerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "request.MixerService",
	HandlerType: (*MixerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoWork",
			Handler:       _MixerService_DoWork_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "experiment/logviewer/mixer/request/request.proto",
}