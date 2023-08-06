package models

type CreateStaffTransaction struct {
	SaleId     int
	StaffId    int
	Type       string
	SourceType string
	Amount     int
	Text       string
}

type StaffTransaction struct {
	Id         int
	SaleId     int
	StaffId    int
	Type       string
	SourceType string
	Amount     int
	Text       string
	CreatedAt  string
}

type GetAllStaffTransactionRequest struct {
	Page  int
	Limit int
}

type GetAllStaffTransaction struct {
	StaffTransactions []StaffTransaction
	Count             int
}
