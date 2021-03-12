package common

// Range is a templace-specific helper to create range of numbers
func Range(from, count int) []int {
	rng := []int{}
	for i := from; i < count; i++ {
		rng = append(rng, i)
	}
	return rng
}
