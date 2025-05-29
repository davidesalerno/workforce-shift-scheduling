package api

type ShiftRequest struct {
    Date     string `json:"date"`     // "YYYY-MM-DD"
    Slot     string `json:"slot"`     // "morning"/"evening"
    Required int    `json:"required"` // numero di persone
}

type WorkerRequest struct {
    ID            string              `json:"id"`
    Name          string              `json:"name"`
    MaxShifts     int                 `json:"max_shifts"`     // max turni consentiti
    WorkedShifts  int                 `json:"worked_shifts"`  // già lavorati
    Availability  map[string][]string `json:"availability"`   // "2025-07-10": ["morning"]
}

type ScheduleRequest struct {
    Shifts  []ShiftRequest  `json:"shifts"`
    Workers []WorkerRequest `json:"workers"`
}

type AssignedWorker struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

type ShiftResponse struct {
    Date     string           `json:"date"`
    Slot     string           `json:"slot"`
    Assigned []AssignedWorker `json:"assigned"`
}