package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostCommentLogic {
	return &PostCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  发表评论
func (l *PostCommentLogic) PostComment(in *pb.PostCommentReq) (*pb.PostCommentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PostCommentResp{}, nil
}
