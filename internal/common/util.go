package common

import (
	"github.com/xh-polaris/meowchat-post-rpc/internal/model"
	"github.com/xh-polaris/meowchat-post-rpc/pb"
)

func PostTransform(in *model.Post) *pb.Post {
	return &pb.Post{
		Id:       in.ID.Hex(),
		CreateAt: in.CreateAt.Unix(),
		UpdateAt: in.UpdateAt.Unix(),
		Title:    in.Title,
		Text:     in.Text,
		CoverUrl: in.CoverUrl,
		Tags:     in.Tags,
		UserId:   in.UserId,
	}
}
