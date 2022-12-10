package tag_db

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TagModel = (*customTagModel)(nil)

type (
	// TagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTagModel.
	TagModel interface {
		tagModel
		Insertx(data *Tag) (sql.Result, error)
		Deletex(data *Tag) error
		FindOnex(data *Tag) (*Tag, error)
	}

	customTagModel struct {
		*defaultTagModel
	}
)

// NewTagModel returns a model for the database table.
func NewTagModel(conn sqlx.SqlConn, c cache.CacheConf) TagModel {
	return &customTagModel{
		defaultTagModel: newTagModel(conn, c),
	}
}
