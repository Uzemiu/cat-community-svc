package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCarouselLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCarouselLogic {
	return &ListCarouselLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  Carousel
func (l *ListCarouselLogic) ListCarousel(in *pb.ListCarouselReq) (*pb.ListCarouselResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListCarouselResp{}, nil
}
