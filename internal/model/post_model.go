package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		Find(ctx context.Context, orderBy string, skip int64, limit int64) ([]Post, error)
		FindByUserAndPvtAndStatReq(ctx context.Context, userId string, privacy bool, status int64, skip int64, limit int64) ([]Post, error)
	}

	customPostModel struct {
		*defaultPostModel
	}
)

func (c customPostModel) FindByUserAndPvtAndStatReq(ctx context.Context, userId string, isAnonymous bool, status int64, skip int64, limit int64) ([]Post, error) {
	var data []Post
	err := c.conn.Find(ctx, &data, bson.M{
		"userId":      userId,
		"isAnonymous": isAnonymous,
		"status":      status,
	}, &mopt.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	return data, err
}

func parse(order string) (string, int64) {
	args := strings.Split(order, " ")
	if args[1] == "asc" {
		return args[0], 1
	} else if args[1] == "desc" {
		return args[0], -1
	} else {
		return "", 0
	}
}

func (c customPostModel) Find(ctx context.Context, orderBy string, skip int64, limit int64) ([]Post, error) {
	var data []Post
	opts := mopt.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	if len(orderBy) != 0 {
		w, o := parse(orderBy)
		opts.Sort = bson.D{
			{w, o},
		}
	}
	err := c.conn.Find(ctx, &data, &bson.M{}, &opts)
	return data, err
}

// NewPostModel returns a model for the mongo.
func NewPostModel(url, db, collection string, c cache.CacheConf) PostModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customPostModel{
		defaultPostModel: newDefaultPostModel(conn),
	}
}
