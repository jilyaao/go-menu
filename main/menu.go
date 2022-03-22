package main

import "fmt"

func main() {
	var command string
	for {
		fmt.Println("----------------------------")
		fmt.Println("p：print words")
		fmt.Println("q：exit")
		fmt.Print("please input command:")
		fmt.Scanln(&command)
		fmt.Println("----------------------------")
		switch command {
		case "p":
			fmt.Println("I print something!")
		case "q":
			return
		default:
			fmt.Println("Command is undefined!")
		}

	}
}
