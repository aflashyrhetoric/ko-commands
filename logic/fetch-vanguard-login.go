package logic

import (
	"time"

	"github.com/go-rod/rod"
)

func FetchVanguardLogin() ([]byte, error) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://logon.vanguard.com/logon?site=pi")
	page.MustWaitIdle()
	time.Sleep(3 * time.Second) // Wait for the page to load completely
	screenshot := page.MustScreenshot()

	return screenshot, nil
}
