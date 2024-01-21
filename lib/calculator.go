package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Calculator() {
	fmt.Println("Contoh 1 atau 1 2 3\nMasukan nomer: ")

	// Untuk mengambil user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Mengubah string slice menjadi int slice
	converter := strings.Split(scanner.Text(), " ") // Misah nomer yang ada setelah space
	ints := make([]int, len(converter))

	for i := 1; i < len(converter); i += 2 {

		fmt.Println(converter[i])
	}

	var total int
	for i, s := range converter {
		ints[i], _ = strconv.Atoi(s)
		total += ints[i]
	}

	fmt.Printf("%d", total)
}
