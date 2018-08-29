package starwarsdice

import (
	"fmt"
	"sync"
)

// Check represents the dice we need to roll to see if we accomplished a task.
type Check struct {
	Ability     int
	Proficiency int
	Boost       int

	Difficulty int
	Challenge  int
	Setback    int
}

// Roll peforms the check and reports the result.
func (c Check) Roll() {
	done := make(chan Result)
	rolls := make(chan Roll, 100)
	dice := make(chan Die, 100)

	go buildDicePool(c, dice)
	go roll(dice, rolls)
	go printAndDisplay(rolls, done)

	result := <-done
	fmt.Println(RESET, "\n\n", result)
}

func buildDicePool(c Check, dice chan Die) {
	buildDiePool(Ability, c.Ability, dice)
	buildDiePool(Proficiency, c.Proficiency, dice)
	buildDiePool(Boost, c.Boost, dice)
	buildDiePool(Difficulty, c.Difficulty, dice)
	buildDiePool(Challenge, c.Challenge, dice)
	buildDiePool(Setback, c.Setback, dice)
	close(dice)
}

func buildDiePool(d Die, count int, dice chan Die) {
	for index := 0; index < count; index++ {
		dice <- d
	}
}

func roller(rollersGroup *sync.WaitGroup, dice chan Die, rolls chan Roll) {
	defer rollersGroup.Done()
	for die := range dice {
		rolls <- die.Roll()
	}
}

func roll(dice chan Die, rolls chan Roll) {
	var rollersGroup sync.WaitGroup
	for index := 0; index < 100; index++ {
		rollersGroup.Add(1)
		go roller(&rollersGroup, dice, rolls)
	}
	rollersGroup.Wait()
	close(rolls)
}

func printAndDisplay(rolls chan Roll, finalResult chan Result) {
	var result Result
	var count uint
	for roll := range rolls {
		if count%40 == 0 {
			fmt.Print("\n")
		}
		fmt.Print(roll.Die, " ")
		count++
		result = result.Add(roll.Result)
	}
	finalResult <- result
}
