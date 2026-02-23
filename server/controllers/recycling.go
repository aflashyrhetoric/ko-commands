package controllers

import (
	"encoding/json"
	"net/http"

	"ko-commands/logic"
)

type recyclingResponse struct {
	Type string `json:"type"`
	In   int    `json:"in_days"`
}

func Recycling(w http.ResponseWriter, r *http.Request) {
	currentDay := logic.GetDayOfTheYear()
	days := logic.GetRecyclingDays()

	nextPaperDay := logic.GetNextRecyclingDay(currentDay, days["paper"])
	nextGlassDay := logic.GetNextRecyclingDay(currentDay, days["glass"])

	w.Header().Set("Content-Type", "application/json")

	if nextPaperDay != -1 && (nextGlassDay == -1 || nextPaperDay < nextGlassDay) {
		json.NewEncoder(w).Encode(recyclingResponse{Type: "paper", In: nextPaperDay - currentDay})
		return
	}
	if nextGlassDay != -1 && (nextPaperDay == -1 || nextGlassDay < nextPaperDay) {
		json.NewEncoder(w).Encode(recyclingResponse{Type: "glass/metal/plastic", In: nextGlassDay - currentDay})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "No more recycling days this year."})
}
