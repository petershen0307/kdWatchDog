package core

import (
	"fmt"
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
					}
					return rMin, rMax
				})
			if uint(min) != test.expectMin || uint(max) != test.expectMax {
				t.Errorf("\nInput slice:%v\nGot min:(%v) expect min:(%v)\nGot max:(%v) expect max:(%v)", test.slice, min, test.expectMin, max, test.expectMax)
			}
		})
	}
}
