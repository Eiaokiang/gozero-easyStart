package biz

import (
	"context"
	"easyStart/internal/logic/ek/model"
	"easyStart/internal/svc"
)

type BizEasyDb struct {
	IExampleStore IExampleStore
	svcCtx        *svc.ServiceContext
}

func NewBizEasyDb(svcCtx *svc.ServiceContext, iExampleStore IExampleStore) *BizEasyDb {
	return &BizEasyDb{
		svcCtx:        svcCtx,
		IExampleStore: iExampleStore,
	}
}

func (d *BizEasyDb) GetOne(ctx context.Context) (*model.Example, error) {
	result, err := d.IExampleStore.GetOne(ctx)
	return result, err
}
