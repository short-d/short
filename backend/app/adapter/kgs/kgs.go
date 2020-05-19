package kgs

import (
	"context"

	"github.com/short-d/app/fw/rpc"
	"github.com/short-d/kgs/app/adapter/rpc/proto"
	"github.com/short-d/short/backend/app/usecase/keygen"
)

var _ keygen.KeyFetcher = (*RPC)(nil)

// RPC represents remote procedure calls which interact with key generation
// service.
type RPC struct {
	gRPCClient proto.KeyGenClient
}

// FetchKeys retrieves keys in batch from key generation service.
func (k RPC) FetchKeys(maxCount int) ([]keygen.Key, error) {
	req := proto.AllocateKeysRequest{
		MaxKeyCount: uint32(maxCount),
	}
	ctx := context.Background()
	res, err := k.gRPCClient.AllocateKeys(ctx, &req)

	if err != nil {
		return nil, err
	}

	keys := make([]keygen.Key, 0)
	for _, key := range res.Keys {
		keys = append(keys, keygen.Key(key))
	}
	return keys, nil
}

// NewRPC initializes GRPC client for key generation service APIs.
func NewRPC(hostname string, port int) (RPC, error) {
	connection, err := rpc.
		NewClientConnBuilder(hostname, port).
		InsecureTLS().
		Build()
	if err != nil {
		return RPC{}, err
	}
	gRPCClient := proto.NewKeyGenClient(connection)
	return RPC{gRPCClient: gRPCClient}, nil
}
