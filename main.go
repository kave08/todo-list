package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello to TODO app")

	command := flag.String("create-todo-task", "no", "createing the todo task")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()

	if *command == "create-todo-task" {
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
	}

}
