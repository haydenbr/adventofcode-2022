package util

func Fold[T any, R any](initial R, s []T, f func(R, T, int) R) R {
	result := initial

	for i, v := range s {
		result = f(result, v, i)
	}

	return result
}

func Map[T any, R any](s []T, f func(T, int) R) []R {
	return Fold(make([]R, len(s)), s, func(result []R, element T, index int) []R {
		result[index] = f(element, index)
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
	return Fold(0, numbers, func(result T, element T, _ int) T {
		return result + element
	})
}