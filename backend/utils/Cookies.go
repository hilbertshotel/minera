package utils

import (
	"math/rand"
    "time"
	"net/http"
)

const CookieName = "Rumi"

func MakeCookie(name string, value string) http.Cookie {
	cookie := http.Cookie {
		Name: name,
		Value: value,
	}
	return cookie
}

var pool = "1234567890abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ:|?$%@][{}#&/()*-_"

func GenerateCookie() (string, string) {
	rand.Seed(time.Now().UnixNano())
	value := make([]byte, 20)

    for i := 0; i < 20; i++ {
        value[i] = pool[rand.Intn(len(pool))]
    }

	return CookieName, string(value)
}
