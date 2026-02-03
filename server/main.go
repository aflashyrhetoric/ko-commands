package main

import (
	"encoding/json"
	"log"
	"net/http"

	"ko-commands/logic"
)

func main() {
	http.HandleFunc("/recycling", func(w http.ResponseWriter, r *http.Request) {
		currentDay := logic.GetDayOfTheYear()
		days := logic.GetRecyclingDays()

		nextPaperDay := logic.GetNextRecyclingDay(currentDay, days["paper"])
		nextGlassDay := logic.GetNextRecyclingDay(currentDay, days["glass"])

		type recyclingResponse struct {
			Type string `json:"type"`
			In   int    `json:"in_days"`
		}

		if nextPaperDay != -1 && (nextGlassDay == -1 || nextPaperDay < nextGlassDay) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(recyclingResponse{
				Type: "paper",
				In:   nextPaperDay - currentDay,
			})
			return
		}
		if nextGlassDay != -1 && (nextPaperDay == -1 || nextGlassDay < nextPaperDay) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(recyclingResponse{
				Type: "glass/metal/plastic",
				In:   nextGlassDay - currentDay,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "No more recycling days this year.",
		})
	})

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
