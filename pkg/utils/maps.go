package utils

import "golang.org/x/exp/constraints"

func MaxValueKey[M ~map[K]V, K comparable, V constraints.Ordered](m M) K {
	var (
		key K
		max V
	)
	for k, v := range m {
		if v > max {
			max = v
			key = k
		}
	}
	return key
}
