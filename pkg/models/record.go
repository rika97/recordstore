package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rika97/recordstore/pkg/config"
)

var db *gorm.DB

type Record struct {
	gorm.Model
	Name   string `gorm:""json:"name"`
	Artist string `json:"artist"`
	Label  string `json:"label"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Record{})
}

func (b *Record) CreateRecord() *Record {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllRecords() []Record {
	var Records []Record
	db.Find(&Records)
	return Records
}

func GetRecordById(Id int64) (*Record, *gorm.DB) {
	var getRecord Record
	db := db.Where("ID=?", Id).Find(&getRecord)
	return &getRecord, db
}

func DeleteRecord(ID int64) Record {
	var record Record
	db.Where("ID=?", ID).Delete(record)
	return record
}
