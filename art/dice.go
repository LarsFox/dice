package art

import (
	"fmt"
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

const (
	// NL is a new line
	NL = "\n"
)

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
	asciiLength = len(strings.Split(one, NL))
)

// Dice returns a wide ASCII image of all the dice.
func Dice(nums ...int) (string, error) {
	total := make([]string, asciiLength)
	for _, num := range nums {
		dice, ok := dice[num]
		if !ok {
			return "", fmt.Errorf("wrong dice number: %d", num)
		}

		for i, line := range strings.Split(dice, NL) {
			total[i] += " " + line
		}
	}
	return strings.Join(total, NL), nil
}
