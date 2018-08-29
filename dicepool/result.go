package starwarsdice

import (
	"fmt"
	"strings"
)

// Result represents the cumalitive outcome of rolling dice.
type Result struct {
	Successes  uint
	Advantages uint
	Triumphs   uint

	Failures uint
	Threats  uint
	Despairs uint
}

// Add returns a new Result adding the fields of the original Result with the other Result passed in.
func (r Result) Add(other Result) Result {
	return Result{
		Successes:  r.Successes + other.Successes,
		Advantages: r.Advantages + other.Advantages,
		Triumphs:   r.Triumphs + other.Triumphs,
		Failures:   r.Failures + other.Failures,
		Threats:    r.Threats + other.Threats,
		Despairs:   r.Despairs + other.Despairs,
	}
}

func (r Result) String() string {
	var s strings.Builder
	s.WriteString(r.successOrFailureMessage())
	s.WriteString(r.flavorMessage())
	s.WriteString(r.triumphMessage())
	s.WriteString(r.despairMessage())
	return s.String()
}

func (r Result) successOrFailureMessage() string {
	hits := r.hits()

	if hits == 0 {
		return "Did not fail or succeed"
	}

	result := "Fail"
	if hits > 0 {
		result = "Succeed"
	}

	by := ""
	if hits != 0 {
		by = fmt.Sprintf(" by %d", Abs(hits))
	}

	return fmt.Sprintf("%s%s", result, by)
}

func (r Result) hits() int {
	return int(r.Successes+r.Triumphs) - int(r.Failures+r.Despairs)
}

func (r Result) flavorMessage() string {
	flavor := r.flavor()

	if flavor == 0 {
		return ""
	}

	var result string
	if flavor > 0 {
		result = "advantage"
	} else {
		result = "disadvantage"
	}

	return Message("with", Abs(flavor), result)
}

func (r Result) flavor() int {
	return int(r.Advantages - r.Threats)
}

func (r Result) triumphMessage() string {
	return BlankOrAndMessage(r.Triumphs, "triumph")
}

func (r Result) despairMessage() string {
	return BlankOrAndMessage(r.Despairs, "despair")
}
