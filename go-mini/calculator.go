package gomini

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convert(n []string) ([]string, []float64) {
	var sum []string
	var num []float64

	for i, s := range n {
		if i%2 != 0 {
			sum = append(sum, s)
		} else {
			temp, _ := strconv.ParseFloat(n[i], 64)
			num = append(num, temp)
		}
	}
	return sum, num
}

func Calculator() {
	fmt.Println("Contoh 1 + 2 - 3\nMasukan nomer: ")

	// Untuk mengambil user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	converter := strings.Split(scanner.Text(), " ")
	sum, num := convert(converter)
	total := num[0]

	for i := 1; i < len(num); i++ {
		switch sum[i-1] {
		case "+":
			fmt.Printf("%g %s %g = ", total, sum[i-1], num[i])
			total += num[i]
			fmt.Printf("%g \n", total)

		case "-":
			fmt.Printf("%g %s %g = ", total, sum[i-1], num[i])
			total -= num[i]
			fmt.Printf("%g \n", total)

		case "/":
			fmt.Printf("%g %s %g = ", total, sum[i-1], num[i])
			total /= num[i]
			fmt.Printf("%g \n", total)

		case "*":
			fmt.Printf("%g %s %g = ", total, sum[i-1], num[i])
			total *= num[i]
			fmt.Printf("%g \n", total)
		}
	}
	fmt.Printf("Hasil akhir : %g", total)
}
