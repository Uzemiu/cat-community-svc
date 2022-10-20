package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminDetailLogic {
	return &GetAdminDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  Admin
func (l *GetAdminDetailLogic) GetAdminDetail(in *pb.GetAdminDetailReq) (*pb.GetAdminDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetAdminDetailResp{}, nil
}
