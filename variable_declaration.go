package main 
import "fmt"
func main() {
var a, b,c=1,7.5,"anuj"
d:="thakur"      //its type determined at run time
const LENGTH int =5
fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)
fmt.Println(LENGTH)

fmt.Printf("a is of the type %T\n",a)
fmt.Printf("b is of the type %T\n",b)
fmt.Printf("c is of the type %T\n",c)
fmt.Printf("d is of the type %T\n",d)	
fmt.Printf("LENGTH is of the type %T\n",LENGTH)
}