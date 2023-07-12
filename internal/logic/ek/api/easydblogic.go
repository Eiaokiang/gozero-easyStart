package api

import (
	"context"
	"easyStart/internal/logic/ek/biz"
	"easyStart/internal/logic/ek/store"
	"errors"

	"easyStart/internal/svc"
	"easyStart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EasyDbLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	BizEasyDb *biz.BizEasyDb
}

func NewEasyDbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EasyDbLogic {
	return &EasyDbLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		BizEasyDb: biz.NewBizEasyDb(svcCtx, store.NewExampleStore(svcCtx.DbStore)),
	}
}

func (l *EasyDbLogic) EasyDb() (resp *types.Response, err error) {
	logx.Info("aaaaabbbbbbbbbbbb")
	one, err := l.BizEasyDb.GetOne(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("查询错误")
	}

	return &types.Response{
		Message: one.Name,
	}, nil
}
