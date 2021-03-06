package node

import (
	"context"
	"errors"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/spiffe/sri/control_plane/api/node/proto"
	oldcontext "golang.org/x/net/context"
)

type grpcServer struct {
	fetchBaseSVID        grpctransport.Handler
	fetchSVID            grpctransport.Handler
	fetchCPBundle        grpctransport.Handler
	fetchFederatedBundle grpctransport.Handler
}

// MakeGRPCServer makes a set of endpoints available as a gRPC server.
func MakeGRPCServer(endpoints Endpoints) (req control_plane_proto.NodeServer) {
	req = &grpcServer{
		fetchBaseSVID: grpctransport.NewServer(
			endpoints.FetchBaseSVIDEndpoint,
			DecodeGRPCFetchBaseSVIDRequest,
			EncodeGRPCFetchBaseSVIDResponse,
		),

		fetchSVID: grpctransport.NewServer(
			endpoints.FetchSVIDEndpoint,
			DecodeGRPCFetchSVIDRequest,
			EncodeGRPCFetchSVIDResponse,
		),

		fetchCPBundle: grpctransport.NewServer(
			endpoints.FetchCPBundleEndpoint,
			DecodeGRPCFetchCPBundleRequest,
			EncodeGRPCFetchCPBundleResponse,
		),

		fetchFederatedBundle: grpctransport.NewServer(
			endpoints.FetchFederatedBundleEndpoint,
			DecodeGRPCFetchFederatedBundleRequest,
			EncodeGRPCFetchFederatedBundleResponse,
		),
	}
	return req
}

// DecodeGRPCFetchBaseSVIDRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCFetchBaseSVIDRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'FetchBaseSVID' Decoder is not impelement")
	return req, err
}

// EncodeGRPCFetchBaseSVIDResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCFetchBaseSVIDResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'FetchBaseSVID' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) FetchBaseSVID(ctx oldcontext.Context, req *control_plane_proto.FetchBaseSVIDRequest) (rep *control_plane_proto.FetchBaseSVIDResponse, err error) {
	_, rp, err := s.fetchBaseSVID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*control_plane_proto.FetchBaseSVIDResponse)
	return rep, err
}

// DecodeGRPCFetchSVIDRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCFetchSVIDRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'FetchSVID' Decoder is not impelement")
	return req, err
}

// EncodeGRPCFetchSVIDResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCFetchSVIDResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'FetchSVID' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) FetchSVID(ctx oldcontext.Context, req *control_plane_proto.FetchSVIDRequest) (rep *control_plane_proto.FetchSVIDResponse, err error) {
	_, rp, err := s.fetchSVID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*control_plane_proto.FetchSVIDResponse)
	return rep, err
}

// DecodeGRPCFetchCPBundleRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCFetchCPBundleRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'FetchCPBundle' Decoder is not impelement")
	return req, err
}

// EncodeGRPCFetchCPBundleResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCFetchCPBundleResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'FetchCPBundle' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) FetchCPBundle(ctx oldcontext.Context, req *control_plane_proto.FetchCPBundleRequest) (rep *control_plane_proto.FetchCPBundleResponse, err error) {
	_, rp, err := s.fetchCPBundle.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*control_plane_proto.FetchCPBundleResponse)
	return rep, err
}

// DecodeGRPCFetchFederatedBundleRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCFetchFederatedBundleRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'FetchFederatedBundle' Decoder is not impelement")
	return req, err
}

// EncodeGRPCFetchFederatedBundleResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCFetchFederatedBundleResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'FetchFederatedBundle' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) FetchFederatedBundle(ctx oldcontext.Context, req *control_plane_proto.FetchFederatedBundleRequest) (rep *control_plane_proto.FetchFederatedBundleResponse, err error) {
	_, rp, err := s.fetchFederatedBundle.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*control_plane_proto.FetchFederatedBundleResponse)
	return rep, err
}
