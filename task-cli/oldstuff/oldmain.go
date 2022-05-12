package notmain

import (
	"bufio"
	"fmt"
	"os"
	//"flag"
)

var tasks []Task

const (
	AddTask   string = "add"
	DoTask    string = "do"
	ListTasks string = "list"
)

func Init() {
	tasks = []Task{
		{
			TaskId:      1,
			IsCompleted: false,
			Description: "Eat cookies",
		},
	}
}

//TODO: add commmand structs
func main() {
	Init()
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
	//fmt.Println(input.Text())
	HandleInput(input.Text())
	fmt.Println(tasks)

}

func HandleInput(inputStr string) {
	//Seperate the string into three parts
	//First part is the command
	//Second part is the description or the ID
	switch inputStr {
	case AddTask:
		//add()
	case DoTask:
	case ListTasks:
	default:
		fmt.Println("This is unsupported")
	}
}

type Task struct {
	TaskId      int
	IsCompleted bool `default:false`
	Description string
}

func add(description string) {
	t := Task{
		TaskId:      len(tasks) + 1,
		Description: description,
	}
	tasks = append(tasks, t)
	fmt.Println(tasks)
	// TODO return sth / error handling
}

func cist() {

}

func completeTask() {

}
