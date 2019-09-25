package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Dealer Struct
type Dealer struct {
	num  int // 카드수
	card []int
}

func (dealer *Dealer) makeCard() {
	// fmt.Println(dealer.num)
	for i := 0; i < dealer.num; i++ {
		dealer.card[i] = (i + 1) % 10
		if dealer.card[i] == 0 {
			dealer.card[i] = 10
		}
	}
}

func (dealer *Dealer) doShuffle() {

	var myCard = make([]int, dealer.num)
	var index = 0
	for {
		if index == dealer.num {
			break
		}
		s1 := rand.NewSource(time.Now().UnixNano())
		rand := rand.New(s1)
		randomNumber := rand.Intn(len(dealer.card))
		if dealer.card[randomNumber] != 0 {
			myCard[index] = dealer.card[randomNumber]
			dealer.card = append(dealer.card[:randomNumber], dealer.card[randomNumber+1:]...)
			// fmt.Println(card, myCard)
			index++
		} else {
			continue
		}
	}
	dealer.card = append(dealer.card, myCard...)
	fmt.Println(dealer.card)
}

func (dealer *Dealer) doDistribute() {

}

type Player struct {
	card    [2]int // 받은 카드
	nRun    int    // 게임 횟수
	winRate []int  // 승률
}

func main() {
	// 딜러 만들고
	dealer := Dealer{}
	dealer.num = 20
	dealer.card = make([]int, dealer.num)

	// 플레이어 만들고
	player_1 := Player{}
	player_1.card = [2]int{0, 0}
	player_1.nRun = 0

	player_2 := Player{}
	player_2.card = [2]int{0, 0}
	player_2.nRun = 0

	// 카드 만들고
	dealer.makeCard()
	// 카드 섞고
	dealer.doShuffle()

	// 게임을 진행
	player_1.card = dealer.card[0:2]
	player_2.card = dealer.card[2:4]

}

func competition(p1 *Player, p2 *Player, d *Dealer) (int, int, int) {

	player_1_result := (player_1[0] + player_1[1]) % 10
	player_2_result := (player_2[0] + player_2[1]) % 10
	winner := 0
	if player_1_result > player_2_result {
		winner = 1
	} else if player_1_result == player_2_result {
		winner = 0
	} else {
		winner = 2
	}
	return player_1_result, player_2_result, winner
}
