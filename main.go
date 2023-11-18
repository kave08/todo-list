package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Email    string
	Password string
}

var userStorage []User

func main() {
	fmt.Println("hello to TODO app")

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command :")
		scanner.Scan()
		scanner.Text()
	}

	fmt.Printf("userStorage : %v \n", userStorage)

}

func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCaategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		os.Exit(1)
	default:
		fmt.Println("Command is not valid", command)
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)

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
}
func createCaategory() {
	scanner := bufio.NewScanner(os.Stdin)

	var title, color string

	fmt.Println("please enter the task title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task color: ")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("category", title, color)
}
func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)

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

	fmt.Println(" user :", id, email, password)

	user := User{
		ID:       len(userStorage),
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

}
func login() {
	// scanner := bufio.NewScanner(os.Stdin)
}
