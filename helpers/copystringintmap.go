package helpers

// CopyStringIntMap copies a map[string]int
func CopyStringIntMap(originalMap map[string]int) map[string]int {
	newMap := make(map[string]int, len(originalMap))
	for k, v := range originalMap {
		newMap[k] = v
	}
	return newMap
}
