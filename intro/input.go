package main
import "fmt"

func main(){
	var name string 
	var age int
	fmt.Println("Enter the name & age")
	fmt.Scan(&name, &age)

	fmt.Println("Name:",name)
	fmt.Println("age:",age)
}