package info

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"lightsvr/internal/svc"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增设备
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create() error {
	// todo: add your logic here and delete this line

	return nil
}
