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


func getVM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmName := vars["vmName"]

	vm := buki.GetVM(vmName)
	b, err := json.Marshal(vm)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-type", "application/json") // temporary
	fmt.Fprintln(w, string(b))

}

func startVM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmName := vars["vmName"]

	buki.StartVM(vmName)
	w.Header().Set("Content-type", "application/json") // temporary
	fmt.Fprintln(w, "ok")

}

func stopVM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmName := vars["vmName"]

	buki.StopVM(vmName)
	w.Header().Set("Content-type", "application/json") // temporary
	fmt.Fprintln(w, "ok")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)

	// VMs
	router.HandleFunc("/vms", getVMs)
	router.HandleFunc("/vm/{vmName}/info", getVM)
	router.HandleFunc("/vm/{vmName}/start", startVM)
	router.HandleFunc("/vm/{vmName}/stop", stopVM)

	// Images
	router.HandleFunc("/images", getImages)

	// Networks
	router.HandleFunc("/networks", getNetworks)

	// router.HandleFunc("/networks/{todoId}", TodoShow)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
