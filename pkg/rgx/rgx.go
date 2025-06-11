// Package rgx provides functions for working with regular expressions.
package rgx

import (
	"errors"
	"regexp"
)

// ErrMismatch is returned when a string does not match the provided regular expression.
// It indicates that the input string does not conform to the expected pattern defined by the regular expression.
var ErrMismatch = errors.New("string does not match regular expression")

// Group matches the input string against the provided regular expression and returns
// a map of named capture groups to their corresponding matched values.
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

// Groups matches the input string against the provided regular expression and returns
// a slice of maps, each containing named capture groups and their corresponding matched values.
// If the string does not match the regular expression, it returns nil and ErrMismatch.
//
// Example:
//
//	r := regexp.MustCompile(`(?P<name>\w+)-(?P<id>\d+)`)
//	Groups(r, "alice-42\nbob-43") // returns: []map[string]string{{"name":"alice", "id":"42"}, {"name":"bob", "id":"43"}}
func Groups(r *regexp.Regexp, str string) ([]map[string]string, error) {
	matches := r.FindAllStringSubmatch(str, -1)
	if matches == nil {
		return nil, ErrMismatch
	}

	result := make([]map[string]string, len(matches))
	for i, match := range matches {
		result[i] = make(map[string]string)
		for j, name := range r.SubexpNames() {
			if j != 0 && name != "" {
				result[i][name] = match[j]
			}
		}
	}

	return result, nil
}
