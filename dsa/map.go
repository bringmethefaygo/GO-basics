package main 
import "fmt"


func main(){
	// change value of map in go-lang

	capital := map[string]string{"nepal":"kathmandu","us":"new york"}
	fmt.Println(capital)
	capital["us"] = "washington dc"
	fmt.Println(capital)

	// add element to go map
	student:=map[int]string{1:"john",2:"lily"}
	student[3] = "robin"
	student[4] = "julie"
	fmt.Println(student)

	onemore :=map[string]string{"one":"okay","two":"best","three":"great"}
	onemore["zour"] = "superb"
	onemore["zive"] = "excellent"
	fmt.Println(onemore)

	// delete element of go map
	person:=map[string]int{"Hermione":21,"Harry":20,"John":25}
	delete(person,"John")
	fmt.Println(person)


}