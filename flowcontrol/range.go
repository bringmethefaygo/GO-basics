package main
import "fmt"

func main(){
	fmt.Println("USING ARRAY")
	number := [5]int{21,24,27,30,33}
	for index,item:=range number{
		fmt.Printf("number[%d]=%d\n",index,item)
	}
	fmt.Println("USING MAP")
	subjectmarks:=map[string]float32{"JAVA":80,"PYTHON":81,"GOLANG":85}
	for subject,marks:=range subjectmarks{
		fmt.Println(subject,":",marks)
	}
}