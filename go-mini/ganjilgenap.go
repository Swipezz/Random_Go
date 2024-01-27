package gomini

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GanjilGenap() {
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

	for i, s := range converter {
		ints[i], _ = strconv.Atoi(s)
	}

	for _, intss := range ints {
		if intss%2 != 0 {
			fmt.Printf("%d = Ganjil \n", intss)
		} else {
			fmt.Printf("%d = Genap \n", intss)
		}
	}
}
