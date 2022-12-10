package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DestroyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDestroyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DestroyLogic {
	return &DestroyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DestroyLogic) Destroy(req *types.DestroyReq) (resp *types.DestroyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
