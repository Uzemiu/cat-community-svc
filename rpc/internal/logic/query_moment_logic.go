package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMomentLogic {
	return &QueryMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  查询动态列表
func (l *QueryMomentLogic) QueryMoment(in *pb.QueryMomentReq) (*pb.QueryMomentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryMomentResp{}, nil
}
