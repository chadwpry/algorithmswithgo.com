package module01

func charactersToSet(chars string) map[int]string {
	output := make(map[int]string, len(chars))
	for i, char := range chars {
		output[i] = string(char)
	}
	return output
}

var (
	characters                = "0123456789ABCDEF" // GHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	charSet    map[int]string = charactersToSet(characters)
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {
	if dec <= 0 {
		return ""
	}

	return DecToBase(dec/base, base) + charSet[dec%base]
}
