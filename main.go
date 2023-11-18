package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello to TODO app")

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()

	if *command == "create-task" {
		var name, duedate, category string

		fmt.Println("Please Enter the task title")
		scanner.Scan()
		name = scanner.Text()

		fmt.Println("Please Enter the taske duration")
		scanner.Scan()
		duedate = scanner.Text()

		fmt.Println("Please Enter the task category")
		scanner.Scan()
		category = scanner.Text()
		fmt.Println("task :", name, category, duedate)
	} else if *command == "create-category" {
		var title, color string

		fmt.Println("please enter the task title: ")
		scanner.Scan()
		title = scanner.Text()

		fmt.Println("please enter the task color: ")
		scanner.Scan()
		color = scanner.Text()

		fmt.Println("category", title, color)
	} else if *command == "register-user" {
		var id, email, password string

		fmt.Println("please enter the task id: ")
		scanner.Scan()
		id = scanner.Text()

		fmt.Println("please enter the task email: ")
		scanner.Scan()
		email = scanner.Text()

		fmt.Println("please enter the task password: ")
		scanner.Scan()
		password = scanner.Text()

		fmt.Println("register :", id, email, password)

	}
}
