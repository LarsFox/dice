package art

import (
	"fmt"
	"strings"
)

const (
	// NL is a new line
	NL = "\n"
)

var (
	// dices is a map of dice ASCII images
	dices = map[int]string{
		0: one,
		1: two,
		2: three,
		3: four,
		4: five,
		5: six,
	}

	// number of lines in one dice image
	asciiLength = len(strings.Split(one, NL))
)

// Dices returns a wide ASCII image of all the dices
func Dices(nums ...int) (string, error) {
	total := make([]string, asciiLength)
	for _, num := range nums {
		dice, ok := dices[num]
		if !ok {
			return "", fmt.Errorf("wrong dice number: %d", num)
		}

		for i, line := range strings.Split(dice, NL) {
			total[i] += " " + line
		}
	}
	return strings.Join(total, NL), nil
}
