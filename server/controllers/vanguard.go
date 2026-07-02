package controllers

import (
	"ko-commands/logic"
	"net/http"
)

func VanguardLoginQR(w http.ResponseWriter, r *http.Request) {
	screenshot, err := logic.FetchVanguardLogin()
	if err != nil {
		http.Error(w, "Failed to fetch screenshot", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", string(rune(len(screenshot))))
	w.Write(screenshot)
}
