package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID       int
	Title    string
	DueDate  string
	Category string
	IsDone   bool
	UserID   int
}

func (u User) print() {

	fmt.Println("User:", u.ID, u.Name, u.Email)
}

var userStorage []User
var taskStorage []Task
var authenticatedUser *User

func main() {
	fmt.Println("hello to TODO app")

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	for {

		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command :")
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(command string) {

	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		if authenticatedUser == nil {
			return
		}
	}
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCaategory()
	case "register-user":
		registerUser()
	case "list-task":
		listTask()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Command is not valid", command)
	}
}

func createTask() {
	if authenticatedUser != nil {
		authenticatedUser.print()
	}

	scanner := bufio.NewScanner(os.Stdin)

	var title, duedate, category string

	fmt.Println("Please Enter the task title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Please Enter the taske duration: ")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("Please Enter the task category: ")
	scanner.Scan()
	category = scanner.Text()

	task := Task{
		ID:       len(taskStorage) + 1,
		Title:    title,
		DueDate:  duedate,
		Category: category,
		IsDone:   false,
		UserID:   authenticatedUser.ID,
	}
	taskStorage = append(taskStorage, task)

	fmt.Println("task :", title, category, duedate)
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

	var id, email, password, name string

	fmt.Println("Please Enter your name: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter password: ")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println(" user :", id, email, password)

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

}

func login() {
	fmt.Println("login process")
	scanner := bufio.NewScanner(os.Stdin)

	var email, password string

	fmt.Println("please enter email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter password: ")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			authenticatedUser = &user

			break
		}
		fmt.Println("You are successfully loged in!!")
	}

	fmt.Println("user :", email, password)

	// user := User{
	// 	ID:       len(userStorage),
	// 	Email:    email,
	// 	Password: password,
	// }

	// userStorage = append(userStorage, user)

	if authenticatedUser == nil {
		fmt.Println("the email or password is not correct!")
	}

}

func listTask() {
	for _, task := range taskStorage {
		if task.ID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}
