package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"posty/pkg/entities"
	"posty/pkg/post"
	"posty/pkg/user"
	"posty/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPost(postService post.Service, userService user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestPost entities.RequestPost

		err := json.NewDecoder(r.Body).Decode(&requestPost)
		if err != nil {
			NewErrorResponse(http.StatusBadRequest, err.Error(), w)
			return
		}

		var post entities.Post
		post.Caption = requestPost.Caption
		post.Image = requestPost.Image
		post.UserID, err = primitive.ObjectIDFromHex(requestPost.UserID)
		if err != nil {
			NewErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}

		userPassHash, err := userService.FetchUser(post.UserID)
		if err != nil {
			NewErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}

		// fmt.Println(userPassHash.Password)

		if !utils.VerifyHash(requestPost.Password, userPassHash.Password) {
			NewErrorResponse(http.StatusUnauthorized, "Unauthorized", w)
			return
		}

		// fmt.Println(requestPost)

		res, dberr := postService.InsertPost(&post)
		if dberr != nil {
			fmt.Println(dberr.Error())
			NewErrorResponse(http.StatusInternalServerError, dberr.Error(), w)
			return
		}

		NewSuccessResponse(http.StatusOK, "Post successfully created", res, w)
	}
}

func GetPost(postService post.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/posts/")
		// fmt.Println(id)
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			NewErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}

		// fmt.Println(objID)

		fetched, dberr := postService.FetchPost(objID)
		if dberr != nil {
			fmt.Println(dberr.Error())

			NewErrorResponse(http.StatusInternalServerError, dberr.Error(), w)
			return
		}

		// fmt.Println(fetched)

		NewSuccessResponse(http.StatusOK, "Post successfully fetched", fetched, w)
	}
}

// Get all posts by a particular user
func GetPostsByUser(postService post.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/posts/users/")
		// fmt.Println(id)
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			NewErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}

		// fmt.Println(objID)

		fetched, dberr := postService.FetchAllPostsByUser(objID)
		if dberr != nil {
			fmt.Println(dberr.Error())

			NewErrorResponse(http.StatusInternalServerError, dberr.Error(), w)
			return
		}

		// fmt.Println(fetched)

		NewSuccessResponse(http.StatusOK, "Post successfully fetched", fetched, w)
	}
}
