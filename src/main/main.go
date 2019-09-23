package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, Card world!")

	const num int = 20
	var card = [num]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	const nPlayer int = 2
	var user [nPlayer][2]int
	var result [nPlayer]int

	// user := make([][]int, nPlayer, 2)

	// for i := 0; i < num; i++ {
	// 	card = append(card, rand.Intn(10)+1)
	// }

	for i := 0; i < nPlayer; i++ {
		for {
			a := rand.Intn(num)
			// fmt.Println(a)
			if card[a] != 0 {
				user[i][0] = card[a]
				card[a] = 0
				break
			}
		}
		for {
			b := rand.Intn(num)
			if card[b] != 0 {
				user[i][1] = card[b]
				card[b] = 0
				break
			}
		}
	}

	fmt.Println(card)
	fmt.Println(user)

	for i := 0; i < nPlayer; i++ {
		result[i] = (user[i][0] + user[i][1]) % 10
	}

	k := 0
	v := 0
	for i, val := range result {
		fmt.Println(i, val)
		if v < result[i] {
			k = i
			v = result[i]
		}
	}

	fmt.Printf("Winner %d th, Score : %d", k+1, v)
}
