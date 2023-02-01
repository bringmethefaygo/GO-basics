package main 

import "fmt"

func main(){
	// slicing array
	numbers:=[8]int{10,20,30,40,50,60,70,80}
	slicenumber := numbers[4:7]
	fmt.Println(slicenumber)

	// slice functions

	// append function
	primenumber :=[]int {2,3}
	primenumber=append(primenumber,5,7)
	fmt.Println(primenumber)

	// copy function
	p :=[]int{2,3,4,5}
	n :=[]int{1,2,3}

	copy(p,n)
	fmt.Println(p)

	


}