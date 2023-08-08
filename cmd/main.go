package main

import (
	"bufio"
	"fmt"
	"os"
	"shop/config"
	"shop/handler"
	"shop/storage/memory"
)

var id string
var typ int

func main() {
	// fmt.Println(os.Getwd())
	// _, err := os.ReadFile("../data/branches.json")
	// fmt.Println(err)
	// return
	cfg := config.Load()
	strg := memory.NewStorage("../data/branches.json")
	h := handler.NewHandler(strg, *cfg)

	fmt.Println("Welcome to my program!")
	fmt.Println("Aviable methods:")
	for _, m := range cfg.Methods {
		fmt.Println("-", m)
	}
	fmt.Println("Available objects:")
	for _, o := range cfg.Objects {
		fmt.Println("-", o)
	}

	for {
		fmt.Println("Enter method and object:")
		method, object := "", ""
		fmt.Scan(&method, &object)

		if object == "branch" {
			if method == "create" {
				fmt.Print("Enter Branch Name: ")
				name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				name = name[:len(name)-1]
				fmt.Print("Enter Branch Address: ")
				address, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				address = address[:len(address)-1]
				fmt.Print("Enter Branch Founded At: ")
				var foundedAt string
				fmt.Scan(&foundedAt)
				h.CreateBranch(name, address, foundedAt)
			} else if method == "update" {
				fmt.Print("Enter Branch Id: ")
				fmt.Scan(&id)
				fmt.Print("Enter Branch Name: ")
				name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				name = name[:len(name)-1]
				fmt.Print("Enter Branch Address: ")
				address, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				address = address[:len(address)-1]
				fmt.Print("Enter Branch Founded at: ")
				foundedAt, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				h.UpdateBranch(id, name, address, foundedAt)
			} else if method == "get" {
				fmt.Print("Enter Branch Id: ")
				fmt.Scan(&id)
				h.GetBranch(id)
			} else if method == "getAll" {
				page, limit, search := 1, 10, ""
				fmt.Print("Enter page number: ")
				fmt.Scan(&page)
				fmt.Print("Enter page limit: ")
				fmt.Scan(&limit)
				fmt.Print("Enter search name: ")
				search, _ = bufio.NewReader(os.Stdin).ReadString('\n')
				search = search[:len(search)-1]
				h.GetAllBranch(page, limit, search)
			} else if method == "delete" {
				fmt.Print("Enter Branch Id: ")
				fmt.Scan(&id)
				h.DeleteBranch(id)
			}
		} else if object == "staff" {
			if method == "create" {
				fmt.Print("Enter Staff Name: ")
				name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				name = name[:len(name)-1]
				fmt.Print("Enter Staff Type: ")
				fmt.Scan(&typ)
				fmt.Print("Enter Staff Birthdate: ")
				birthDate, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				birthDate = birthDate[:len(birthDate)-1]
				var balance float64
				fmt.Print("Enter Staff Balance: ")
				fmt.Scan(&balance)
				h.CreateStaff(name, birthDate, typ, balance)
			} else if method == "update" {
				fmt.Print("Enter Staff Id: ")
				fmt.Scan(&id)
				fmt.Print("Enter Staff Name: ")
				name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				fmt.Print("Enter Staff Type: ")
				fmt.Scan(&typ)
				fmt.Print("Enter Staff Birthdate: ")
				birthDate, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				var balance float64
				fmt.Scan(&balance)
				h.UpdateStaff(typ, id, name, birthDate, balance)
			} else if method == "get" {
				fmt.Print("Enter Staff Id: ")
				fmt.Scan(&id)
				h.GetStaff(id)
			} else if method == "getAll" {
				page, limit, search := 1, 10, ""
				fmt.Print("Enter page number: ")
				fmt.Scan(&page)
				fmt.Print("Enter page limit: ")
				fmt.Scan(&limit)
				fmt.Print("Enter search name: ")
				search, _ = bufio.NewReader(os.Stdin).ReadString('\n')
				h.GetAllStaff(page, limit, search)
			} else if method == "delete" {
				fmt.Print("Enter Staff Id: ")
				fmt.Scan(&id)
				h.DeleteStaff(id)
			}
		} else if object == "staffTarif" {
			if method == "create" {
				fmt.Print("Enter StaffTarif Name: ")
				name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				fmt.Print("Enter StaffTarif Type: ")
				fmt.Scan(&typ)
				fmt.Print("Enter StaffTarif foundedAt: ")
				foundedAt, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				var amountCash, amountCard float64
				fmt.Print("Enter AmountForCash: ")
				fmt.Scan(&amountCash)
				fmt.Print("Enter AmountForCard: ")
				fmt.Scan(&amountCard)
				h.CreateStaffTarif(name, foundedAt, typ, amountCash, amountCard)
			} else if method == "update" {
				fmt.Print("Enter StaffTarif Id: ")
				fmt.Scan(&id)
				fmt.Print("Enter StaffTarif Type: ")
				fmt.Scan(&typ)
				fmt.Print("Enter StaffTarif Name: ")
				name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				fmt.Print("Enter StaffTarif FoundedAt: ")
				foundedAt, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				var amountCash, amountCard float64
				fmt.Print("Enter AmountForCash: ")
				fmt.Scan(&amountCash)
				fmt.Print("Enter AmountForCard: ")
				fmt.Scan(&amountCard)
				h.UpdateStaffTarif(id, typ, name, foundedAt, amountCash, amountCard)
			} else if method == "get" {
				fmt.Print("Enter StaffTarif Id: ")
				fmt.Scan(&id)
				h.GetStaffTarif(id)
			} else if method == "getAll" {
				page, limit, search := 1, 10, ""
				fmt.Print("Enter page number: ")
				fmt.Scan(&page)
				fmt.Print("Enter page limit: ")
				fmt.Scan(&limit)
				fmt.Print("Enter search name: ")
				search, _ = bufio.NewReader(os.Stdin).ReadString('\n')
				h.GetAllStaffTarif(page, limit, search)
			} else if method == "delete" {
				fmt.Print("Enter StaffTarif Id: ")
				fmt.Scan(&id)
				h.DeleteStaffTarif(id)
			}
		} else if object == "staffTransaction" {
			if method == "create" {
				amount, typ := 0.0, 0
				fmt.Print("Enter Amount: ")
				fmt.Scan(&amount)
				fmt.Print("Enter Transaction Type: ")
				fmt.Scan(&typ)
				fmt.Print("Enter sourceType: ")
				sourceType, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				fmt.Print("Enter text for transaction: ")
				text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				h.CreateStaffTransaction(amount, typ, sourceType, text)
			} else if method == "update" {
				fmt.Print("Enter transaction Id: ")
				fmt.Scan(&id)
				amount, typ := 0.0, 0
				fmt.Print("Enter Amount: ")
				fmt.Scan(&amount)
				fmt.Print("Enter Transaction Type: ")
				fmt.Scan(&typ)
				fmt.Print("Enter sourceType: ")
				sourceType, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				fmt.Print("Enter text for transaction: ")
				text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				h.UpdateStaffTransaction(id, amount, typ, sourceType, text)
			} else if method == "get" {
				fmt.Print("Enter transaction Id: ")
				fmt.Scan(&id)
				h.GetStaffTransaction(id)
			} else if method == "getAll" {
				page, limit := 1, 10
				fmt.Print("Enter page number: ")
				fmt.Scan(&page)
				fmt.Print("Enter page limit: ")
				fmt.Scan(&limit)
				h.GetAllStaffTransaction(page, limit)
			} else if method == "delete" {
				fmt.Print("Enter transaction Id: ")
				fmt.Scan(&id)
				h.DeleteStaffTransaction(id)
			}
		} else if object == "sale" {
			if method == "create" {
				var (
					paymentType, status int
					price               float64
				)
				fmt.Print("Enter payment type: ")
				fmt.Scan(&paymentType)
				fmt.Print("Enter status: ")
				fmt.Scan(&status)
				fmt.Print("Enter Price: ")
				fmt.Scan(&price)
				fmt.Print("Enter client name: ")
				clientName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				clientName = clientName[:len(clientName)-1]
				fmt.Print("Enter StaffId: ")
				staffId, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				staffId = staffId[:len(staffId)-1]
				h.CreateSale(paymentType, status, price, clientName, staffId)
			} else if method == "update" {
				fmt.Print("Enter Sale Id: ")
				fmt.Scan(&id)
				var (
					paymentType, status int
					price               float64
				)
				fmt.Print("Enter payment type: ")
				fmt.Scan(&paymentType)
				fmt.Print("Enter status: ")
				fmt.Scan(&status)
				fmt.Print("Enter Price: ")
				fmt.Scan(&price)
				fmt.Print("Enter client name: ")
				clientName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				h.UpdateSale(id, paymentType, status, price, clientName)
			} else if method == "get" {
				fmt.Print("Enter Sale Id: ")
				fmt.Scan(&id)
				h.GetSale(id)
			} else if method == "getAll" {
				page, limit := 1, 10
				fmt.Print("Enter page number: ")
				fmt.Scan(&page)
				fmt.Print("Enter page limit: ")
				fmt.Scan(&limit)
				h.GetAllSale(page, limit)
			} else if method == "delete" {
				fmt.Print("Enter Sale Id: ")
				fmt.Scan(&id)
				h.DeleteSale(id)
			}
		}
	}
}
