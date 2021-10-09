package handler

import (
	"fmt"
	"net/http"
	"posty/pkg/user"
)

func AddUser(userService user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("create user here")
	}
}
