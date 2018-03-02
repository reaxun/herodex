package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func addRoutes(router *mux.Router) {
	router.HandleFunc("/hero/{name}", getHero).Methods("GET")
	router.HandleFunc("/hero/{name}/maxstats", getMaxStats).Methods("GET")
	router.HandleFunc("/hero/{name}/maxstats/{rarity}", getMaxStats).Methods("GET")
	router.HandleFunc("/skill/{name}", getSkill).Methods("GET")
}

// StartAPI initalizes the API to listen on port 12345
func StartAPI() {
	router := mux.NewRouter()
	addRoutes(router)
	log.Fatal(http.ListenAndServe(":12345", router))
}
