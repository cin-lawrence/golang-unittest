package planner

import (
	"context"
	"testing"
	"time"

	"ch24/fakes/domain/recipe"
	"ch24/fakes/internal/expect"
	"ch24/fakes/internal/plannertest"
)

type Calendar interface {
	ScheduleMeal(context.Context, recipe.Recipe, time.Time) error
	GetSchedule(context.Context) (map[time.Time]recipe.Recipes, error)
}

type CalendarContract struct {
	NewCalendar func() Calendar
}

func (c *CalendarContract) Test(t *testing.T) {
	t.Run("it returns what is put in", func(t *testing.T) {
		var (
			ctx         = context.Background()
			someRecipes = plannertest.RandomRecipes()
			tomorrow    = time.Now()
			sut         = c.NewCalendar()
		)

		for _, r := range someRecipes {
			expect.NoErr(t, sut.ScheduleMeal(ctx, r, tomorrow))
		}
		got, err := sut.GetSchedule(ctx)

		expect.NoErr(t, err)
		expect.DeepEqual(t, got[tomorrow], someRecipes)
	})
}
