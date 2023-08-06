package memory

import (
	"errors"
	"shop/models"
	"time"
)

type staffTarifRepo struct {
	staffTarifs []models.StaffTarif
}

func NewStaffTarifRepo() *staffTarifRepo {
	return &staffTarifRepo{staffTarifs: make([]models.StaffTarif, 0)}
}

func (x *staffTarifRepo) CreateStaffTarif(req models.CreateStaffTarif) (int, error) {
	var id int
	if len(x.staffTarifs) == 0 {
		id = 1
	} else {
		id = x.staffTarifs[len(x.staffTarifs)-1].Id + 1
	}

	x.staffTarifs = append(x.staffTarifs, models.StaffTarif{
		Id:            id,
		Name:          req.Name,
		Type:          req.Type,
		AmountForCash: req.AmountForCash,
		AmountForCard: req.AmountForCard,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		FoundedAt:     req.FoundedAt,
	})
	return id, nil
}

func (s *staffTarifRepo) UpdateStaffTarif(req models.StaffTarif) (msg string, err error) {
	for i, v := range s.staffTarifs {
		if v.Id == req.Id {
			s.staffTarifs[i] = req
			msg = "updated successfully"
			return
		}
	}
	return "", errors.New("not found")
}

func (s *staffTarifRepo) GetStaffTarif(req models.IdRequest) (resp models.StaffTarif, err error) {
	for _, v := range s.staffTarifs {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.StaffTarif{}, errors.New("not found")
}

func (s *staffTarifRepo) GetAllStaffTarif(req models.GetAllStaffTarifRequest) (resp models.GetAllStaffTarif, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(s.staffTarifs) {
		resp.StaffTarifs = []models.StaffTarif{}
		resp.Count = len(s.staffTarifs)
		return resp, nil
	} else if end > len(s.staffTarifs) {
		return models.GetAllStaffTarif{
			StaffTarifs: s.staffTarifs[start:],
			Count:       len(s.staffTarifs),
		}, nil
	}

	return models.GetAllStaffTarif{
		StaffTarifs: s.staffTarifs[start:end],
		Count:       len(s.staffTarifs)}, nil
}

func (s *staffTarifRepo) DeleteStaffTarif(req models.IdRequest) (string, error) {
	for i, v := range s.staffTarifs {
		if v.Id == req.Id {
			if i == (len(s.staffTarifs) - 1) {
				s.staffTarifs = s.staffTarifs[:i]
			} else {
				s.staffTarifs = append(s.staffTarifs[:i], s.staffTarifs[i+1:]...)
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")
}
