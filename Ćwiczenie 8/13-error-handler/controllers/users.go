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
func (u *Users) List(w http.ResponseWriter, r *http.Request) error{
	var user models.User
	list, err := user.List(u.Db)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}

	var listResponse []models.UserResponse
	for _, user := range list {
		var userResponse models.UserResponse
		userResponse.Transform(&user)
		listResponse = append(listResponse, userResponse)
	}

	return api.Response(w, listResponse, http.StatusOK)
}

func (u *Users) View(w http.ResponseWriter, r *http.Request) error{
	paramID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return errors.Wrap(err, "type casting")
	}
	var user models.User
	user.ID = uint(id)
	err = user.Get(u.Db)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	data, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}

	var response response.UserResponse
	response.Transform(&user)
	data, err = json.Marshal(response)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

//Create : http handler for create new user
func (u *Users) Create(w http.ResponseWriter, r *http.Request) error{
	var userRequest request.NewUserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userRequest)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	if userRequest.Password != userRequest.RePassword {
		return errors.Wrap(err, "getting list user")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	userRequest.Password = string(pass)
	user := userRequest.Transform()
	err = user.Create(u.Db)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	var response response.UserResponse
	response.Transform(user)
	data, err := json.Marshal(response)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

//Update : http handler for update user by id
func (u *Users) Update(w http.ResponseWriter, r *http.Request) error{
	paramID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	var user models.User
	user.ID = uint64(id)
	err = user.Get(u.Db)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	var userRequest request.UserRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&userRequest)
	if err != nil {
		return errors.Wrap(err, "getting list user")
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
			return errors.Wrap(err, "getting list user")
		}
		userRequest.Password = string(pass)
	}
	if userRequest.ID <= 0 {
		userRequest.ID = user.ID
	}
	userUpdate := userRequest.Transform(&user)
	err = userUpdate.Update(u.Db)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	var response response.UserResponse
	response.Transform(userUpdate)
	data, err := json.Marshal(response)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

//Delete : http handler for delete user by id
func (u *Users) Delete(w http.ResponseWriter, r *http.Request) error{
	paramID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	var user models.User
	user.ID = uint64(id)
	err = user.Get(u.Db)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	isDelete, err := user.Delete(u.Db)
	if err != nil {
		return errors.Wrap(err, "getting list user")
	}
	if !isDelete {
		return errors.Wrap(err, "getting list user")
	}
	w.WriteHeader(http.StatusNoContent)
}
