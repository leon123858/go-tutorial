package orm

import (
	"database/sql"
	"db/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

var gormDB *gorm.DB

func GetDB() *gorm.DB {
	if gormDB != nil {
		return gormDB
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	gormDB = db
	return gormDB
}

func Close() {
	db, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
	gormDB = nil
}

func CreateTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func InsertRecord(db *sql.DB, name, email string) error {
	user := model.User{Name: "John Doe", Email: "john@example.com"}
	result := gormDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateRecord(db *sql.DB, id int, email string) error {
	user := model.User{}
	result := gormDB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.Email = email
	result = gormDB.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteRecord(db *sql.DB, id int) error {
	user := model.User{}
	result := gormDB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = gormDB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func QueryRecords(db *sql.DB) ([]*model.User, error) {
	users := make([]*model.User, 0)
	result := gormDB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
