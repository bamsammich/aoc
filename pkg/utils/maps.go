package utils

import (
	"slices"

	"golang.org/x/exp/constraints"
)

func MaxValueKey[M ~map[K]V, K constraints.Ordered, V constraints.Ordered](m M) K {
	var (
		key K
		max V
	)
	for k, v := range m {
		if v > max || (v == max && k > key) {
			max = v
			key = k
		}
	}
	return key
}

func UniqueSlice[S []E, E comparable](s S) S {
	out := S{}

	for _, e := range s {
		if !slices.Contains(out, e) {
			out = append(out, e)
		}
	}
	return out
}
