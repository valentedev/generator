package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	app := &application{
		config: cfg,
	}

	fmt.Printf("Serving on port: %v\n", app.config.port)

	err := app.serve()
	if err != nil {
		log.Println(err)
	}
}

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/", app.generateHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5500", "http://127.0.0.1", "https://localhost"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders: []string{"Authorization", "Content-Type", "Access-Control-Allow-Origin"},
		Debug:          true,
	})

	handler := c.Handler(router)

	return handler
}

func (app *application) serve() error {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Server connected!")

	return nil
}
