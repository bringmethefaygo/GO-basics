package main
import "fmt"

func main(){


	var number1 int = 10
	fmt.Println(number1)

	var number2 = 20
	fmt.Println(number2)

	number3 := 30
	fmt.Println(number3)

	// creating multiple variable at once
	var name,age = "Palistha",22
	fmt.Println(name)
	fmt.Println(age)

	name,age = "Abhishek",20
	fmt.Println(name)
	fmt.Println(age)

	// contants
	const speed = 2000000000//initial value
	// speed  = 30//error!!!!
}