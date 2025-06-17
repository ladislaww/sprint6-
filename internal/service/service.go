package service

import (
	
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertMorseOrText(data string) string {
	
	var result string

	if strings.ContainsAny(strings.ToLower(data), " абвгдеёжзийклмнопрстуфхцчшщъыьэюя") {
		result = morse.ToMorse(data)

	} else {
		result = morse.ToText(data)
	}

	return result
}
