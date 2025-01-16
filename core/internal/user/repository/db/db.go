package db

import (
	"context"

	"github.com/yumin00/go-hexagonal/core/domain"
	"gorm.io/gorm"
)

type dbUserRepository struct {
	db *gorm.DB
}

func NewDBUserRepository(db *gorm.DB) domain.UserRepository {
	return &dbUserRepository{
		db: db,
	}
}

func (d *dbUserRepository) GetUserInfoById(ctx context.Context, id int32) (*domain.UserInfo, error) {
	var userInfo *domain.UserInfo

	result := d.db.WithContext(ctx).Table("users").
		Where("id = ?", id).
		First(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	return userInfo, nil
}
