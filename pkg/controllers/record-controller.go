package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rika97/recordstore/pkg/models"
	"github.com/rika97/recordstore/pkg/utils"
)

var NewRecord models.Record

func GetRecord(w http.ResponseWriter, r *http.Request) {
	newRecords := models.GetAllRecords()
	res, _ := json.Marshal(newRecords)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRecordById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recordId := vars["recordId"]
	ID, err := strconv.ParseInt(recordId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	recordDetails, _ := models.GetRecordById(ID)
	res, _ := json.Marshal(recordDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	CreateRecord := &models.Record{}
	utils.ParseBody(r, CreateRecord)
	b := CreateRecord.CreateRecord()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recordId := vars["recordId"]
	ID, err := strconv.ParseInt(recordId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	record := models.DeleteRecord(ID)
	res, _ := json.Marshal(record)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateRecord(w http.ResponseWriter, r *http.Request) {
	var updateRecord = &models.Record{}
	utils.ParseBody(r, updateRecord)
	vars := mux.Vars(r)
	recordId := vars["recordId"]
	ID, err := strconv.ParseInt(recordId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	recordDetails, db := models.GetRecordById(ID)
	if updateRecord.Name != "" {
		recordDetails.Name = updateRecord.Name
	}
	if updateRecord.Artist != "" {
		recordDetails.Artist = updateRecord.Artist
	}
	if updateRecord.Label != "" {
		recordDetails.Label = updateRecord.Label
	}
	db.Save(&recordDetails)
	res, _ := json.Marshal(recordDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
