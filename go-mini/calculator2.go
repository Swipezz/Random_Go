package gomini

import (
	"fmt"
	"unicode"
)

func Calculator2() {
	var stringNumber string
	var stringFunc string
	var intNumber []int

	var kunci int

	fmt.Println("Format calculator | 1+2 |")
	fmt.Scanln(&stringNumber)
	for i, char := range stringNumber {
		if !unicode.IsDigit(char) {
			if string(char) != "+" || string(char) != "-"  {
				
			}
			stringFunc = string(char)
		} else {
			if kunci == 0 {
				intNumber[0] += int(char) 
			}
			if kunci == 1 {
				intNumber[1] += int(char) 
			}
		} 		
	}
}
