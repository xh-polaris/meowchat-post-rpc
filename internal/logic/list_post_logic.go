package logic

import (
	"context"
	"postRpc/internal/common"

	"postRpc/internal/svc"
	"postRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListpostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListpostLogic {
	return &ListpostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListpostLogic) ListPost(in *pb.ListPostReq) (*pb.ListPostResp, error) {
	find, err := l.svcCtx.PostModel.Find(l.ctx, in.OrderBy, in.Skip, in.Limit)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.Post, 0, len(find))
	for _, val := range find {
		res = append(res, common.PostTransform(&val))
	}
	return &pb.ListPostResp{Posts: res}, nil
}
