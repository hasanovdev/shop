package handler

import (
	"fmt"
	"shop/models"
)

func (h *handler) CreateStaff(name, birthDate string, typ int, balance float64) {
	resp, err := h.strg.Staff().CreateStaff(models.CreateStaff{
		Name:      name,
		Type:      typ,
		Balance:   balance,
		BirthDate: birthDate,
	})
	if err != nil {
		fmt.Println("error from CreateStaff:", err.Error())
		return
	}
	fmt.Println("created new staff with id:", resp)
}

func (h *handler) UpdateStaff(typ int, id, name, birthDate string, balance float64) {
	resp, err := h.strg.Staff().UpdateStaff(models.Staff{
		Id:        id,
		Name:      name,
		Type:      typ,
		Balance:   balance,
		BirthDate: birthDate,
	})
	if err != nil {
		fmt.Println("error from UpdateStaff:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetStaff(id string) {
	resp, err := h.strg.Staff().GetStaff(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetStaff:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetAllStaff(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Staff().GetAllStaff(models.GetAllStaffRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})
	if err != nil {
		fmt.Println("error from GetAllStaff:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) DeleteStaff(id string) {
	resp, err := h.strg.Staff().DeleteStaff(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteStaff:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) ChangeBalance(id string, money float64) {
	resp, err := h.strg.Staff().ChangeBalance(models.ChangeBalanceStaff{Id: id, AddMoney: money})
	if err != nil {
		fmt.Println("error from ChangeBalance: ", err.Error())
		return
	}
	fmt.Println(resp)
}
