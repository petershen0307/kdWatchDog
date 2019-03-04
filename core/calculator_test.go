package core

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func Test_minMaxIntSlice_WhenCall_GetMinMaxValue(t *testing.T) {
	testTable := []struct {
		slice     interface{}
		expectMin uint
		expectMax uint
	}{
		{[]uint{1, 2, 3, 4, 5}, 1, 5},
		{[]dailyInfo{{Date: 1}, {Date: 2}, {Date: 3}, {Date: 4}, {Date: 5}}, 1, 5},
	}

	for i, test := range testTable {
		t.Run(fmt.Sprintf("test: %v", i+1), func(t *testing.T) {
			min, max := minMaxIntSlice(test.slice,
				func(element reflect.Value, min, max uint) (tMin, tMax) {
					rMin := tMin(min)
					rMax := tMax(max)
					switch element.Type().Name() {
					case "uint":
						rv, ok := element.Interface().(uint)
						if ok {
							if rv < min {
								rMin = tMin(rv)
							}
							if rv > max {
								rMax = tMax(rv)
							}
						}
					case "dailyInfo":
						rv, ok := element.Interface().(dailyInfo)
						if ok {
							if rv.Date < min {
								rMin = tMin(rv.Date)
							}
							if rv.Date > max {
								rMax = tMax(rv.Date)
							}
						}
					}
					return rMin, rMax
				})
			if uint(min) != test.expectMin || uint(max) != test.expectMax {
				t.Errorf("\nInput slice:%v\nGot min:(%v) expect min:(%v)\nGot max:(%v) expect max:(%v)", test.slice, min, test.expectMin, max, test.expectMax)
			}
		})
	}
}

func Test_minMaxFloatSlice_WhenCall_GetMinMaxValue(t *testing.T) {
	testTable := []struct {
		slice     interface{}
		expectMin float64
		expectMax float64
	}{
		{[]float64{1.1, 2.1, 3.1, 4.1, 5.1}, 1.1, 5.1},
		{[]dailyInfo{{HighPrice: 1.1}, {HighPrice: 2.1}, {HighPrice: 3.1}, {HighPrice: 4.1}, {HighPrice: 5.1}}, 1.1, 5.1},
	}

	for i, test := range testTable {
		t.Run(fmt.Sprintf("test: %v", i+1), func(t *testing.T) {
			min, max := minMaxFloatSlice(test.slice,
				func(element reflect.Value) (float64, float64) {
					r := float64(0)
					switch element.Type().Name() {
					case "float64":
						v, ok := element.Interface().(float64)
						if ok {
							r = v
						}
					case "dailyInfo":
						v, ok := element.Interface().(dailyInfo)
						if ok {
							r = v.HighPrice
						}
					}
					return r, r
				})
			if min != test.expectMin || max != test.expectMax {
				t.Errorf("\nInput slice:%v\nGot min:(%v) expect min:(%v)\nGot max:(%v) expect max:(%v)", test.slice, min, test.expectMin, max, test.expectMax)
			}
		})
	}
}

func Test_KDCalculator_Always_Success(t *testing.T) {
	testTable := []struct {
		stockDailyInfo []dailyInfo
		n              int
		expect         []KDResult
	}{
		{
			stockDailyInfo: []dailyInfo{
				{Date: 20181029, OpenPrice: 42.1, HighPrice: 42.2, LowPrice: 41.85, ClosePrice: 42.15},
				{Date: 20181030, OpenPrice: 42.05, HighPrice: 42.15, LowPrice: 41.65, ClosePrice: 41.8},
				{Date: 20181031, OpenPrice: 42, HighPrice: 42.2, LowPrice: 41.9, ClosePrice: 42},
				{Date: 20181101, OpenPrice: 42, HighPrice: 42.7, LowPrice: 42, ClosePrice: 42.5},
				{Date: 20181102, OpenPrice: 42.7, HighPrice: 42.8, LowPrice: 42.2, ClosePrice: 42.25},
				{Date: 20181105, OpenPrice: 42, HighPrice: 42.4, LowPrice: 41.9, ClosePrice: 42.35},
				{Date: 20181106, OpenPrice: 42.4, HighPrice: 42.75, LowPrice: 42.05, ClosePrice: 42.3},
				{Date: 20181107, OpenPrice: 42.4, HighPrice: 43.4, LowPrice: 42.35, ClosePrice: 43.4},
				{Date: 20181108, OpenPrice: 44, HighPrice: 44, LowPrice: 43.45, ClosePrice: 43.7},
				{Date: 20181109, OpenPrice: 43.45, HighPrice: 43.6, LowPrice: 43.25, ClosePrice: 43.35},
				{Date: 20181112, OpenPrice: 43.15, HighPrice: 43.65, LowPrice: 43.15, ClosePrice: 43.4},
				{Date: 20181113, OpenPrice: 43, HighPrice: 43.1, LowPrice: 42.75, ClosePrice: 43},
			},
			n: 9,
			expect: []KDResult{
				{Date: 20181029, ClosePrice: 42.15, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20181030, ClosePrice: 41.8, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20181031, ClosePrice: 42, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20181101, ClosePrice: 42.5, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20181102, ClosePrice: 42.25, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20181105, ClosePrice: 42.35, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20181106, ClosePrice: 42.3, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20181107, ClosePrice: 43.4, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20181108, ClosePrice: 43.7, NHighPrice: 44, NLowPrice: 41.65, RSV: 87.23404255, K: 62.41134752, D: 54.13711584},
				{Date: 20181109, ClosePrice: 43.35, NHighPrice: 44, NLowPrice: 41.65, RSV: 72.34042553, K: 65.72104019, D: 57.99842396},
				{Date: 20181112, ClosePrice: 43.4, NHighPrice: 44, NLowPrice: 41.9, RSV: 71.42857143, K: 67.6235506, D: 61.2067995},
				{Date: 20181113, ClosePrice: 43, NHighPrice: 44, NLowPrice: 41.9, RSV: 52.38095238, K: 62.54268453, D: 61.65209451},
			},
		},
	}
	const floatPrecise = 0.000001
	for _, v := range testTable {
		kd := KDCalculator(v.stockDailyInfo, v.n)
		for i, r := range kd {
			if v.expect[i].Date != r.Date {
				t.Errorf("Expect Date(%v), result(%v)", v.expect[i].Date, r.Date)
				break
			}
			if v.expect[i].ClosePrice != r.ClosePrice {
				t.Errorf("Expect ClosePrice(%v), result(%v)", v.expect[i].ClosePrice, r.ClosePrice)
				break
			}
			if math.Abs(v.expect[i].K-r.K) > floatPrecise || math.Abs(v.expect[i].D-r.D) > floatPrecise {
				t.Errorf("KD not match Expect [K:D](%v:%v), result [K:D](%v:%v)", v.expect[i].K, v.expect[i].D, r.K, r.D)
				break
			}
		}
	}
}
