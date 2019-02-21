package routes

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"os"
	//"strconv"
)

func HandleRequests()  {
	router := mux.NewRouter().StrictSlash(true)
	urls(router)
	server(router)
}

func server(router *mux.Router)  {
	server := http.ListenAndServe(os.Getenv("HOST"), router)
	log.Fatal(server)
}

func urls(router *mux.Router)  {
	router.HandleFunc("/", Index)
	router.HandleFunc("/messages", Messages)
}

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("We are on index")
}

func Messages(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("We are on Messages")
	fmt.Println(r.Form)
	//fmt.Println(strconv.Atoi(r))
	fmt.Println(r.URL.Query().Get("hub.challenge"))
	fmt.Fprintf(w, r.URL.Query().Get("hub.challenge"))
}