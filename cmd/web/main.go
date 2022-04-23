package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/callsamu/lovecalc/pkg/core"
)

type application struct {
	calculator    core.Calculator
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.Int("addr", 4000, "Port address to listen on")
	algorithm := flag.String("algorithm", core.HashWithFNV, "Algorithm to be used by the calculator")
	flag.Parse()

	c, err := core.NewCalculator(*algorithm)
	if err != nil {
		log.Fatal(err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache("./ui/template/")
	fmt.Println(templateCache)
	if err != nil {
		log.Fatal(err)
	}

	app := application{
		calculator:    c,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", *addr),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.infoLog.Printf("Listening on port %d", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
