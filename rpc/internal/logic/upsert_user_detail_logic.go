package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpsertUserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpsertUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertUserDetailLogic {
	return &UpsertUserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  插入或更新用户信息
func (l *UpsertUserDetailLogic) UpsertUserDetail(in *pb.UpsertUserReq) (*pb.UpsertUserDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpsertUserDetailResp{}, nil
}
