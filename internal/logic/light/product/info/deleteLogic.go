package info

import (
	"context"
	"lightsvr/internal/repo/relationDB"

	"lightsvr/internal/svc"
	"lightsvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除产品详情
func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.WithID) error {
	err := relationDB.NewProductInfoRepo(l.ctx).Delete(l.ctx, req.ID)

	return err
}
