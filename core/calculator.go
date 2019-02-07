package core

import (
	"math"
	"reflect"
	"sort"
)

type tMin uint
type tMax uint

type kdResult struct {
	Date       uint
	ClosePrice float64
	NHighPrice float64
	NLowPrice  float64
	RSV        float64
	K          float64
	D          float64
}

func minMaxIntSlice(slice interface{}, minMax func(element reflect.Value, min, max uint) (tMin, tMax)) (tMin, tMax) {
	min := tMin(^uint(0)) // give uint max value
	max := tMax(0)        // give uint min value
	rv := reflect.ValueOf(slice)
	length := rv.Len()
	for i := 0; i < length; i++ {
		min, max = minMax(rv.Index(i), uint(min), uint(max))
	}
	return min, max
}

// return min, max
func minMaxFloatSlice(slice interface{}, getValue func(element reflect.Value) float64) (float64, float64) {
	min := float64(0)
	max := float64(0)
	rv := reflect.ValueOf(slice)
	length := rv.Len()
	for i := 0; i < length; i++ {
		value := getValue(rv.Index(i))
		if 0 == i {
			min, max = value, value
		}
		min = math.Min(value, min)
		max = math.Max(value, max)
	}
	return min, max
}

func kdCalculator(stockDailyInfo []dailyInfo, n int) []kdResult {
	sort.Slice(stockDailyInfo, func(i, j int) bool {
		return stockDailyInfo[i].Date < stockDailyInfo[j].Date
	})
	result := []kdResult{}
	// 找最高跟最低的收盤價, 在近n天
	for i, v := range stockDailyInfo {
		// index start from 0
		if i < n-1 {
			// go default float value is 0
			result = append(result, kdResult{Date: v.Date, ClosePrice: v.ClosePrice, K: 50, D: 50})
			continue
		}
		// n days high low price
		min, max := minMaxFloatSlice(stockDailyInfo[i-(n-1):i+1],
			func(element reflect.Value) float64 {
				if v, ok := element.Interface().(dailyInfo); ok {
					return v.ClosePrice
				}
				return float64(0)
			})
		rsv := float64(100) * (v.ClosePrice - min) / (max - min)
		k := (float64(1)/3)*rsv + (float64(2)/3)*result[i-1].K
		d := (float64(1)/3)*k + (float64(2)/3)*result[i-1].D
		result = append(result, kdResult{Date: v.Date, ClosePrice: v.ClosePrice, NLowPrice: min, NHighPrice: max, RSV: rsv, K: k, D: d})
	}
	return result
}
