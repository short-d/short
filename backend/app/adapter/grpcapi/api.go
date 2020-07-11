package grpcapi

import (
	"github.com/short-d/app/fw/rpc"
	"github.com/short-d/short/backend/app/adapter/grpcapi/proto"
	"google.golang.org/grpc"
)

var _ rpc.API = (*ShortGRPCApi)(nil)

// ShortGRPCApi provides an efficient way for remote systems to interact with Short backend.
type ShortGRPCApi struct {
	metaTagServer proto.MetaTagServiceServer
}

// RegisterServers registers gRPC servers that handle user requests.
func (s ShortGRPCApi) RegisterServers(server *grpc.Server) {
	proto.RegisterMetaTagServiceServer(server, s.metaTagServer)
}

// NewShortGRPCApi creates ShortGRPCApi.
func NewShortGRPCApi(metaTagServer proto.MetaTagServiceServer) ShortGRPCApi {
	return ShortGRPCApi{metaTagServer: metaTagServer}
}
