package info

import (
	"context"
	"gitee.com/unitedrhino/share/utils"
	"lightsvr/internal/repo/relationDB"

	"lightsvr/internal/svc"
	"lightsvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加产品详情
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.ProductInfo) (resp *types.WithID, err error) {
	po := utils.Copy[relationDB.LightProductInfo](req)
	po.ID = 0
	err = relationDB.NewProductInfoRepo(l.ctx).Insert(l.ctx, po)
	return &types.WithID{ID: po.ID}, err
}
