package sumarr

import (
	"testing"
)

func sum(input interface{}) int {
	arr, ok := input.([]interface{})

	if !ok {
		v := input.(int)
		return v
	}

	var s int
	for i := 0; i < len(arr); i++ {
		s = s + sum(arr[i])
	}

	return s
}

func TestSum(t *testing.T) {

	cases := []struct {
		name string
		in   []interface{}
		want int
	}{
		{"first", []interface{}{1}, 1},
		{"second", []interface{}{1, []interface{}{2, 3}}, 6},
		{"third", []interface{}{
			-1,
			[]interface{}{-2, -3},
			[]interface{}{1, []interface{}{2, 3}}},
			0},
	}
	for _, tt := range cases {
		if got := sum(tt.in); got != tt.want {
			t.Errorf("sum(%v) got: %v, want: %v", tt.name, got, tt.want)
		}
	}
}