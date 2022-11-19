package mysql

import (
	"context"
	"errors"
	"exercise/domain"
	"exercise/domain/entity"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) *mysqlUserRepository {
	return &mysqlUserRepository{DB: DB}
}

func (m *mysqlUserRepository) FindByEmail(c context.Context, email string) (*entity.User, error) {
	var res entity.User
	err := m.DB.WithContext(c).Where("email = ?", email).Take(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = domain.ErrNotFound
	}
	return &res, err
}

func (m *mysqlUserRepository) Create(c context.Context, user *entity.User) error {
	err := m.DB.WithContext(c).Create(&user).Error
	return err
}
