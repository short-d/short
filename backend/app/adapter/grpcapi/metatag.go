package grpcapi

import (
	"context"

	"github.com/short-d/short/backend/app/adapter/grpcapi/proto"
	"github.com/short-d/short/backend/app/usecase/shortlink"
)

type MetaTagServer struct {
	metaTag shortlink.MetaTag
}

var _ proto.MetaTagServiceServer = (*MetaTagServer)(nil)

func (m MetaTagServer) GetOGTags(ctx context.Context, req *proto.GetOGTagsRequest) (*proto.GetOGTagsResponse, error) {
	ogMetaTags, err := m.metaTag.GetOpenGraphTags(req.GetAlias())
	if err != nil {
		return &proto.GetOGTagsResponse{}, err
	}

	return &proto.GetOGTagsResponse{
		Title:       *ogMetaTags.Title,
		Description: *ogMetaTags.Description,
		ImageUrl:    *ogMetaTags.ImageURL,
	}, nil
}

func (m MetaTagServer) GetTwitterTags(ctx context.Context, req *proto.GetTwitterTagsRequest) (*proto.GetTwitterTagsResponse, error) {
	twitterMetaTags, err := m.metaTag.GetTwitterTags(req.GetAlias())
	if err != nil {
		return &proto.GetTwitterTagsResponse{}, err
	}

	return &proto.GetTwitterTagsResponse{
		Title:       *twitterMetaTags.Title,
		Description: *twitterMetaTags.Description,
		ImageUrl:    *twitterMetaTags.ImageURL,
	}, nil
}

func NewMetaTagServer(metaTag shortlink.MetaTag) proto.MetaTagServiceServer {
	return MetaTagServer{metaTag: metaTag}
}
