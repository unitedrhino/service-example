package info

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"lightsvr/internal/svc"
)

type ControlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 设备控制
func NewControlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ControlLogic {
	return &ControlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ControlLogic) Control() error {
	// todo: add your logic here and delete this line

	return nil
}
