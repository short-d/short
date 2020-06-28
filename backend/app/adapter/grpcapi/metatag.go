package grpcapi

import (
	"context"

	"github.com/short-d/short/backend/app/adapter/grpcapi/proto"
	"github.com/short-d/short/backend/app/usecase/shortlink"
)

// MetaTagServer allows the client to retrieve the SEO meta tags for a short link.
type MetaTagServer struct {
	metaTag shortlink.MetaTag
}

var _ proto.MetaTagServiceServer = (*MetaTagServer)(nil)

// GetOpenGraphTags fetches Open Graph tags for a given short link.
func (m MetaTagServer) GetOpenGraphTags(ctx context.Context, req *proto.GetOpenGraphTagsRequest) (*proto.GetOpenGraphTagsResponse, error) {
	ogMetaTags, err := m.metaTag.GetOpenGraphTags(req.GetAlias())
	if err != nil {
		return &proto.GetOpenGraphTagsResponse{}, err
	}

	return &proto.GetOpenGraphTagsResponse{
		Title:       *ogMetaTags.Title,
		Description: *ogMetaTags.Description,
		ImageUrl:    *ogMetaTags.ImageURL,
	}, nil
}

// GetTwitterTags fetches Twitter tags for a given short link.
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

// NewMetaTagServer creates MetaTag gRPC server
func NewMetaTagServer(metaTag shortlink.MetaTag) proto.MetaTagServiceServer {
	return MetaTagServer{metaTag: metaTag}
}
