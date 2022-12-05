package utils

type ArrayInterface[Type any] interface {
	Slice(arr []Type, startIdx int, endIdx int) []Type
}

func Slice[Type any](arr []Type, startIdx int, endIdx int) []Type {
	return append(arr[:startIdx], arr[endIdx:]...)
}
