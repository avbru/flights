package main

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"net/http"
)

var (
	errMethodNotAllowed = errors.New("method not allowed")
)

func handlerErr(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, errMethodNotAllowed):
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
