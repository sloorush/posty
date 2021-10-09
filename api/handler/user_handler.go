package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"posty/pkg/entities"
	"posty/pkg/user"
	"posty/utils"
)

func AddUser(userService user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestUser entities.RequestUser

		err := json.NewDecoder(r.Body).Decode(&requestUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(requestUser)

		if !utils.IsEmailValid(requestUser.Email) {
			err1 := errors.New("invalid email")
			if err1 != nil {
				fmt.Println(err1.Error())
				var data entities.Error

				data.Status = http.StatusBadRequest
				data.Message = err1.Error()

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(data.Status)
				json.NewEncoder(w).Encode(data)
			}
		}

		res, dberr := userService.InsertUser(&requestUser)
		if dberr != nil {
			fmt.Println(dberr.Error())
			var data entities.Error

			data.Status = http.StatusInternalServerError
			data.Message = dberr.Error()

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(data.Status)
			json.NewEncoder(w).Encode(data)
		}

		var data entities.Success

		data.Status = http.StatusAccepted
		data.Message = "User successfully created"
		data.Data = res

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(data.Status)
		json.NewEncoder(w).Encode(data)
	}
}
