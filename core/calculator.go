package core

import (
	"reflect"
	"sort"
)

type tMin uint
type tMax uint

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

func kdCalculator(stockDailyInfo []dailyInfo, n uint) {
	sort.Slice(stockDailyInfo, func(i, j int) bool {
		return stockDailyInfo[i].Date < stockDailyInfo[j].Date
	})
}
