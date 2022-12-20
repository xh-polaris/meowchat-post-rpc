// Code generated by goctl. DO NOT EDIT!
// Source: post.proto

package server

import (
	"context"

	"postRpc/internal/logic"
	"postRpc/internal/svc"
	"postRpc/pb"
)

type PostRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPostRpcServer
}

func NewPostRpcServer(svcCtx *svc.ServiceContext) *PostRpcServer {
	return &PostRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *PostRpcServer) CreatePost(ctx context.Context, in *pb.CreatePostReq) (*pb.CreatePostResp, error) {
	l := logic.NewCreatePostLogic(ctx, s.svcCtx)
	return l.CreatePost(in)
}

func (s *PostRpcServer) RetrievePost(ctx context.Context, in *pb.RetrievePostReq) (*pb.RetrievePostResp, error) {
	l := logic.NewRetrievePostLogic(ctx, s.svcCtx)
	return l.RetrievePost(in)
}

func (s *PostRpcServer) UpdatePost(ctx context.Context, in *pb.UpdatePostReq) (*pb.UpdatePostResp, error) {
	l := logic.NewUpdatePostLogic(ctx, s.svcCtx)
	return l.UpdatePost(in)
}

func (s *PostRpcServer) DeletePost(ctx context.Context, in *pb.DeletePostReq) (*pb.DeletePostResp, error) {
	l := logic.NewDeletePostLogic(ctx, s.svcCtx)
	return l.DeletePost(in)
}

func (s *PostRpcServer) ListPost(ctx context.Context, in *pb.ListPostReq) (*pb.ListPostResp, error) {
	l := logic.NewListPostLogic(ctx, s.svcCtx)
	return l.ListPost(in)
}

func (s *PostRpcServer) ListPostByUserAndPvtAndStat(ctx context.Context, in *pb.ListPostByUserAndPvtAndStatReq) (*pb.ListPostByUserAndPvtAndStatResp, error) {
	l := logic.NewListPostByUserAndPvtAndStatLogic(ctx, s.svcCtx)
	return l.ListPostByUserAndPvtAndStat(in)
}
