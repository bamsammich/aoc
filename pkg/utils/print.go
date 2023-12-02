package utils

import "fmt"

func Printlns[T any](lines ...T) {
	for _, l := range lines {
		fmt.Println(l)
	}
}
