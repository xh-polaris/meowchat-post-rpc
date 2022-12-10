package svc

import (
	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/config"
	"github.com/xh-polaris/meowchat-post-rpc/postrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	PostRpc postrpc.PostRPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		PostRpc: postrpc.NewPostRPC(zrpc.MustNewClient(c.PostRpc)),
	}
}
