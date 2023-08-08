package memory

import (
	"errors"
	"shop/models"
	"time"

	"github.com/google/uuid"
)

type staffRepo struct {
	staffs []models.Staff
}

func NewStaffRepo() *staffRepo {
	return &staffRepo{staffs: make([]models.Staff, 0)}
}

func (x *staffRepo) CreateStaff(req models.CreateStaff) (string, error) {
	id := uuid.New().String()
	date, _ := time.Parse("2006-01-02", req.BirthDate)
	x.staffs = append(x.staffs, models.Staff{
		Id:        id,
		Name:      req.Name,
		Type:      req.Type,
		Balance:   req.Balance,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		BirthDate: req.BirthDate,
		Age:       time.Now().Year() - date.Year(),
	})
	return id, nil
}

func (s *staffRepo) UpdateStaff(req models.Staff) (msg string, err error) {
	for i, v := range s.staffs {
		if v.Id == req.Id {
			s.staffs[i] = req
			msg = "updated successfully"
			return
		}
	}
	return "", errors.New("not found")
}

func (s *staffRepo) GetStaff(req models.IdRequest) (resp models.Staff, err error) {
	for _, v := range s.staffs {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Staff{}, errors.New("not found")
}

func (s *staffRepo) GetAllStaff(req models.GetAllStaffRequest) (resp models.GetAllStaff, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(s.staffs) {
		resp.Staffs = []models.Staff{}
		resp.Count = len(s.staffs)
		return resp, nil
	} else if end > len(s.staffs) {
		return models.GetAllStaff{
			Staffs: s.staffs[start:],
			Count:  len(s.staffs),
		}, nil
	}

	return models.GetAllStaff{
		Staffs: s.staffs[start:end],
		Count:  len(s.staffs)}, nil
}

func (s *staffRepo) DeleteStaff(req models.IdRequest) (string, error) {
	for i, v := range s.staffs {
		if v.Id == req.Id {
			if i == (len(s.staffs) - 1) {
				s.staffs = s.staffs[:i]
			} else {
				s.staffs = append(s.staffs[:i], s.staffs[i+1:]...)
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")
}

func (s *staffRepo) ChangeBalance(req models.Staff) (string, error) {
	for i, v := range s.staffs {
		if v.Id == req.Id {
			s.staffs[i].Balance = req.Balance
			return "changed successfully", nil
		}
	}
	return "", errors.New("not found")
}
