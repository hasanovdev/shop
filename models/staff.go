package models

type CreateStaff struct {
	Name      string
	Type      int
	Balance   float64
	BirthDate string
}

type Staff struct {
	Id        string
	Name      string
	Type      int
	Balance   float64
	CreatedAt string
	BirthDate string
	Age       int
}

type GetAllStaffRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllStaff struct {
	Staffs []Staff
	Count  int
}
