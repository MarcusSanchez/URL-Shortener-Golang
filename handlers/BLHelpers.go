package handlers

import (
	"math/rand"
	"net/url"
	"time"
	"urlShortener/database"
)

const BASE62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortenedUrl() string {
	return database.Route + "/s/" + encodeBase62(generateRandomInt())
}

func generateRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10000) + 1
	return randomNumber
}

func encodeBase62(number int) string {
	if number == 0 {
		return string(BASE62[0])
	}
	base62 := ""
	for number > 0 {
		number = number / 62
		remainder := number % 62
		base62 = string(BASE62[remainder]) + base62
	}
	return base62
}

func verifyUrl(urlString string) error {
	// Parse the URL
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return err
	}
	// Check if the URL is valid
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return err
	}
	return nil
}
