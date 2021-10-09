package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"posty/pkg/entities"
	"posty/pkg/user"
	"posty/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(userService user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestUser entities.RequestUser

		err := json.NewDecoder(r.Body).Decode(&requestUser)
		if err != nil {
			NewErrorResponse(http.StatusBadRequest, err.Error(), w)
			return
		}

		// fmt.Println(requestUser)

		if !utils.IsEmailValid(requestUser.Email) {
			validationErr := errors.New("invalid email")
			if validationErr != nil {
				fmt.Println(validationErr.Error())
				NewErrorResponse(http.StatusBadRequest, validationErr.Error(), w)
				return
			}
		}

		res, dberr := userService.InsertUser(&requestUser)
		if dberr != nil {
			fmt.Println(dberr.Error())
			NewErrorResponse(http.StatusInternalServerError, dberr.Error(), w)
			return
		}

		NewSuccessResponse(http.StatusOK, "User successfully created", res, w)
	}
}

func GetUser(userService user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		// fmt.Println(id)
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			NewErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}

		// fmt.Println(objID)

		fetched, dberr := userService.FetchUser(objID)
		if dberr != nil {
			fmt.Println(dberr.Error())

			NewErrorResponse(http.StatusInternalServerError, dberr.Error(), w)
			return
		}

		// fmt.Println(fetched)

		NewSuccessResponse(http.StatusOK, "User successfully fetched", fetched, w)
	}
}
