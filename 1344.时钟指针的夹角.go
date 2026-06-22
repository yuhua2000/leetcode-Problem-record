package leetcode

import (
	"math"
)

func angleClock(hour int, minutes int) float64 {
	hour %= 12
	minutes %= 60
	hourAngle := 360*(float64(hour)/12) + 30*(float64(minutes)/60)

	minutesAngle := 360 * (float64(minutes) / 60)
	result := math.Abs(hourAngle - minutesAngle)
	if result > 180 {
		result = 360 - result
	}
	return result
}
