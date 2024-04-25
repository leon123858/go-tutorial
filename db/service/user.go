package service

import (
	"db/model"
	"db/repository/orm"
	"db/repository/pg"
)

type UserServiceType int

const (
	Postgress UserServiceType = iota
	PgOrm                     // ORM
)

type UserService interface {
	InitTable() error
	GetType() UserServiceType
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
	GetUserList() ([]*model.User, error)
}

func NewUserService(t UserServiceType) *UserService {
	var ret UserService
	switch t {
	case Postgress:
		ret = new(PostgressUserService)
		ret.(*PostgressUserService).alias = t
		ret.(*PostgressUserService).db = pg.GetDB()
	case PgOrm:
		ret = new(PgOrmUserService)
		ret.(*PgOrmUserService).alias = t
		ret.(*PgOrmUserService).gorm = orm.GetDB()
	}
	return &ret
}

func Close(service UserService) {
	switch service.GetType() {
	case Postgress:
		pg.Close()
	case PgOrm:
		orm.Close()
	}
}
