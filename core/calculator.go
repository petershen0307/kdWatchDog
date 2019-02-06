package core

import (
	"reflect"
	"sort"
)

type tMin int
type tMax int

func minMaxIntSlice(slice interface{}, minMax func(element reflect.Value, min, max int) (tMin, tMax)) (tMin, tMax) {
	min := tMin(int((^uint(0)) >> 1))         // give int max value
	max := tMax(int(-int((^uint(0))>>1) - 1)) // give int min value
	rv := reflect.ValueOf(slice)
	length := rv.Len()
	for i := 0; i < length; i++ {
		min, max = minMax(rv.Index(i), int(min), int(max))
	}
	return min, max
}

func kdCalculator(stockDailyInfo []dailyInfo, n uint) {
	sort.Slice(stockDailyInfo, func(i, j int) bool {
		return stockDailyInfo[i].Date < stockDailyInfo[j].Date
	})
}
