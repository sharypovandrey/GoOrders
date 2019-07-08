package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/kr/pretty"

	"googlemaps.github.io/maps"
)

type Order struct {
	ID          int       `json:"id"`
	Origin      [2]string `json:"origin"`
	Destination [2]string `json:"destination"`
	Distance    int       `json:"distance"`
	Status      string    `json:"status"`
}

type OrderResponse struct {
	ID       int    `json:"id"`
	Distance int    `json:"distance"`
	Status   string `json:"status"`
}

type OrderStatus struct {
	Status string `json:"status"`
}

func (o *Order) getOrder(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id, distance, status FROM orders WHERE id=%d", o.ID)
	return db.QueryRow(statement).Scan(&o.ID, &o.Distance, &o.Status)
}

func (o *Order) getOrderStatus(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT status FROM orders WHERE id=%d", o.ID)
	return db.QueryRow(statement).Scan(&o.Status)
}

func (o *Order) updateOrder(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE orders SET status='%s' WHERE id=%d", o.Status, o.ID)
	_, err := db.Exec(statement)
	return err
}

func getOrders(db *sql.DB, page, limit int) ([]OrderResponse, error) {
	start := (page - 1) * limit
	statement := fmt.Sprintf("SELECT id, distance, status FROM orders LIMIT %d OFFSET %d", limit, start)
	rows, err := db.Query(statement)
	// pretty.Println("rows", rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []OrderResponse{}
	for rows.Next() {
		var o OrderResponse
		if err := rows.Scan(&o.ID, &o.Distance, &o.Status); err != nil {
			pretty.Println("scan", err)
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (o *Order) placeOrder(db *sql.DB) error {
	statement := fmt.Sprintf(
		"INSERT INTO orders(start_latitude, start_longtitude, end_latitude, end_longtitude, distance, status) VALUES('%s', '%s', '%s', '%s', %d, '%s')",
		o.Origin[0], o.Origin[1], o.Destination[0], o.Destination[1], o.Distance, o.Status,
	)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&o.ID)
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) validateCoordinates() error {
	if err := CheckCoordinate(o.Origin[0], 90.); err != nil {
		return err
	}
	if err := CheckCoordinate(o.Origin[1], 180.); err != nil {
		return err
	}
	if err := CheckCoordinate(o.Destination[0], 90.); err != nil {
		return err
	}
	if err := CheckCoordinate(o.Destination[1], 180.); err != nil {
		return err
	}
	return nil
}

func CheckCoordinate(coordinate string, limit float32) error {
	v, err := strconv.ParseFloat(coordinate, 32)
	if err != nil {
		return err
	}
	if float32(v) > limit || float32(v) < -limit {
		return errors.New("The coordinate is out of bounds")
	}
	return nil
}

func (o *Order) CalculateDistance() (int, error) {
	key := os.Getenv("GOOGLE_MAPS_KEY")
	c, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		return 0, err
	}
	r := &maps.DistanceMatrixRequest{
		Origins:      []string{fmt.Sprintf("%s,%s", o.Origin[0], o.Origin[1])},
		Destinations: []string{fmt.Sprintf("%s,%s", o.Destination[0], o.Destination[1])},
	}
	res, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		return 0, err
	}
	var distance int
	for _, r := range res.Rows {
		for _, e := range r.Elements {
			distance += e.Distance.Meters
		}
	}
	fmt.Println("DISTANCE", distance)
	return distance, nil
}
