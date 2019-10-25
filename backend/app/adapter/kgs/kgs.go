package kgs

import (
	"context"
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc/credentials"
	"short/app/usecase/service"

	"github.com/byliuyang/kgs/app/entity"

	"github.com/byliuyang/kgs/app/adapter/rpc/proto"
	"google.golang.org/grpc"
)

var _ service.KeyGen = (*Rpc)(nil)

type Rpc struct {
	gRpcClient proto.KeyGenClient
}

func (k Rpc) FetchKeys(maxCount int) ([]entity.Key, error) {
	req := proto.AllocateKeysRequest{
		MaxKeyCount: uint32(maxCount),
	}
	ctx := context.Background()
	res, err := k.gRpcClient.AllocateKeys(ctx, &req)

	if err != nil {
		return nil, err
	}

	keys := make([]entity.Key, 0)
	for _, key := range res.Keys {
		keys = append(keys, entity.Key(key))
	}
	return keys, nil
}

func NewRpc(hostname string, port int) (Rpc, error) {
	target := fmt.Sprintf("%s:%d", hostname, port)

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	gRpcTLS := credentials.NewTLS(config)
	gRpcCredentials := grpc.WithTransportCredentials(gRpcTLS)
	connection, err := grpc.Dial(target, gRpcCredentials)
	if err != nil {
		return Rpc{}, err
	}
	gRpcClient := proto.NewKeyGenClient(connection)
	return Rpc{gRpcClient: gRpcClient}, nil
}
