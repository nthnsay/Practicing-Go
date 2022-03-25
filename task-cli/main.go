package main

import (
	"bufio"
	"fmt"
	"os"
	//"flag"
)

func main() {
	greet := `task is a CLI for managing your TODOs.

	Usage:
	  task [command]
	
	Available Commands:
	  add         Add a new task to your TODO list
	  do          Mark a task on your TODO list as complete
	  list        List all of your incomplete tasks
	
	Use "task [command] --help" for more information about a command.`
	fmt.Println(greet)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
}
