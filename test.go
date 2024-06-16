package main
import (
	"net/http"
	"encoding/json"

)
func get(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{"message":"Hello from docker debayan"})
}