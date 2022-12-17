package common

import (
	"postRpc/internal/model"
	"postRpc/pb"
)

func PostTransform(in *model.Post) *pb.Post {
	return &pb.Post{
		Id:          in.ID.Hex(),
		CreateAt:    in.CreateAt.Unix(),
		UpdateAt:    in.UpdateAt.Unix(),
		IsAnonymous: in.IsAnonymous,
		Title:       in.Title,
		Text:        in.Text,
		CoverUrl:    in.CoverUrl,
		Tags:        in.Tags,
		UserId:      in.UserId,
		Status:      in.Status,
	}
}
