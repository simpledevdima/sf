// Package sf is a kit of simple functions for often used actions
package sf

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

// GetHash make and return random hash code
func GetHash() string {
	rand.Seed(time.Now().UnixNano())
	data := []byte(time.Now().String() + string(rune(rand.Intn(99999))))
	hash := fmt.Sprintf("%x", md5.Sum(data))
	return hash
}

// Contains search string in slice of strings and return true if found or false if not found
func Contains(elements []string, str string) bool {
	for _, element := range elements {
		if element == str {
			return true
		}
	}
	return false
}
