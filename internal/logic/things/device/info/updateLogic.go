package info

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"lightsvr/internal/svc"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新设备
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update() error {
	// todo: add your logic here and delete this line

	return nil
}
