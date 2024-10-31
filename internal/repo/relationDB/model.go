package relationDB

import (
	"gitee.com/unitedrhino/share/stores"
	"time"
)

type LightExample struct {
	ID int64 `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"` // id编号
}

// 设备信息表
type LightDeviceMsgCount struct {
	ID   int64     `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	Num  int64     `gorm:"column:num;type:bigint;default:0"`           //数量
	Date time.Time `gorm:"column:date;NOT NULL;uniqueIndex:date_type"` //统计的日期
	stores.OnlyTime
}

func (m *LightDeviceMsgCount) TableName() string {
	return "light_device_msg_count"
}
