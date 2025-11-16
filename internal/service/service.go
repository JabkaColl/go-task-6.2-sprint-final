package service

import (
	"strings"

	"github.com/JabkaColl/go-task-6.2-sprint-final/pkg/morse"
)

func ConvertString(input string) (string, error) {
	strNew := strings.TrimSpace(input)

	if strNew == "" {
		return "", nil
	}

	if isMorse(strNew) {
		// Морзе → Текст
		result := morse.ToText(strNew)
		return result, nil
	} else {
		// Текст → Морзе
		result := morse.ToMorse(strNew)
		return result, nil
	}
}

func isMorse(s string) bool {
	if s == "" {
		return false
	}

	// Проверяем, что строка содержит только символы Морзе
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '/' && r != '\n' && r != '\r' {
			return false
		}
	}

	// Дополнительная проверка: должны быть точки/тире (не только пробелы)
	hasMorseChars := false
	for _, r := range s {
		if r == '.' || r == '-' {
			hasMorseChars = true
			break
		}
	}

	return hasMorseChars
}
