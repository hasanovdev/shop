package main

import (
	"bufio"
	"fmt"
	"os"
	"shop/config"
	"shop/handler"
	"shop/storage/memory"
)

func main() {
	cfg := config.Load()
	strg := memory.NewStorage()
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
				fmt.Print("Enter Branch Address: ")
				address, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				fmt.Print("Enter Branch Founded At: ")
				foundedAt, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				h.CreateBranch(name, address, foundedAt)
			} else if method == "update" {
				id := 0
				fmt.Print("Enter Branch Id: ")
				fmt.Scan(&id)
				fmt.Print("Enter Branch Name: ")
				name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				fmt.Print("Enter Branch Address: ")
				address, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				fmt.Print("Enter Branch Founded at: ")
				foundedAt, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				h.UpdateBranch(id, name, address, foundedAt)
			} else if method == "get" {
				id := 0
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
				h.GetAllBranch(page, limit, search)
			} else if method == "delete" {
				id := 0
				fmt.Print("Enter Branch Id: ")
				fmt.Scan(&id)
				h.DeleteBranch(id)
			}
		} else if object == "staff" {
			if method == "create" {
				fmt.Print("Enter Staff Name: ")
				name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				typ := 0
				fmt.Print("Enter Staff Type: ")
				fmt.Scan(&typ)
				fmt.Print("Enter Staff Birthdate: ")
				birthDate, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				var balance float64
				fmt.Print("Enter Staff Balance: ")
				fmt.Scan(&balance)
				h.CreateStaff(name, birthDate, typ, balance)
			} else if method == "update" {
				id, typ := 0, 0
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
				h.UpdateStaff(id, typ, name, birthDate, balance)
			} else if method == "get" {
				id := 0
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
				id := 0
				fmt.Print("Enter Staff Id: ")
				fmt.Scan(&id)
				h.DeleteStaff(id)
			}
		}
	}
}
