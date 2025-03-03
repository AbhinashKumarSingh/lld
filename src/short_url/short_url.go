package short_url

import (
	"fmt"
	"strings"
	"sync"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var store = make(map[string]string) // Maps short codes to long URLs
var mutex = sync.RWMutex{}
var lastProcessedID int64

func generateShortURL() string {
	mutex.Lock()
	lastProcessedID++
	id := lastProcessedID
	mutex.Unlock()
	return encodeBase62(id)

}

func encodeBase62(id int64) string {
	if id == 0 {
		return "0"
	}

	var encodeBase62 strings.Builder
	for id > 0 {
		rem := id % 62
		encodeBase62.WriteByte(base62Chars[rem])
		id = id / 62
	}
	return reverseString(encodeBase62.String())

}

// Reverse string utility
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Init() {
	mutex.Lock()
	lastProcessedID = 35000000000000 // Limit ID size
	mutex.Unlock()

	// Store mapping
	for i := 1000000000000000; i < 10000000000000000; i++ {
		shortCode := generateShortURL()
		mutex.Lock()
		store[shortCode] = fmt.Sprintf("for short code:%s long url is %d", shortCode, i)
		fmt.Println(store[shortCode])
		mutex.Unlock()
	}

}
