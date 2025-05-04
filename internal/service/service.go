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
		// Если это код Морзе, пытаемся конвертировать в текст
		result = morse.ToText(data)
		if result == "" { // Проверяем на пустую строку как индикатор ошибки
			return "", errors.New("the response does not contain the expected Morse code")
		}
	} else {
		// Если это текст, пытаемся конвертировать в код Морзе
		result = morse.ToMorse(data)
		if result == "" { // Проверяем на пустую строку как индикатор ошибки
			return "", errors.New("failed to convert text to morse code")
		}
	}
	return result, nil
}
