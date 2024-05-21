package module01

import "unicode/utf8"

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	if len(word) == 0 {
		return ""
	}

	b := []byte(word)
	r, w := utf8.DecodeRune(b)

	return Reverse(word[w:]) + string(r)
}
