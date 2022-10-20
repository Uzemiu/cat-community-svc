package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCatDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCatDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCatDetailLogic {
	return &GetCatDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  Cat
func (l *GetCatDetailLogic) GetCatDetail(in *pb.CatDetailReq) (*pb.CatDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CatDetailResp{}, nil
}
