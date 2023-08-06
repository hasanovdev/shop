package handler

import (
	"fmt"
	"shop/models"
)

func (h *handler) CreateStaffTransaction(saleId, staffId, amount int, typ, sourceType, text string) {
	resp, err := h.strg.StaffTransaction().CreateStaffTransaction(models.CreateStaffTransaction{
		SaleId:     saleId,
		StaffId:    staffId,
		Type:       typ,
		SourceType: sourceType,
		Amount:     amount,
		Text:       text,
	})
	if err != nil {
		fmt.Println("error from CreateStaffTransaction:", err.Error())
		return
	}
	fmt.Println("created new staffTransaction with id:", resp)
}

func (h *handler) UpdateStaffTransaction(id, saleId, staffId, amount int, typ, sourceType, text string) {
	resp, err := h.strg.StaffTransaction().UpdateStaffTransaction(models.StaffTransaction{
		Id:         id,
		SaleId:     saleId,
		StaffId:    staffId,
		Type:       typ,
		SourceType: sourceType,
		Amount:     amount,
		Text:       text,
	})
	if err != nil {
		fmt.Println("error from UpdateStaffTransaction:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetStaffTransaction(id int) {
	resp, err := h.strg.StaffTransaction().GetStaffTransaction(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetStaffTransaction:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetAllStaffTransaction(page, limit int) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.StaffTransaction().GetAllStaffTransaction(models.GetAllStaffTransactionRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		fmt.Println("error from GetAllStaffTransaction:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) DeleteStaffTransaction(id int) {
	resp, err := h.strg.StaffTransaction().DeleteStaffTransaction(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteStaffTransaction:", err.Error())
		return
	}
	fmt.Println(resp)
}
