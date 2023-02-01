package main 
import "fmt"

func main(){
	// creating an array in go
	var arrayint = [5]int{1,2,3,4,5}
	fmt.Println(arrayint)

	// initalizing array in go
	var arrayofint[3] int
	arrayofint[0] = 12
	arrayofint[1] = 11
	arrayofint[2] = 83
	fmt.Println(arrayofint)
	var a[5] int 
	fmt.Println(a)

	// initializing the elements of index 0 and 3 only
	arr := [5]int{0:7,3:9}
	fmt.Println(arr)

	// length of an array
	length :=len(arr)
	fmt.Println(length)

	// looping through the array
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

	// multidimensional array
	marray := [2][2]int{{1,2},{3,4}}
	fmt.Println(marray)
}