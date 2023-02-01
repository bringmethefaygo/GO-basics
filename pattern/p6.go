package main 
import "fmt"

func main(){
	for i:=0;i<5;i++{
		for j:=0;j<5-i;j++{
			fmt.Print("* ")
		}
		for j:=0;j<i*2;j++{
			fmt.Print("  ")
		}
		for j:=0;j<5-i;j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}

	for i:=1;i<5;i++{
		for j:=0;j<=i;j++{
			fmt.Print("* ")
		}
		for j:=0;j<(5-i-1)*2;j++{
			fmt.Print("  ")
		}
		for j:=0;j<=i;j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

