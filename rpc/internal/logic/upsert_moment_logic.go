package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpsertMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpsertMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertMomentLogic {
	return &UpsertMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  上传或更新动态
func (l *UpsertMomentLogic) UpsertMoment(in *pb.UpsertMomentReq) (*pb.UpsertMomentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpsertMomentResp{}, nil
}
