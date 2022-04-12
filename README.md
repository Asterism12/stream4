# Stream4
Stream4 is a JAVA-stream-style Go library using generics, allows streaming operation with up to 4 types.

## Quick Start
```go
package main

import (
    "fmt"
    "github.com/asterbao/stream4"
)

func main() {
	even := stream4.Slice([]int{1, 2, 3, 4, 5}).
		Filter(func(i int) bool {
			return i%2 == 0
		}).
		ToList()
	fmt.Println(even)
	// output : [2 4]

	acc := stream4.Slice2[int, float32]([]int{1, 2, 3, 4, 5}).
		Reduce2(func(v int, acc float32, i int) float32 {
			return acc + 0.1*float32(v)
		}, 0)
	fmt.Println(acc)
	// output : 1.5

	stream4.Slice3[int, float32, string]([]int{1, 2, 3}).
		Map2(func(v int) float32 {
			return float32(v) + 0.5
		}).
		Map(func(v float32) float32 {
			return v + 0.1
		}).
		Map3(func(v float32) string {
			return fmt.Sprintf("%f", v)
		}).
		ForEach(func(v string) {
			fmt.Printf("%s ", v)
		})
	// output : 1.600000 2.600000 3.600000

	min := stream4.Slice4[int, int, int, int]([]int{5, 4, 3, 2, 1}).
		Limit(3).
		Min(func(i int, j int) bool {
			return i < j
		})
	fmt.Println(min)
	// output : 3
}
```

## Supported Operation
* Map 
* Reduce 
* Filter 
* ForEach 
* Limit 
* Min 
* GroupBy(not streaming) 
