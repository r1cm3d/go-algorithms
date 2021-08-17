package main

import (
	"math"
	"reflect"
	"testing"
)

func calcmaxdd(series [][]float64, window int) map[int]float64 {
	if window == 0 || series == nil || len(series) < window {
		return nil
	}

	results := make(map[int]float64, len(series)-window)

	k := window
	for i := window; i < len(series); i++ {
		maxdd := maxdd(series, i-window, window)
		results[k] = maxdd
		k++
	}

	return results
}

func maxdd(series [][]float64, idx, window int) float64 {
	troughValueIndex := idx
	peakValueIndex := idx
	highestValueIndex := idx
	troughValue := series[idx][1]
	peakValue := series[idx][1]
	highestValue := series[idx][1]

	for i := idx + 1; i <= (idx + window); i++ {
		if series[i][1] < troughValue {
			troughValue = series[i][1]
			troughValueIndex = i
		}

		if series[i][1] > peakValue {
			highestValue = series[i][1]
			highestValueIndex = i
		}

		if series[i][1] > peakValue && i < troughValueIndex {
			peakValue = series[i][1]
			peakValueIndex = i
		}

		if highestValue > peakValue && highestValueIndex < troughValueIndex {
			peakValue = highestValue
		}
	}

	if highestValue > peakValue && highestValueIndex < troughValueIndex {
		peakValue = highestValue
	}

	if peakValue != 0 && troughValueIndex > peakValueIndex {
		return math.Abs(troughValue-peakValue) / peakValue
	} else {
		return 0
	}
}

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
		if got := calcmaxdd(tt.in.series, tt.in.window); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Create(%v) got: %v, want: %v", tt.name, got, tt.want)
		}
	}
}
