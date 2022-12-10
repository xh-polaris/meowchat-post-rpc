// Code generated by goctl. DO NOT EDIT!
// Source: post.proto

package server

import (
	"context"

	"github.com/xh-polaris/meowchat-post-rpc/internal/logic"
	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/pb/pb"
)

type PostRPCServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPostRPCServer
}

func NewPostRPCServer(svcCtx *svc.ServiceContext) *PostRPCServer {
	return &PostRPCServer{
		svcCtx: svcCtx,
	}
}

func (s *PostRPCServer) List(ctx context.Context, in *pb.ListReq) (*pb.ListResp, error) {
	l := logic.NewListLogic(ctx, s.svcCtx)
	return l.List(in)
}

func (s *PostRPCServer) Retrieve(ctx context.Context, in *pb.RetrieveReq) (*pb.RetrieveResp, error) {
	l := logic.NewRetrieveLogic(ctx, s.svcCtx)
	return l.Retrieve(in)
}

func (s *PostRPCServer) Creat(ctx context.Context, in *pb.CreatReq) (*pb.CreatResp, error) {
	l := logic.NewCreatLogic(ctx, s.svcCtx)
	return l.Creat(in)
}

func (s *PostRPCServer) Update(ctx context.Context, in *pb.UpdateReq) (*pb.UpdateResp, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *PostRPCServer) Destroy(ctx context.Context, in *pb.DestroyReq) (*pb.DestroyResp, error) {
	l := logic.NewDestroyLogic(ctx, s.svcCtx)
	return l.Destroy(in)
}
