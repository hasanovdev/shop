package memory

import (
	"errors"
	"shop/models"
	"time"
)

type saleRepo struct {
	sales []models.Sale
}

func NewSaleRepo() *saleRepo {
	return &saleRepo{sales: make([]models.Sale, 0)}
}

func (x *saleRepo) CreateSale(req models.CreateSale) (int, error) {
	var id int
	if len(x.sales) == 0 {
		id = 1
	} else {
		id = x.sales[len(x.sales)-1].Id + 1
	}

	x.sales = append(x.sales, models.Sale{
		Id:              id,
		BranchID:        req.BranchID,
		ShopAssistantID: req.ShopAssistantID,
		CashierID:       req.CashierID,
		Price:           req.Price,
		PaymentType:     req.PaymentType,
		Status:          req.Status,
		ClientName:      req.ClientName,
		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
	})
	return id, nil
}

func (s *saleRepo) UpdateSale(req models.Sale) (msg string, err error) {
	for i, v := range s.sales {
		if v.Id == req.Id {
			s.sales[i] = req
			msg = "updated successfully"
			return
		}
	}
	return "", errors.New("not found")
}

func (s *saleRepo) GetSale(req models.IdRequest) (resp models.Sale, err error) {
	for _, v := range s.sales {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Sale{}, errors.New("not found")
}

func (s *saleRepo) GetAllSale(req models.GetAllSaleRequest) (resp models.GetAllSale, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(s.sales) {
		resp.Sales = []models.Sale{}
		resp.Count = len(s.sales)
		return resp, nil
	} else if end > len(s.sales) {
		return models.GetAllSale{
			Sales: s.sales[start:],
			Count: len(s.sales),
		}, nil
	}

	return models.GetAllSale{
		Sales: s.sales[start:end],
		Count: len(s.sales)}, nil
}

func (s *saleRepo) DeleteSale(req models.IdRequest) (string, error) {
	for i, v := range s.sales {
		if v.Id == req.Id {
			if i == (len(s.sales) - 1) {
				s.sales = s.sales[:i]
			} else {
				s.sales = append(s.sales[:i], s.sales[i+1:]...)
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")
}
