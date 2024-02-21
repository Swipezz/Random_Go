package gomini

import (
	"fmt"
	"math/rand"
	"strconv"
)

func Roulette() {
	var chance int
	slot := []string{"O", "O", "O", "O", "O", "O"}
	var temp string

	for i := 1; i < 7; i++ {
		chance += 16
		if i == 5 {
			chance += 4
		}
		fmt.Printf(`

 1   2   3   4   5   6
(%s) (%s) (%s) (%s) (%s) (%s)
Masukan peluru! `, slot[0], slot[1], slot[2], slot[3], slot[4], slot[5])
		fmt.Scanln(&temp)

		nomer, _ := strconv.Atoi(temp)

		if nomer > 6 || slot[(nomer-1)] == "X" {
			fmt.Println("Salah Goblok!")
			i--
			chance -= 16
		} else {
			slot[(nomer - 1)] = "X"

			fmt.Printf(`


(?) (?) (?) (?) (?) (?)
%d%% Mati! | Siap? `, chance)
			fmt.Scanln()

			if slot[rand.Intn(len(slot))] == "X" {
				break
			}
		}
	}
	fmt.Println("\nDUAR!")
}
