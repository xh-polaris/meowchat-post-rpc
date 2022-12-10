package logic

import (
	"context"
	"strconv"

	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DestroyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDestroyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DestroyLogic {
	return &DestroyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DestroyLogic) Destroy(in *pb.DestroyReq) (*pb.DestroyResp, error) {
	destroyId, _ := strconv.ParseInt(in.PostId, 10, 64)
	err := l.svcCtx.DataBase.PostDB.Deletex(destroyId)
	if err != nil {
		return nil, err
	}
	return &pb.DestroyResp{}, nil
}
