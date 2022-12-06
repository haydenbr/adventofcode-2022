package util

func Fold[T any, R any](initial R, s []T, f func(R, T) R) R {
	result := initial

	for _, v := range s {
		result = f(result, v)
	}

	return result
}

func FoldIndex[T any, R any](initial R, s []T, f func(R, T, int) R) R {
	result := initial

	for i, v := range s {
		result = f(result, v, i)
	}

	return result
}

func Map[T any, R any](s []T, f func(T) R) []R {
	return FoldIndex(make([]R, len(s)), s, func(result []R, element T, index int) []R {
		result[index] = f(element)
		return result
	})
}

func Last[T any](s []T) T {
	return s[len(s)-1]
}

func Lastn[T any](s []T, n int) []T {
	return s[len(s)-n:]
}

type Number interface {
	int | int64 | float64
}

func Sum[T Number](numbers []T) T {
	return Fold(0, numbers, func(result T, element T) T {
		return result + element
	})
}

func PartitionEveryN[T any](source []T, groupSize int) [][]T {
	totalGroups := len(source) / groupSize

	if len(source)%groupSize > 0 {
		totalGroups += 1
	}

	groups := make([][]T, totalGroups)

	for groupIndex := range groups {
		group := make([]T, groupSize)
		groups[groupIndex] = group

		for itemIndex := 0; itemIndex < groupSize; itemIndex++ {
			nextIndex := (groupIndex * groupSize) + itemIndex

			if nextIndex < len(source) {
				group[itemIndex] = source[nextIndex]
			} else {
				break
			}
		}
	}

	return groups
}

func PartitionHalf[T any](s []T) ([]T, []T) {
	halfLength := len(s) / 2
	return s[:halfLength], s[halfLength:]
}
