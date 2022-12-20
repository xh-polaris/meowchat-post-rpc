package logic

import (
	"context"
	"postRpc/internal/common"

	"postRpc/internal/svc"
	"postRpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListpostbyuserandpvtandstatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPostByUserAndPvtAndStatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListpostbyuserandpvtandstatLogic {
	return &ListpostbyuserandpvtandstatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListpostbyuserandpvtandstatLogic) ListPostByUserAndPvtAndStat(in *pb.ListPostByUserAndPvtAndStatReq) (*pb.ListPostByUserAndPvtAndStatResp, error) {
	find, err := l.svcCtx.PostModel.FindByUserAndPvtAndStatReq(l.ctx, in.UserId, in.IsAnonymous, in.Status, in.Skip, in.Limit)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.Post, 0, len(find))
	for _, val := range find {
		res = append(res, common.PostTransform(&val))
	}
	return &pb.ListPostByUserAndPvtAndStatResp{Posts: res}, nil
}
