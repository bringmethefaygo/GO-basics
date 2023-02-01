package main
import "fmt"

func main(){
	for i:=0;i<5;i++{
		for j:=0;j<5-1-i;j++{
			fmt.Print(" ")
		}
		if(i==0 || i==1 ||i==4){
			for j:=0;j<=i;j++{
				fmt.Print("* ")
			}
		}else{
			for j:=0;j<=i;j++{
				if(j==0 || j==i){
					fmt.Print("* ")
				}else{
					fmt.Print("  ")
				}
			}
		}
		fmt.Println()
	}
}