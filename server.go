// server.go

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("game_data/state.txt")
	if err != nil {
		http.Error(w, "Unable to read game state", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file content", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Game State: %s", string(data))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	state := r.FormValue("state")

	file, err := os.Create("game_data/state.txt")
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(state))
	if err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Game state saved")
}

func main() {
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/post", handlePost)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}