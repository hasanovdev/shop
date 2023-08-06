package memory

import (
	"errors"
	"shop/models"
	"time"
)

type staffTransactionRepo struct {
	staffTransactions []models.StaffTransaction
}

func NewStaffTransactionsRepo() *staffTransactionRepo {
	return &staffTransactionRepo{staffTransactions: make([]models.StaffTransaction, 0)}
}

func (x *staffTransactionRepo) CreateStaffTransaction(req models.CreateStaffTransaction) (int, error) {
	var id int
	if len(x.staffTransactions) == 0 {
		id = 1
	} else {
		id = x.staffTransactions[len(x.staffTransactions)-1].Id + 1
	}

	x.staffTransactions = append(x.staffTransactions, models.StaffTransaction{
		Id:         id,
		SaleId:     req.SaleId,
		StaffId:    req.StaffId,
		Type:       req.Type,
		SourceType: req.SourceType,
		Amount:     req.Amount,
		Text:       req.Text,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	})
	return id, nil
}

func (s *staffTransactionRepo) UpdateStaffTransaction(req models.StaffTransaction) (msg string, err error) {
	for i, v := range s.staffTransactions {
		if v.Id == req.Id {
			s.staffTransactions[i] = req
			msg = "updated successfully"
			return
		}
	}
	return "", errors.New("not found")
}

func (s *staffTransactionRepo) GetStaffTransaction(req models.IdRequest) (resp models.StaffTransaction, err error) {
	for _, v := range s.staffTransactions {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.StaffTransaction{}, errors.New("not found")
}

func (s *staffTransactionRepo) GetAllStaffTransaction(req models.GetAllStaffTransactionRequest) (resp models.GetAllStaffTransaction, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(s.staffTransactions) {
		resp.StaffTransactions = []models.StaffTransaction{}
		resp.Count = len(s.staffTransactions)
		return resp, nil
	} else if end > len(s.staffTransactions) {
		return models.GetAllStaffTransaction{
			StaffTransactions: s.staffTransactions[start:],
			Count:             len(s.staffTransactions),
		}, nil
	}

	return models.GetAllStaffTransaction{
		StaffTransactions: s.staffTransactions[start:end],
		Count:             len(s.staffTransactions)}, nil
}

func (s *staffTransactionRepo) DeleteStaffTransaction(req models.IdRequest) (string, error) {
	for i, v := range s.staffTransactions {
		if v.Id == req.Id {
			if i == (len(s.staffTransactions) - 1) {
				s.staffTransactions = s.staffTransactions[:i]
			} else {
				s.staffTransactions = append(s.staffTransactions[:i], s.staffTransactions[i+1:]...)
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")
}
