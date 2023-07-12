package biz

import (
	"context"
	"easyStart/internal/logic/ek/model"
)

type IExampleStore interface {
	GetOne(ctx context.Context) (*model.Example, error)
}
