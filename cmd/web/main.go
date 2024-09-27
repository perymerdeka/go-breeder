package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config AppConfig
}

type AppConfig struct {
	useCache bool
}

func main() {
 app := application{
	templateMap: make(map[string]*template.Template),
 }

 flag.BoolVar(&app.config.useCache, "cache", false, "use template cache")
 flag.Parse()


 server := &http.Server{
	Addr: port,
	Handler: app.routes(),
	IdleTimeout: 30 * time.Second,
	ReadTimeout: 30 * time.Second,
	WriteTimeout: 30 * time.Second,
 }

 fmt.Printf("Starting server on port %s\n", port)

 err := server.ListenAndServe()
 if err != nil {
	log.Panic(err)
 }

}