package main

import (
	"database/sql"
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func TestApp_placeOrder(t *testing.T) {
	type fields struct {
		Router *mux.Router
		DB     *sql.DB
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Router: tt.fields.Router,
				DB:     tt.fields.DB,
			}
			a.placeOrder(tt.args.w, tt.args.r)
		})
	}
}

func TestApp_takeOrder(t *testing.T) {
	type fields struct {
		Router *mux.Router
		DB     *sql.DB
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Router: tt.fields.Router,
				DB:     tt.fields.DB,
			}
			a.takeOrder(tt.args.w, tt.args.r)
		})
	}
}

func TestApp_listOrders(t *testing.T) {
	type fields struct {
		Router *mux.Router
		DB     *sql.DB
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Router: tt.fields.Router,
				DB:     tt.fields.DB,
			}
			a.listOrders(tt.args.w, tt.args.r)
		})
	}
}

func Test_respondWithError(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		code    int
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondWithError(tt.args.w, tt.args.code, tt.args.message)
		})
	}
}

func Test_respondWithJSON(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		code    int
		payload interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondWithJSON(tt.args.w, tt.args.code, tt.args.payload)
		})
	}
}
