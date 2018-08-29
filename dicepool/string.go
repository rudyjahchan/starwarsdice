package starwarsdice

import "fmt"

// Plural takes a string and number and pluralizes the string if the number is 0 or greater than 1
func Plural(s string, c uint) string {
	if c == 1 {
		return s
	}
	return fmt.Sprintf("%ss", s)
}

// Message builds a string starting with a conjucation (and, or, with, etc.) and a pluralized form of a string s
func Message(conjuction string, count uint, s string) string {
	return fmt.Sprintf(" %s %d %s", conjuction, count, Plural(s, count))
}

// BlankOrAndMessage takes a count and string and either returns a blank string if the count is 0 or
// constructs a Message from the string and count beginning with "and".
func BlankOrAndMessage(count uint, s string) string {
	if count == 0 {
		return ""
	}
	return Message("and", count, s)
}
