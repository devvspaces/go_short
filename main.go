package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"shortener/src"
)

type NewType struct{}

func (v NewType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/devvspaces", http.StatusPermanentRedirect)
}

func init() {
	flag.StringVar(&yaml_file, "y", "", "File path for yaml file with short links")
}

func RunMap() {

	var mux = map[string]string{
		"/urlshort":       "https://github.com/gophercises/urlshort",
		"/urlshort-final": "https://github.com/gophercises/urlshort/tree/solution",
	}

	var fallback NewType

	src.MapHandler(mux, fallback)

}

var yaml_file string

func RunYaml() {

	flag.Parse()

	if yaml_file == "" {
		log.Fatal("You must enter a valid file name")
	}

	var fallback NewType

	src.YAMLHandler(src.ReadFile(yaml_file), fallback)

}

func main() {
	RunMap()
	fmt.Println("Starting server on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
