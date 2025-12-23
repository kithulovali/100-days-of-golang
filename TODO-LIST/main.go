package main

import (
	"fmt"
)

// To do structure
type Tasks struct {
	Title   string
	Content string
	Author  string
}

// Showing the added taks
func (t *Tasks) addedTask() {

	// Formatting the output into a key value format
	output := map[string]string{
		"Title":   t.Title,
		"Content": t.Content,
		"Author":  t.Author,
	}
	fmt.Println("The added Task is :", output)

}

// Getting user tasks for to do list
func (t *Tasks) intask() {
	fmt.Print("Task Title : ")
	fmt.Scanln(&t.Title)
	fmt.Print("Task Content : ")
	fmt.Scanln(&t.Content)
	fmt.Print("Task Auther : ")
	fmt.Scanln(&t.Author)
}

func main() {

	task1 := Tasks{}
	task1.intask()
	task1.addedTask()

}
