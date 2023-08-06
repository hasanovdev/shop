package memory

import "shop/storage"

type store struct {
	branches *branchRepo
	staffs   *staffRepo
}

func NewStorage() storage.StorageI {
	return &store{
		branches: NewBranchRepo(),
		staffs:   NewStaffRepo(),
	}
}

func (s *store) Branch() storage.BranchesI {
	return s.branches
}

func (s *store) Staff() storage.StaffsI {
	return s.staffs
}
