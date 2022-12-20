package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	IsAnonymous bool               `bson:"isAnonymous,omitempty" json:"isAnonymous,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Text        string             `bson:"text,omitempty" json:"text,omitempty"`
	CoverUrl    string             `bson:"coverUrl,omitempty" json:"coverUrl,omitempty"`
	Tags        []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	UserId      string             `bson:"userId,omitempty" json:"userId,omitempty"`
	Status      int64              `bson:"status,omitempty" json:"status,omitempty"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
