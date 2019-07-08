package main

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestOrder_getOrder(t *testing.T) {
	type fields struct {
		ID          int
		Origin      [2]string
		Destination [2]string
		Distance    int
		Status      string
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				ID:          tt.fields.ID,
				Origin:      tt.fields.Origin,
				Destination: tt.fields.Destination,
				Distance:    tt.fields.Distance,
				Status:      tt.fields.Status,
			}
			if err := o.getOrder(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Order.getOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_getOrderStatus(t *testing.T) {
	type fields struct {
		ID          int
		Origin      [2]string
		Destination [2]string
		Distance    int
		Status      string
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				ID:          tt.fields.ID,
				Origin:      tt.fields.Origin,
				Destination: tt.fields.Destination,
				Distance:    tt.fields.Distance,
				Status:      tt.fields.Status,
			}
			if err := o.getOrderStatus(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Order.getOrderStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_updateOrder(t *testing.T) {
	type fields struct {
		ID          int
		Origin      [2]string
		Destination [2]string
		Distance    int
		Status      string
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				ID:          tt.fields.ID,
				Origin:      tt.fields.Origin,
				Destination: tt.fields.Destination,
				Distance:    tt.fields.Distance,
				Status:      tt.fields.Status,
			}
			if err := o.updateOrder(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Order.updateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getOrders(t *testing.T) {
	type args struct {
		db    *sql.DB
		page  int
		limit int
	}
	tests := []struct {
		name    string
		args    args
		want    []OrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getOrders(tt.args.db, tt.args.page, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_placeOrder(t *testing.T) {
	type fields struct {
		ID          int
		Origin      [2]string
		Destination [2]string
		Distance    int
		Status      string
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				ID:          tt.fields.ID,
				Origin:      tt.fields.Origin,
				Destination: tt.fields.Destination,
				Distance:    tt.fields.Distance,
				Status:      tt.fields.Status,
			}
			if err := o.placeOrder(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Order.placeOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_validateCoordinates(t *testing.T) {
	type fields struct {
		ID          int
		Origin      [2]string
		Destination [2]string
		Distance    int
		Status      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				ID:          tt.fields.ID,
				Origin:      tt.fields.Origin,
				Destination: tt.fields.Destination,
				Distance:    tt.fields.Distance,
				Status:      tt.fields.Status,
			}
			if err := o.validateCoordinates(); (err != nil) != tt.wantErr {
				t.Errorf("Order.validateCoordinates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckCoordinate(t *testing.T) {
	type args struct {
		coordinate string
		limit      float32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckCoordinate(tt.args.coordinate, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("CheckCoordinate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_CalculateDistance(t *testing.T) {
	type fields struct {
		ID          int
		Origin      [2]string
		Destination [2]string
		Distance    int
		Status      string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				ID:          tt.fields.ID,
				Origin:      tt.fields.Origin,
				Destination: tt.fields.Destination,
				Distance:    tt.fields.Distance,
				Status:      tt.fields.Status,
			}
			got, err := o.CalculateDistance()
			if (err != nil) != tt.wantErr {
				t.Errorf("Order.CalculateDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Order.CalculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
