package art

import (
	"errors"
	"strings"
)

const (
	one = `
     _________    
    /        /\   
   /        /  \  
  /        /    \ 
 /________/      \
 \        \      /
  \   @    \    / 
   \        \  /  
    \________\/   `

	two = `
     _________    
    /        /\   
   /        /  \  
  /        /    \ 
 /________/      \
 \ @      \      /
  \        \    / 
   \     @  \  /  
    \________\/   `

	three = `
     _________    
    /        /\   
   /        /  \  
  /        /    \ 
 /________/      \
 \ @      \      /
  \   @    \    / 
   \     @  \  /  
    \________\/   `

	four = `
     _________    
    /        /\   
   /        /  \  
  /        /    \ 
 /________/      \
 \ @    @ \      /
  \        \    / 
   \ @    @ \  /  
    \________\/   `

	five = `
     _________    
    /        /\   
   /        /  \  
  /        /    \ 
 /________/      \
 \ @    @ \      /
  \   @    \    / 
   \ @    @ \  /  
    \________\/   `

	six = `
     _________    
    /        /\   
   /        /  \  
  /        /    \ 
 /________/      \
 \ @    @ \      /
  \ @    @ \    / 
   \ @    @ \  /  
    \________\/   `
)

const nl = "\n"

var (
	// dice is a map of dice ASCII images.
	dice = map[int]string{
		0: one,
		1: two,
		2: three,
		3: four,
		4: five,
		5: six,
	}

	// number of lines in one dice image.
	asciiLength = len(strings.Split(one, nl))

	// ErrInvalidNum is returned when there is no art for the die.
	ErrInvalidNum = errors.New("invalid number")
)

// Dice returns a wide ASCII image of all the dice.
func Dice(nums ...int) (string, error) {
	total := make([]string, asciiLength)
	for _, num := range nums {
		dice, ok := dice[num]
		if !ok {
			return "", ErrInvalidNum
		}

		for i, line := range strings.Split(dice, nl) {
			total[i] += " " + line
		}
	}
	return strings.Join(total, nl), nil
}
