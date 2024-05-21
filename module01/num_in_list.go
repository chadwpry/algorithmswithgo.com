package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	if len(list) == 0 {
		return false
	}

	if list[0] == num {
		return true
	}

	return NumInList(list[1:], num)
}
