package main
import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

)




func main(){
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/hello",get).Methods("GET","OPTIONS")





	port := ":8080"
	log.Printf("Server is listening on port %s", port)
	log.Fatal(http.ListenAndServe(port,r))
}
