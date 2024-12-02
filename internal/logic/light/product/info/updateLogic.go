package info

import (
	"context"
	"gitee.com/unitedrhino/share/utils"
	"lightsvr/internal/repo/relationDB"

	"lightsvr/internal/svc"
	"lightsvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新产品详情
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.ProductInfo) error {
	old, err := relationDB.NewProductInfoRepo(l.ctx).FindOne(l.ctx, req.ID)
	if err != nil {
		return err
	}
	newPo := utils.Copy[relationDB.LightProductInfo](req)
	newPo.NoDelTime = old.NoDelTime
	newPo.DeletedTime = old.DeletedTime
	err = relationDB.NewProductInfoRepo(l.ctx).Update(l.ctx, newPo)
	return err
}
