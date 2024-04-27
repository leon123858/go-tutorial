package main

import (
	"db/service/user"
	"testing"
)

func Test_processUser(t *testing.T) {
	type args struct {
		userService *user.UserService
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test processUser: pg",
			args: args{
				userService: user.NewUserService(user.Postgress),
			},
		},
		{
			name: "Test processUser: pg orm",
			args: args{
				userService: user.NewUserService(user.PgOrm),
			},
		},
		{
			name: "Test processUser: mongo",
			args: args{
				userService: user.NewUserService(user.Mongo),
			},
		},
		{
			name: "Test processUser: redis",
			args: args{
				userService: user.NewUserService(user.Rdb),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processUser(tt.args.userService)
		})
	}
}
