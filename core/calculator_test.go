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
				func(element reflect.Value, min, max float64) (float64, float64) {
					rMin := min
					rMax := max
					compareValue := float64(0)
					switch element.Type().Name() {
					case "float64":
						rv, ok := element.Interface().(float64)
						if ok {
							compareValue = rv
						}
					case "dailyInfo":
						rv, ok := element.Interface().(dailyInfo)
						if ok {
							compareValue = rv.HighPrice
						}
					}
					rMax = math.Max(compareValue, max)
					rMin = math.Min(compareValue, min)
					return rMin, rMax
				})
			if min != test.expectMin || max != test.expectMax {
				t.Errorf("\nInput slice:%v\nGot min:(%v) expect min:(%v)\nGot max:(%v) expect max:(%v)", test.slice, min, test.expectMin, max, test.expectMax)
			}
		})
	}
}
