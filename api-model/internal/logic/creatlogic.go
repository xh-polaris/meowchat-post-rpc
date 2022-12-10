package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatLogic {
	return &CreatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatLogic) Creat(req *types.CreatReq) (resp *types.CreatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
