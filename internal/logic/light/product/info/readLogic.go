package info

import (
	"context"
	"gitee.com/unitedrhino/share/utils"
	"lightsvr/internal/repo/relationDB"

	"lightsvr/internal/svc"
	"lightsvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取产品详情详情
func NewReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadLogic {
	return &ReadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadLogic) Read(req *types.WithID) (resp *types.ProductInfo, err error) {
	po, err := relationDB.NewProductInfoRepo(l.ctx).FindOne(l.ctx, req.ID)
	return utils.Copy[types.ProductInfo](po), err
}
