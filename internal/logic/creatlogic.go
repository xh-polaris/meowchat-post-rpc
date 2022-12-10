package logic

import (
	"context"
	post_db "github.com/xh-polaris/meowchat-post-rpc/database/post-db"
	postTag_db "github.com/xh-polaris/meowchat-post-rpc/database/postTag-db"
	"strconv"

	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatLogic {
	return &CreatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatLogic) Creat(in *pb.CreatReq) (*pb.CreatResp, error) {
	insert, err := l.svcCtx.DataBase.PostDB.Insertx(&post_db.Post{
		Title: in.Post.Title,
		Text:  in.Post.Text,
	})
	if err != nil {
		return nil, err
	}
	insertId, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}
	for _, tag := range in.Post.Tags {
		tagId, _ := strconv.ParseInt(tag.Id, 10, 64)
		if _, err = l.svcCtx.DataBase.PostTagDB.Insertx(&postTag_db.PostTag{
			TagId:  tagId,
			PostId: insertId,
		}); err != nil {
			return nil, err
		}
	}
	insertIdString := strconv.FormatInt(insertId, 10)
	return &pb.CreatResp{PostId: insertIdString}, nil
}
