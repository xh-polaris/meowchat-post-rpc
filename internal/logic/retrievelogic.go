package logic

import (
	"context"
	"strconv"

	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveLogic {
	return &RetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveLogic) Retrieve(in *pb.RetrieveReq) (*pb.RetrieveResp, error) {
	data, err := l.svcCtx.DataBase.PostDB.FindOnex(in.PostId)
	if err != nil {
		return nil, err
	}
	tagx, _ := l.svcCtx.DataBase.PostTagDB.Findx(in.PostId)
	var tag []*pb.Tag
	for _, idx := range tagx {
		_tag, _ := l.svcCtx.DataBase.TagDB.FindOne(context.Background(), int64(idx))
		tag = append(tag, &pb.Tag{
			Id:   strconv.FormatInt(_tag.Id, 10),
			Name: _tag.Name,
		})
	}

	return &pb.RetrieveResp{
		Post: &pb.Post{
			Id:          strconv.FormatInt(data.Id, 10),
			CreateAt:    data.CreateAt.Unix(),
			DeleteAt:    data.DeleteAt.Time.Unix(),
			IsDelete:    data.IsDeleted != 0,
			IsAnonymous: data.IsAnonymous != 0,
			Title:       data.Title,
			Text:        data.Text,
			CoverUrl:    data.CoverUrl,
			UserId:      data.UserId,
			Status:      data.Status,
			Tags:        tag,
		},
	}, nil
}
