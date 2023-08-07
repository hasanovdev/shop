package models

type CreateSale struct {
	BranchID        int
	ShopAssistantID int
	CashierID       int
	Price           float64
	PaymentType     int //card, cash
	Status          int // success,cancel
	ClientName      string
}

type Sale struct {
	Id              int
	BranchID        int
	ShopAssistantID int
	CashierID       int
	Price           float64
	PaymentType     int //card, cash
	Status          int // success,cancel
	ClientName      string
	CreatedAt       string
}

type GetAllSaleRequest struct {
	Page  int
	Limit int
}

type GetAllSale struct {
	Sales []Sale
	Count int
}
