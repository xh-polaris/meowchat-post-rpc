package logic

import (
	"context"
	"postRpc/internal/common"
	"postRpc/internal/svc"
	"postRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrievepostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrievePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrievepostLogic {
	return &RetrievepostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrievepostLogic) RetrievePost(in *pb.RetrievePostReq) (*pb.RetrievePostResp, error) {
	one, err := l.svcCtx.PostModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.RetrievePostResp{Post: common.PostTransform(one)}, nil
}
