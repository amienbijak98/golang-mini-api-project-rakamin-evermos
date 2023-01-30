package models

import "time"

type User struct {
	ID              int       `gorm:"primary_key;auto_increment" json:"id"`
	Name            string    `json:"name"`
	Email           string    `gorm:"unique" json:"email"`
	Password        string    `json:"password"`
	MobileNumber    string    `gorm:"unique" json:"mobile_number"`
	BirthDate       time.Time `json:"birth_date"`
	Gender          string    `json:"gender"`
	About           string    `json:"about"`
	Profession      string    `json:"profession"`
	ProvinceID      string    `json:"province_id"`
	RegencyID       string    `json:"regency_id"`
	IsAdmin         bool      `json:"isAdmin"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time `json:"created_at"`
	Address         Address
	Store           Store
	UserCredentials UserCredentials
	Transaction     []Transaction
}

type Address struct {
	ID            int           `gorm:"primary_key;auto_increment" json:"id"`
	UserID        int           `json:"user_id"`
	AddressTitle  string        `json:"address_title"`
	RecipientName string        `json:"recipient_name"`
	MobileNumber  string        `json:"mobile_number"`
	AddressDetail string        `json:"address_detail"`
	UpdatedAt     time.Time     `json:"updated_at"`
	CreatedAt     time.Time     `json:"created_at"`
	Transaction   []Transaction `gorm:"foreignKey:ShippingAddress"`
}

type Store struct {
	ID                int       `gorm:"primary_key;auto_increment" json:"id"`
	UserID            int       `gorm:"unique" json:"user_id"`
	StoreName         string    `json:"store_name"`
	PhotoUrl          string    `json:"photo_url"`
	UpdatedAt         time.Time `json:"updated_at"`
	CreatedAt         time.Time `json:"created_at"`
	Product           []Product
	ProductLog        []ProductLog
	TransactionDetail []TransactionDetail
}

type Category struct {
	ID           int       `gorm:"primary_key;auto_increment" json:"id"`
	CategoryName string    `json:"category_name"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
	Product      []Product
	ProductLog   []ProductLog
}

type Product struct {
	ID            int       `gorm:"primary_key;auto_increment" json:"id"`
	ProductName   string    `json:"product_name"`
	Slug          string    `json:"slug"`
	ResellerPrice int       `json:"reseller_price"`
	ConsumenPrice int       `json:"consumen_price"`
	Stock         int       `json:"stock"`
	Description   string    `json:"description"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
	StoreID       int       `json:"store_id"`
	CategoryID    int       `json:"category_id"`
	ProductLog    ProductLog
	ProductImage  []ProductImage
}

type ProductLog struct {
	ID                int       `gorm:"primary_key;auto_increment" json:"id"`
	ProductID         int       `json:"product_id"`
	ProductName       string    `json:"product_name"`
	Slug              string    `json:"slug"`
	ResellerPrice     int       `json:"reseller_price"`
	ConsumerPrice     int       `json:"consumer_price"`
	Description       string    `json:"description"`
	UpdatedAt         time.Time `json:"updated_at"`
	CreatedAt         time.Time `json:"created_at"`
	StoreID           int       `json:"store_id"`
	CategoryID        int       `json:"category_id"`
	TransactionDetail []TransactionDetail
}

type ProductImage struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	ProductID int       `json:"product_id"`
	URL       string    `json:"url"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Transaction struct {
	ID                int       `json:"id" gorm:"primary_key"`
	UserID            int       `json:"user_id"`
	ShippingAddress   int       `json:"shipping_address"`
	TotalPrice        int       `json:"total_price"`
	InvoiceNumber     string    `json:"invoice_number"`
	PaymentMethod     string    `json:"payment_method"`
	UpdatedAt         time.Time `json:"updated_at"`
	CreatedAt         time.Time `json:"created_at"`
	TransactionDetail []TransactionDetail
}

type TransactionDetail struct {
	ID            int       `json:"id" gorm:"primary_key"`
	TransactionID int       `json:"transaction_id"`
	ProductLogID  int       `json:"product_log_id"`
	StoreID       int       `json:"store_id"`
	Qty           int       `json:"qty"`
	TotalPrice    int       `json:"total_price"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}

type UserCredentials struct {
	ID     int    `json:"id" gorm:"primary_key"`
	UserID int    `json:"user_id"`
	UUID   string `json:"uuid" gorm:"unique"`
}
