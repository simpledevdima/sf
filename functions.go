// Package sf is a kit of simple functions for often used actions
package sf

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

// GetHash generates and returns a random hash code of the specified length
// it is also possible to add salt as second argument
func GetHash(length int, salt ...string) string {
	rand.Seed(time.Now().UnixNano())
	data := make([][]byte, 3)
	data[0] = []byte(time.Now().String())
	data[1] = []byte(string(rune(rand.Intn(99999))))
	if len(salt) > 0 {
		data[2] = []byte(salt[0])
	}
	full := bytes.Join(data, []byte(""))
	hash := fmt.Sprintf("%x", md5.Sum(full))
	b := make([]byte, length+2-len(hash))
	rand.Read(b)
	add := fmt.Sprintf("%x", b)[2 : length+2-len(hash)]
	return hash + add
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
