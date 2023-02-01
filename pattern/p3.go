package main
import "fmt"

func main(){
	for i:=0;i<5;i++{
		for j:=0;j<5-1-i;j++{
			fmt.Print(" ")
		}
		for j:=0;j<=i;j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}