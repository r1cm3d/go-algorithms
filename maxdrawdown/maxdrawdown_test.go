package maxdrawdown

import (
	"math"
	"reflect"
	"testing"
)

func Calc(series map[int]float64, window int) map[int]float64 {
	if window == 0 || series == nil || len(series) < window {
		return nil
	}

	results := make(map[int]float64, len(series)-window)

	curKey := window
	for i := window; i < len(series); i++ {
		initialKey := i-window
		mdd := subMaxDD(series, initialKey, window)
		results[curKey] = mdd
		curKey++
	}

	return results
}

func subMaxDD(series map[int]float64, initialKey, window int) float64 {
	troughKey, highestKey := initialKey, initialKey
	trough, peak, highest := series[initialKey], series[initialKey], series[initialKey]

	for i := initialKey + 1; i <= (initialKey + window); i++ {
		curVal := series[i]

		if curVal > peak {
			highest = curVal
			highestKey = i

			if i < troughKey {
				peak = curVal
			}
		}

		if curVal < trough {
			trough = curVal
			troughKey = i
		}

		if highestKey < troughKey && highest > peak {
			peak = highest
		}
	}

	return math.Abs(trough-peak) / peak
}

func TestCalcMaxDD(t *testing.T) {
	type in struct {
		series map[int]float64
		window int
	}
	cases := []struct {
		name string
		in
		want map[int]float64
	}{
		{"nil", in{series: nil, window: 5}, nil},
		{"series length lower than window", in{series: map[int]float64{
			0: 500,
			1: 750,
			2: 400,
			3: 600,
		}, window: 5}, nil},
		{"assessment example", in{series: map[int]float64{
			0: 12,
			1: 11,
			2: 13,
			3: 16,
			4: 10,
			5: 11,
			6: 8,
			7: 9,
			8: 12,
		}, window: 3}, map[int]float64{
			3: 0.08333333333333333,
			4: 0.375,
			5: 0.375,
			6: 0.5,
			7: 0.2727272727272727,
			8: 0.2727272727272727,
		}},
		{"investopedia example", in{series: map[int]float64{
			0: 500,
			1: 750,
			2: 400,
			3:  600,
			4:  350,
			5:  800,
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