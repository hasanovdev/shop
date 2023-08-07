package memory

import "shop/storage"

type store struct {
	branches          *branchRepo
	staffs            *staffRepo
	staffTarifs       *staffTarifRepo
	staffTransactions *staffTransactionRepo
	sales             *saleRepo
}

func NewStorage() storage.StorageI {
	return &store{
		branches:          NewBranchRepo(),
		staffs:            NewStaffRepo(),
		staffTarifs:       NewStaffTarifRepo(),
		staffTransactions: NewStaffTransactionsRepo(),
		sales:             NewSaleRepo(),
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

func (s *store) StaffTransaction() storage.StaffTransactionsI {
	return s.staffTransactions
}

func (s *store) Sale() storage.SalesI {
	return s.sales
}
