package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"postRpc/internal/svc"
	"postRpc/pb"
)

type UpdatepostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatepostLogic {
	return &UpdatepostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatepostLogic) UpdatePost(in *pb.UpdatePostReq) (*pb.UpdatePostResp, error) {

	data, err := l.svcCtx.PostModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	data.IsAnonymous = in.IsAnonymous
	data.Title = in.Title
	data.Text = in.Text
	data.CoverUrl = in.CoverUrl
	data.Tags = in.Tags
	data.Status = in.Status
	err = l.svcCtx.PostModel.Update(l.ctx, data)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePostResp{}, nil
}
