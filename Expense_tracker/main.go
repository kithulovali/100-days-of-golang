package main

import (
	"Expense_tracker/types"
	"fmt"
)

func main() {
	fname := "goffart"
	lname := "jean marc"
	fullName := types.RenderName(fname, lname)
	fmt.Println(fullName)
}
