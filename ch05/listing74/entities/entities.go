package entities
// user defines a user in the program.
type user struct {
	Name  string
	Email string
}

// Admin defines an admin in the program.
type Admin struct {
	//嵌入的类型是未公开的
	user   // The embedded type is unexported.
	Rights int
}