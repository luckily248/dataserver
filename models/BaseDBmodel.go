package models

import (
	"gopkg.in/mgo.v2"
	"log"
)

type BaseDBmodel struct {
	session *mgo.Session
	db      *mgo.Database
	c       *mgo.Collection
}

func (this *BaseDBmodel) DBname() string {
	return "dataserver"
}

func (this *BaseDBmodel) init() {
	//mongodburl := beego.AppConfig.String("mongodburl")
	//log.Printf("mgourl:%s\n", mongodburl)
	newsession, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Printf("mgo init err:%s\n", err.Error())
		panic(err)
	}
	this.session = newsession
	this.session.SetMode(mgo.Monotonic, true)
	this.db = this.session.DB(this.DBname())
}
