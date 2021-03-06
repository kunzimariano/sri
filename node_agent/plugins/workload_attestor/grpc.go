package workloadattestor

import (
	common "github.com/spiffe/sri/node_agent/plugins/common/proto"
	"github.com/spiffe/sri/node_agent/plugins/workload_attestor/proto"
	"golang.org/x/net/context"
)

type GRPCServer struct {
	WorkloadAttestorImpl WorkloadAttestor
}

func (m *GRPCServer) Attest(ctx context.Context, req *node_agent_proto.AttestRequest) (*node_agent_proto.AttestResponse, error) {
	response, err := m.WorkloadAttestorImpl.Attest(req)
	return response, err
}

func (m *GRPCServer) Configure(ctx context.Context, req *common.ConfigureRequest) (*common.ConfigureResponse, error) {
	response, err := m.WorkloadAttestorImpl.Configure(req)
	return response, err
}

func (m *GRPCServer) GetPluginInfo(ctx context.Context, req *common.GetPluginInfoRequest) (*common.GetPluginInfoResponse, error) {
	response, err := m.WorkloadAttestorImpl.GetPluginInfo(req)
	return response, err
}

type GRPCClient struct {
	client node_agent_proto.WorkloadAttestorClient
}

func (m *GRPCClient) Attest(req *node_agent_proto.AttestRequest) (*node_agent_proto.AttestResponse, error) {
	res, err := m.client.Attest(context.Background(), req)
	return res, err
}

func (m *GRPCClient) Configure(req *common.ConfigureRequest) (*common.ConfigureResponse, error) {
	res, err := m.client.Configure(context.Background(), req)
	return res, err
}

func (m *GRPCClient) GetPluginInfo(req *common.GetPluginInfoRequest) (*common.GetPluginInfoResponse, error) {
	res, err := m.client.GetPluginInfo(context.Background(), req)
	return res, err
}
