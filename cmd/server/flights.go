package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/avbru/flights/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Service struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) flightsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.getHandler(w, r)
	case "POST":
		s.postHandler(w, r)
	default:
		handlerErr(w, errMethodNotAllowed)
	}
}

func (s *Service) getHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		handlerErr(w, err)
		return
	}
	fmt.Println(r.URL.Query())

	const query = `SELECT number, origin, destination, departure, arrival from flights`

	whereClause, values := filters(r.URL.Query())

	rows, err := s.db.Query(r.Context(), query+whereClause, values...)
	if err != nil {
		handlerErr(w, err)
		return
	}
	defer rows.Close()

	var flights model.Flights
	for rows.Next() {
		var f model.Flight
		err := rows.Scan(&f.Number, &f.Origin, &f.Destination, &f.Departure, &f.Arrival)
		if err != nil {
			handlerErr(w, err)
			return
		}
		flights = append(flights, f)
	}

	if len(flights) == 0 {
		handlerErr(w, pgx.ErrNoRows)
		return
	}

	flights.OrderBy(r.Form.Get("order_by"), r.Form.Get("order"))

	enc := json.NewEncoder(w)
	if err := enc.Encode(flights); err != nil {
		handlerErr(w, err)
		return
	}
}

func filters(req url.Values) (string, []interface{}) {
	allowedFilters := []string{"destination", "origin"}
	var values []interface{}
	var keys []string

	k := 1
	for _, v := range allowedFilters {
		if req.Has(v) {
			values = append(values, req.Get(v))
			keys = append(keys, v+"=$"+strconv.Itoa(k))
			k++
		}
	}
	if len(keys) == 0 {
		return "", nil
	}

	clause := " WHERE " + strings.Join(keys, " AND ")
	return clause, values
}

func (s *Service) postHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var flights model.Flights
	if err := json.NewDecoder(r.Body).Decode(&flights); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		handlerErr(w, err)
		return
	}
	defer tx.Rollback(ctx) //nolint:errcheck

	query := `INSERT INTO flights (number, origin, destination, departure, arrival) VALUES ($1,$2,$3,$4,$5)`

	for _, f := range flights {
		if _, err := tx.Exec(ctx, query, f.Number, f.Origin, f.Destination, f.Departure, f.Arrival); err != nil {
			handlerErr(w, err)
			if err := tx.Rollback(ctx); err != nil {
				handlerErr(w, err)
			}
			return
		}
	}

	if err := tx.Commit(ctx); err != nil {
		handlerErr(w, err)
	}
}
