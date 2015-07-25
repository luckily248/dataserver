package models

import (
	"gopkg.in/mgo.v2"
	"beego"
)

type BaseDBmodel struct {
	session *mgo.Session
	db      *mgo.Database
	c       *mgo.Collection
}

func (this *BaseDBmodel) DBname() string {
	return "dataserver"
}

func (this *BaseDBmodel) init() (err error){
	mgourl:=beego.AppConfig.String("mgourl")
	newsession, err := mgo.Dial(mgourl)
	if err!=nil{
		return
	}
	this.session = newsession
	this.session.SetMode(mgo.Monotonic, true)
	this.db = this.session.DB(this.DBname())
	return
}
