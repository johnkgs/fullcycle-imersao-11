package utils

type ArrayInterface[Type any] interface {
	Slice(arr []Type, startIdx int, endIdx int) []Type
}

func Slice[Type any](arr []Type, startIdx int, endIdx int) []Type {
	return append(arr[:startIdx], arr[endIdx:]...)
}

func Contains[Type comparable](arr []Type, value Type) bool {
	for _, item := range arr {
		if value == item {
			return true
		}
	}
	return false
}
