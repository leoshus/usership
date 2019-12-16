package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"usership/model"
	"usership/util"
)

//ListUser: list all users
func ListUser(w http.ResponseWriter,r *http.Request)  {
	users, err := model.ListUsers()
	if err != nil {
		log.Fatal(err)
		util.ResponseWithError(w,http.StatusInternalServerError, err.Error())
		return
	}
	util.ResponseSuccess(w,users)
}

//CreateUser: create a user
func CreateUser(w http.ResponseWriter,r *http.Request)  {
	var u model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		util.ResponseWithError(w,http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	err := model.InsertUser(&u)
	if err != nil {
		util.ResponseWithError(w,http.StatusInternalServerError, err.Error())
		return
	}
	util.ResponseSuccess(w,u)
}