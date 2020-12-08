package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/micro/micro/v3/service/errors"

	pb "go.imgur.com/comments/proto"
)

// Comment struct.
type Comment struct{}

// NewComment struct.
func NewComment() *Comment {
	return &Comment{}
}

// New handler
func (c *Comment) New(ctx context.Context, req *pb.NewRequest, rsp *pb.NewResponse) error {
	return errors.InternalServerError("request.id.1", "failed to create new comment")
}

// List handler
func (c *Comment) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
	rsp.Comments = []*pb.CommentItem{
		{
			Id:      uuid.New().String(),
			Post:    "post1",
			Author:  "imgur",
			Message: "this is a comment",
			Created: 1604650806,
		},
	}
	return nil
}
