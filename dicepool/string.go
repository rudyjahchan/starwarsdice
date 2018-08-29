package starwarsdice

import "fmt"

func Plural(s string, c uint) string {
	if c == 1 {
		return s
	}
	return fmt.Sprintf("%ss", s)
}

func Message(conjuction string, count uint, s string) string {
	return fmt.Sprintf(" %s %d %s", conjuction, count, Plural(s, count))
}

func BlankOrAndMessage(count uint, s string) string {
	if count == 0 {
		return ""
	}
	return Message("and", count, s)
}
