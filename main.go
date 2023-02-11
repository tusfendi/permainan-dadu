package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"strings"
)

type Player struct {
	PlayerOrder int
	Dice        int
	Point       int
	DiceResult  []string
}

func main() {
	var player int
	var dice int

	fmt.Println("===== PERMAINAN DADU ====== ")
	fmt.Print("Input jumlah pemain: ")
	fmt.Scanln(&player)
	if player < 2 {
		println("Pemain harus minimal 2 orang yaaa")
		return
	}
	fmt.Print("Input jumlah dadu: ")
	fmt.Scanln(&dice)
	fmt.Println("=========================================")
	fmt.Println("=========================================")
	fmt.Println("Pemain =", player, ", Dadu =", dice)

	// Push Mapping
	play := list.New()
	for i := 1; i <= player; i++ {
		play.PushBack(&Player{
			PlayerOrder: i,
			Dice:        dice,
		})
	}

	PlayGame(1, play, play.Len())
}

func PlayGame(i int, play *list.List, len int) {
	playerNow := play.Front()
	fmt.Println("=========================================")
	fmt.Println("Giliran ", i, " lempar dadu:")
	for i := 0; i < len; i++ {
		player := playerNow.Value.(*Player)
		dice := player.Dice
		player.Dice = 0
		oldPlayerPoint := player.Point
		var diceRandom []string
		var diceResult []string
		for i := 0; i < dice; i++ {
			random := rand.Intn(7-1) + 1
			if random == 1 {
				checkNextPlayer := playerNow.Next()
				if checkNextPlayer == nil {
					checkNextPlayer = play.Front()
				}
				// 1 masih belum masuk di nextnya yaa
				nextPlayer := checkNextPlayer.Value.(*Player)
				nextPlayer.DiceResult = append(nextPlayer.DiceResult, "1")
				nextPlayer.Dice++
			} else if random == 6 {
				player.Point++
			} else {
				player.Dice++
				diceResult = append(diceResult, fmt.Sprint(random))
			}
			diceRandom = append(diceRandom, fmt.Sprint(random))
		}
		player.DiceResult = diceResult
		print("    Pemain # ", player.PlayerOrder, "(", oldPlayerPoint, "): ")
		if strings.Join(diceRandom, ", ") != "" {
			println(strings.Join(diceRandom, ", "))
		} else {
			println("_ (Berhenti bermain karena tidak memiliki dadu)")
		}
		playerNow = playerNow.Next()
	}
	// EVALUASI
	println("Setelah evaluasi:")
	playerNow = play.Front()
	diceChecker := 0
	var winners []string
	var lost int
	lastPoint := 0
	for i := 0; i < len; i++ {
		player := playerNow.Value.(*Player)
		if player.Dice > 0 {
			diceChecker++
			lost = player.PlayerOrder
		}

		if player.Point == lastPoint {
			winners = append(winners, fmt.Sprint("#", player.PlayerOrder))
		} else if player.Point > lastPoint {
			lastPoint = player.Point
			winners = []string{
				fmt.Sprint("#", player.PlayerOrder),
			}
		}
		print("    Pemain # ", player.PlayerOrder, "(", player.Point, "): ", strings.Join(player.DiceResult, ", "))
		println()
		playerNow = playerNow.Next()
	}
	i++
	if diceChecker > 1 {
		PlayGame(i, play, len)
	} else {
		fmt.Println("=========================================")
		fmt.Print(" Game berakhir karena hanya pemain #", lost, " yang memiliki dadu.")
		fmt.Println()
		fmt.Println(" Game dimenangkan oleh pemain ", strings.Join(winners, ", "), " karena memiliki poin lebih banyak dari pemain lainnya.")

	}
}

// println(play.Len())

// println("player : ", play.Front().Value.(*Player).PlayerOrder)
// println("player : ", play.Front().Next().Value.(*Player).PlayerOrder)
// // play.Front().Next().Value.(*Player).PlayerOrder = 200
// println("player : ", play.Front().Next().Next().Value.(*Player).PlayerOrder)
// if play.Front().Next().Next().Next() == nil {
// 	println("koosoonggg balik ke 1 ya nak")
// 	println("player : ", play.Front().Value.(*Player).PlayerOrder)
// }
// println(play.Front().Next().Value.(*Player).PlayerOrder)
// println(play.Front().Next().Next().Value.(*Player).PlayerOrder)
// println(play.Back().Value.(*Player).PlayerOrder)

// playerNow := play.Front()
// for i := 0; i < play.Len(); i++ {
// 	// if i == 1 {
// 	println("player : ", playerNow.Value.(*Player).PlayerOrder)
// 	// } else {
// 	// 	println("player : ", playerNow.Next().Value.(*Player).PlayerOrder)
// 	// }
// 	playerNow = playerNow.Next()
// }
