package memory

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"shop/models"
	"time"

	"github.com/google/uuid"
)

type branchRepo struct {
	fileName string
}

func NewBranchRepo(fn string) *branchRepo {
	return &branchRepo{fileName: fn}
}

func (x *branchRepo) CreateBranch(req models.CreateBranch) (string, error) {
	branches, err := x.read()
	if err != nil {
		return "", err
	}

	id := uuid.New().String()
	date, err := time.Parse("2006-01-02", req.FoundedAt)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	branches = append(branches, models.Branch{
		Id:        id,
		Name:      req.Name,
		Address:   req.Address,
		FoundedAt: req.FoundedAt,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Year:      time.Now().Year() - date.Year(),
	})

	err = x.write(branches)
	if err != nil {
		return "", err
	}

	return id, nil
	// x.branches = append(x.branches, models.Branch{
	// 	Id:        id,
	// 	Name:      req.Name,
	// 	Address:   req.Address,
	// 	FoundedAt: req.FoundedAt,
	// 	CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	// 	Year:      time.Now().Year() - date.Year(),
	// })
	// return id, nil
}

func (b *branchRepo) UpdateBranch(req models.Branch) (msg string, err error) {
	branches, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range branches {
		if v.Id == req.Id {
			branches[i] = req
			err = b.write(branches)
			if err != nil {
				return "", err
			}
			msg = "updated successfully"
			return msg, nil
		}
	}
	return "", errors.New("not found")
}

func (b *branchRepo) GetBranch(req models.IdRequest) (resp models.Branch, err error) {
	branches, err := b.read()
	if err != nil {
		return models.Branch{}, err
	}
	for _, v := range branches {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Branch{}, errors.New("not found")
}

func (b *branchRepo) GetAllBranch(req models.GetAllBranchRequest) (resp models.GetAllBranch, err error) {
	branches, err := b.read()
	if err != nil {
		return models.GetAllBranch{}, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(branches) {
		resp.Branches = []models.Branch{}
		resp.Count = len(branches)
		return resp, nil
	} else if end > len(branches) {
		return models.GetAllBranch{
			Branches: branches[start:],
			Count:    len(branches),
		}, nil
	}

	return models.GetAllBranch{
		Branches: branches[start:end],
		Count:    len(branches)}, nil
}

func (b *branchRepo) DeleteBranch(req models.IdRequest) (string, error) {
	branches, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range branches {
		if v.Id == req.Id {
			if i == (len(branches) - 1) {
				branches = branches[:i]
			} else {
				branches = append(branches[:i], branches[i+1:]...)
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")
}

func (u *branchRepo) read() ([]models.Branch, error) {
	var (
		branches []models.Branch
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return branches, nil
}

func (u *branchRepo) write(branches []models.Branch) error {

	body, err := json.Marshal(branches)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
