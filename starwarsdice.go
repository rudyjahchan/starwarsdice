package main

import (
	"fmt"
	starwarsdice "starwarsdice/dicepool"
	"time"
)

func scanDice(d starwarsdice.Die, c *int) {
	fmt.Printf("Number of %s dice? ", d)
	fmt.Scanf("%d", c)
}

func main() {
	var check starwarsdice.Check
	scanDice(starwarsdice.Ability, &(check.Ability))
	scanDice(starwarsdice.Proficiency, &(check.Proficiency))
	scanDice(starwarsdice.Boost, &(check.Boost))
	scanDice(starwarsdice.Difficulty, &(check.Difficulty))
	scanDice(starwarsdice.Challenge, &(check.Challenge))
	scanDice(starwarsdice.Setback, &(check.Setback))

	startTime := time.Now()
	check.Roll()
	endTime := time.Now()

	diff := endTime.Sub(startTime)
	fmt.Println("Time: ", diff.Seconds(), "s")
}
