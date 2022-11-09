package array

// DiffArray the difference between two arrays
//
// return`s b exists but a does not exist
func Diff(a, b []string) []string {
	m := make(map[string]bool)
	for _, item := range a {
		m[item] = true
	}

	var diff []string
	for _, item := range b {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}
