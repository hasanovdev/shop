package models

type CreateStaffTransaction struct {
	Type       int
	SourceType string
	Amount     int
	Text       string
}

type StaffTransaction struct {
	Id         string
	SaleId     string
	StaffId    string
	Type       int
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
