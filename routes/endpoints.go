package routes

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"bot/app/services/facebook"
	"fmt"
	"os"
	"io/ioutil"
)

func HandleRequests()  {
	router := mux.NewRouter().StrictSlash(true)
	urls(router)
	server(router)
}

func server(router *mux.Router)  {
	fmt.Println("Server up. Running at " + os.Getenv("HOST"))
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

func Messages(w http.ResponseWriter, req *http.Request)  {
	body, _ := ioutil.ReadAll(req.Body)
	sender, text, postback := facebook.GetMessageData(string(body))
	facebook.HandleResponses(sender, text, postback)
	//fmt.Println(req.URL.Query().Get("hub.challenge"))
	//fmt.Fprintf(w, req.URL.Query().Get("hub.challenge"))
}