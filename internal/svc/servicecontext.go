package svc

import (
	post_db "github.com/xh-polaris/meowchat-post-rpc/database/post-db"
	postTag_db "github.com/xh-polaris/meowchat-post-rpc/database/postTag-db"
	tag_db "github.com/xh-polaris/meowchat-post-rpc/database/tag-db"
	"github.com/xh-polaris/meowchat-post-rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DataBasex struct {
	PostDB    post_db.PostModel
	TagDB     tag_db.TagModel
	PostTagDB postTag_db.PostTagModel
}

type ServiceContext struct {
	Config   config.Config
	DataBase DataBasex
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DataBase: DataBasex{
			PostDB:    post_db.NewPostModel(sqlx.NewMysql(c.DBTable["post"]), c.CacheRedis),
			TagDB:     tag_db.NewTagModel(sqlx.NewMysql(c.DBTable["tag"]), c.CacheRedis),
			PostTagDB: postTag_db.NewPostTagModel(sqlx.NewMysql(c.DBTable["post_tag"]), c.CacheRedis),
		},
	}
}
