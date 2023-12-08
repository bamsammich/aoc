package utils

import "golang.org/x/exp/constraints"

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
