package maxdrawdown

import "math"

func Calc(series [][]float64, window int) map[int]float64 {
	if window == 0 || series == nil || len(series) < window {
		return nil
	}

	results := make(map[int]float64, len(series)-window)

	k := window
	for i := window; i < len(series); i++ {
		mdd := subMaxDD(series, i-window, window)
		results[k] = mdd
		k++
	}

	return results
}

func subMaxDD(series [][]float64, idx, window int) float64 {
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
	}

	return 0
}
