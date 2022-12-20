package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"postRpc/internal/model"
	"postRpc/internal/svc"
	"postRpc/pb"
)

type CreatepostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatepostLogic {
	return &CreatepostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatepostLogic) CreatePost(in *pb.CreatePostReq) (*pb.CreatePostResp, error) {
	var modelPost = model.Post{
		IsAnonymous: in.IsAnonymous,
		Title:       in.Title,
		Text:        in.Text,
		CoverUrl:    in.CoverUrl,
		Tags:        in.Tags,
		UserId:      in.UserId,
		Status:      in.Status,
	}
	err := l.svcCtx.PostModel.Insert(l.ctx, &modelPost)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePostResp{Id: modelPost.ID.Hex()}, nil
}
