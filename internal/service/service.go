package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var ErrInvalidInput = errors.New("invalid input: not Morse code or plain text")

func Convert(input string) (string, error) {

	input = strings.TrimSpace(input)

	// Проверяем, является ли строка кодом Морзе
	if isMorseCode(input) {
		return morse.ToText(input), nil
	}

	// Если это не код Морзе, пробуем конвертировать его в код Морзе
	return morse.ToMorse(input), nil
}

func isMorseCode(s string) bool {
	// Разделяем строку на слова по пробелам
	words := strings.Fields(s)

	for _, word := range words {
		// Проверяем каждый символ в слове
		for _, char := range word {
			if char != '.' && char != '-' && char != ' ' {
				return false // Если есть недопустимый символ, возвращаем false
			}
		}
	}
	return true // Все символы допустимы для кода Морзе
}
