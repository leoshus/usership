package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"usership/model"
	"usership/util"
)

//QueryRelationShip: query specific user's relationships by userId
func QueryRelationShip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseInt(vars["user_id"], 10, 64)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	//query relationship by userId from database
	relationship, err := model.QueryRelationShips(userId)
	if err != nil {
		util.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.ResponseSuccess(w, relationship)
}
//ChangeState: match two user and change relationship state
func ChangeState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseInt(vars["user_id"], 10, 64)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	otherUserId, err := strconv.ParseInt(vars["other_user_id"], 10, 64)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	relationShip := &model.RelationShips{}
	if err := decoder.Decode(relationShip); err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	switch relationShip.State {
	case "disliked":
	case "liked":
	default:
		util.ResponseWithError(w, http.StatusBadRequest, "invalid state")
		return
	}
	relationShip.Id = userId
	relationShip.OtherUserId = otherUserId
	relationShip.RelationType = "relationship"
	err = model.UpdateState(relationShip)
	if err != nil {
		util.ResponseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	ship, _ := model.QueryRelationShip(userId, otherUserId)
	util.ResponseSuccess(w, ship)

}
