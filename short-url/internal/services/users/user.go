package users

import "short-url/pkg/pg"

var dsn = pg.Dsn{
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "mysecretpassword",
	DB:       "postgres",
}

type IUserService interface {
	// CreateUser create a new user
	CreateUser(email string) (string, error)
	// GetUserStatistics User Statistics
	GetUserStatistics(email string) ([]pg.Event, error)
	// CreateEvent create a new event for a user
	CreateEvent(event pg.Event) error
}

type UserService struct {
	db *UserPgImpl
}

type UserPgImpl struct {
	db *pg.Client
}

func (up *UserPgImpl) CreateUser(email string) (string, error) {
	pwd, err := up.db.CreateUser(email)
	if err != nil {
		return "", err
	}
	return pwd, nil
}

func (up *UserPgImpl) GetUserStatistics(email string) ([]pg.Event, error) {
	events, err := up.db.FindEventByUser(email)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (up *UserPgImpl) CreateEvent(event pg.Event) error {
	err := up.db.CreateEvent(event)
	if err != nil {
		return err
	}
	return nil
}

func NewUserService() *UserService {
	db, err := pg.NewClient(dsn)
	if err != nil {
		panic(err)
	}
	err = db.Migration()
	if err != nil {
		return nil
	}
	return &UserService{
		db: &UserPgImpl{
			db: db,
		},
	}
}

func (us *UserService) CreateUser(email string) (string, error) {
	return us.db.CreateUser(email)
}

func (us *UserService) GetUserStatistics(email string) ([]pg.Event, error) {
	return us.db.GetUserStatistics(email)
}

func (us *UserService) CreateEvent(event pg.Event) error {
	return us.db.CreateEvent(event)
}
