package starwarsdice

import (
	"fmt"
)

type Check struct {
	Ability     int
	Proficiency int
	Boost       int

	Difficulty int
	Challenge  int
	Setback    int
}

func (c Check) Roll() Result {
	var result Result
	result = result.Add(roll(Ability, c.Ability))
	result = result.Add(roll(Proficiency, c.Proficiency))
	result = result.Add(roll(Boost, c.Boost))
	result = result.Add(roll(Difficulty, c.Difficulty))
	result = result.Add(roll(Challenge, c.Challenge))
	result = result.Add(roll(Setback, c.Setback))
	fmt.Println(RESET, "\n\n", result)
	return result
}

func roll(d Die, c int) Result {
	var result Result
	for index := 0; index < c; index++ {
		roll := d.Roll()
		if index%40 == 0 {
			fmt.Print("\n")
		}
		fmt.Print(d, " ")
		result = result.Add(roll.Result)
	}

	return result
}
