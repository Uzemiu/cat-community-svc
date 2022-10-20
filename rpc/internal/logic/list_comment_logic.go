package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentLogic {
	return &ListCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  Comment
func (l *ListCommentLogic) ListComment(in *pb.ListCommentReq) (*pb.QueryMomentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryMomentResp{}, nil
}
