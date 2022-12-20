package model

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/xh-polaris/meowchat-post-rpc/internal/config"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mitchellh/mapstructure"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PostCollectionName = "post"
const PostIndexName = "meowchat_post.post-alias"

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		FindMany(ctx context.Context, skip int64, count int64) ([]*Post, error)
		FindManyByUserId(ctx context.Context, userId string, status int64, skip int64, count int64) ([]*Post, error)
		Search(ctx context.Context, keyword string, count, skip int64) ([]*Post, error)
	}

	customPostModel struct {
		*defaultPostModel
		es *elasticsearch.Client
	}
)

// NewPostModel returns a model for the mongo.
func NewPostModel(url, db string, c cache.CacheConf, es config.ElasticsearchConf) PostModel {
	conn := monc.MustNewModel(url, db, PostCollectionName, c)
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: es.Addresses,
		Username:  es.Username,
		Password:  es.Password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return &customPostModel{
		defaultPostModel: newDefaultPostModel(conn),
		es:               esClient,
	}
}

func (m *customPostModel) FindManyByUserId(ctx context.Context, userId string, status int64, skip int64, count int64) ([]*Post, error) {
	var data []*Post
	err := m.conn.Find(ctx, &data, bson.M{
		"userId": userId,
		"status": status,
	}, &options.FindOptions{
		Skip:  &skip,
		Limit: &count,
	})
	return data, err
}

func (m *customPostModel) FindMany(ctx context.Context, skip int64, count int64) ([]*Post, error) {
	var data []*Post
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &count,
	}
	err := m.conn.Find(ctx, &data, &bson.M{}, &opts)
	return data, err
}

func (m *customPostModel) Search(ctx context.Context, keyword string, count, skip int64) ([]*Post, error) {
	search := m.es.Search
	query := map[string]any{
		"from": skip,
		"size": count,
		"query": map[string]any{
			"bool": map[string]any{
				"must": []any{
					map[string]any{
						"multi_match": map[string]any{
							"query":  keyword,
							"fields": []string{"title", "text", "tags"},
						},
					},
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}
	res, err := search(
		search.WithIndex(PostIndexName),
		search.WithContext(ctx),
		search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		} else {
			logx.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var r map[string]any
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}
	hits := r["hits"].(map[string]any)["hits"].([]any)
	posts := make([]*Post, 0, 10)
	for i := range hits {
		hit := hits[i].(map[string]any)
		post := &Post{}
		source := hit["_source"].(map[string]any)
		if source["createAt"], err = time.Parse("2006-01-02T15:04:05Z07:00", source["createAt"].(string)); err != nil {
			return nil, err
		}
		if source["updateAt"], err = time.Parse("2006-01-02T15:04:05Z07:00", source["updateAt"].(string)); err != nil {
			return nil, err
		}
		hit["_source"] = source
		err := mapstructure.Decode(hit["_source"], post)
		if err != nil {
			return nil, err
		}
		oid := hit["_id"].(string)
		id, err := primitive.ObjectIDFromHex(oid)
		if err != nil {
			return nil, err
		}
		post.ID = id
		posts = append(posts, post)
	}
	return posts, nil
}
