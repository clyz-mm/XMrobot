package instructions

import (
	"XMrobot/dict"
	"math/rand"
	"time"
)

var jokeList = dict.GetJokeAll()

func GetOneJoke() string {
	rand.Seed(time.Now().UnixNano())
	return jokeList[rand.Intn(len(jokeList))][1]
}
