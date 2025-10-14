package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

type Point struct {
	X float64
	Y float64
}

func NewPointFromAngle(angle float64) Point {
	return Point{
		math.Sin(angle),
		math.Cos(angle),
	}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / secondsInHalfClock * float64(t.Second())
}

func minutesInRadians(t time.Time) float64 {
	return math.Pi/minutesInHalfClock*float64(t.Minute()) + secondsInRadians(t)/minutesInClock
}

func hoursInRadians(t time.Time) float64 {
	return math.Pi/hoursInHalfClock*float64(t.Hour()%hoursInClock) +
		minutesInRadians(t)/hoursInClock
}

func secondHandPoint(t time.Time) Point {
	return NewPointFromAngle(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return NewPointFromAngle(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return NewPointFromAngle(hoursInRadians(t))
}
