package model

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"usership/db"
)

type RelationShips struct {
	Id           int64  `json:"-"`
	OtherUserId  int64  `json:"user_id"`
	State        string `json:"state"`
	RelationType string `json:"type"`
}

//format RelationShips when print
func (r RelationShips) String() string {
	return fmt.Sprintf("RelationShips<%d %d %s %s>", r.Id, r.OtherUserId, r.State, r.RelationType)
}

//QueryRelationShips:
func QueryRelationShips(userId int64) ([]*RelationShips, error) {
	//check user if exists
	err := ExistUser(userId)
	if err != nil {
		return nil, err
	}
	var relationShip []*RelationShips
	err = db.DB().Model(&relationShip).Where("id=?", userId).Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return relationShip, err
}

//QueryRelationShip: query relationship between two users
func QueryRelationShip(userId, otherUserId int64) (relationShip *RelationShips, err error) {
	err = ExistUser(userId)
	if err != nil {
		return nil, err
	}
	err = ExistUser(otherUserId)
	if err != nil {
		return nil, err
	}
	relationShip = &RelationShips{}
	err = db.DB().Model(relationShip).Where("id=? and other_user_id=?", userId, otherUserId).Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return relationShip, err
}

//updateState: update user relationship's state
func UpdateState(relationShip *RelationShips) error {
	return db.DB().RunInTransaction(func(tx *pg.Tx) error {

		err := ExistUser(relationShip.Id)
		if err != nil {
			return err
		}
		err = ExistUser(relationShip.OtherUserId)
		if err != nil {
			return err
		}
		otherShip, err := QueryRelationShip(relationShip.OtherUserId, relationShip.Id)
		switch {
		//single like,insert directly
		case otherShip == nil:
			err = insertOrUpdateRelation(relationShip)
			return err
		//has matched ignore
		case otherShip.State == "matched":
			return nil
		//like each other so matched
		case otherShip.State == "liked" && relationShip.State == "liked":
			relationShip.State = "matched"
			otherShip.State = "matched"
			_, err := db.DB().Model(otherShip).Column("state").WherePK().Update()
			if err != nil {
				return err
			}
		default:
			//do nothing
		}
		err = insertOrUpdateRelation(relationShip)
		return err
	})
}
//insert or update RelationShip
func insertOrUpdateRelation(ships *RelationShips) error {
	err := db.DB().Insert(ships)
	_, err = db.DB().Model(ships).OnConflict("(id,other_user_id) DO UPDATE").Set("state=EXCLUDED.state").Insert()
	return err
}