package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type API struct {
	port   string
	host   string
	router *httprouter.Router
}

func NewApi(host, port string) *API {
	return &API{
		host: host,
		port: port,
		//router: httprouter.New(),
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func (api *API) Start() {

	// router := httprouter.New()
	// router.GET("/", Index)
	// router.GET("/hello/:name", Hello)

	// hostPort := fmt.Sprintf(":%s", api.port)
	// log.Printf("Listening to: %s", hostPort)
	// err := http.ListenAndServe(":8080", router)
	// log.Fatal(err)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func (api *API) Stop() {

}
