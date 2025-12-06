package utils

type IdRange struct {
	Min, Max, MinLength, MaxLength int
}

func (r IdRange) ExpandIds() []int {
	size := r.TotalIds()
	ids := make([]int, size)

	for i := 0; i < size; i++ {
		ids[i] = r.Min + i
	}
	return ids
}

func (r IdRange) TotalIds() int {
	return r.Max - r.Min + 1
}
