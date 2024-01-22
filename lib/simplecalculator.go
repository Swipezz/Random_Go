package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convert(n []string) (float64, float64) {
	sum1, _ := strconv.ParseFloat(n[0], 64)
	sum2, _ := strconv.ParseFloat(n[2], 64)

	return sum1, sum2
}

func Calculator() {
	fmt.Println("Contoh 1 + 1\nMasukan nomer: ")

	// Untuk mengambil user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	converter := strings.Split(scanner.Text(), " ")
	num1, num2 := convert(converter)
	switch converter[1] {
	case "-":
		fmt.Println(num1 - num2)
	case "+":
		fmt.Println(num1 + num2)
	case "/":
		fmt.Println(num1 / num2)
	case "*":
		fmt.Println(num1 * num2)
	}
}
