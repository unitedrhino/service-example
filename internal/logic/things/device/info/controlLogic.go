package info

import (
	"context"
	"gitee.com/unitedrhino/share/stores"
	"gorm.io/gorm"
	"lightsvr/internal/repo/relationDB"
	"time"

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
	//普通使用模式
	err := relationDB.NewDeviceMsgCountRepo(l.ctx).Insert(l.ctx, &relationDB.LightDeviceMsgCount{
		Num:  6,
		Date: time.Now(),
	})
	//事务示例
	err = stores.GetCommonConn(l.ctx).Transaction(func(tx *gorm.DB) error {
		//初始化的时候一定要传tx,如果传的是ctx,则不会在这个事务里
		po, err := relationDB.NewDeviceMsgCountRepo(tx).FindOne(l.ctx, 66)
		if err != nil {
			return err
		}
		po.Num++
		return relationDB.NewDeviceMsgCountRepo(tx).Update(l.ctx, po)
	})
	return err
}
