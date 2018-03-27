# Dice

Lots of the board games use dice as a calculating mechanism. It is fine if you play for an eternity, but usually people win after a certain amount of rolls. In my case, some of the games end up with a quarrel over the random behaviour of dice.

This program is an attempt to eliminate the ruthless randomness of dice.

## How it works

All possible permutations of dice are multiplied by 3 and then are given in a random order. For example, with one die, each number will appear every 18 turns:
`6 5 2 1 1 4 2 3 3 2 5 1 5 3 4 6 6 4`

Then a new random sequence is generated and used.

## The reason for multiplication

Without the multiplication the sequence exhausts too soon. With one die it is too easy to predict the last, 6th number. Consider this situation after 5 rolls:
`6 5 2 1 3 ?`
**?** is definitely going to be 4, and the turn can be planned in advance. With the multiplication every other number is possible, though 4 has the highest chances.

With 18 rolls for one die and 108 rolls for two dice, it becomes too hard to remember all the rolls and to calculate the remaining. To make it even harder, in case with more than one die the output is also shuffled.

Of course, it is possible to have something like this:
`1 1 1 ?`
In this case all players understand that number 1 will not roll in at least 15 turns. But since this situation _is not likely_ to happen, it is fair that the number will not appear again.

Also, multiplication grants a possibility of up to six repetitions of a single number:

```markdown
First sequence
4 2 1 2 3 6 3 4 3 1 1 4 6 2 6 5 5 5

Second sequence:
5 5 5 ?
```

## Compilation

Windows
`env GOOS=windows GOARCH=amd64 go build`

Mac
`env GOOS=darwin GOARCH=amd64 go build`

Linux
`env GOOS=linux GOARCH=amd64 go build`

## Bugs

On Windows, for some reason, prints the first two rolls at once.
