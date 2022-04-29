package utils

// Partition splits the array of elements into two groups:
// the left partition contains elements that the predicate returns `true` for.
// the right partition contains elements that the predicate returns `false` for.
//
// The predicate is invoked with each element
func Partition[V any](elements []V, predicate func(V) bool) ([]V, []V) {
	leftPartition := []V{}
	rightPartition := []V{}

	for _, v := range elements {
		if predicate(v) {
			leftPartition = append(leftPartition, v)
		} else {
			rightPartition = append(rightPartition, v)
		}
	}

	return leftPartition, rightPartition
}