package main
import "fmt"
//var p = fmt.Println

type Person struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

func main() {
	fmt.Println("Starting the application...")
}