package helpers

// SliceContains checks if all the strings in the second slice are found in the first slice
func SliceContains(a, b []string) bool {
	if len(a) < len(b) {
		return false
	}

	aMap := make(map[string]bool)
	for _, str := range a {
		aMap[str] = true
	}
	bMap := make(map[string]bool)
	for _, str := range b {
		bMap[str] = true
	}

	for str := range bMap {
		if !aMap[str] {
			return false
		}
	}
	return true
}
