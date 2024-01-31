package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(scanner.Text())
}
