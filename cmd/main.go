package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/LarsFox/dice/art"
	"github.com/LarsFox/dice/iter"
)

const (
	die   = 6
	loop  = 3
	limit = 6
)

func main() {
	fmt.Println("Yo! Type in a number of dice to roll, but no more than 6.")

	var number int
	for number == 0 {
		fmt.Scan(&number)
		if number > limit {
			number = 0
		}
	}

	fmt.Println("\nOk, let’s roll some dice!\n\n(Type in anything to exit.)")

	var i int
	var enough bool
	scanner := bufio.NewScanner(os.Stdin)
	for !enough {
		seq := iter.Sequence(die, number, loop)
		for _, perm := range seq {
			dice, _ := art.Dice(perm...)
			fmt.Println(dice)
			i++

			if scanner.Scan() {
				if enough = scanner.Text() != ""; enough {
					break
				}
			}
		}
	}

	fmt.Printf("\nTotal rolls: %d!\n\nHave a nice day!\n", i+1)
}
