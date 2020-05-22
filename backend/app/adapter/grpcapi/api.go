package grpcapi

import (
	"github.com/short-d/app/fw/rpc"
	"github.com/short-d/short/backend/app/adapter/grpcapi/proto"
	"google.golang.org/grpc"
)

var _ rpc.API = (*ShortGRPCApi)(nil)

type ShortGRPCApi struct {
	metaTagServer proto.MetaTagServiceServer
}

func (s ShortGRPCApi) RegisterServers(server *grpc.Server) {
	proto.RegisterMetaTagServiceServer(server, s.metaTagServer)
}

func NewShortGRPCApi(metaTagServer proto.MetaTagServiceServer) ShortGRPCApi {
	return ShortGRPCApi{metaTagServer: metaTagServer}
}
