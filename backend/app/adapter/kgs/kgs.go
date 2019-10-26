package kgs

import (
	"context"
	"crypto/tls"
	"fmt"
	"short/app/usecase/service"

	"github.com/byliuyang/kgs/app/adapter/rpc/proto"
	"github.com/byliuyang/kgs/app/entity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ service.KeyFetcher = (*RPC)(nil)

// RPC represents remote procedure calls which interact with key generation
// service.
type RPC struct {
	gRPCClient proto.KeyGenClient
}

// FetchKeys retrieves keys in batch from key generation service.
func (k RPC) FetchKeys(maxCount int) ([]entity.Key, error) {
	req := proto.AllocateKeysRequest{
		MaxKeyCount: uint32(maxCount),
	}
	ctx := context.Background()
	res, err := k.gRPCClient.AllocateKeys(ctx, &req)

	if err != nil {
		return nil, err
	}

	keys := make([]entity.Key, 0)
	for _, key := range res.Keys {
		keys = append(keys, entity.Key(key))
	}
	return keys, nil
}

// NewRPC initializes GRPC client for key generation service APIs.
func NewRPC(hostname string, port int) (RPC, error) {
	target := fmt.Sprintf("%s:%d", hostname, port)

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	gRPCTls := credentials.NewTLS(config)
	gRPCCredentials := grpc.WithTransportCredentials(gRPCTls)
	connection, err := grpc.Dial(target, gRPCCredentials)
	if err != nil {
		return RPC{}, err
	}
	gRPCClient := proto.NewKeyGenClient(connection)
	return RPC{gRPCClient: gRPCClient}, nil
}
