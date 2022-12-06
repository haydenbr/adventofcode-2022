package util

import sets "github.com/deckarep/golang-set/v2"

func NewSetFromSlice[T comparable](slice []T) sets.Set[T] {
	newSet := sets.NewSet[T]()

	for _, v := range slice {
		newSet.Add(v)
	}

	return newSet
}
