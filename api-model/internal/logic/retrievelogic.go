package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveLogic {
	return &RetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveLogic) Retrieve(req *types.RetrieveReq) (resp *types.RetrieveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
