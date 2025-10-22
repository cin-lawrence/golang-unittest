package inmemory_test

import (
	"testing"

	"ch24/fakes/adapters/inmemory"
	"ch24/fakes/domain/planner"
)

func TestInMemoryCalendar(t *testing.T) {
	(&planner.CalendarContract{
		NewCalendar: func() planner.Calendar {
			return inmemory.NewCalendar()
		},
	}).Test(t)
}
