package util

type IntRange struct {
	Start int
	End   int
}

func (r IntRange) InsideInclusive(n int) bool {
	return n <= r.End && n >= r.Start
}

// Implements sort inteface for range
type ByStart []IntRange

func (r ByStart) Len() int { return len(r) }
func (r ByStart) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r ByStart) Less(i, j int) bool { return r[i].Start < r[j].Start }
