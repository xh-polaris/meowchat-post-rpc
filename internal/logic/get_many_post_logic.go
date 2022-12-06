package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetManyPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetManyPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetManyPostLogic {
	return &GetManyPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetManyPostLogic) GetManyPost(in *pb.GetManyPostReq) (*pb.GetManyPostResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetManyPostResp{}, nil
}
