package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/orders", a.placeOrder).Methods("POST")
	a.Router.HandleFunc("/orders/{id:[0-9]+}", a.takeOrder).Methods("PATCH")
	a.Router.HandleFunc("/orders", a.listOrders).Methods("GET")
}

func (a *App) Initialize(user, password, dbname string) {
	// connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	connectionString := fmt.Sprintf("%s:%s@tcp(172.21.0.2:3306)/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) placeOrder(w http.ResponseWriter, r *http.Request) {
	var o Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&o); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := o.validateCoordinates(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	distance, err := o.CalculateDistance()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Distance error: %s", err.Error()))
		return
	}
	o.Distance = distance
	o.Status = "UNASSIGNED"
	defer r.Body.Close()
	if err := o.placeOrder(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	res := OrderResponse{
		ID:       o.ID,
		Distance: o.Distance,
		Status:   o.Status,
	}
	respondWithJSON(w, http.StatusCreated, res)
}

func (a *App) takeOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	o := Order{ID: id}
	if err := o.getOrder(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Order not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	if o.Status == "TAKEN" {
		respondWithError(w, http.StatusNotAcceptable, "Order is already taken")
		return
	}
	o.Status = "TAKEN"
	if err := o.updateOrder(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	res := OrderStatus{
		Status: o.Status,
	}
	respondWithJSON(w, http.StatusOK, res)
}

func (a *App) listOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CALL")
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page", page, "limit", limit)
	if limit < 1 {
		respondWithError(w, http.StatusInternalServerError, "Limit should be more than 0")
		return
	}
	if page < 1 {
		respondWithError(w, http.StatusInternalServerError, "Page should be more than 0")
		return
	}
	orders, err := getOrders(a.DB, page, limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, orders)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
