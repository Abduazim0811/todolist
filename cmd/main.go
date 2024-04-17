package main

import (
	"fmt"
	"os"
	add "todolist/internal/Add"
	amal "todolist/internal/amallar"
	user "todolist/models"
)

func main() {
	var (
		users user.Users
		son   int
	)
	fmt.Printf("Sur_name: ")
	fmt.Scanln(&users.Surname)
	fmt.Printf("Name: ")
	fmt.Scanln(&users.Name)
	i := 0
	for i == 0 {
		fmt.Println("[1]Amallar royxati\n[2]Amallar qoshish\n[3]Exit")
		fmt.Scanln(&son)
		switch son {
		case 1:
			amal.Amallar(users)
		case 2:
			add.Add()
		case 3:
			os.Exit(0)
		}
	}

}
