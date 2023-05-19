package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	app "go-crud/src/application"
	"go-crud/src/db"
	ihttp "go-crud/src/http"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(fmt.Errorf("error - server failed to start. err: %v", err))
	}
}

func run() error {
	// tying up all the components together and running the server
	db, err := db.NewRDS()
	if err != nil {
		return errors.Wrap(err, "unable to intialize db")
	}
	svc := app.NewPetSvc(db)
	h := ihttp.NewHandler(svc)
	r := chi.NewRouter()
	ihttp.Routes(r, h)
	return http.ListenAndServe(":80", r)
}
