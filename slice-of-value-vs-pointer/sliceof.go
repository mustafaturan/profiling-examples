package sliceofvaluevspointer

type T struct{}

func SliceOfValue(n int) []T {
	ts := make([]T, n)
	return ts
}

func SliceOfPointer(n int) []*T {
	ts := make([]*T, n)
	return ts
}
