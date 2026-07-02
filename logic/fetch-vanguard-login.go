package logic

import (
	"github.com/go-rod/rod"
)

func FetchVanguardLogin() ([]byte, error) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://logon.vanguard.com/logon?site=pi")
	screenshot := page.MustScreenshot()

	return screenshot, nil
}
