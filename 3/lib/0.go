package lib

import (
    "fmt"
)

//	Iterative function to  
//	reverse digits of number
func RevNum(n int) int {
    reverseNum := 0
    for n > 0 {
        remainder := n % 10
        reverseNum *= 10
        reverseNum += remainder
        n /= 10
    }
    return reverseNum 
}

//	Condition function for to know 
//	This number palindrome or not
func IsPal(n int) bool {
	reverseNum := RevNum(n)
	if ((n >= 0 && n < 10) || n == reverseNum){
		return true
	}else{
		return false
	}
}

//	Function for collect palindrome number 
//	into an array
func SearchPal(n1, n2 int) []int {
	pal := []int{}
	if (n1 < n2 && n1 > 0 && n2 < 1000000001){		//	1000000001 is 10^9 - 1
		//	Looping for decomposes 
		//	numbers from n1 to n2
		for n1<=n2 {
			//	If a palindrome number is detected
			//	the number will be store to array int
			if (IsPal(n1)){
				pal = append(pal, n1)
			}
			n1++
		}
	}else{
		fmt.Println("ANGKA YANG DIMASUKKAN TIDAK MEMENUHI KUALIFIKASI")
	}
	return pal
}

//	Function for counting 
//	The palindrome numbers
func SumPal(n1, n2 int) int {
	return len(SearchPal(n1, n2))
}