package scheduler

import (
	"sort"
	"time"
)

type Shift struct {
	Date     time.Time
	Slot     string
	Required int
	Assigned []*Worker
}

type Worker struct {
	ID           string
	Name         string
	WorkedShifts int
	MaxShifts    int
	Availability map[string][]string
}

func AssignShifts(shifts []*Shift, workers []*Worker) {
	for _, shift := range shifts {
		dateStr := shift.Date.Format("2006-01-02")
		available := []*Worker{}
		for _, w := range workers {
			if w.WorkedShifts >= w.MaxShifts {
				continue
			}
			if slots, ok := w.Availability[dateStr]; ok {
				for _, s := range slots {
					if s == shift.Slot {
						available = append(available, w)
						break
					}
				}
			}
		}

		sort.SliceStable(available, func(i, j int) bool {
			return available[i].WorkedShifts < available[j].WorkedShifts
		})

		for i := 0; i < shift.Required && i < len(available); i++ {
			w := available[i]
			shift.Assigned = append(shift.Assigned, w)
			w.WorkedShifts++
		}
	}
}
