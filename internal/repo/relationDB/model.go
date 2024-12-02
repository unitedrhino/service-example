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

// 产品信息表
type LightProductInfo struct {
	ID           int64  `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	ProductID    string `gorm:"column:product_id;type:varchar(100);uniqueIndex:pd;NOT NULL"` // 产品id
	ProductName  string `gorm:"column:product_name;type:varchar(100);NOT NULL"`              // 产品名称
	ProductImg   string `gorm:"column:product_img;type:varchar(200)"`                        // 产品图片
	ProductType  int64  `gorm:"column:product_type;type:smallint;default:1"`                 // 产品状态:1:开发中,2:审核中,3:已发布
	AuthMode     int64  `gorm:"column:auth_mode;type:smallint;default:1"`                    // 认证方式:1:账密认证,2:秘钥认证
	AutoRegister int64  `gorm:"column:auto_register;type:smallint;default:1"`                // 动态注册:1:关闭,2:打开,3:打开并自动创建设备
	Secret       string `gorm:"column:secret;type:varchar(50)"`                              // 动态注册产品秘钥
	Desc         string `gorm:"column:description;type:varchar(200)"`                        // 描述
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:pd"`
	//Devices []*DmDeviceInfo    `gorm:"foreignKey:ProductID;references:ProductID"` // 添加外键
}

func (m *LightProductInfo) TableName() string {
	return "light_product_info"
}
