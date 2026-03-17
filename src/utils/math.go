package utils

import (
	"math"
)

func CalculateMA(data []float64, period int) float64 {
	var sum float64
	for i := len(data) - period; i < len(data); i++ {
		sum += data[i]
	}
	return sum / float64(period)
}
