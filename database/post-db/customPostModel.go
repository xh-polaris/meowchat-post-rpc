package post_db

import (
	"database/sql"
	"fmt"
	"github.com/xh-polaris/meowchat-post-rpc/database/error_type"
	"github.com/xh-polaris/meowchat-post-rpc/pb/pb"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/net/context"
	"time"
)

func (m *customPostModel) Insertx(data *Post) (sql.Result, error) {
	return m.Insert(context.Background(), data)
}

func (m *customPostModel) Deletex(id int64) error {
	iyuaPostIdKey := fmt.Sprintf("%s%v", cacheIyuaPostIdPrefix, id)
	_, err := m.ExecCtx(context.Background(), func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `delete_at` = ?, `is_deleted` = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, time.Now(), 1, id)
	}, iyuaPostIdKey)
	if err == ErrNotFound {
		err = error_type.ErrNotFound
	}
	return err
}

func (m *customPostModel) Updatex(data *Post) error {
	return m.Update(context.Background(), data)
}

func (m *customPostModel) FindOnex(id int64) (*Post, error) {
	data, err := m.FindOne(context.Background(), id)
	if err == ErrNotFound {
		err = error_type.ErrNotFound
	}
	return data, err
}
func (m *defaultPostModel) Listx(ctx context.Context, skip int64, count int64) ([]*pb.Post, error) {
	var resp []*pb.Post
	query := fmt.Sprintf("select %s from %s limit ?,?", postRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, skip, count)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
