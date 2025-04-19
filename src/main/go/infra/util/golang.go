package util

func ToPointerSlice[S any](slice []S) []*S {
	result := make([]*S, len(slice))
	for index, value := range slice {
		result[index] = &value
	}
	return result
}
