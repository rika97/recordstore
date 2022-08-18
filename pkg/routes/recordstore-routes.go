package routes

import (
	"github.com/gorilla/mux"
	"github.com/rika97/recordstore/pkg/controllers"
)

var RegisterRecordStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/record/", controllers.CreateRecord).Methods("POST")
	router.HandleFunc("/record", controllers.GetRecord).Methods("GET")
	router.HandleFunc("/record/{recordId}", controllers.GetRecordById).Methods("GET")
	router.HandleFunc("/record/{recordId}", controllers.UpdateRecord).Methods("PUT")
	router.HandleFunc("/record/{recordId}", controllers.DeleteRecord).Methods("DELETE")
}
