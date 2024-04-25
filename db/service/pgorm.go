package service

import (
	"db/model"
	"gorm.io/gorm"
)

type PgOrmUserService struct {
	alias UserServiceType
	gorm  *gorm.DB
}

func (pgOrm *PgOrmUserService) InitTable() error {
	err := pgOrm.gorm.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}

func (pgOrm *PgOrmUserService) GetType() UserServiceType {
	return pgOrm.alias
}

func (pgOrm *PgOrmUserService) CreateUser(user *model.User) error {
	result := pgOrm.gorm.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pgOrm *PgOrmUserService) UpdateUser(user *model.User) error {
	result := pgOrm.gorm.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pgOrm *PgOrmUserService) DeleteUser(id int) error {
	result := pgOrm.gorm.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pgOrm *PgOrmUserService) GetUserList() ([]*model.User, error) {
	var users []*model.User
	result := pgOrm.gorm.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
