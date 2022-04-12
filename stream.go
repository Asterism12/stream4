package stream4

import "github.com/Asterism12/stream4/function"

// Slice returns a sequential Stream supporting 1 type with the specified slice as its source
func Slice[T any](vs []T) Stream1[T, any, any, any] {
	return Stream1[T, any, any, any]{vs: vs}
}

// Slice2 returns a sequential Stream supporting 2 types with the specified slice as its source
func Slice2[T1 any, T2 any](vs []T1) Stream1[T1, T2, any, any] {
	return Stream1[T1, T2, any, any]{vs: vs}
}

// Slice3 returns a sequential Stream supporting 3 types with the specified slice as its source
func Slice3[T1 any, T2 any, T3 any](vs []T1) Stream1[T1, T2, T3, any] {
	return Stream1[T1, T2, T3, any]{vs: vs}
}

// Slice4 returns a sequential Stream supporting 4 types with the specified slice as its source
func Slice4[T1 any, T2 any, T3 any, T4 any](vs []T1) Stream1[T1, T2, T3, T4] {
	return Stream1[T1, T2, T3, T4]{vs: vs}
}

// GroupBy grouping elements according to a classification function, and returning the results in a Map
func GroupBy[K comparable, V any](vs []V, f func(v V) K) map[K][]V {
	return function.GroupBy(vs, f)
}
