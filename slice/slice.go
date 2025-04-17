package slice

func Contains[T comparable](values []T, target T) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

// Extract は、任意の型 T のスライスから R 型のフィールドを抽出する汎用関数です。
func Extract[T any, R any](items []T, selector func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = selector(item)
	}
	return result
}
