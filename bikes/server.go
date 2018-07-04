package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func GetAllBikes(w http.ResponseWriter, r *http.Request) {
	
	// Parse Resquest ? OR Not

	// We will send responses to this channel.
	ch := make(chan []storage.Bike)
	go storage.ReturnAllBikes(ch)
	
	
	
	select {
		case listAllBike := <- ch:
		send := json.Encode(listAllBike)
		w.Write(send)
	}
	

}

func GetBike(w http.ResponseWriter, r *http.Request) {
	// make a SQL Request for return  bike

	id := r.URL // must complete this ...
  	ch := make(chan []storage.Bike)
	
	go storage.ReturnBike(id)
	
	
	
	select {
		case resultBike := <- ch:
		send := json.Encode(resultBike)
		w.Write(send)
	}
	send := json.Encode(Result)
	w.Write(send)	

}
func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/listbikes", GetAllBikes).Methods("GET")
	router.HandleFunc("/bike/{id}", GetBike).Methods("GET")
    log.Fatal(http.ListenAndServe(":8081", router))
}