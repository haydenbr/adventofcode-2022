package util

func Fold[T any, R any](initial R, s []T, f func(R, T) R) R {
	result := initial

	for _, v := range s {
		result = f(result, v)
	}

	return result
}

func Filter[T any](s []T, predicate func(T) bool) []T {
	result := make([]T, 0)

	// for i, e: = range s {
	// 	if
	// }
	for _, e := range s {
		if predicate(e) {
			result = append(result, e)
		}
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

func ProjectTo[T any, R any](s []T, f func(T) R) []R {
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

func SumWith[TSlice any, TResult Number](s []TSlice, fn func(TSlice) TResult) TResult {
	return Fold(0, s, func(result TResult, element TSlice) TResult {
		return result + fn(element)
	})
}

func SliceToMap[TSlice any, TMapKey comparable, TMapValue any](s []TSlice, keySelector func(TSlice) TMapKey, valueSelector func(TSlice) TMapValue) map[TMapKey]TMapValue {
	m := make(map[TMapKey]TMapValue)

	for _, e := range s {
		m[keySelector(e)] = valueSelector(e)
	}

	return m
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

func NthIndexOf[T any](s []T, n int, predicate func(T) bool) int {
	count := 0

	for i, e := range s {
		match := predicate(e)

		if match {
			count++
		}

		if count == n {
			return i
		}
	}

	return -1
}

func FirstIndexOf[T any](s []T, predicate func(T) bool) int {
	return NthIndexOf(s, 1, predicate)
}

func PopN[T any](s []T, n int) ([]T, []T) {
	popped := make([]T, n)

	for i := 0; i < n; i++ {
		popped[i] = s[len(s)-1-i]
	}

	return s[:len(s)-n], popped
}

func CutN[T any](s []T, n int) ([]T, []T) {
	cutPoint := len(s) - n
	return s[:cutPoint], s[cutPoint:]
}

func Peek[T any](s []T) T {
	return s[len(s)-1]
}

func Push[T any](s []T, e ...T) []T {
	return append(s, e...)
}

func TrimLastN[T any](s []T, n int) []T {
	return s[:len(s)-n]
}

func TrimLast[T any](s []T) []T {
	return TrimLastN(s, 1)
}

func ForEachReverse[T any](s []T, forEachFn func(e T)) {
	for i := range s {
		i = len(s) - 1 - i
		forEachFn(s[i])
	}
}
