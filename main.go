package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/LarsFox/dice/art"
	"github.com/LarsFox/dice/iter"
)

const loop = 3

func main() {
	fmt.Println("Yo! Type in a number of dices to roll, but no more than 6.")

	var number int
	for number == 0 {
		fmt.Scan(&number)
		if number > 6 {
			number = 0
		}
	}

	pairs := make([][]int, 0, number*loop)
	for i := 0; i < loop; i++ {
		pairs = append(pairs, iter.Product(6, number)...)
	}

	fmt.Println("\nOk, let’s roll some dices!\n\n(Type in anything to exit.)")

	var enough string
	var i int
	for enough == "" {
		rand.Seed(time.Now().UnixNano())
		iter.ShuffleIntInt(pairs)

		for _, pair := range pairs {
			dices, _ := art.Dices(pair...)
			fmt.Println(dices)

			fmt.Scanln(&enough)
			if enough != "" {
				break
			}
			i++
		}
	}

	fmt.Printf("\nTotal rolls: %d!\n\nNice rolling!\n", i+1)
	fmt.Scanln(&enough)
}
