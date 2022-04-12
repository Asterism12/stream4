package function

// Map returns the results of applying the given function to the elements of this slice
func Map[T any, R any](vs []T, f func(T) R) []R {
	nvs := make([]R, len(vs))
	for i := range vs {
		nvs[i] = f(vs[i])
	}
	return nvs
}

// Reduce performs a reduction on the elements of this slice,
// using the provided identity, accumulation and combining functions
func Reduce[T any, R any](vs []T, f func(T, R, int) R, initial R) R {
	for i := range vs {
		initial = f(vs[i], initial, i)
	}
	return initial
}

// Filter returns the results of the elements of this slice that match the given predicate
func Filter[T any](vs []T, f func(T) bool) []T {
	var nvs []T
	for i := range vs {
		if f(vs[i]) {
			nvs = append(nvs, vs[i])
		}
	}
	return nvs
}

// ForEach performs an action for each element of this slice
func ForEach[T any](vs []T, f func(T)) {
	for i := range vs {
		f(vs[i])
	}
}

// Limit returns the results of the elements of this slice, truncated to be no longer than maxSize in length
func Limit[T any](vs []T, limit int) []T {
	if len(vs) < limit {
		return vs
	}
	return vs[0:limit]
}

// Min returns the minimum element of this stream according to the provided less function
func Min[T any](vs []T, less func(T, T) bool) T {
	min := vs[0]
	for i := range vs {
		if less(vs[i], min) {
			min = vs[i]
		}
	}
	return min
}

// GroupBy grouping elements according to a classification function, and returning the results in a Map
func GroupBy[K comparable, V any](vs []V, f func(V) K) map[K][]V {
	mp := make(map[K][]V, len(vs))
	for i := range vs {
		mp[f(vs[i])] = append(mp[f(vs[i])], vs[i])
	}
	return mp
}
