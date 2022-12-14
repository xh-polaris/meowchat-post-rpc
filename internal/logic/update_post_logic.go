package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-post-rpc/errorx"
	"github.com/xh-polaris/meowchat-post-rpc/internal/model"
	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePostLogic) UpdatePost(in *pb.UpdatePostReq) (*pb.UpdatePostResp, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, errorx.ErrInvalidObjectId
	}
	err = l.svcCtx.PostModel.Update(l.ctx, &model.Post{
		ID:       oid,
		Title:    in.Title,
		Text:     in.Text,
		CoverUrl: in.CoverUrl,
		Tags:     in.Tags,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePostResp{}, nil
}
