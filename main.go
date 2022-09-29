package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"shortener/src"
)

// Parse out yaml from cli
var yaml_file string
var db_name string

func init() {
	flag.StringVar(&yaml_file, "y", "", "File path for yaml file with short links")
	flag.StringVar(&db_name, "db", "", "File name for sqlite3 db containing urls")
}

func main() {
	flag.Parse()

	mux := createMux()

	// Using map handler by default
	var pathToUrls = map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	Handler := src.MapHandler(pathToUrls, mux)

	// Only use yaml handler if yaml file passed
	if yaml_file != "" {
		fmt.Printf("Reading yaml file: %v\n", yaml_file)
		data := src.ReadFile(yaml_file)
		yamlHandler, err := src.YAMLHandler(data, Handler)
		if err != nil {
			panic(err)
		}
		Handler = yamlHandler
	}

	// Only use db if database file provided
	if db_name != "" {
		fmt.Printf("Reading DB: %v\n", db_name)

		data, err := src.ReadDb(db_name)
		if err != nil {
			panic(err)
		}

		DBHandler := src.MapHandler(data, Handler)
		Handler = DBHandler
	}

	// Starts a connection to the server
	startServer(Handler)

}

func startServer(handler http.Handler) {
	fmt.Println("Starting server on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", fallback)
	return mux
}

func fallback(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World. Url Shortner")
}
