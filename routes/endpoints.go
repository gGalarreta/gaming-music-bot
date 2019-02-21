package routes

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"os"
)

func AppSetup()  {
	router := mux.NewRouter().StrictSlash(true)
	urls(router)
	server(router)
}

func urls(router *mux.Router)  {
	router.HandleFunc("/", Index)
}

func server(router *mux.Router)  {
	server := http.ListenAndServe(os.Getenv("SERVER_HOST"), router)
	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello this is an index")
}