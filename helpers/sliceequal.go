package helpers

// SliceEqual checks if two slices contains the same strings, regardless of order
func SliceEqual(a, b []string) bool {
	if len(a) != len(b) {
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

	for str := range aMap {
		if !bMap[str] {
			return false
		}
	}

	return true
}
