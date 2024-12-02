package info

import (
	"context"
	"gitee.com/unitedrhino/share/stores"
	"gitee.com/unitedrhino/share/utils"
	"lightsvr/internal/repo/relationDB"

	"lightsvr/internal/svc"
	"lightsvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取产品详情列表
func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.ProductInfoIndexReq) (resp *types.ProductInfoIndexResp, err error) {
	f := relationDB.ProductInfoFilter{ProductName: req.ProductName}
	total, err := relationDB.NewProductInfoRepo(l.ctx).CountByFilter(l.ctx, f)
	if err != nil {
		return nil, err
	}
	pos, err := relationDB.NewProductInfoRepo(l.ctx).FindByFilter(l.ctx, f, utils.Copy[stores.PageInfo](req.Page))
	if err != nil {
		return nil, err
	}
	return &types.ProductInfoIndexResp{
		List:  utils.CopySlice[types.ProductInfo](pos),
		Total: total,
	}, err
}
