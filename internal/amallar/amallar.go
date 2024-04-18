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

func Bajarilgan_amallar(soz string) {
	file3, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/bajarilgan_amallar.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()

	yozuvchi := bufio.NewWriter(file3)
	fmt.Fprintln(yozuvchi, soz)
	yozuvchi.Flush()

}

func Bajarilmagan_amallar(soz string) {
	file4, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/bajarilmagan_amallar.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file4.Close()

	yozuvchi := bufio.NewWriter(file4)
	fmt.Fprintln(yozuvchi, soz)
	yozuvchi.Flush()
}

func Amallar(user models.Users) {
	var (
		son int
	)

	current := time.Now()
	file, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/users.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, user.Surname) && !strings.Contains(line, user.Name) {
			soz := user.Name + user.Surname
			yozuvchi := bufio.NewWriter(file)
			fmt.Fprintln(yozuvchi, soz)
			yozuvchi.Flush()
			fmt.Println("Siz yangi user siz sizda amallar yoq!!!!!!!!!!!!")
			return
		}
	}

	file2, err := os.Open("/home/abduazim/Projects/Golang/todolist/amallar.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		line := scanner2.Text()
		// fmt.Println(line)
		if strings.Contains(line, "Amal vaqti:") {
			res := strings.Split(line, ",")
			if len(res) < 2 {
				fmt.Println("Invalid line format:", line)
				continue
			}
			res2 := strings.Split(res[0], "Amallar vaqti:")
			if len(res2) < 2 {
				fmt.Println("Invalid line format:", line)
				continue
			}
			tm2, err := time.Parse("2006-01-02 08:00:00", res2[1])
			if err != nil {
				fmt.Println("Error parsing time:", err)
			}
			fmt.Println("salom")
			if current.Before(tm2) {
				rs := res[0] + "\n" + res[1]
				Bajarilgan_amallar(rs)
			} else {
				rs := res[0] + "\n" + res[1]
				Bajarilmagan_amallar(rs)
			}
		}
	}
	if err := scanner2.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}

	fmt.Println("[1]Bajarilgan amallar\t[2]Bajarilmagan amallar")
	fmt.Scanln(&son)
	switch son {
	case 1:
		file3, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/bajarilgan_amallar.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file3.Close()

		scanner := bufio.NewScanner(file3)
		fmt.Println("Bajarilgan amallar")
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	case 2:
		file4, err := os.OpenFile("/home/abduazim/Projects/Golang/todolist/bajarilmagan_amallar.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file4.Close()

		scanner := bufio.NewScanner(file4)
		fmt.Println("Bajarilmagan amallar")
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}
	fmt.Println()
}
