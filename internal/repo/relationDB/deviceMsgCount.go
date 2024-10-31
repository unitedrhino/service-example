package relationDB

import (
	"context"
	"gitee.com/unitedrhino/share/stores"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
这个是参考样例
使用教程:
1. 将DeviceMsgCount全局替换为模型的表名
2. 完善todo
*/

type DeviceMsgCountRepo struct {
	db *gorm.DB
}

func NewDeviceMsgCountRepo(in any) *DeviceMsgCountRepo {
	return &DeviceMsgCountRepo{db: stores.GetCommonConn(in)}
}

type DeviceMsgCountFilter struct {
	//todo 添加过滤字段
}

func (p DeviceMsgCountRepo) fmtFilter(ctx context.Context, f DeviceMsgCountFilter) *gorm.DB {
	db := p.db.WithContext(ctx)
	//todo 添加条件
	return db
}

func (p DeviceMsgCountRepo) Insert(ctx context.Context, data *LightDeviceMsgCount) error {
	result := p.db.WithContext(ctx).Create(data)
	return stores.ErrFmt(result.Error)
}

func (p DeviceMsgCountRepo) FindOneByFilter(ctx context.Context, f DeviceMsgCountFilter) (*LightDeviceMsgCount, error) {
	var result LightDeviceMsgCount
	db := p.fmtFilter(ctx, f)
	err := db.First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}
func (p DeviceMsgCountRepo) FindByFilter(ctx context.Context, f DeviceMsgCountFilter, page *stores.PageInfo) ([]*LightDeviceMsgCount, error) {
	var results []*LightDeviceMsgCount
	db := p.fmtFilter(ctx, f).Model(&LightDeviceMsgCount{})
	db = page.ToGorm(db)
	err := db.Find(&results).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return results, nil
}

func (p DeviceMsgCountRepo) CountByFilter(ctx context.Context, f DeviceMsgCountFilter) (size int64, err error) {
	db := p.fmtFilter(ctx, f).Model(&LightDeviceMsgCount{})
	err = db.Count(&size).Error
	return size, stores.ErrFmt(err)
}

func (p DeviceMsgCountRepo) Update(ctx context.Context, data *LightDeviceMsgCount) error {
	err := p.db.WithContext(ctx).Where("id = ?", data.ID).Save(data).Error
	return stores.ErrFmt(err)
}

func (p DeviceMsgCountRepo) DeleteByFilter(ctx context.Context, f DeviceMsgCountFilter) error {
	db := p.fmtFilter(ctx, f)
	err := db.Delete(&LightDeviceMsgCount{}).Error
	return stores.ErrFmt(err)
}

func (p DeviceMsgCountRepo) Delete(ctx context.Context, id int64) error {
	err := p.db.WithContext(ctx).Where("id = ?", id).Delete(&LightDeviceMsgCount{}).Error
	return stores.ErrFmt(err)
}
func (p DeviceMsgCountRepo) FindOne(ctx context.Context, id int64) (*LightDeviceMsgCount, error) {
	var result LightDeviceMsgCount
	err := p.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}

// 批量插入 LightStrategyDevice 记录
func (p DeviceMsgCountRepo) MultiInsert(ctx context.Context, data []*LightDeviceMsgCount) error {
	err := p.db.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Model(&LightDeviceMsgCount{}).Create(data).Error
	return stores.ErrFmt(err)
}

func (d DeviceMsgCountRepo) UpdateWithField(ctx context.Context, f DeviceMsgCountFilter, updates map[string]any) error {
	db := d.fmtFilter(ctx, f)
	err := db.Model(&LightDeviceMsgCount{}).Updates(updates).Error
	return stores.ErrFmt(err)
}
