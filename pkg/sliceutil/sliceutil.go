package sliceutil

// RemoveEmpty removes empty string elements from slice
func RemoveEmpty(slice []string) []string {
	ret := make([]string, 0, len(slice))

	for i := range slice {
		if slice[i] == "" {
			continue
		}

		ret = append(ret, slice[i])
	}

	return ret
}
