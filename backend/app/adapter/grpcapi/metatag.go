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
	openGraphMetaTags, err := m.metaTag.GetOpenGraphTags(req.GetAlias())
	if err != nil {
		return &proto.GetOpenGraphTagsResponse{}, err
	}

	emptyString := ""
	if openGraphMetaTags.Title == nil {
		openGraphMetaTags.Title = &emptyString
	}
	if openGraphMetaTags.Description == nil {
		openGraphMetaTags.Description = &emptyString
	}
	if openGraphMetaTags.ImageURL == nil {
		openGraphMetaTags.ImageURL = &emptyString
	}

	return &proto.GetOpenGraphTagsResponse{
		Title:       *openGraphMetaTags.Title,
		Description: *openGraphMetaTags.Description,
		ImageUrl:    *openGraphMetaTags.ImageURL,
	}, nil
}

// GetTwitterTags fetches Twitter tags for a given short link.
func (m MetaTagServer) GetTwitterTags(ctx context.Context, req *proto.GetTwitterTagsRequest) (*proto.GetTwitterTagsResponse, error) {
	twitterMetaTags, err := m.metaTag.GetTwitterTags(req.GetAlias())
	if err != nil {
		return &proto.GetTwitterTagsResponse{}, err
	}

	emptyString := ""
	if twitterMetaTags.Title == nil {
		twitterMetaTags.Title = &emptyString
	}
	if twitterMetaTags.Description == nil {
		twitterMetaTags.Description = &emptyString
	}
	if twitterMetaTags.ImageURL == nil {
		twitterMetaTags.ImageURL = &emptyString
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
