package regex

import (
	"regexp"
)

var (
	UrlRegex     = `^(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`
	SoftUrlRegex = `https?://[^\s]+`
	ImageRegex   = `[^\s]+\.(?:jpg|jpeg|png|gif|bmp|webp)`
)

func UrlIsValid(url string) bool {
	urlRegex := regexp.MustCompile(UrlRegex)
	match := urlRegex.MatchString(url)
	return match
}

func CountUrls(body string) int {
	urlRegex := regexp.MustCompile(SoftUrlRegex)
	matches := urlRegex.FindAllString(body, -1)
	return len(matches)
}

func CountImages(body string) int {
	imageRegex := regexp.MustCompile(ImageRegex)
	matches2 := imageRegex.FindAllString(body, -1)
	return len(matches2)
}
