package postTag_db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/xh-polaris/meowchat-post-rpc/database/error_type"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func (m *customPostTagModel) Insertx(data *PostTag) (sql.Result, error) {
	return m.Insert(context.Background(), data)
}
func (m *customPostTagModel) Deletex(data *PostTag) error {
	return m.Delete(context.Background(), data.Id)
}
func (m *customPostTagModel) Findx(postId int64) ([]int, error) {
	iyuaPostTagIdKey := fmt.Sprintf("%s%s%v", cacheIyuaPostTagIdPrefix, "postid:", postId)
	var resp []int
	err := m.QueryRowCtx(context.Background(), resp, iyuaPostTagIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select `tag_id` from %s where `post_id` = ?", m.table)
		return conn.QueryRowCtx(ctx, v, query, postId)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, error_type.ErrNotFound
	default:
		return nil, err
	}
}
