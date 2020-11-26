package model

// Sale it's an item of the sale
type Sale struct {
	ID          int
	Amounts     []int
	TotalAmount int
	Status      SaleStatus
}

// SaleStatus enum
type SaleStatus string

// Different status of the sale
const (
	PENDING  SaleStatus = "pending"
	FINISH   SaleStatus = "finish"
	CANCELED SaleStatus = "canceled"
)
