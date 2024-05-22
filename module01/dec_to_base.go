package module01

func genDecToStrMap(chars string) map[int]string {
	output := make(map[int]string, len(chars))
	for i, rn := range chars {
		output[i] = string(rn)
	}
	return output
}

func genStrToDecMap(chars string) map[string]int {
	output := make(map[string]int, len(chars))
	for i, rn := range chars {
		output[string(rn)] = i
	}
	return output
}

var (
	characters                 = "0123456789ABCDEF" // GHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	decToStrMap map[int]string = genDecToStrMap(characters)
	strToDecMap map[string]int = genStrToDecMap(characters)
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

	return DecToBase(dec/base, base) + decToStrMap[dec%base]
}
