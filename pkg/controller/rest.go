package controller

import (
	"net/http"
	"time"

	api "github.com/davidesalerno/workforce-shift-scheduling/pkg/api"
	scheduler "github.com/davidesalerno/workforce-shift-scheduling/pkg/scheduler"
	"github.com/gin-gonic/gin"
)

func RestScheduleHandler(c *gin.Context) {
	var req api.ScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Converte input in strutture interne
	shifts := []*scheduler.Shift{}
	for _, s := range req.Shifts {
		date, err := time.Parse("2006-01-02", s.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato data non valido"})
			return
		}
		shifts = append(shifts, &scheduler.Shift{
			Date:     date,
			Slot:     s.Slot,
			Required: s.Required,
		})
	}

	workers := []*scheduler.Worker{}
	for _, w := range req.Workers {
		workers = append(workers, &scheduler.Worker{
			ID:           w.ID,
			Name:         w.Name,
			WorkedShifts: w.WorkedShifts,
			MaxShifts:    w.MaxShifts,
			Availability: w.Availability,
		})
	}

	scheduler.AssignShifts(shifts, workers)

	// Costruisce la risposta
	var resp []api.ShiftResponse
	for _, shift := range shifts {
		sr := api.ShiftResponse{
			Date: shift.Date.Format("2006-01-02"),
			Slot: shift.Slot,
		}
		for _, w := range shift.Assigned {
			sr.Assigned = append(sr.Assigned, api.AssignedWorker{ID: w.ID, Name: w.Name})
		}
		resp = append(resp, sr)
	}

	c.JSON(http.StatusOK, resp)
}
