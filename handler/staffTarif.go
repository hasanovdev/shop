package handler

import (
	"fmt"
	"shop/models"
)

func (h *handler) CreateStaffTarif(name, foundedAt string, typ int, amountCash, amountCard float64) {
	resp, err := h.strg.StaffTarif().CreateStaffTarif(models.CreateStaffTarif{
		Name:          name,
		Type:          typ,
		AmountForCash: amountCash,
		AmountForCard: amountCard,
		FoundedAt:     foundedAt,
	})
	if err != nil {
		fmt.Println("error from CreateStaffTarif:", err.Error())
		return
	}
	fmt.Println("created new staffTarif with id:", resp)
}

func (h *handler) UpdateStaffTarif(id string, typ int, name, foundedAt string, amountCash, amountCard float64) {
	resp, err := h.strg.StaffTarif().UpdateStaffTarif(models.StaffTarif{
		Id:            id,
		Name:          name,
		Type:          typ,
		AmountForCash: amountCash,
		AmountForCard: amountCard,
		FoundedAt:     foundedAt,
	})
	if err != nil {
		fmt.Println("error from UpdateStaff:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetStaffTarif(id string) {
	resp, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetStaffTarif:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetAllStaffTarif(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.StaffTarif().GetAllStaffTarif(models.GetAllStaffTarifRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})
	if err != nil {
		fmt.Println("error from GetAllStaffTarif:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) DeleteStaffTarif(id string) {
	resp, err := h.strg.StaffTarif().DeleteStaffTarif(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteStaffTarif:", err.Error())
		return
	}
	fmt.Println(resp)
}
