package handler

import (
	"fmt"
	"shop/models"
)

func (h *handler) CreateSale(branchId, shopAssistantId, cashierId, paymentType, status int, price float64, clientName string) {
	resp, err := h.strg.Sale().CreateSale(models.CreateSale{
		BranchID:        branchId,
		ShopAssistantID: shopAssistantId,
		CashierID:       cashierId,
		Price:           price,
		PaymentType:     paymentType,
		Status:          status,
		ClientName:      clientName,
	})
	if err != nil {
		fmt.Println("error from CreateSale:", err.Error())
		return
	}
	fmt.Println("created new sale with id:", resp)
}

func (h *handler) UpdateSale(id, branchId, shopAssistantId, cashierId, paymentType, status int, price float64, clientName string) {
	resp, err := h.strg.Sale().UpdateSale(models.Sale{
		Id:              id,
		BranchID:        branchId,
		ShopAssistantID: shopAssistantId,
		CashierID:       cashierId,
		Price:           price,
		PaymentType:     paymentType,
		Status:          status,
		ClientName:      clientName,
	})
	if err != nil {
		fmt.Println("error from UpdateSale:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetSale(id int) {
	resp, err := h.strg.Sale().GetSale(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetSale:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetAllSale(page, limit int) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Sale().GetAllSale(models.GetAllSaleRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		fmt.Println("error from GetAllSale:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) DeleteSale(id int) {
	resp, err := h.strg.Sale().DeleteSale(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteSale:", err.Error())
		return
	}
	fmt.Println(resp)
}