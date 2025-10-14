package clockface

import (
	"math"
	"testing"
	"time"
)

const equalityThreshold = 1e-7

func roughlyEqualFloat64(a, b float64) bool {
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2027, time.October, 24, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), math.Pi / 2 * 3},
		{simpleTime(0, 0, 7), math.Pi / 30 * 7},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := secondsInRadians(tc.time)
			if !roughlyEqualFloat64(got, tc.angle) {
				t.Fatalf("Wanted %v radians, but got %v", tc.angle, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := secondHandPoint(tc.time)
			if !roughlyEqualPoint(got, tc.point) {
				t.Fatalf("Wanted %v Point, but got %v", tc.point, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 7, 0), math.Pi / 30 * 7},
		{simpleTime(0, 0, 7), (math.Pi / (30 * 60)) * 7},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := minutesInRadians(tc.time)
			if !roughlyEqualFloat64(got, tc.angle) {
				t.Fatalf("Wanted %v radians, but got %v", tc.angle, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 15, 0), Point{1, 0}},
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := minuteHandPoint(tc.time)
			if !roughlyEqualPoint(got, tc.point) {
				t.Fatalf("Wanted %v Point, but got %v", tc.point, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi / 2 * 3},
		{simpleTime(0, 1, 30), math.Pi / (6 * 60 * 60) * 90},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := hoursInRadians(tc.time)
			if !roughlyEqualFloat64(got, tc.angle) {
				t.Fatalf("Wanted %v radians, but got %v", tc.angle, got)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(3, 0, 0), Point{1, 0}},
		{simpleTime(18, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := hourHandPoint(tc.time)
			if !roughlyEqualPoint(got, tc.point) {
				t.Fatalf("Wanted %v Point, but got %v", tc.point, got)
			}
		})
	}
}
