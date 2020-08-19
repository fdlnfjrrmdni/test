package main

import (
    "fmt"
    "strings"
    "strconv"
    "sort"
)

func OrderThis(str string) string {
	//	Ordering by category
	category := []int{6,7,0,9,4,8,1,2,5,3}

	//	Processing input string data into array
	arr := strings.Split(str, " ")

	//	Declares an empty multidimensional array 
	//	to store the results
	result := make([][]string, len(category))

	//	Processing input string data 
	//	into multidimensional array

	//	To match the notation one by one
	//	With category to be sorted by category
	for i:=0; i<len(arr); i++ {
		//	Takes the first notation
		arr2 := arr[i]

		//	Takes the first character 
		//	of the notation
		stStr := arr[i][:1]

		//	Convert data string to int
		//	In order for conditioning (below)
		stInt, err := strconv.Atoi(stStr)
		if err != nil {
	        fmt.Println(err)
	    }

		for j:=0; j<len(category); j++ {
			//	Initialize the multidimensional array with 
			//	notation based on its category/index of category
			temp := result[j]

			// 	Matching first char of 
			//	Notation with category for 
			//	Get the index for fill multidimensional array 
			//	According to the index obtained
			if (stInt == category[j]) {
				//	Manipulate variable with notation
				temp = append(temp, arr2)
			}
			//	Fill multidimensional array
			//	According to the index obtained
			//	Then the results will be sorted 
			//	According to category
			result[j] = temp
		}
	}

	//	The result will be an multidemensional array
	//	with a length of 10
	//	The contents are the input value
	//	And the contents have been sorted according to category

	// 	In DeleteSameId() 
	//	The data is filtered again 
	//	To delete the book notation,
	//	If there is the same identity and more than 2
	//	And converting multidimensional array data to strings
	finalResult := DeleteSameId(result)

	return finalResult
}

func DeleteSameId(str [][]string) string {
	category := []int{6,7,0,9,4,8,1,2,5,3}
	result := []string{}
	var resultStr string
	arrRemove := make([][]string, len(category))
	
	for k:=0; k<len(category); k++ {
		temp := str[k]

		//	DESCENDING ORDER BY THE HEIGHT
		if (len(temp) > 1) {
			sort.Slice(temp, func(i, j int) bool {	//	descending [3N21 3N20 3N14 3A13]
			    return temp[i][2:] > temp[j][2:]
			})
		}

		//	GET CATEGORY IF HAVE THE SAME IDENTITY MORE THAN 2
		if (len(temp) > 2) {
			for i:=0; i<len(temp); i++ {
				stChar := temp[i][:2]
				// fmt.Println("FOR 1 :",temp[i])
				sumTemp := 0
				for j:=0; j<len(temp); j++ {
					stChar2 := temp[j][:2]
					if (stChar == stChar2) {
						sumTemp = sumTemp+1
					}
				}

				//	IF IN ARRAY MULTIDIMENTION THERE IS MORE THAN 2 ELEMENT
				if (sumTemp > 2) {	//sum > 2
					stStr := temp[i][:1]
					stInt, err := strconv.Atoi(stStr)		//  3 from [3]A13 (int)
					
					if err != nil {
				        fmt.Println(err)
				    }

					for k:=0; k<len(category); k++ {
						tempRemove := arrRemove[k]
						if (stInt == category[k]) {
							tempRemove = append(tempRemove, temp[i])
						}
						arrRemove[k] = tempRemove
					}
				}
			}
		}
		
		//	CONVERT [][]string TO []string
		if (temp != nil) {
			for i:=0; i<len(temp); i++ {			//	len(temp) is [[2][4]] from [[5X19 5G14][3A13 3N20 3N21 3N14]]	
				result = append(result, temp[i])	//	i is index of  ^  ^			 (1)  (2)   (1)  (2)  (3)  (4)
			}
		}
	}

	//	DELETE NOTATION(LOWEST HEIGHT) HAVE THE SAME IDENTITY MORE THAN 2
	remove := make([][]string, 0)
	for i:=0; i<len(arrRemove); i++ {
		if (arrRemove[i] != nil) {
			leng := len(arrRemove[i])
			tempRemove := arrRemove[i][2:leng]
			// notRemove := arrRemove[i][:2]
			remove = append(remove, tempRemove)
		}
	}
	// fmt.Println(remove)

	// CONVERT remove [][]string TO []string
	var resultRemove []string
	for i:=0; i<len(remove); i++ {
		// fmt.Println(remove[i])
		for j:=0; j<len(remove[i]); j++ {
			// fmt.Println(remove[i][j])
			resultRemove = append(resultRemove, remove[i][j])
		}
	}
	// fmt.Println(resultRemove)

	//	CONVERT []string TO string
	for j:=0; j<len(result); j++ {
		resultStr = resultStr + result[j] + " "
	}
	
	// fmt.Println(resultStr)
	for l:=0; l<len(resultRemove); l++ {
		resultStr = strings.ReplaceAll(resultStr, resultRemove[l]+" ", "")
	}


	//	DELETE LAST SPACE
	return resultStr[:len(resultStr)-1]
}


func main() {
	input := "3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14"
	fmt.Println(OrderThis(input))
}