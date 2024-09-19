package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	intersect := intersectServiceImpl{}

	r.HandleFunc("/intersect", cubicHandler(intersect)).Methods("POST")

	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Handler for the POST request
func cubicHandler(service IntersectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the JSON request
		var cubicRequest CubicRequest
		if err := json.NewDecoder(r.Body).Decode(&cubicRequest); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Check for intersection
		if service.Intersects(cubicRequest.First, cubicRequest.Second) {
			volume := service.IntersectedVolume(cubicRequest.First, cubicRequest.Second)
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(CubicResponse{Success: true, Volume: volume})
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := json.NewEncoder(w).Encode(CubicResponse{Success: false})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
