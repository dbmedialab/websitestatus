package main

import (
	"net/http"
	"text/template"

	"github.com/egreb/sitesstatus/broker"
	"github.com/egreb/sitesstatus/site"
	"github.com/egreb/sitesstatus/utils"
	"github.com/egreb/sitesstatus/worker"
	"github.com/julienschmidt/httprouter"
)

// set routes

func main() {
	var router = httprouter.New()
	broker := broker.NewServer()
	router.NotFound = http.FileServer(http.Dir("static"))
	router.GET("/", index)
	router.GET("/events", broker.ServeHTTP)
	worker.Worker(broker)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := site.Site{}
	t, err := template.ParseFiles("./static/index.html")
	utils.Check(err)
	t.Execute(w, data)
}
