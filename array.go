package main

import "fmt"

func main() {
	//var array [10]float32
	//var array1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//var array2 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}   //slice of unspecifide size
	/* an array with 5 rows and 2 columns*/

	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	//output each arry element's value
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d]=%d\n", i, j, a[i][j])
		}
	}
}
