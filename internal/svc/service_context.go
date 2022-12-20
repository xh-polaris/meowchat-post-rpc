package svc

import (
	"postRpc/internal/config"
	"postRpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	PostModel model.PostModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		PostModel: model.NewPostModel(c.MongoConf.Source, c.MongoConf.DataBase, c.MongoConf.CollPost, c.RedisConf),
	}
}
