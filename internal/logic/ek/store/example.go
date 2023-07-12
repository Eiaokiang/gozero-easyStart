package store

import (
	"context"
	"easyStart/internal/logic/ek/biz"
	"easyStart/internal/logic/ek/model"
	"gorm.io/gorm"
)

type ExampleStore struct {
	Store *gorm.DB
}

func NewExampleStore(store *gorm.DB) biz.IExampleStore {
	return &ExampleStore{
		Store: store,
	}
}

func (d *ExampleStore) GetOne(ctx context.Context) (*model.Example, error) {
	var result model.Example
	take := d.Store.Table("example").Take(&result)
	if take.Error != nil {
		return nil, take.Error
	}
	return &result, nil
}
