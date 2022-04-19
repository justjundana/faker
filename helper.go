package faker

import (
	"math/rand"
	"strings"
	"time"
)

// Random Pick Data
func randomPick(s []string) string {
	rand.Seed(time.Now().UnixNano())
	return s[rand.Intn(len(s))]
}

// toProperCase Returns a string s with an upper-case on the first word.
func toProperCase(s string) string {
	return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
}
