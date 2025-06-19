package service

import (
	
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)







func isMorse(s string) bool {
	for _, ch := range s {
		if ch != '.' && ch != '-' && ch != ' ' && ch != '\n' && ch != '\r' {
			return false
		}
	}
	return true
}



func ConvertMorseOrText(data string) string {
	
	data = strings.TrimSpace(data)
	
	if isMorse(data){
		return morse.ToText(data)
	} 
	return morse.ToMorse(data)
}
