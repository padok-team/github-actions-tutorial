package foobar

import (
	"errors"
	"strconv"
)

// Sequence returns a sequence of numbers from 1 to length, where multiples of
// 3 are replaced by "foo", multiples of 5 are replaced by "bar", and multiples
// of both 3 and 5 are rpelaced by "foobar".
func Sequence(length int) ([]string, error) {
	if length < 0 {
		return nil, errors.New("length is negative")
	}

	seq := make([]string, length)

	for i := range seq {
		n := i + 1
		switch {
		case n%3 == 0 && n%5 == 0:
			seq[i] = "foobar"
		case n%7 == 0:
			seq[i] = "bar"
		case n%3 == 0:
			seq[i] = "foo"
		default:
			seq[i] = strconv.Itoa(n)
		}
	}

	return seq, nil
}
