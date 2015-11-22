package main

import (

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"github.com/mustafaakin/buki"
	"encoding/json"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Buki Remote API")
}

func getVMs(w http.ResponseWriter, r *http.Request) {
	vms := buki.ListVM()
	b, err := json.Marshal(vms)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-type", "application/json") // temporary
	fmt.Fprintln(w, string(b))
}

func getImages(w http.ResponseWriter, r *http.Request) {
	vms := buki.GetAvailableImages()
	b, err := json.Marshal(vms)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-type", "application/json") // temporary
	fmt.Fprintln(w, string(b))
}

func getNetworks(w http.ResponseWriter, r *http.Request) {
	vms := buki.GetNetworks()
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
	router.HandleFunc("/", index)
	router.HandleFunc("/vms", getVMs)
	router.HandleFunc("/images", getImages)
	router.HandleFunc("/networks", getNetworks)

	// router.HandleFunc("/networks/{todoId}", TodoShow)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
