package router

import (
	core "AMCO/server/core/handler"

	"github.com/gorilla/mux"
)

func ServeHandle() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/amco/api", core.Index)
	router.HandleFunc("/amco/api/producer", core.ProduceHandle).Methods("POST")
	router.HandleFunc("/amco/api/consumer", core.ConsumeHandle).Methods("POST")

	return router
}
