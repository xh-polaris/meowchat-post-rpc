package logic

import (
	"context"
	"database/sql"
	post_db "github.com/xh-polaris/meowchat-post-rpc/database/post-db"
	postTag_db "github.com/xh-polaris/meowchat-post-rpc/database/postTag-db"
	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/pb/pb"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *pb.UpdateReq) (*pb.UpdateResp, error) {
	if in.Post == nil {
		return nil, nil
	}
	// TODO remove old post_tag
	postId, _ := strconv.ParseInt(in.Post.Id, 10, 64)
	err := l.svcCtx.DataBase.PostTagDB.Deletex(in.Post.Id)
	if err != nil {
		return nil, err
	}
	for _, tag := range in.Post.Tags {
		tagId, _ := strconv.ParseInt(tag.Id, 10, 64)
		if _, err := l.svcCtx.DataBase.PostTagDB.Insertx(&postTag_db.PostTag{
			PostId: postId,
			TagId:  tagId,
		}); err != nil {
			return nil, err
		}
	}
	var isDelete int64
	if in.Post.IsDelete {
		isDelete = 1
	}
	var isAnonymous byte
	if in.Post.IsAnonymous {
		isAnonymous = 1
	}
	return &pb.UpdateResp{}, l.svcCtx.DataBase.PostDB.Updatex(&post_db.Post{
		Id:       postId,
		CreateAt: time.Unix(in.Post.CreateAt, 0),
		DeleteAt: sql.NullTime{
			time.Unix(in.Post.DeleteAt, 0),
			in.Post.DeleteAt != 0,
		},
		IsDeleted:   isDelete,
		IsAnonymous: isAnonymous,
		Title:       in.Post.Title,
		Text:        in.Post.Text,
		CoverUrl:    in.Post.CoverUrl,
		Status:      in.Post.Status,
		UserId:      in.Post.UserId,
	})
}
