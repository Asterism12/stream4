package stream4

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	even := Slice([]int{1, 2, 3, 4, 5}).
		Filter(func(i int) bool {
			return i%2 == 0
		}).
		ToList()

	fmt.Println(even)
	// output : [2 4]
}

func TestExample2(t *testing.T) {
	acc := Slice2[int, float32]([]int{1, 2, 3, 4, 5}).
		Reduce2(func(v int, acc float32, i int) float32 {
			return acc + 0.1*float32(v)
		}, 0)

	fmt.Println(acc)
	// output : 1.5
}

func TestExample3(t *testing.T) {
	Slice3[int, float32, string]([]int{1, 2, 3}).
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
}

func TestExample4(t *testing.T) {
	min := Slice4[int, int, int, int]([]int{5, 4, 3, 2, 1}).
		Limit(3).
		Min(func(i int, j int) bool {
			return i < j
		})

	fmt.Println(min)
	// output : 3
}
