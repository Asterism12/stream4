package stream4

import "github.com/Asterism12/stream4/function"

// Stream2 a sequence of elements of type 2
type Stream2[T1 any, T2 any, T3 any, T4 any] struct {
	vs []T2
}

// Map1 returns the results of applying the given function to the elements of this slice
//
// convert result to type 1
func (s Stream2[T1, T2, T3, T4]) Map1(f func(T2) T1) Stream1[T1, T2, T3, T4] {
	return Stream1[T1, T2, T3, T4]{vs: function.Map(s.vs, f)}
}

// Map2 returns the results of applying the given function to the elements of this slice
//
// convert result to type 2
func (s Stream2[T1, T2, T3, T4]) Map2(f func(T2) T2) Stream2[T1, T2, T3, T4] {
	return Stream2[T1, T2, T3, T4]{vs: function.Map(s.vs, f)}
}

// Map3 returns the results of applying the given function to the elements of this slice
//
// convert result to type 3
func (s Stream2[T1, T2, T3, T4]) Map3(f func(T2) T3) Stream3[T1, T2, T3, T4] {
	return Stream3[T1, T2, T3, T4]{vs: function.Map(s.vs, f)}
}

// Map4 returns the results of applying the given function to the elements of this slice
//
// convert result to type 4
func (s Stream2[T1, T2, T3, T4]) Map4(f func(T2) T4) Stream4[T1, T2, T3, T4] {
	return Stream4[T1, T2, T3, T4]{vs: function.Map(s.vs, f)}
}

// Map returns the results of applying the given function to the elements of this slice
func (s Stream2[T1, T2, T3, T4]) Map(f func(T2) T2) Stream2[T1, T2, T3, T4] {
	return s.Map2(f)
}

// Reduce1 performs a reduction on the elements of this slice,
// using the provided identity, accumulation and combining functions
//
// convert result to type 1
func (s Stream2[T1, T2, T3, T4]) Reduce1(f func(v T2, acc T1, i int) T1, initial T1) T1 {
	return function.Reduce(s.vs, f, initial)
}

// Reduce2 performs a reduction on the elements of this slice,
// using the provided identity, accumulation and combining functions
//
// convert result to type 2
func (s Stream2[T1, T2, T3, T4]) Reduce2(f func(v T2, acc T2, i int) T2, initial T2) T2 {
	return function.Reduce(s.vs, f, initial)
}

// Reduce3 performs a reduction on the elements of this slice,
// using the provided identity, accumulation and combining functions
//
// convert result to type 3
func (s Stream2[T1, T2, T3, T4]) Reduce3(f func(v T2, acc T3, i int) T3, initial T3) T3 {
	return function.Reduce(s.vs, f, initial)
}

// Reduce4 performs a reduction on the elements of this slice,
// using the provided identity, accumulation and combining functions
//
// convert result to type 4
func (s Stream2[T1, T2, T3, T4]) Reduce4(f func(v T2, acc T4, i int) T4, initial T4) T4 {
	return function.Reduce(s.vs, f, initial)
}

// Reduce performs a reduction on the elements of this slice,
// using the provided identity, accumulation and combining functions
func (s Stream2[T1, T2, T3, T4]) Reduce(f func(v T2, acc T2, i int) T2, initial T2) T2 {
	return s.Reduce2(f, initial)
}

// Filter returns the results of the elements of this slice that match the given predicate
func (s Stream2[T1, T2, T3, T4]) Filter(f func(T2) bool) Stream2[T1, T2, T3, T4] {
	return Stream2[T1, T2, T3, T4]{vs: function.Filter(s.vs, f)}
}

// ForEach performs an action for each element of this slice
func (s Stream2[T1, T2, T3, T4]) ForEach(f func(T2)) {
	function.ForEach(s.vs, f)
}

// Limit returns the results of the elements of this slice, truncated to be no longer than maxSize in length
func (s Stream2[T1, T2, T3, T4]) Limit(limit int) Stream2[T1, T2, T3, T4] {
	return Stream2[T1, T2, T3, T4]{vs: function.Limit(s.vs, limit)}
}

// Min returns the minimum element of this stream according to the provided less function
func (s Stream2[T1, T2, T3, T4]) Min(less func(T2, T2) bool) T2 {
	return function.Min(s.vs, less)
}

// ToList returns the slice
func (s Stream2[T1, T2, T3, T4]) ToList() []T2 {
	return s.vs
}
