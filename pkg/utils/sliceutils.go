package utils

// RemoveDuplicates remove dublicates in slice.
func RemoveDuplicates[T comparable](arr []T) []T {
	seen := make(map[T]struct{})
	unique := make([]T, 0, len(arr))

	for _, item := range arr {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			unique = append(unique, item)
		}
	}

	return unique
}
