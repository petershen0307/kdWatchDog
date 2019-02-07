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
				func(element reflect.Value) float64 {
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
					return r
				})
			if min != test.expectMin || max != test.expectMax {
				t.Errorf("\nInput slice:%v\nGot min:(%v) expect min:(%v)\nGot max:(%v) expect max:(%v)", test.slice, min, test.expectMin, max, test.expectMax)
			}
		})
	}
}

func Test_kdCalculator_Always_Success(t *testing.T) {
	testTable := []struct {
		stockDailyInfo []dailyInfo
		n              int
		expect         []kdResult
	}{
		{
			stockDailyInfo: []dailyInfo{
				{Date: 20150121, ClosePrice: 67.25},
				{Date: 20150122, ClosePrice: 67.6},
				{Date: 20150123, ClosePrice: 68.7},
				{Date: 20150126, ClosePrice: 68.7},
				{Date: 20150127, ClosePrice: 69.15},
				{Date: 20150128, ClosePrice: 69.15},
				{Date: 20150129, ClosePrice: 68.3},
				{Date: 20150130, ClosePrice: 68},
				{Date: 20150202, ClosePrice: 68.15},
				{Date: 20150203, ClosePrice: 68.65},
				{Date: 20150204, ClosePrice: 69.4},
				{Date: 20150205, ClosePrice: 69.1},
				{Date: 20150206, ClosePrice: 68.75},
				{Date: 20150209, ClosePrice: 68.5},
				{Date: 20150210, ClosePrice: 68.25},
				{Date: 20150211, ClosePrice: 69},
				{Date: 20150212, ClosePrice: 69},
				{Date: 20150213, ClosePrice: 69.45},
			},
			n: 9,
			expect: []kdResult{
				{Date: 20150121, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20150122, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20150123, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20150126, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20150127, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20150128, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20150129, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20150130, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
				{Date: 20150202, NHighPrice: 69.15, NLowPrice: 67.25, RSV: 47.36842105, K: 49.12280702, D: 49.70760234},
				{Date: 20150203, NHighPrice: 69.15, NLowPrice: 67.6, RSV: 67.74193548, K: 55.32918317, D: 51.58146262},
				{Date: 20150204, NHighPrice: 69.4, NLowPrice: 68, RSV: 100, K: 70.21945545, D: 57.79412689},
				{Date: 20150205, NHighPrice: 69.4, NLowPrice: 68, RSV: 78.57142857, K: 73.00344649, D: 62.86390009},
				{Date: 20150206, NHighPrice: 69.4, NLowPrice: 68, RSV: 53.57142857, K: 66.52610718, D: 64.08463579},
				{Date: 20150209, NHighPrice: 69.4, NLowPrice: 68, RSV: 35.71428571, K: 56.25550003, D: 61.47492387},
				{Date: 20150210, NHighPrice: 69.4, NLowPrice: 68, RSV: 17.85714286, K: 43.45604764, D: 55.46863179},
				{Date: 20150211, NHighPrice: 69.4, NLowPrice: 68, RSV: 71.42857143, K: 52.78022223, D: 54.57249527},
				{Date: 20150212, NHighPrice: 69.4, NLowPrice: 68.15, RSV: 68, K: 57.85348149, D: 55.66615734},
				{Date: 20150213, NHighPrice: 69.45, NLowPrice: 68.25, RSV: 100, K: 71.90232099, D: 61.07821189},
			},
		},
	}
	const floatPrecise = 0.000001
	for _, v := range testTable {
		kd := kdCalculator(v.stockDailyInfo, v.n)
		for i, r := range kd {
			if v.expect[i].Date != r.Date {
				t.Errorf("Expect Date(%v), result(%v)", v.expect[i].Date, r.Date)
				break
			}
			if math.Abs(v.expect[i].K-r.K) > floatPrecise || math.Abs(v.expect[i].D-r.D) > floatPrecise {
				t.Errorf("KD not match Expect [K:D](%v:%v), result [K:D](%v:%v)", v.expect[i].K, v.expect[i].D, r.K, r.D)
				break
			}
		}
	}
}
