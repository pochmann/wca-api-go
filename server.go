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

	// Return this cuber's basic info or results
	if len(cuberId) == 10 {
		cuber, ok := getCuber(cuberId)
		writeJson(cuber, ok, w, r)
	} else {
		cuber, _ := getCuber(cuberId[:10])
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Content-Encoding", "gzip")
		w.Write(cuber.resultsJsonGzip)
	}
}

func handle_ranking(w http.ResponseWriter, r *http.Request) {
	eventId := r.URL.Path[len("/rankings/"):]
	i := getStr32(eventId) // TODO check for existence
	writeJson(rankingEntries[i][:100], true, w, r)
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
	http.HandleFunc("/rankings/", handle_ranking)
	http.ListenAndServe(":8080", nil)
}
