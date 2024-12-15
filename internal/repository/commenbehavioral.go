package repository

import (
	"FoodFound/internal/entity"
	"context"
	"gorm.io/gorm"
)

type commonBehavioral[T entity.DBModel] struct {
	db *gorm.DB
}

func NewCommonBehavioral[T entity.DBModel](db *gorm.DB) CommonBehaviourRepository[T] {
	return commonBehavioral[T]{
		db: db,
	}
}

func (c commonBehavioral[T]) ByID(ctx context.Context, id uint) (T, error) {
	return c.ByField(ctx, "id", id)
}

func (c commonBehavioral[T]) ByField(ctx context.Context, field string, id uint) (T, error) {
	var t T
	err := c.db.WithContext(ctx).Model(t).Where("id=?", id).First(t).Error
	return t, err
}

func (c commonBehavioral[T]) save(ctx context.Context, model T) error {
	return c.db.WithContext(ctx).Save(model).Error
}
