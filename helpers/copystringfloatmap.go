package helpers

// CopyStringFloatMap copies a map[string]float64
func CopyStringFloatMap(originalMap map[string]float64) map[string]float64 {
	newMap := make(map[string]float64, len(originalMap))
	for k, v := range originalMap {
		newMap[k] = v
	}
	return newMap
}
