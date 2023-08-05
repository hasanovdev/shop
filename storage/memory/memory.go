package memory

import "shop/storage"

type store struct {
	branches *branchRepo
}

func NewStorage() storage.StorageI {
	return &store{
		branches: NewBranchRepo(),
	}
}

func (s *store) Branch() storage.BranchesI {
	return s.branches
}
