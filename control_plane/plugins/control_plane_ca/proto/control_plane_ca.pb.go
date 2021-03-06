// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control_plane_ca.proto

/*
Package control_plane_proto is a generated protocol buffer package.

It is generated from these files:
	control_plane_ca.proto

It has these top-level messages:
	SignCsrRequest
	SignCsrResponse
	GenerateCsrRequest
	GenerateCsrResponse
	FetchCertificateRequest
	FetchCertificateResponse
	LoadCertificateRequest
	LoadCertificateResponse
*/
package control_plane_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import control_plane_proto1 "github.com/spiffe/sri/control_plane/plugins/common/proto"

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

// ConfigureRequest from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type ConfigureRequest control_plane_proto1.ConfigureRequest

func (m *ConfigureRequest) Reset() { (*control_plane_proto1.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string {
	return (*control_plane_proto1.ConfigureRequest)(m).String()
}
func (*ConfigureRequest) ProtoMessage() {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*control_plane_proto1.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type ConfigureResponse control_plane_proto1.ConfigureResponse

func (m *ConfigureResponse) Reset() { (*control_plane_proto1.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string {
	return (*control_plane_proto1.ConfigureResponse)(m).String()
}
func (*ConfigureResponse) ProtoMessage() {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*control_plane_proto1.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type GetPluginInfoRequest control_plane_proto1.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset() { (*control_plane_proto1.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string {
	return (*control_plane_proto1.GetPluginInfoRequest)(m).String()
}
func (*GetPluginInfoRequest) ProtoMessage() {}

// GetPluginInfoResponse from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type GetPluginInfoResponse control_plane_proto1.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset() { (*control_plane_proto1.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).String()
}
func (*GetPluginInfoResponse) ProtoMessage() {}
func (m *GetPluginInfoResponse) GetName() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetName()
}
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetType()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetCompany()
}

// *Represents a request with a certificate signing request.
type SignCsrRequest struct {
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *SignCsrRequest) Reset()                    { *m = SignCsrRequest{} }
func (m *SignCsrRequest) String() string            { return proto.CompactTextString(m) }
func (*SignCsrRequest) ProtoMessage()               {}
func (*SignCsrRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignCsrRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

// *Represents a response with a signed certificate.
type SignCsrResponse struct {
	SignedCertificate []byte `protobuf:"bytes,1,opt,name=signedCertificate,proto3" json:"signedCertificate,omitempty"`
}

func (m *SignCsrResponse) Reset()                    { *m = SignCsrResponse{} }
func (m *SignCsrResponse) String() string            { return proto.CompactTextString(m) }
func (*SignCsrResponse) ProtoMessage()               {}
func (*SignCsrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignCsrResponse) GetSignedCertificate() []byte {
	if m != nil {
		return m.SignedCertificate
	}
	return nil
}

// *Represents an empty request.
type GenerateCsrRequest struct {
}

func (m *GenerateCsrRequest) Reset()                    { *m = GenerateCsrRequest{} }
func (m *GenerateCsrRequest) String() string            { return proto.CompactTextString(m) }
func (*GenerateCsrRequest) ProtoMessage()               {}
func (*GenerateCsrRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// *Represents a response with a certificate signing request.
type GenerateCsrResponse struct {
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *GenerateCsrResponse) Reset()                    { *m = GenerateCsrResponse{} }
func (m *GenerateCsrResponse) String() string            { return proto.CompactTextString(m) }
func (*GenerateCsrResponse) ProtoMessage()               {}
func (*GenerateCsrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GenerateCsrResponse) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

// *Represents an empty request.
type FetchCertificateRequest struct {
}

func (m *FetchCertificateRequest) Reset()                    { *m = FetchCertificateRequest{} }
func (m *FetchCertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchCertificateRequest) ProtoMessage()               {}
func (*FetchCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// *Represents a response with a stored intermediate certificate.
type FetchCertificateResponse struct {
	StoredIntermediateCert []byte `protobuf:"bytes,1,opt,name=storedIntermediateCert,proto3" json:"storedIntermediateCert,omitempty"`
}

func (m *FetchCertificateResponse) Reset()                    { *m = FetchCertificateResponse{} }
func (m *FetchCertificateResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchCertificateResponse) ProtoMessage()               {}
func (*FetchCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FetchCertificateResponse) GetStoredIntermediateCert() []byte {
	if m != nil {
		return m.StoredIntermediateCert
	}
	return nil
}

// *Represents a request with a signed intermediate certificate.
type LoadCertificateRequest struct {
	SignedIntermediateCert []byte `protobuf:"bytes,1,opt,name=signedIntermediateCert,proto3" json:"signedIntermediateCert,omitempty"`
}

func (m *LoadCertificateRequest) Reset()                    { *m = LoadCertificateRequest{} }
func (m *LoadCertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*LoadCertificateRequest) ProtoMessage()               {}
func (*LoadCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LoadCertificateRequest) GetSignedIntermediateCert() []byte {
	if m != nil {
		return m.SignedIntermediateCert
	}
	return nil
}

// *Represents an empty response.
type LoadCertificateResponse struct {
}

func (m *LoadCertificateResponse) Reset()                    { *m = LoadCertificateResponse{} }
func (m *LoadCertificateResponse) String() string            { return proto.CompactTextString(m) }
func (*LoadCertificateResponse) ProtoMessage()               {}
func (*LoadCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*SignCsrRequest)(nil), "control_plane_proto.SignCsrRequest")
	proto.RegisterType((*SignCsrResponse)(nil), "control_plane_proto.SignCsrResponse")
	proto.RegisterType((*GenerateCsrRequest)(nil), "control_plane_proto.GenerateCsrRequest")
	proto.RegisterType((*GenerateCsrResponse)(nil), "control_plane_proto.GenerateCsrResponse")
	proto.RegisterType((*FetchCertificateRequest)(nil), "control_plane_proto.FetchCertificateRequest")
	proto.RegisterType((*FetchCertificateResponse)(nil), "control_plane_proto.FetchCertificateResponse")
	proto.RegisterType((*LoadCertificateRequest)(nil), "control_plane_proto.LoadCertificateRequest")
	proto.RegisterType((*LoadCertificateResponse)(nil), "control_plane_proto.LoadCertificateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ControlPlaneCA service

type ControlPlaneCAClient interface {
	// * Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *control_plane_proto1.ConfigureRequest, opts ...grpc.CallOption) (*control_plane_proto1.ConfigureResponse, error)
	// * Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *control_plane_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*control_plane_proto1.GetPluginInfoResponse, error)
	// * Interface will take in a CSR and sign it with the stored intermediate certificate.
	SignCsr(ctx context.Context, in *SignCsrRequest, opts ...grpc.CallOption) (*SignCsrResponse, error)
	// * Used for generating a CSR for the intermediate signing certificate. The CSR will then be submitted to the CA plugin for signing.
	GenerateCsr(ctx context.Context, in *GenerateCsrRequest, opts ...grpc.CallOption) (*GenerateCsrResponse, error)
	// * Used to read the stored Intermediate CP cert.
	FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error)
	// * Used for setting/storing the signed intermediate certificate.
	LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error)
}

type controlPlaneCAClient struct {
	cc *grpc.ClientConn
}

func NewControlPlaneCAClient(cc *grpc.ClientConn) ControlPlaneCAClient {
	return &controlPlaneCAClient{cc}
}

func (c *controlPlaneCAClient) Configure(ctx context.Context, in *control_plane_proto1.ConfigureRequest, opts ...grpc.CallOption) (*control_plane_proto1.ConfigureResponse, error) {
	out := new(control_plane_proto1.ConfigureResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.ControlPlaneCA/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) GetPluginInfo(ctx context.Context, in *control_plane_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*control_plane_proto1.GetPluginInfoResponse, error) {
	out := new(control_plane_proto1.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.ControlPlaneCA/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) SignCsr(ctx context.Context, in *SignCsrRequest, opts ...grpc.CallOption) (*SignCsrResponse, error) {
	out := new(SignCsrResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.ControlPlaneCA/SignCsr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) GenerateCsr(ctx context.Context, in *GenerateCsrRequest, opts ...grpc.CallOption) (*GenerateCsrResponse, error) {
	out := new(GenerateCsrResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.ControlPlaneCA/GenerateCsr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error) {
	out := new(FetchCertificateResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.ControlPlaneCA/FetchCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error) {
	out := new(LoadCertificateResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.ControlPlaneCA/LoadCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ControlPlaneCA service

type ControlPlaneCAServer interface {
	// * Responsible for configuration of the plugin.
	Configure(context.Context, *control_plane_proto1.ConfigureRequest) (*control_plane_proto1.ConfigureResponse, error)
	// * Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *control_plane_proto1.GetPluginInfoRequest) (*control_plane_proto1.GetPluginInfoResponse, error)
	// * Interface will take in a CSR and sign it with the stored intermediate certificate.
	SignCsr(context.Context, *SignCsrRequest) (*SignCsrResponse, error)
	// * Used for generating a CSR for the intermediate signing certificate. The CSR will then be submitted to the CA plugin for signing.
	GenerateCsr(context.Context, *GenerateCsrRequest) (*GenerateCsrResponse, error)
	// * Used to read the stored Intermediate CP cert.
	FetchCertificate(context.Context, *FetchCertificateRequest) (*FetchCertificateResponse, error)
	// * Used for setting/storing the signed intermediate certificate.
	LoadCertificate(context.Context, *LoadCertificateRequest) (*LoadCertificateResponse, error)
}

func RegisterControlPlaneCAServer(s *grpc.Server, srv ControlPlaneCAServer) {
	s.RegisterService(&_ControlPlaneCA_serviceDesc, srv)
}

func _ControlPlaneCA_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(control_plane_proto1.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.ControlPlaneCA/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).Configure(ctx, req.(*control_plane_proto1.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(control_plane_proto1.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.ControlPlaneCA/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).GetPluginInfo(ctx, req.(*control_plane_proto1.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_SignCsr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCsrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).SignCsr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.ControlPlaneCA/SignCsr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).SignCsr(ctx, req.(*SignCsrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_GenerateCsr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateCsrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).GenerateCsr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.ControlPlaneCA/GenerateCsr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).GenerateCsr(ctx, req.(*GenerateCsrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_FetchCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).FetchCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.ControlPlaneCA/FetchCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).FetchCertificate(ctx, req.(*FetchCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_LoadCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).LoadCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.ControlPlaneCA/LoadCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).LoadCertificate(ctx, req.(*LoadCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlPlaneCA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "control_plane_proto.ControlPlaneCA",
	HandlerType: (*ControlPlaneCAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _ControlPlaneCA_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _ControlPlaneCA_GetPluginInfo_Handler,
		},
		{
			MethodName: "SignCsr",
			Handler:    _ControlPlaneCA_SignCsr_Handler,
		},
		{
			MethodName: "GenerateCsr",
			Handler:    _ControlPlaneCA_GenerateCsr_Handler,
		},
		{
			MethodName: "FetchCertificate",
			Handler:    _ControlPlaneCA_FetchCertificate_Handler,
		},
		{
			MethodName: "LoadCertificate",
			Handler:    _ControlPlaneCA_LoadCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control_plane_ca.proto",
}

func init() { proto.RegisterFile("control_plane_ca.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x8b, 0x5a, 0xb5, 0xea, 0xb4, 0x05, 0xba, 0x54, 0x94, 0xfa, 0x54, 0xb9, 0x7f, 0xa0,
	0x2d, 0xb5, 0xa5, 0x44, 0xca, 0x35, 0x8a, 0xac, 0x04, 0x21, 0xe5, 0x60, 0x91, 0x9c, 0x72, 0x21,
	0xc6, 0x1e, 0x9b, 0x95, 0xec, 0x5d, 0x67, 0x77, 0xfd, 0xa9, 0xf3, 0x25, 0x22, 0xdb, 0x0b, 0xc1,
	0xd8, 0x0e, 0xdc, 0x60, 0xe6, 0xcd, 0xef, 0xcd, 0x7a, 0x1e, 0x0c, 0x7d, 0xce, 0x94, 0xe0, 0xf1,
	0x32, 0x8d, 0x3d, 0x86, 0x4b, 0xdf, 0xb3, 0x52, 0xc1, 0x15, 0x27, 0x83, 0x6a, 0xbd, 0x28, 0x1a,
	0x97, 0x11, 0x55, 0xeb, 0x6c, 0x65, 0xf9, 0x3c, 0xb1, 0x65, 0x4a, 0xc3, 0x10, 0x6d, 0x29, 0xa8,
	0x5d, 0x91, 0xda, 0x69, 0x9c, 0x45, 0x94, 0x49, 0xdb, 0xe7, 0x49, 0xc2, 0x99, 0x5d, 0x4c, 0xea,
	0x3f, 0x25, 0xdb, 0x34, 0xa1, 0x7b, 0x43, 0x23, 0xe6, 0x48, 0xb1, 0xc0, 0x87, 0x0c, 0xa5, 0x22,
	0x7d, 0x78, 0xed, 0x4b, 0x31, 0xea, 0x7c, 0xef, 0x4c, 0x3e, 0x2e, 0xf2, 0x9f, 0xe6, 0x39, 0xf4,
	0xb6, 0x1a, 0x99, 0x72, 0x26, 0x91, 0x4c, 0xe1, 0xb3, 0xa4, 0x11, 0xc3, 0xc0, 0x41, 0xa1, 0x68,
	0x48, 0x7d, 0x4f, 0xa1, 0x1e, 0xa9, 0x37, 0xcc, 0x2f, 0x40, 0x66, 0xc8, 0x50, 0x78, 0x0a, 0x9f,
	0x8d, 0xcc, 0x31, 0x0c, 0x2a, 0x55, 0x8d, 0xae, 0xfb, 0x7f, 0x83, 0xaf, 0x57, 0xa8, 0xfc, 0xf5,
	0x0e, 0x72, 0xc3, 0x58, 0xc0, 0xa8, 0xde, 0xd2, 0xa0, 0x33, 0x18, 0x4a, 0xc5, 0x05, 0x06, 0x73,
	0xa6, 0x50, 0x24, 0x18, 0xd0, 0xdc, 0x09, 0x85, 0xd2, 0xec, 0x96, 0xae, 0xe9, 0xc2, 0xf0, 0x9a,
	0x7b, 0x41, 0xdd, 0xad, 0x20, 0x16, 0x8f, 0x6b, 0x25, 0x36, 0x76, 0xf3, 0x07, 0xd4, 0x88, 0xe5,
	0x92, 0x27, 0x8f, 0x6f, 0xa0, 0xeb, 0x94, 0x37, 0x73, 0xf3, 0x93, 0x39, 0x17, 0xe4, 0x0e, 0xde,
	0x3b, 0x9c, 0x85, 0x34, 0xca, 0x04, 0x92, 0x5f, 0x56, 0xc3, 0xf1, 0xad, 0x6d, 0x5f, 0x6f, 0x66,
	0xfc, 0x3e, 0x24, 0xd3, 0xdf, 0x24, 0x84, 0x4f, 0x33, 0x54, 0x6e, 0x11, 0x8b, 0x39, 0x0b, 0x39,
	0xf9, 0xd3, 0x38, 0x58, 0xd1, 0x6c, 0x3c, 0xfe, 0x1e, 0x23, 0xd5, 0x3e, 0xb7, 0xf0, 0x4e, 0x47,
	0x86, 0xfc, 0x68, 0x1c, 0xab, 0x86, 0xce, 0xf8, 0xf9, 0xb2, 0x48, 0x53, 0xef, 0xe1, 0xc3, 0x4e,
	0x62, 0xc8, 0xb8, 0x65, 0xa1, 0xfd, 0xa4, 0x19, 0x93, 0xc3, 0x42, 0xed, 0xc0, 0xa1, 0xbf, 0x9f,
	0x27, 0x32, 0x6d, 0x9c, 0x6e, 0x49, 0xa4, 0xf1, 0xff, 0x48, 0xb5, 0x36, 0x8c, 0xa1, 0xb7, 0x17,
	0x0d, 0xf2, 0xaf, 0x91, 0xd0, 0x1c, 0x49, 0x63, 0x7a, 0x9c, 0xb8, 0x74, 0x73, 0x5f, 0xad, 0xde,
	0x16, 0x82, 0xd3, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xa4, 0xc3, 0x4f, 0x6c, 0x04, 0x00,
	0x00,
}
