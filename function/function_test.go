package function

import (
	"reflect"
	"strconv"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		vs []int
		f  func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "even",
			args: args{
				vs: []int{1, 2, 3, 4, 5},
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.vs, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupBy(t *testing.T) {
	type args struct {
		vs []int
		f  func(int) int
	}
	tests := []struct {
		name string
		args args
		want map[int][]int
	}{
		{
			name: "remainder divided by 3",
			args: args{
				vs: []int{1, 2, 3, 4, 5},
				f: func(i int) int {
					return i % 3
				},
			},
			want: map[int][]int{
				0: {3}, 1: {1, 4}, 2: {2, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.vs, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLimit(t *testing.T) {
	type args struct {
		vs    []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "limit",
			args: args{
				vs:    []int{1, 2, 3, 4, 5},
				limit: 3,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Limit(tt.args.vs, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		vs []int
		f  func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "is even",
			args: args{
				vs: []int{1, 2, 3, 4, 5},
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []bool{false, true, false, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.vs, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		vs   []int
		less func(int, int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "min",
			args: args{
				vs: []int{2, 4, 1, 5, 3},
				less: func(i int, j int) bool {
					return i < j
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.vs, tt.args.less); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		vs      []int
		f       func(int, string, int) string
		initial string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string join",
			args: args{
				vs: []int{1, 2, 3, 4, 5},
				f: func(v int, s string, _ int) string {
					return s + strconv.Itoa(v)
				},
				initial: "0",
			},
			want: "012345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.vs, tt.args.f, tt.args.initial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
