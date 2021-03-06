package controllers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/jacky-htg/go-services/models"
	"github.com/jacky-htg/go-services/payloads/request"
	"github.com/jacky-htg/go-services/payloads/response"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

//Users : struct for set Users Dependency Injection
type Users struct {
	Db  *sqlx.DB
	Log *log.Logger
}

//List : http handler for returning list of users
func (u *Users) List(w http.ResponseWriter, r *http.Request) {
	var user models.User
	list, err := user.List(u.Db)
	if err != nil {
		u.Log.Printf("error call list users: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var listResponse []models.UserResponse
	for _, user := range list {
		var userResponse models.UserResponse
		userResponse.Transform(&user)
		listResponse = append(listResponse, userResponse)
	}

	data, err := json.Marshal(list)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

func (u *Users) View(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("error type casting paramID: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user models.User
	user.ID = uint(id)
	err = user.Get(u.Db)
	if err != nil {
		u.Log.Printf("error call list user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response response.UserResponse
	response.Transform(&user)
	data, err = json.Marshal(response)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

//Create : http handler for create new user
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.NewUserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userRequest)
	if err != nil {
		u.Log.Printf("error decode user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userRequest.Password != userRequest.RePassword {
		err = errors.New("Password not match")
		u.Log.Printf("error : %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Printf("error generate password: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userRequest.Password = string(pass)
	user := userRequest.Transform()
	err = user.Create(u.Db)
	if err != nil {
		u.Log.Printf("error call create user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var response response.UserResponse
	response.Transform(user)
	data, err := json.Marshal(response)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

//Update : http handler for update user by id
func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("error type casting paramID: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user models.User
	user.ID = uint64(id)
	err = user.Get(u.Db)
	if err != nil {
		u.Log.Printf("error call list user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var userRequest request.UserRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&userRequest)
	if err != nil {
		u.Log.Printf("error decode user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(userRequest.Password) > 0 && userRequest.Password != userRequest.RePassword {
		err = errors.New("Password not match")
		u.Log.Printf("error : %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(userRequest.Password) > 0 {
		pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			u.Log.Printf("error generate password: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		userRequest.Password = string(pass)
	}
	if userRequest.ID <= 0 {
		userRequest.ID = user.ID
	}
	userUpdate := userRequest.Transform(&user)
	err = userUpdate.Update(u.Db)
	if err != nil {
		u.Log.Printf("error call update user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var response response.UserResponse
	response.Transform(userUpdate)
	data, err := json.Marshal(response)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

//Delete : http handler for delete user by id
func (u *Users) Delete(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("error type casting paramID: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user models.User
	user.ID = uint64(id)
	err = user.Get(u.Db)
	if err != nil {
		u.Log.Printf("error call list user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	isDelete, err := user.Delete(u.Db)
	if err != nil {
		u.Log.Printf("error call delete user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isDelete {
		err = errors.New("error delete user")
		u.Log.Printf("error call delete user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
