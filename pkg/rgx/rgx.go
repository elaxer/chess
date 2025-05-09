// Package rgx provides functions for working with regular expressions.
package rgx

import (
	"errors"
	"regexp"
)

var ErrMismatch = errors.New("string does not match regular expression")

// Group matches the input string against the provided regular expression and returns
// a map of named capture groups to their corresponding matched values.
//
// If the string does not match the regular expression, it returns ErrMismatch.
//
// Example:
//
//	r := regexp.MustCompile(`(?P<name>\w+)-(?P<id>\d+)`)
//	Group(r, "alice-42") // returns: map["name":"alice", "id":"42"]
func Group(r *regexp.Regexp, str string) (map[string]string, error) {
	match := r.FindStringSubmatch(str)
	if match == nil {
		return nil, ErrMismatch
	}

	result := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result, nil
}
