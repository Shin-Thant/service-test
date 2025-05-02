package internal

import (
	"context"

	"github.com/Shin-Thant/service-test/internal/pb/feed/v1"
)

type FeedServiceHandler struct {
	feed.UnimplementedFeedServiceServer
}

func NewFeedServiceHandler() *FeedServiceHandler {
	return &FeedServiceHandler{}
}

func (s *FeedServiceHandler) FindOne(context.Context, *feed.FindOneRequest) (*feed.FindOneResponse, error) {
	return &feed.FindOneResponse{
		Id:   1,
		Name: "Shin Thant",
	}, nil
}
