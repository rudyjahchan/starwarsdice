package starwarsdice

import "testing"

func TestAdd(t *testing.T) {
	aresult := Result{3, 1, 1, 0, 0, 0}
	anotherresult := Result{0, 0, 0, 2, 2, 0}

	result := aresult.Add(anotherresult)
	if result.Successes != 3 {
		t.Errorf("%q.Add(%q).Successes == %q, want %q", aresult, anotherresult, result.Successes, 3)
	}
}
