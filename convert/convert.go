package convert

func ToPtr[T any](v T) *T {
	return &v
}

func ToVal[T any](p *T) T {
	return *p
}

func ToMap[K comparable, T any](slice []T, f func(T) K) map[K]T {
	result := make(map[K]T)
	for _, v := range slice {
		result[f(v)] = v
	}
	return result
}
