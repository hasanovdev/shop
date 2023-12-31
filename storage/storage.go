package storage

import "shop/models"

type StorageI interface {
	Branch() BranchesI
	Staff() StaffsI
	StaffTarif() StaffTarifsI
	StaffTransaction() StaffTransactionsI
	Sale() SalesI
}

type BranchesI interface {
	CreateBranch(models.CreateBranch) (string, error)
	UpdateBranch(models.Branch) (string, error)
	GetBranch(models.IdRequest) (models.Branch, error)
	GetAllBranch(models.GetAllBranchRequest) (models.GetAllBranch, error)
	DeleteBranch(models.IdRequest) (string, error)
}

type StaffsI interface {
	CreateStaff(models.CreateStaff) (string, error)
	UpdateStaff(models.Staff) (string, error)
	GetStaff(models.IdRequest) (models.Staff, error)
	GetAllStaff(models.GetAllStaffRequest) (models.GetAllStaff, error)
	DeleteStaff(models.IdRequest) (string, error)
	ChangeBalance(models.ChangeBalanceStaff) (string, error)
}

type StaffTarifsI interface {
	CreateStaffTarif(models.CreateStaffTarif) (string, error)
	UpdateStaffTarif(models.StaffTarif) (string, error)
	GetStaffTarif(models.IdRequest) (models.StaffTarif, error)
	GetAllStaffTarif(models.GetAllStaffTarifRequest) (models.GetAllStaffTarif, error)
	DeleteStaffTarif(models.IdRequest) (string, error)
}

type StaffTransactionsI interface {
	CreateStaffTransaction(models.CreateStaffTransaction) (string, error)
	UpdateStaffTransaction(models.StaffTransaction) (string, error)
	GetStaffTransaction(models.IdRequest) (models.StaffTransaction, error)
	GetAllStaffTransaction(models.GetAllStaffTransactionRequest) (models.GetAllStaffTransaction, error)
	DeleteStaffTransaction(models.IdRequest) (string, error)
}

type SalesI interface {
	CreateSale(models.CreateSale) (string, error)
	UpdateSale(models.Sale) (string, error)
	GetSale(models.IdRequest) (models.Sale, error)
	GetAllSale(models.GetAllSaleRequest) (models.GetAllSale, error)
	DeleteSale(models.IdRequest) (string, error)
}
