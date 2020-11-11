package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/lesnikyan/go_http_example/routes"
	// "strings"
)

// "github.com/lesnikyan/go_http_example/routes"
/*
testing:
$ ab -n 20000 -c 10 http://localhost:82/

*/

func main() {

	// use custom handler

	xhandler := routes.RegexpHandler{}

	xhandler.HandleRegexp("/abc\\=.+", handleAbc)

	xhandler.HandleRegexp("/", handleRoot)

	http.Handle("/", xhandler)

	// use net/http handler

	http.HandleFunc("/zzz/", handleSpec)

	http.HandleFunc("/sublvl", func(w http.ResponseWriter, rq *http.Request) {
		handleRoot(w, rq)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TplData Index page content
type TplData struct {
	Title   string
	Content string
	Page    int
}

func handleAbc(w http.ResponseWriter, rq *http.Request) {

	uri := rq.RequestURI

	tpdata := TplData{
		Title:   "Hello page",
		Content: "ABC URL: " + uri,
		Page:    1,
	}

	render(w, tpdata)
}

// specific url handler
func handleSpec(w http.ResponseWriter, rq *http.Request) {

	uri := rq.RequestURI

	tpdata := TplData{
		Title:   "Hello page",
		Content: "Special URL: " + uri,
		Page:    1,
	}

	render(w, tpdata)
}

// root handler
func handleRoot(w http.ResponseWriter, rq *http.Request) {

	tpdata := TplData{
		Title:   "Hello page",
		Content: "Page content",
		Page:    1,
	}

	render(w, tpdata)
}

// tpl render
func render(w http.ResponseWriter, data TplData) {
	tpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Fatal(err)
	}

	tpl.Execute(w, data)
}

// service functions
func pwr(w http.ResponseWriter, msg string) {
	fmt.Fprintf(w, msg)
}

func p(msg string) {
	fmt.Println(msg)
}
