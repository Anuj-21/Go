package main 
import "fmt"
func main(){
	var b int=3
	var a int

	numbers:=[6]int{1,2,3,5}

	for a:=0;a<3;a++{
		fmt.Printf("value of a: %d\n",a)
	}
 fmt.Printf("#####################################################################")
	for a<b{
		a++
		fmt.Printf("value of a:%d\n",a)
	}
 fmt.Printf("#####################################################################")
	for i,x:=range numbers {
		fmt.Printf("value of x= %d at %d\n",x,i)
	}
}