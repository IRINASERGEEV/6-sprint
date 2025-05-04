package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	if len(data) == 0 {
		return "", errors.New("no data to convert")
	}

	var result string

	if strings.ContainsAny(data, ".- ") {
		result = morse.ToText(data)
	} else {
		result = morse.ToMorse(data)
	}
	return result, nil
}
