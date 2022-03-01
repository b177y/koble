package validator

import "strings"

var ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"
var MAXLEN = 32

func IsValidName(name string) bool {
	if len(name) > MAXLEN {
		return false
	}
	for _, char := range name {
		if !strings.Contains(ALPHABET, string(char)) {
			return false
		}
	}
	return true
}
