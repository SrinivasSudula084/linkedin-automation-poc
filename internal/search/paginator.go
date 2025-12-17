package search

import (
	"github.com/go-rod/rod"
)

// GoToNextPage attempts to navigate to next search page
func GoToNextPage(page *rod.Page) bool {
	nextBtn, err := page.Element(`button[aria-label="Next"]`)
	if err != nil {
		return false
	}

	nextBtn.MustClick()
	return true
}
