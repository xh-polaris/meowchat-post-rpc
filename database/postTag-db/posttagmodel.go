package postTag_db

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostTagModel = (*customPostTagModel)(nil)

type (
	// PostTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostTagModel.
	PostTagModel interface {
		postTagModel
		Insertx(data *PostTag) (sql.Result, error)
		Deletex(postId string) error
		Findx(postId int64) ([]int, error)
	}

	customPostTagModel struct {
		*defaultPostTagModel
	}
)

// NewPostTagModel returns a model for the database table.
func NewPostTagModel(conn sqlx.SqlConn, c cache.CacheConf) PostTagModel {
	return &customPostTagModel{
		defaultPostTagModel: newPostTagModel(conn, c),
	}
}
