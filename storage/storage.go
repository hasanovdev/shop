package storage

import "shop/models"

type StorageI interface {
	Branch() BranchesI
	Staff() StaffsI
	StaffTarif() StaffTarifsI
}

type BranchesI interface {
	CreateBranch(models.CreateBranch) (int, error)
	UpdateBranch(models.Branch) (string, error)
	GetBranch(models.IdRequest) (models.Branch, error)
	GetAllBranch(models.GetAllBranchRequest) (models.GetAllBranch, error)
	DeleteBranch(models.IdRequest) (string, error)
}

type StaffsI interface {
	CreateStaff(models.CreateStaff) (int, error)
	UpdateStaff(models.Staff) (string, error)
	GetStaff(models.IdRequest) (models.Staff, error)
	GetAllStaff(models.GetAllStaffRequest) (models.GetAllStaff, error)
	DeleteStaff(models.IdRequest) (string, error)
}

type StaffTarifsI interface {
	CreateStaffTarif(models.CreateStaffTarif) (int, error)
	UpdateStaffTarif(models.StaffTarif) (string, error)
	GetStaffTarif(models.IdRequest) (models.StaffTarif, error)
	GetAllStaffTarif(models.GetAllStaffTarifRequest) (models.GetAllStaffTarif, error)
	DeleteStaffTarif(models.IdRequest) (string, error)
}
