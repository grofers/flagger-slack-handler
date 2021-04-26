package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	wait               time.Duration
	port, loadTesterNs string
)

func main() {
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "The duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&port, "port", "8080", "The port on which the web server should run")
	flag.StringVar(&loadTesterNs, "loadtester-namespace", "ci", "Namespace where flagger-loadtester is installed")
	flag.Parse()

	r := newRegisteredRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      loggedRouter,
	}

	log.Printf("Starting server on port %s", port)
	log.Printf("Passing gating instructions to `flagger-loadtester` in `namespace/%s`\n", loadTesterNs)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Gracefully shutting down server")
	os.Exit(0)
}

func newRegisteredRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/handler", handler)

	return r
}
