package routes

import (
	"main/controllers"
	"sync"

	"github.com/gorilla/mux"
)

var router *mux.Router
var once sync.Once

func GetRouter() *mux.Router {

	once.Do(func() {

		router = mux.NewRouter()
		initRouter()
	})

	return router
}

func initRouter() {

	// Add your routes here.

	router.HandleFunc("/receipts/process", controllers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", controllers.GetPoints).Methods("GET")
}
