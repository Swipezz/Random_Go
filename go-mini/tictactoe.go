package gomini

import (
	"fmt"
	"math/rand"
	"strings"
)

func playerMove(player string, move string, save []string) int {
	index := getKotak(move)
	if index < 0 || save[index] != " " {
		fmt.Println("Move tidak valid!")
		index = -1
		return index
	}

	save[index] = player
	return index
}

func getKotak(move string) int {
	move = strings.ToLower(move)

	switch move {
	case "a1":
		return 0
	case "b1":
		return 1
	case "c1":
		return 2
	case "a2":
		return 3
	case "b2":
		return 4
	case "c2":
		return 5
	case "a3":
		return 6
	case "b3":
		return 7
	case "c3":
		return 8
	default:
		return -1
	}
}

func enemyMove(enemy string, save []string) {
	pick := rand.Intn(len(save))
	var temp string
	for _, a := range save {
		temp += a
	}

	tempF := strings.Contains(temp, " ")

	if !tempF {
		return
	}

	if save[pick] != " " {
		enemyMove(enemy, save)
		return
	}
	save[pick] = enemy
}

func validasi(player string, enemy string, save []string) bool {
	if save[0] == player && save[1] == player && save[2] == player {
		fmt.Println("Kamu menang!")
		return true
	}
	if save[3] == player && save[4] == player && save[5] == player {
		fmt.Println("Kamu menang!")
		return true
	}
	if save[6] == player && save[7] == player && save[8] == player {
		fmt.Println("Kamu menang!")
		return true
	}
	if save[0] == player && save[3] == player && save[6] == player {
		fmt.Println("Kamu menang!")
		return true
	}
	if save[1] == player && save[4] == player && save[7] == player {
		fmt.Println("Kamu menang!")
		return true
	}
	if save[2] == player && save[5] == player && save[8] == player {
		fmt.Println("Kamu menang!")
		return true
	}
	if save[0] == player && save[4] == player && save[8] == player {
		fmt.Println("Kamu menang!")
		return true
	}
	if save[2] == player && save[4] == player && save[6] == player {
		fmt.Println("Kamu menang!")
		return true
	}

	if save[0] == enemy && save[1] == enemy && save[2] == enemy {
		fmt.Println("Kamu kalah!")
		return true
	}
	if save[3] == enemy && save[4] == enemy && save[5] == enemy {
		fmt.Println("Kamu kalah!")
		return true
	}
	if save[6] == enemy && save[7] == enemy && save[8] == enemy {
		fmt.Println("Kamu kalah!")
		return true
	}
	if save[0] == enemy && save[3] == enemy && save[6] == enemy {
		fmt.Println("Kamu kalah!")
		return true
	}
	if save[1] == enemy && save[4] == enemy && save[7] == enemy {
		fmt.Println("Kamu kalah!")
		return true
	}
	if save[2] == enemy && save[5] == enemy && save[8] == enemy {
		fmt.Println("Kamu kalah!")
		return true
	}
	if save[0] == enemy && save[4] == enemy && save[8] == enemy {
		fmt.Println("Kamu kalah!")
		return true
	}
	if save[2] == enemy && save[4] == enemy && save[6] == enemy {
		fmt.Println("Kamu kalah!")
		return true
	}
	return false
}

func TicTacToe() {
	save := []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	total := 5
	var player string
	var enemy string

	fmt.Print("Pertama / Kedua : ")
	fmt.Scanln(&player)
	player = strings.ToLower(player)

	if player == "pertama" {
		player = "X"
		enemy = "O"
		fmt.Printf(`
|%s|%s|%s|1
|%s|%s|%s|2
|%s|%s|%s|3
 A B C

`, save[0], save[1], save[2], save[3], save[4], save[5], save[6], save[7], save[8])
	} else {
		player = "O"
		enemy = "X"
		enemyMove(enemy, save)
		total -= 1
		fmt.Printf(`
|%s|%s|%s|1
|%s|%s|%s|2
|%s|%s|%s|3
 A B C

`, save[0], save[1], save[2], save[3], save[4], save[5], save[6], save[7], save[8])
	}

	for i := 0; i < total; i++ {
		var move string
		fmt.Print("Giliranmu : ")
		fmt.Scanln(&move)

		for playerMove(player, move, save) == -1 {
			fmt.Print("Giliranmu : ")
			fmt.Scanln(&move)
		}
		enemyMove(enemy, save)
		fmt.Printf(`
|%s|%s|%s|1
|%s|%s|%s|2
|%s|%s|%s|3
 A B C

`, save[0], save[1], save[2], save[3], save[4], save[5], save[6], save[7], save[8])

		if validasi(player, enemy, save) {
			break
		}
	}
}
