// Package sm2 implements functions for SM2 algorithm calculations
package sm2

import "math"

func GetInterval(ef float64, repetition int, prevInterval int) int {
	switch repetition {
	case 0:
		return 1
	case 1:
		return 6
	}

	return int(float64(prevInterval) * ef)
}

func GetEF(ef float64, q int) float64 {
	return math.Max(ef+(0.1-float64(5-q)*(0.08+float64(5-q)*0.02)), 1.3)
}
