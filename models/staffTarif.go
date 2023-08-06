package models

type CreateStaffTarif struct {
	Name          string
	Type          int
	AmountForCash float64
	AmountForCard float64
	FoundedAt     string
}

type StaffTarif struct {
	Id            int
	Name          string
	Type          int
	AmountForCash float64
	AmountForCard float64
	CreatedAt     string
	FoundedAt     string
}

type GetAllStaffTarifRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllStaffTarif struct {
	StaffTarifs []StaffTarif
	Count       int
}
