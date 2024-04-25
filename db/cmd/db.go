package main

import (
	"db/model"
	"db/service"
	"fmt"
)

func main() {
	// init Tables
	userService := service.NewUserService(service.PgOrm)
	//userService := service.NewUserService(service.Postgress)
	if err := (*userService).InitTable(); err != nil {
		return
	}

	// Create
	for i := 0; i < 10; i++ {
		newName := fmt.Sprintf("test%d", i)
		newEmail := fmt.Sprintf("test%d@abc.com", i)
		user := &model.User{
			Name:  newName,
			Email: newEmail,
		}
		if err := (*userService).CreateUser(user); err != nil {
			return
		}
	}

	// Read
	users, err := (*userService).GetUserList()
	if err != nil {
		return
	}
	for _, user := range users {
		println(user.Name, user.Email)
	}

	// Update
	users[0].Email = "newEmail"
	if err := (*userService).UpdateUser(users[0]); err != nil {
		return
	}

	// Read
	users, err = (*userService).GetUserList()
	if err != nil {
		return
	}
	for _, user := range users {
		println(user.Name, user.Email)
	}

	// Delete
	for _, user := range users {
		if err := (*userService).DeleteUser(user.ID); err != nil {
			return
		}
	}

	// Read
	println("After delete")
	users, err = (*userService).GetUserList()
	if err != nil {
		return
	}
	for _, user := range users {
		println(user.Name)
	}

	// Close
	service.Close(*userService)
}
