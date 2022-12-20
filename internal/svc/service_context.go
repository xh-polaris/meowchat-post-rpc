package svc

import (
	"github.com/xh-polaris/meowchat-post-rpc/internal/config"
	"github.com/xh-polaris/meowchat-post-rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	PostModel model.PostModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		PostModel: model.NewPostModel(c.Mongo.URL, c.Mongo.DB, c.Cache, c.Elasticsearch),
	}
}
