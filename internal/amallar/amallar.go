package amallar

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	models "todolist/models"
)

func Amallar(user models.Users) {
	var (
		son       int
	)

	current := time.Now()
	file, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/users.txt",os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, user.Surname) && !strings.Contains(line, user.Name) {
			fmt.Println("Siz yangi user siz sizda amallar yoq!!!!!!!!!!!!")
			return
		}
	}

	file2, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/amallar.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	file3, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/bajarilgan_amallar.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()

	file4, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/bajarilmagan_amallar.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file4.Close()
	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		line := scanner2.Text()
		fmt.Println(line)
		if strings.Contains(line, "Amal vaqti:") {
			res := strings.Split(line, ",")
			res2 := strings.Split(res[0], "Amallar vaqti:")
			tm2, err := time.Parse("2006-01-02 08:00:00", res2[1])
			if err != nil {
				fmt.Println("Sana va vaqtni o'qishda xatolik bor: ", err)
			}
			fmt.Println("salom")
			if current.Before(tm2) {
				rs:=res[0] + "\n" + res[1]
				yozuvchi:=bufio.NewWriter(file3)
				fmt.Fprintln(yozuvchi,rs)
				yozuvchi.Flush()
			} else {
				rs:=res[0] + "\n" + res[1]
				yozuvchi:=bufio.NewWriter(file4)
				fmt.Fprintln(yozuvchi,rs)
				yozuvchi.Flush()
			}
		}

	}

	fmt.Println("[1]Bajarilgan amallar\t[2]Bajarilmagan amallar")
	fmt.Scanln(&son)
	switch son {
	case 1:
		// file3, err := os.Open("/home/abduazim/Projects/Golang/todolist/bajarilgan_amallar.txt")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer file3.Close()

		scanner := bufio.NewScanner(file3)
		fmt.Println("Bajarilgan amallar")
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	case 2:
		// file4, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/bajarilmagan_amallar.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer file4.Close()

		scanner:=bufio.NewScanner(file4)
		fmt.Println("Bajarilmagan amallar")
		for scanner.Scan(){
			line:=scanner.Text()
			fmt.Println(line)
		}
	}
	fmt.Println()
}
