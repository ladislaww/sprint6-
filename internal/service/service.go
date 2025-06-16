package service

import (
	"os"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func FileContent(path string) (string, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	var result string
	text := string(content)

	if strings.ContainsAny(strings.ToLower(text), "абвгдеёжзийклмнопрстуфхцчшщъыьэюя") {
		result = morse.ToMorse(text)

	} else {
		result = morse.ToText(text)
	}

	if err := os.WriteFile(path, []byte(result), 0755); err != nil {
		return "", err
	}
	return path, nil
}
