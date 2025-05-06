// Description: Пакет rgx предоставляет функции для работы с регулярными выражениями.
package rgx

import (
	"errors"
	"regexp"
)

var ErrNotMatch = errors.New("строка не совпадает с регулярным выражением")

// Group возвращает именованные группы и их значения из регулярного выражения и строки.
// Возвращает ошибку ErrNotMatch, если строка не совпадает с регулярным выражением.
func Group(r *regexp.Regexp, str string) (map[string]string, error) {
	match := r.FindStringSubmatch(str)
	if match == nil {
		return nil, ErrNotMatch
	}

	result := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result, nil
}
