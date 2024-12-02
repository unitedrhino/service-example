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
1. 将ProductInfo全局替换为模型的表名
2. 完善todo
*/

type ProductInfoRepo struct {
	db *gorm.DB
}

func NewProductInfoRepo(in any) *ProductInfoRepo {
	return &ProductInfoRepo{db: stores.GetCommonConn(in)}
}

type ProductInfoFilter struct {
	ProductName string
}

func (p ProductInfoRepo) fmtFilter(ctx context.Context, f ProductInfoFilter) *gorm.DB {
	db := p.db.WithContext(ctx)
	if f.ProductName != "" {
		db = db.Where("product_name like ?", "%"+f.ProductName+"%")
	}
	return db
}

func (p ProductInfoRepo) Insert(ctx context.Context, data *LightProductInfo) error {
	result := p.db.WithContext(ctx).Create(data)
	return stores.ErrFmt(result.Error)
}

func (p ProductInfoRepo) FindOneByFilter(ctx context.Context, f ProductInfoFilter) (*LightProductInfo, error) {
	var result LightProductInfo
	db := p.fmtFilter(ctx, f)
	err := db.First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}
func (p ProductInfoRepo) FindByFilter(ctx context.Context, f ProductInfoFilter, page *stores.PageInfo) ([]*LightProductInfo, error) {
	var results []*LightProductInfo
	db := p.fmtFilter(ctx, f).Model(&LightProductInfo{})
	db = page.ToGorm(db)
	err := db.Find(&results).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return results, nil
}

func (p ProductInfoRepo) CountByFilter(ctx context.Context, f ProductInfoFilter) (size int64, err error) {
	db := p.fmtFilter(ctx, f).Model(&LightProductInfo{})
	err = db.Count(&size).Error
	return size, stores.ErrFmt(err)
}

func (p ProductInfoRepo) Update(ctx context.Context, data *LightProductInfo) error {
	err := p.db.WithContext(ctx).Where("id = ?", data.ID).Save(data).Error
	return stores.ErrFmt(err)
}

func (p ProductInfoRepo) DeleteByFilter(ctx context.Context, f ProductInfoFilter) error {
	db := p.fmtFilter(ctx, f)
	err := db.Delete(&LightProductInfo{}).Error
	return stores.ErrFmt(err)
}

func (p ProductInfoRepo) Delete(ctx context.Context, id int64) error {
	err := p.db.WithContext(ctx).Where("id = ?", id).Delete(&LightProductInfo{}).Error
	return stores.ErrFmt(err)
}
func (p ProductInfoRepo) FindOne(ctx context.Context, id int64) (*LightProductInfo, error) {
	var result LightProductInfo
	err := p.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}

// 批量插入 LightStrategyDevice 记录
func (p ProductInfoRepo) MultiInsert(ctx context.Context, data []*LightProductInfo) error {
	err := p.db.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Model(&LightProductInfo{}).Create(data).Error
	return stores.ErrFmt(err)
}

func (d ProductInfoRepo) UpdateWithField(ctx context.Context, f ProductInfoFilter, updates map[string]any) error {
	db := d.fmtFilter(ctx, f)
	err := db.Model(&LightProductInfo{}).Updates(updates).Error
	return stores.ErrFmt(err)
}
