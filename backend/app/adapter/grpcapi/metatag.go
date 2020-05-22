package grpcapi

import (
	"context"

	"github.com/short-d/short/backend/app/adapter/grpcapi/proto"
	"github.com/short-d/short/backend/app/usecase/url"
)

type MetaTagServer struct {
	metaTag url.MetaTag
}

var _ proto.MetaTagServiceServer = (*MetaTagServer)(nil)

func (m MetaTagServer) UpdateOGTags(ctx context.Context, req *proto.UpdateOGTagsRequest) (*proto.UpdateOGTagsResponse, error) {
	ogMetaTags, err := m.metaTag.UpdateOGTags(req.GetAlias(), req.GetTitle(), req.GetDescription(), req.GetImageUrl())
	if err != nil {
		return &proto.UpdateOGTagsResponse{}, err
	}

	return &proto.UpdateOGTagsResponse{
		Title:       ogMetaTags.OGTitle,
		Description: ogMetaTags.OGDescription,
		ImageUrl:    ogMetaTags.OGImageURL,
	}, nil
}

func (m MetaTagServer) UpdateTwitterTags(ctx context.Context, req *proto.UpdateTwitterTagsRequest) (*proto.UpdateTwitterTagsResponse, error) {
	twitterMetaTags, err := m.metaTag.UpdateTwitterTags(req.GetAlias(), req.GetTitle(), req.GetDescription(), req.GetImageUrl())
	if err != nil {
		return &proto.UpdateTwitterTagsResponse{}, err
	}

	return &proto.UpdateTwitterTagsResponse{
		Title:       twitterMetaTags.TwitterTitle,
		Description: twitterMetaTags.TwitterDescription,
		ImageUrl:    twitterMetaTags.TwitterImageURL,
	}, nil
}

func (m MetaTagServer) GetOGTags(ctx context.Context, req *proto.GetOGTagsRequest) (*proto.GetOGTagsResponse, error) {
	ogMetaTags, err := m.metaTag.GetOGTags(req.GetAlias())
	if err != nil {
		return &proto.GetOGTagsResponse{}, err
	}

	return &proto.GetOGTagsResponse{
		Title:       ogMetaTags.OGTitle,
		Description: ogMetaTags.OGDescription,
		ImageUrl:    ogMetaTags.OGImageURL,
	}, nil
}

func (m MetaTagServer) GetTwitterTags(ctx context.Context, req *proto.GetTwitterTagsRequest) (*proto.GetTwitterTagsResponse, error) {
	twitterMetaTags, err := m.metaTag.GetTwitterTags(req.GetAlias())
	if err != nil {
		return &proto.GetTwitterTagsResponse{}, err
	}

	return &proto.GetTwitterTagsResponse{
		Title:       twitterMetaTags.TwitterTitle,
		Description: twitterMetaTags.TwitterDescription,
		ImageUrl:    twitterMetaTags.TwitterImageURL,
	}, nil
}

func NewMetaTagServer(metaTag url.MetaTag) proto.MetaTagServiceServer {
	return MetaTagServer{metaTag: metaTag}
}
