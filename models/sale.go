package models

type CreateSale struct {
	Price       float64
	PaymentType int //card, cash
	Status      int // success,cancel
	ClientName  string
	StaffId     string	
}

type Sale struct {
	Id              string
	BranchID        string
	ShopAssistantID string
	CashierID       string
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
