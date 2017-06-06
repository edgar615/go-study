package main

import (
	"fmt"

	"go_study/ch05/listing74/entities"
)

// main is the entry point for the application.
func main() {
	// Create a value of type Admin from the entities package.
	a := entities.Admin{
		Rights: 10,
	}

	// Set the exported fields from the unexported
	// inner type.
	//即使内部类型是未公开的，内部类型里声明的字段依旧是公开的
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}