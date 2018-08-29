package starwarsdice

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// Die reperesents a single type of die used in a pool
type Die struct {
	Symbol Symbol
	Color  Color

	numberOfFaces *big.Int
	faces         []Result
}

// Roll is the outcome of rolling a specific Die
type Roll struct {
	Die    Die
	Result Result
}

// Roll rolls the die and returns the outcome.
func (d Die) Roll() Roll {
	faceIndex, _ := rand.Int(rand.Reader, d.numberOfFaces)
	result := d.faces[faceIndex.Int64()]
	time.Sleep(100 * time.Millisecond)
	return Roll{d, result}
}

func (d Die) String() string {
	return fmt.Sprint(d.Color, d.Symbol, RESET)
}

func newDie(symbol Symbol, color Color, faces []Result) Die {
	numberOfFaces := big.NewInt(int64(len(faces)))
	return Die{symbol, color, numberOfFaces, faces}
}

var nothing = Result{0, 0, 0, 0, 0, 0}

var aHit = Result{1, 0, 0, 0, 0, 0}
var twoHits = Result{2, 0, 0, 0, 0, 0}
var aHitWithAnAdvantage = Result{1, 1, 0, 0, 0, 0}
var withAnAdvantage = Result{0, 1, 0, 0, 0, 0}
var withTwoAdvantages = Result{0, 2, 0, 0, 0, 0}
var andATriumph = Result{0, 0, 1, 0, 0, 0}

var aFailure = Result{0, 0, 0, 1, 0, 0}
var twoFailures = Result{0, 0, 0, 2, 0, 0}
var aFailureWithAThreat = Result{0, 0, 0, 1, 1, 0}
var withAThreat = Result{0, 0, 0, 0, 1, 0}
var withTwoThreats = Result{0, 0, 0, 0, 2, 0}
var andADespair = Result{0, 0, 0, 0, 0, 1}

// Boost dice are used to represent possession of some benefit in the moment.
var Boost = newDie(Cube, Blue, []Result{
	nothing,
	nothing,
	aHit,
	aHitWithAnAdvantage,
	withTwoAdvantages,
	withAnAdvantage,
})

// Setback dice are used to represent additional situational difficulties
var Setback = newDie(Cube, Black, []Result{
	nothing,
	nothing,
	aFailure,
	aFailure,
	withAThreat,
	withAThreat,
})

// Ability dice represent a character's skill in the dice pool.
var Ability = newDie(Diamond, Green, []Result{
	nothing,
	aHit,
	aHit,
	twoHits,
	withAnAdvantage,
	withAnAdvantage,
	aHitWithAnAdvantage,
	withTwoAdvantages,
})

// Difficulty dice represent the difficulty of what the character is attempting.
var Difficulty = newDie(Diamond, Purple, []Result{
	nothing,
	aFailure,
	twoFailures,
	withAThreat,
	withAThreat,
	withAThreat,
	withTwoThreats,
	aFailureWithAThreat,
})

// Proficiency dice represent when a character's specialization in the skill used to resolve an action.
var Proficiency = newDie(Hexagon, Yellow, []Result{
	nothing,
	aHit,
	aHit,
	twoHits,
	twoHits,
	withAnAdvantage,
	aHitWithAnAdvantage,
	aHitWithAnAdvantage,
	aHitWithAnAdvantage,
	twoHits,
	twoHits,
	andATriumph,
})

// Challenge dice represent particularly heightened situations
var Challenge = newDie(Hexagon, Red, []Result{
	nothing,
	aFailure,
	aFailure,
	twoFailures,
	twoFailures,
	withAThreat,
	withAThreat,
	aFailureWithAThreat,
	aFailureWithAThreat,
	twoFailures,
	twoFailures,
	andADespair,
})
