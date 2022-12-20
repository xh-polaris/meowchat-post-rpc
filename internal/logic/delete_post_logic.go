package logic

import (
	"context"

	"postRpc/internal/svc"
	"postRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletepostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletepostLogic {
	return &DeletepostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletepostLogic) DeletePost(in *pb.DeletePostReq) (*pb.DeletePostResp, error) {
	err := l.svcCtx.PostModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePostResp{}, nil
}
