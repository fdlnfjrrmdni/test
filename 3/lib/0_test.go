package lib

import (
	"testing"
)

func TestSumPalindrome(t *testing.T) {
	input := [][]int{{1, 10}, {99, 100}, {21, 31}, {1,100}}
	output := []int{9,1,1,18}
	for i:=0; i<len(output); i++ {
		input2 := input[i]
		n1 := input2[0]
		n2 := input2[1]		

		t.Logf("INPUT : %d %d",n1,n2)

		expected := output[i]
		actual := SumPal(n1, n2)
		
		if actual != expected {
			t.Errorf("OUTPUT: %d ", actual)
			t.Errorf("Test failed, expected: '%d', got: '%d' \n\n", expected, actual)
		}else {
			t.Logf("OUTPUT: %d \n\n", actual)
		}
	}
}