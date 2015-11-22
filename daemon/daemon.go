package main

import (

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"github.com/mustafaakin/buki"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Buki Remote API")
}

func GetVMs(w http.ResponseWriter, r *http.Request) {
	vms := buki.ListVM()
	b, err := json.Marshal(vms)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-type", "application/json") // temporary
	fmt.Fprintln(w, string(b))
}

/*
func GetVMs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
*/

func main() {
	/*


	*/

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/vms", GetVMs)
//	router.HandleFunc("/images", GetImages)
//	router.HandleFunc("/networks", GetNetworks)

	// router.HandleFunc("/networks/{todoId}", TodoShow)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
