package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func writeJson(data interface{}, ok bool, w http.ResponseWriter, r *http.Request) {
	if ok {
		j, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(j)
	} else {
		http.NotFound(w, r)
	}
}

func handle_cuber(w http.ResponseWriter, r *http.Request) {
	cuberId := r.URL.Path[len("/cubers/"):]
	cuber, ok := getCuber(cuberId)
	writeJson(cuber, ok, w, r)
}

func main() {

	// Load the export
	fmt.Print("Loading the WCA export... ")
	t0 := time.Now()
	LoadWcaData()
	fmt.Println("done in", time.Since(t0))

	// Prepare extra data
	fmt.Print("Preparing extra data... ")
	t0 = time.Now()
	PrepareExtraData()
	fmt.Println("done in", time.Since(t0))

	// Start the server
	fmt.Println("Starting the server...")
	http.HandleFunc("/cubers/", handle_cuber)
	http.ListenAndServe(":8080", nil)
}
