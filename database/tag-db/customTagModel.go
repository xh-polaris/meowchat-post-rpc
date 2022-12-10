package tag_db

import (
	"context"
	"database/sql"
	"github.com/xh-polaris/meowchat-post-rpc/database/error_type"
)

func (m *customTagModel) Insertx(data *Tag) (sql.Result, error) {
	return m.Insert(context.Background(), data)
}
func (m *customTagModel) Deletex(data *Tag) error {
	return m.Delete(context.Background(), data.Id)
}
func (m *customTagModel) FindOnex(data *Tag) (*Tag, error) {
	tag, err := m.FindOne(context.Background(), data.Id)
	if err == ErrNotFound {
		err = error_type.ErrNotFound
	}
	return tag, err
}
