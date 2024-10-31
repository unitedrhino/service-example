package info

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"lightsvr/internal/svc"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除设备
func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete() error {
	// todo: add your logic here and delete this line

	return nil
}
