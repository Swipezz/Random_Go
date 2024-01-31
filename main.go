package main

import (
	"bufio"
	"log"
	"os"
)

func user() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	text := scanner.Text()

	return text
}

func main() {
	// id := user()
	// name := user()
	// age, _ := strconv.Atoi(user())
	// grade, _ := strconv.Atoi(user())

}
