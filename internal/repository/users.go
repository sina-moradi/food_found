package repository

import (
	"FoodFound/internal/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	CommonBehaviourRepository[*entity.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		CommonBehaviourRepository: NewCommonBehavioral[*entity.User](db),
	}
}
