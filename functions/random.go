package functions

import (
	"fmt"
	"math/rand"
	"time"
)

// BreakTime is to wait for a ramdom time
func BreakTime() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	waitTime := rand.Intn(20) + 10
	fmt.Printf("Waiting for %v seconds...\n", waitTime)
	time.Sleep(time.Duration(waitTime) * time.Second)
}

// GenerateRamdomString is to generate a ramdom string
func GenerateRamdomString(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randString := make([]rune, length)
	for i := range randString {
		randString[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(randString)
}
