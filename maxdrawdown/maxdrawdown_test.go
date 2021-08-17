package maxdrawdown

import (
	"reflect"
	"testing"
)

func TestCalcMaxDD(t *testing.T) {
	type in struct {
		series [][]float64
		window int
	}
	cases := []struct {
		name string
		in
		want map[int]float64
	}{
		{"nil", in{series: nil, window: 5}, nil},
		{"series length lower than window", in{series: [][]float64{
			{0, 500},
			{1, 750},
			{2, 400},
			{3, 600},
		}, window: 5}, nil},
		{"assessment example", in{series: [][]float64{
			{0, 12},
			{1, 11},
			{2, 13},
			{3, 16},
			{4, 10},
			{5, 11},
			{6, 8},
			{7, 9},
			{8, 12},
		}, window: 3}, map[int]float64{
			3: 0.08333333333333333,
			4: 0.375,
			5: 0.375,
			6: 0.5,
			7: 0.2727272727272727,
			8: 0.2727272727272727,
		}},
		{"investopedia example", in{series: [][]float64{
			{0, 500},
			{1, 750},
			{2, 400},
			{3, 600},
			{4, 350},
			{5, 800},
		}, window: 5}, map[int]float64{
			5: 0.5333333333333333,
		}},
	}

	for _, tt := range cases {
		if got := Calc(tt.in.series, tt.in.window); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Create(%v) got: %v, want: %v", tt.name, got, tt.want)
		}
	}
}