// Code generated by protoc-gen-go.
// source: actionplugin.proto
// DO NOT EDIT!

/*
Package actionplugin is a generated protocol buffer package.

It is generated from these files:
	actionplugin.proto

It has these top-level messages:
	ActionPluginManifest
	ActionQuery
	ActionResult
	WorkerHTTPPortQuery
*/
package actionplugin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type ActionPluginManifest struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version     string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Author      string `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
}

func (m *ActionPluginManifest) Reset()                    { *m = ActionPluginManifest{} }
func (m *ActionPluginManifest) String() string            { return proto.CompactTextString(m) }
func (*ActionPluginManifest) ProtoMessage()               {}
func (*ActionPluginManifest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActionPluginManifest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActionPluginManifest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ActionPluginManifest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ActionPluginManifest) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type ActionQuery struct {
	Options map[string]string `protobuf:"bytes,1,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	JobID   int64             `protobuf:"varint,2,opt,name=jobID" json:"jobID,omitempty"`
}

func (m *ActionQuery) Reset()                    { *m = ActionQuery{} }
func (m *ActionQuery) String() string            { return proto.CompactTextString(m) }
func (*ActionQuery) ProtoMessage()               {}
func (*ActionQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ActionQuery) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *ActionQuery) GetJobID() int64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

type ActionResult struct {
	Status  string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
}

func (m *ActionResult) Reset()                    { *m = ActionResult{} }
func (m *ActionResult) String() string            { return proto.CompactTextString(m) }
func (*ActionResult) ProtoMessage()               {}
func (*ActionResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ActionResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ActionResult) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type WorkerHTTPPortQuery struct {
	Port int32 `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
}

func (m *WorkerHTTPPortQuery) Reset()                    { *m = WorkerHTTPPortQuery{} }
func (m *WorkerHTTPPortQuery) String() string            { return proto.CompactTextString(m) }
func (*WorkerHTTPPortQuery) ProtoMessage()               {}
func (*WorkerHTTPPortQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WorkerHTTPPortQuery) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*ActionPluginManifest)(nil), "actionplugin.ActionPluginManifest")
	proto.RegisterType((*ActionQuery)(nil), "actionplugin.ActionQuery")
	proto.RegisterType((*ActionResult)(nil), "actionplugin.ActionResult")
	proto.RegisterType((*WorkerHTTPPortQuery)(nil), "actionplugin.WorkerHTTPPortQuery")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ActionPlugin service

type ActionPluginClient interface {
	Manifest(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ActionPluginManifest, error)
	Run(ctx context.Context, in *ActionQuery, opts ...grpc.CallOption) (*ActionResult, error)
	WorkerHTTPPort(ctx context.Context, in *WorkerHTTPPortQuery, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type actionPluginClient struct {
	cc *grpc.ClientConn
}

func NewActionPluginClient(cc *grpc.ClientConn) ActionPluginClient {
	return &actionPluginClient{cc}
}

func (c *actionPluginClient) Manifest(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ActionPluginManifest, error) {
	out := new(ActionPluginManifest)
	err := grpc.Invoke(ctx, "/actionplugin.ActionPlugin/Manifest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionPluginClient) Run(ctx context.Context, in *ActionQuery, opts ...grpc.CallOption) (*ActionResult, error) {
	out := new(ActionResult)
	err := grpc.Invoke(ctx, "/actionplugin.ActionPlugin/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionPluginClient) WorkerHTTPPort(ctx context.Context, in *WorkerHTTPPortQuery, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/actionplugin.ActionPlugin/WorkerHTTPPort", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ActionPlugin service

type ActionPluginServer interface {
	Manifest(context.Context, *google_protobuf.Empty) (*ActionPluginManifest, error)
	Run(context.Context, *ActionQuery) (*ActionResult, error)
	WorkerHTTPPort(context.Context, *WorkerHTTPPortQuery) (*google_protobuf.Empty, error)
}

func RegisterActionPluginServer(s *grpc.Server, srv ActionPluginServer) {
	s.RegisterService(&_ActionPlugin_serviceDesc, srv)
}

func _ActionPlugin_Manifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionPluginServer).Manifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actionplugin.ActionPlugin/Manifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionPluginServer).Manifest(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionPlugin_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionPluginServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actionplugin.ActionPlugin/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionPluginServer).Run(ctx, req.(*ActionQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionPlugin_WorkerHTTPPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerHTTPPortQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionPluginServer).WorkerHTTPPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actionplugin.ActionPlugin/WorkerHTTPPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionPluginServer).WorkerHTTPPort(ctx, req.(*WorkerHTTPPortQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActionPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "actionplugin.ActionPlugin",
	HandlerType: (*ActionPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Manifest",
			Handler:    _ActionPlugin_Manifest_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _ActionPlugin_Run_Handler,
		},
		{
			MethodName: "WorkerHTTPPort",
			Handler:    _ActionPlugin_WorkerHTTPPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "actionplugin.proto",
}

func init() { proto.RegisterFile("actionplugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8b, 0xd3, 0x50,
	0x14, 0x6d, 0x26, 0x9d, 0x19, 0xbd, 0x2d, 0xa2, 0xd7, 0x61, 0x88, 0x71, 0x53, 0xdf, 0x42, 0xc7,
	0xcd, 0x1b, 0x18, 0x37, 0xd2, 0x85, 0xd4, 0x62, 0xa1, 0x82, 0xc5, 0x18, 0x0a, 0x82, 0xbb, 0x34,
	0x79, 0x4d, 0x63, 0xd3, 0xbc, 0xf0, 0x3e, 0x0a, 0xd9, 0xf8, 0x5f, 0xfc, 0x75, 0xfe, 0x0d, 0xc9,
	0x7b, 0xc9, 0x90, 0x40, 0xbb, 0x7b, 0xe7, 0xde, 0x73, 0x4f, 0xee, 0xb9, 0x27, 0x80, 0x51, 0xac,
	0x32, 0x5e, 0x94, 0xb9, 0x4e, 0xb3, 0x82, 0x96, 0x82, 0x2b, 0x8e, 0xe3, 0x6e, 0xcd, 0x7f, 0x9d,
	0x72, 0x9e, 0xe6, 0xec, 0xde, 0xf4, 0x36, 0x7a, 0x7b, 0xcf, 0x0e, 0xa5, 0xaa, 0x2c, 0x95, 0xfc,
	0x81, 0x9b, 0xcf, 0x86, 0x1c, 0x18, 0xf2, 0x2a, 0x2a, 0xb2, 0x2d, 0x93, 0x0a, 0x11, 0x86, 0x45,
	0x74, 0x60, 0x9e, 0x33, 0x71, 0xee, 0x9e, 0x86, 0xe6, 0x8d, 0x1e, 0x5c, 0x1f, 0x99, 0x90, 0x19,
	0x2f, 0xbc, 0x0b, 0x53, 0x6e, 0x21, 0x4e, 0x60, 0x94, 0x30, 0x19, 0x8b, 0xac, 0xac, 0xa5, 0x3c,
	0xd7, 0x74, 0xbb, 0x25, 0xbc, 0x85, 0xab, 0x48, 0xab, 0x1d, 0x17, 0xde, 0xd0, 0x34, 0x1b, 0x44,
	0xfe, 0x3a, 0x30, 0xb2, 0x0b, 0xfc, 0xd0, 0x4c, 0x54, 0x38, 0x83, 0x6b, 0x6e, 0x26, 0xa4, 0xe7,
	0x4c, 0xdc, 0xbb, 0xd1, 0xc3, 0x5b, 0xda, 0x33, 0xd8, 0xe1, 0xd2, 0xef, 0x96, 0xb8, 0x28, 0x94,
	0xa8, 0xc2, 0x76, 0x0c, 0x6f, 0xe0, 0xf2, 0x37, 0xdf, 0x7c, 0xfd, 0x62, 0x76, 0x74, 0x43, 0x0b,
	0xfc, 0x29, 0x8c, 0xbb, 0x74, 0x7c, 0x0e, 0xee, 0x9e, 0x55, 0x8d, 0xbd, 0xfa, 0x59, 0xcf, 0x1d,
	0xa3, 0x5c, 0xb3, 0xc6, 0x9b, 0x05, 0xd3, 0x8b, 0x8f, 0x0e, 0x99, 0xc1, 0xd8, 0x7e, 0x36, 0x64,
	0x52, 0xe7, 0xaa, 0xf6, 0x22, 0x55, 0xa4, 0xb4, 0x6c, 0xc6, 0x1b, 0x54, 0xdf, 0x27, 0x61, 0x2a,
	0xca, 0x72, 0xd9, 0xde, 0xa7, 0x81, 0xe4, 0x3d, 0xbc, 0xfc, 0xc9, 0xc5, 0x9e, 0x89, 0xe5, 0x7a,
	0x1d, 0x04, 0x5c, 0x28, 0x6b, 0x16, 0x61, 0x58, 0x72, 0xa1, 0x8c, 0xcc, 0x65, 0x68, 0xde, 0x0f,
	0xff, 0x9c, 0xf6, 0x6b, 0x36, 0x11, 0x5c, 0xc2, 0x93, 0xc7, 0x54, 0x6e, 0xa9, 0xcd, 0x92, 0xb6,
	0x59, 0xd2, 0x45, 0x9d, 0xa5, 0x4f, 0x4e, 0x1d, 0xa9, 0x9f, 0x28, 0x19, 0xe0, 0x27, 0x70, 0x43,
	0x5d, 0xe0, 0xab, 0xb3, 0x17, 0xf5, 0xfd, 0x53, 0x2d, 0xeb, 0x9a, 0x0c, 0x70, 0x05, 0xcf, 0xfa,
	0x2e, 0xf0, 0x4d, 0x9f, 0x7f, 0xc2, 0xa3, 0x7f, 0x66, 0x65, 0x32, 0x98, 0x7f, 0x83, 0x77, 0x31,
	0x3f, 0x50, 0x7e, 0xdc, 0xd1, 0x38, 0x91, 0x54, 0x26, 0x7b, 0x9a, 0x8a, 0x32, 0x6e, 0xb4, 0xba,
	0xc2, 0xf3, 0x17, 0x5d, 0x47, 0x41, 0x2d, 0x14, 0x38, 0xbf, 0x7a, 0x7f, 0xf9, 0xe6, 0xca, 0xe8,
	0x7f, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xf5, 0xc3, 0x76, 0x10, 0x03, 0x00, 0x00,
}
