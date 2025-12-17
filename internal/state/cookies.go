package state

import (
	"os"

	"github.com/go-rod/rod"
)

func SaveCookies(page *rod.Page, path string) error {
	cookies, err := page.Cookies([]string{})
	if err != nil {
		return err
	}
	return os.WriteFile(path, rod.Marshal(cookies), 0644)
}

func LoadCookies(page *rod.Page, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return page.SetCookies(rod.MustUnmarshalCookies(data))
}
