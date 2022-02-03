package simulator

import (
	"math/rand"
	"time"
)

const (
	COOKIE_SID = "connect.sid"
)

var sidChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randSid returns a random 25 character session identifier.
func randSid() string {
	b := make([]rune, 25)
	for i := range b {
		b[i] = sidChars[rand.Intn(len(sidChars))]
	}
	return string(b)
}
