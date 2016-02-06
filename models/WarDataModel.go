package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type WarDataModel struct {
	BaseDBmodel
	Id        bson.ObjectId `bson:"_id" form:"-" `
	TeamA     string        `form:"TeamA"`
	TeamB     string        `form:"TeamB"`
	battles   []Battle
	Timestamp time.Time
}
type Battle struct {
	No int
}

func (this *WarDataModel) Tablename() string {
	return "wardata"
}

func (this *WarDataModel) init() (err error) {
	err = this.BaseDBmodel.init()
	if err != nil {
		return
	}
	this.c = this.db.C(this.Tablename())
	this.Id = bson.NewObjectId()
	this.TeamA = "teamA"
	this.TeamB = "teamB"
	this.Timestamp = time.Now()
	return
}
func (this *WarDataModel) GetWarId() string {
	return this.Id.Hex()

}
func (this *WarDataModel) SetWarId(id string) {
	this.Id = bson.ObjectIdHex(id)
}
func AddWarData(teama string, teamb string) (id string, err error) {
	wardata := &WarDataModel{}
	err = wardata.init()
	if err != nil {
		return
	}
	defer wardata.session.Close()
	wardata.TeamA = teama
	wardata.TeamB = teamb
	id = wardata.GetWarId()
	err = wardata.c.Insert(wardata)
	return
}

func GetWarData(clanname string) (content *WarDataModel, err error) {
	wardata := &WarDataModel{}
	err = wardata.init()
	if err != nil {
		return
	}
	defer wardata.session.Close()
	wardata.TeamA = clanname
	err = wardata.c.Find(bson.M{"teama": wardata.TeamA}).Sort("-timestamp").One(&content)
	return
}

func UpdateWarData(warid string, updata interface{}) (err error) {
	wardata := &WarDataModel{}
	err = wardata.init()
	if err != nil {
		return
	}
	defer wardata.session.Close()
	wardata.SetWarId(warid)
	err = wardata.c.UpdateId(wardata.Id, updata)
	return
}
