package add

import (
	"bufio"
	"fmt"
	"log"
	"os"
	md "todolist/models"
)

func Add() {
	var amal md.Amal

	fmt.Printf("Name: ")
	fmt.Scanln(&amal.Name)
	fmt.Printf("Year: ")
	fmt.Scanln(&amal.Year)
	fmt.Printf("Month: ")
	fmt.Scanln(&amal.Month)
	fmt.Printf("Day: ")
	fmt.Scanln(&amal.Day)
	fmt.Printf("Hours: ")
	fmt.Scanln(&amal.Hours)
	fmt.Printf("Minute: ")
	fmt.Scanln(&amal.Minute)
	fmt.Printf("Secund: ")
	fmt.Scanln(&amal.Secund)
	time := amal.Year + "-" + amal.Month + "-" + amal.Day + "-" + " " + amal.Hours + ":" + amal.Minute + ":" + amal.Secund
	file, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/amallar.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res:="Amal vaqti: " + time + "," + "Amal nomi: " + amal.Name + "\n"

	yozuvchi:=bufio.NewWriter(file)
	fmt.Fprintln(yozuvchi,res)
	yozuvchi.Flush()

	fmt.Println("Amallar qo'shildi")
}
