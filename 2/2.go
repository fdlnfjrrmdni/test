package main

import (
    "fmt"
    "math"
)

func getValue(str string, i int, m int) int {
	if (i + m > len(str)) {
		return -1
	}

	// Find value at index i and length m. 
	value := 0
	for j:=0; j<m; j++ { 
		c := str[i + j] - '0'
		if (c < 0 || c > 9) {
			return -1
		}
		value = value * 10 + c 
	} 
	return value
}

func findMissingNumber(str string) int { 
	// Try all lengths for first number 
	for m:=1; m<=7; m=1+m { 
		// Get value of first number with current 
		// length/ 
		n := getValue(str, 0, m) 
		if (n == -1) {
			break
		}

		// To store missing number of current length 
		missingNo := -1

		// To indicate whether the sequence failed 
		// anywhere for current length. 
		fail := false

		// Find subsequent numbers with previous number as n 
		for i:=m; i!=len(str); i += 1 + math.Log10(n) { 
			// If we haven't yet found the missing number 
			// for current length. Next number is n+2. Note 
			// that we use Log10 as (n+2) may have more 
			// length than n. 
			if ((missingNo == -1) && (getValue(str, i, 1+math.Log10(n+2)) == n+2)) { 
				missingNo = n + 1
				n += 2
			}else if (getValue(str, i, 1+math.Log10(n+1)) == n+1) {	// If next value is (n+1) 
				n++
			}else { 
				fail = true
				break
			} 
		} 

		if (!fail) {
			return missingNo
		}
	} 
	return -1 // not found or no missing number 
} 

func main() {
    fmt.Println(findMissingNumber("99101102"))
}