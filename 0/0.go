package main

import (
    "fmt"
    lib "../3/lib"	 
)

//	The unit test file of this code 
//	are located in ../3/lib
//	I save it here because it will be use
//	in question number 3 which will later be dockerized
func main() {
	var n1,n2 int
	fmt.Println("EXAMPLE")
	fmt.Println("1 100		[input]")
	fmt.Println("18		[output]\n")
	
	//	User can input the desired number
	fmt.Scanf("%d %d",&n1,&n2)
    fmt.Println(lib.SumPal(n1, n2))
}