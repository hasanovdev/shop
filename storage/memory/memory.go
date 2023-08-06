package memory

import "shop/storage"

type store struct {
	branches    *branchRepo
	staffs      *staffRepo
	staffTarifs *staffTarifRepo
}

func NewStorage() storage.StorageI {
	return &store{
		branches:    NewBranchRepo(),
		staffs:      NewStaffRepo(),
		staffTarifs: NewStaffTarifRepo(),
	}
}

func (s *store) Branch() storage.BranchesI {
	return s.branches
}

func (s *store) Staff() storage.StaffsI {
	return s.staffs
}

func (s *store) StaffTarif() storage.StaffTarifsI {
	return s.staffTarifs
}
