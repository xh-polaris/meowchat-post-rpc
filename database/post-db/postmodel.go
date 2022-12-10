package post_db

import (
	"context"
	"database/sql"
	"github.com/xh-polaris/meowchat-post-rpc/pb/pb"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		Insertx(data *Post) (sql.Result, error)
		Deletex(id int64) error
		Updatex(data *Post) error
		FindOnex(id int64) (*Post, error)
		Listx(ctx context.Context, skip int64, count int64) ([]*pb.Post, error)
	}

	customPostModel struct {
		*defaultPostModel
	}
)

// NewPostModel returns a model for the database table.
func NewPostModel(conn sqlx.SqlConn, c cache.CacheConf) PostModel {
	return &customPostModel{
		defaultPostModel: newPostModel(conn, c),
	}
}
