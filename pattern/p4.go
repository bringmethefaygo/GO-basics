package main 
import "fmt"

func main(){
	for i:=0;i<4;i++{
		for j:=3;j>3-i;j--{
			fmt.Print(" ")
		}
		for j:=0;j<4-i;j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}