package repository

import (
	"FoodFound/internal/entity"
	"context"
)

type CommonBehaviourRepository[T entity.DBModel] interface {
	ByID(ctx context.Context, id uint) (T, error)
	ByField(ctx context.Context, field string, id uint) (T, error)
	save(ctx context.Context, model T) error
}

type UserRepository interface {
	CommonBehaviourRepository[*entity.User]
}
