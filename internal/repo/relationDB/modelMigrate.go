package relationDB

import (
	"context"
	"gitee.com/unitedrhino/share/conf"
	"gitee.com/unitedrhino/share/stores"
)

func Migrate(c conf.Database) error {
	if c.IsInitTable == false {
		return nil
	}
	db := stores.GetCommonConn(context.TODO())
	err := db.AutoMigrate(
		&LightDeviceMsgCount{}, //需要初始化的表放这里会自动处理
	)
	if err != nil {
		return err
	}
	return err
}
