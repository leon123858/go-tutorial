package service

import (
	"database/sql"
	"db/model"
	"db/repository/pg"
	"gorm.io/gorm"
)

type PostgressUserService struct {
	db    *sql.DB
	gorm  *gorm.DB
	alias UserServiceType
}

func (pgs *PostgressUserService) InitTable() error {
	err := pg.CreateTable(pgs.db)
	if err != nil {
		return err
	}
	return nil
}

func (pgs *PostgressUserService) GetType() UserServiceType {
	return pgs.alias
}

func (pgs *PostgressUserService) CreateUser(user *model.User) error {
	err := pg.InsertRecord(pgs.db, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (pgs *PostgressUserService) UpdateUser(user *model.User) error {
	err := pg.UpdateRecord(pgs.db, user.ID, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (pgs *PostgressUserService) DeleteUser(id int) error {
	err := pg.DeleteRecord(pgs.db, id)
	if err != nil {
		return err
	}
	return nil
}

func (pgs *PostgressUserService) GetUserList() ([]*model.User, error) {
	users, err := pg.QueryRecords(pgs.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}
