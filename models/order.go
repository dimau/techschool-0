package models

import (
	"time"
)

// Order represents the main order entity
type Order struct {
	ID                int       `json:"id" db:"id"`
	OrderUID          string    `json:"order_uid" db:"order_uid"`
	TrackNumber       string    `json:"track_number" db:"track_number"`
	Entry             string    `json:"entry" db:"entry"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerID        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	ShardKey          string    `json:"shardkey" db:"shardkey"`
	SmID              int       `json:"sm_id" db:"sm_id"`
	DateCreated       time.Time `json:"date_created" db:"date_created"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`

	// Related data
	Delivery *Delivery `json:"delivery,omitempty"`
	Payment  *Payment  `json:"payment,omitempty"`
	Items    []Item    `json:"items,omitempty"`
}

// Delivery represents delivery information for an order
type Delivery struct {
	ID        int       `json:"id" db:"id"`
	OrderID   int       `json:"order_id" db:"order_id"`
	Name      string    `json:"name" db:"name"`
	Phone     string    `json:"phone" db:"phone"`
	Zip       string    `json:"zip" db:"zip"`
	City      string    `json:"city" db:"city"`
	Address   string    `json:"address" db:"address"`
	Region    string    `json:"region" db:"region"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Payment represents payment information for an order
type Payment struct {
	ID           int       `json:"id" db:"id"`
	OrderID      int       `json:"order_id" db:"order_id"`
	Transaction  string    `json:"transaction" db:"transaction"`
	RequestID    string    `json:"request_id" db:"request_id"`
	Currency     string    `json:"currency" db:"currency"`
	Provider     string    `json:"provider" db:"provider"`
	Amount       int       `json:"amount" db:"amount"`
	PaymentDt    int64     `json:"payment_dt" db:"payment_dt"`
	Bank         string    `json:"bank" db:"bank"`
	DeliveryCost int       `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int       `json:"goods_total" db:"goods_total"`
	CustomFee    int       `json:"custom_fee" db:"custom_fee"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Item represents an item in an order
type Item struct {
	ID          int       `json:"id" db:"id"`
	OrderID     int       `json:"order_id" db:"order_id"`
	ChrtID      int64     `json:"chrt_id" db:"chrt_id"`
	TrackNumber string    `json:"track_number" db:"track_number"`
	Price       int       `json:"price" db:"price"`
	Rid         string    `json:"rid" db:"rid"`
	Name        string    `json:"name" db:"name"`
	Sale        int       `json:"sale" db:"sale"`
	Size        string    `json:"size" db:"size"`
	TotalPrice  int       `json:"total_price" db:"total_price"`
	NmID        int64     `json:"nm_id" db:"nm_id"`
	Brand       string    `json:"brand" db:"brand"`
	Status      int       `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// OrderRequest represents the incoming JSON structure for creating/updating orders
type OrderRequest struct {
	OrderUID          string          `json:"order_uid"`
	TrackNumber       string          `json:"track_number"`
	Entry             string          `json:"entry"`
	Delivery          DeliveryRequest `json:"delivery"`
	Payment           PaymentRequest  `json:"payment"`
	Items             []ItemRequest   `json:"items"`
	Locale            string          `json:"locale"`
	InternalSignature string          `json:"internal_signature"`
	CustomerID        string          `json:"customer_id"`
	DeliveryService   string          `json:"delivery_service"`
	ShardKey          string          `json:"shardkey"`
	SmID              int             `json:"sm_id"`
	DateCreated       time.Time       `json:"date_created"`
	OofShard          string          `json:"oof_shard"`
}

// DeliveryRequest represents delivery data in the request
type DeliveryRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

// PaymentRequest represents payment data in the request
type PaymentRequest struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

// ItemRequest represents item data in the request
type ItemRequest struct {
	ChrtID      int64  `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int64  `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
