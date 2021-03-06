package main 
import (
"fmt"
"strings"
)
func main() {
	var greeting="Hello World"

	fmt.Printf("normal string:")
	fmt.Printf("%s",greeting)
	fmt.Printf("\n")
   fmt.Printf("hex bytes: ")
   for i := 0; i < len(greeting); i++ {
       fmt.Printf("%x ", greeting[i])
   }
   fmt.Printf("\n")
   
   const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" 
   /*q flag escapes unprintable characters, with + flag it escapses non-ascii characters as well 
     to make output unambigous  */
   fmt.Printf("quoted string: ")
   fmt.Printf("%+q", sampleText)
   fmt.Printf("\n")  

//concatenation
   greeting1:=[]string{"Hello","world!"}
   fmt.Println(strings.Join(greeting1," *** "))
}