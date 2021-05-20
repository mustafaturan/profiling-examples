package sliceofvaluevspointer

type T struct {
}

type TWithInt struct {
	current int
}

type TWithLargeData struct {
	current int
	name    string
	items   []string
}

func SliceOfValue(n int) []T {
	ts := make([]T, n)
	return ts
}

func SliceOfValueWithInt(n int) []TWithInt {
	ts := make([]TWithInt, n)
	return ts
}

func SliceOfValueWithLargeData(n int) []TWithLargeData {
	ts := make([]TWithLargeData, n)
	return ts
}

func SliceOfPointer(n int) []*T {
	ts := make([]*T, n)
	return ts
}

func SliceOfPointerWithInt(n int) []*TWithInt {
	ts := make([]*TWithInt, n)
	return ts
}

func SliceOfPointerWithLargeData(n int) []*TWithLargeData {
	ts := make([]*TWithLargeData, n)
	return ts
}
