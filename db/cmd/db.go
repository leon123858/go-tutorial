package main

import (
	"db/model"
	"db/service/user"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	// init Tables
	//userService := tmpUser.NewUserService(tmpUser.PgOrm)
	//userService := tmpUser.NewUserService(tmpUser.Postgress)
	userService := user.NewUserService(user.Mongo)
	if err := (*userService).InitTable(); err != nil {
		panic(err)
		return
	}

	// Create
	for i := 0; i < 10; i++ {
		randomStr := uuid.New().String()
		newName := fmt.Sprintf("test%s", randomStr)
		newEmail := fmt.Sprintf("test%s@abc.com", randomStr)
		user := &model.User{
			Name:  newName,
			Email: newEmail,
		}
		if err := (*userService).CreateUser(user.Email, user.Name); err != nil {
			panic(err)
		}
	}

	// Read
	users, err := (*userService).GetUserList()
	if err != nil {
		panic(err)
		return
	}
	for _, tmpUser := range users {
		println(tmpUser.Name, tmpUser.Email)
	}

	// Update
	users[0].Email = "newEmail-" + uuid.NewString()
	if err := (*userService).UpdateUser(users[0].ID, users[0].Email); err != nil {
		panic(err)
		return
	}

	// Read
	users, err = (*userService).GetUserList()
	if err != nil {
		panic(err)
		return
	}
	for _, user := range users {
		println(user.Name, user.Email)
	}

	// Delete
	for _, user := range users {
		if err := (*userService).DeleteUser(user.ID); err != nil {
			panic(err)
			return
		}
	}

	// Read
	println("After delete")
	users, err = (*userService).GetUserList()
	if err != nil {
		panic(err)
		return
	}
	for _, user := range users {
		println(user.Name)
	}

	// Close
	user.Close(*userService)
}
