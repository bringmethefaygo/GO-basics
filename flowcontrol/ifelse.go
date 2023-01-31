package main
import "fmt"

func main(){
	a := 3
	if(a>10){
		fmt.Println("great")
	}else if(a>2 && a<10){
		fmt.Println("okay")
	}else{
		fmt.Println("bad")
	}
}