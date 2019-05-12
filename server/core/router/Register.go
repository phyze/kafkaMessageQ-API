package router

import (
	core "KafkaMessageQ-API/server/core/handler"

	"github.com/gorilla/mux"
)

func ServeHandle() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", core.Index)
	router.HandleFunc("/api/producer", core.ProduceHandle).Methods("POST")
	router.HandleFunc("/api/consumer", core.ConsumeHandle).Methods("POST")

	return router
}
