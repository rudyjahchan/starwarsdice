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
	var dicePoolGroup sync.WaitGroup

	dicePoolGroup.Add(6)

	go buildDiePool(&dicePoolGroup, Ability, c.Ability, dice)
	go buildDiePool(&dicePoolGroup, Proficiency, c.Proficiency, dice)
	go buildDiePool(&dicePoolGroup, Boost, c.Boost, dice)
	go buildDiePool(&dicePoolGroup, Difficulty, c.Difficulty, dice)
	go buildDiePool(&dicePoolGroup, Challenge, c.Challenge, dice)
	go buildDiePool(&dicePoolGroup, Setback, c.Setback, dice)

	dicePoolGroup.Wait()
	close(dice)
}

func buildDiePool(dicePoolWorkgroup *sync.WaitGroup, d Die, count int, dice chan Die) {
	defer dicePoolWorkgroup.Done()

	for index := 0; index < count; index++ {
		dice <- d
	}
}

func roller(rollersGroup *sync.WaitGroup, die Die, rolls chan Roll) {
	defer rollersGroup.Done()
	rolls <- die.Roll()
}

func roll(dice chan Die, rolls chan Roll) {
	var rollersGroup sync.WaitGroup
	for die := range dice {
		rollersGroup.Add(1)
		go roller(&rollersGroup, die, rolls)
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
