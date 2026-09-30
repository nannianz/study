package main

import (
	"fmt"
)

func factorial(b int) int {
	fmt.Println("b:", b)
	if b <= 1 {
		return 1
	}
	return b * factorial(b-1)
}
func Map[T, U any](in []T, f func(T) U) []U {
	out := make([]U, 0, len(in))
	for _, v := range in {
		out = append(out, f(v))
	}
	return out
}
func Filter[T any](in []T, keep func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}
func Contains[T comparable](xs []T, target T) bool {
	for _, v := range xs {
		if v == target {
			return true
		}
	}
	return false
}

type Device struct {
	Name    string
	APCount int
}

func main() {
	fmt.Println(factorial(5))
	// b: 5
	// b: 4
	// b: 3
	// b: 2
	// b: 1
	// 120
	devices := []Device{
		{Name: "phone", APCount: 1},
		{Name: "table", APCount: 2},
		{Name: "pc", APCount: 3},
	}
	name := Map(devices, func(d Device) string {
		return d.Name
	})
	fmt.Println("name:", name)
	apCount := Map(devices, func(d Device) int {
		return d.APCount
	})
	fmt.Println("apCount:", apCount)

	outline := Filter(devices, func(d Device) bool {
		return d.APCount == 2
	})
	fmt.Println("outline:", outline)

	c := Contains(name, "phone")
	fmt.Println("c:", c)
	// name: [phone table pc]
	// apCount: [1 2 3]
	// outline: [{table 2}]
	// c: true
}
