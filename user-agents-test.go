package main

import (
	"github.com/marycka9/go-reverso-api/client"
	"github.com/marycka9/go-reverso-api/entities"
	"github.com/marycka9/go-reverso-api/languages"
	"log"
	"math/rand"
	"time"
)

const trying = 10

var wordsForTranslate = []string{
	"as American as apple pie",
	"boozehound",
	"apples and oranges",
	"guestimate",
	"albatross around one's neck",
	"time and tide wait for no man",
	"ditz",
	"hot air",
	"wannabe",
	"the thought that counts",
	"washed-up",
	"put on an act",
	"put off",
	"hostess with the mostest",
	"there is no accounting for taste",
}

func getWord() string {
	return wordsForTranslate[rand.Intn(len(wordsForTranslate))]
}

func userAgentTests() {
	availableUserAgentsCount := entities.GetUserAgentsLen()

	c := client.NewClient()
	langs := languages.GetLanguages()

	log.Println("Available User Agents Count: ", availableUserAgentsCount)

	for i := 0; i < availableUserAgentsCount; i++ {
		for j := 0; j < len(wordsForTranslate); j++ {
			log.Println("Trying: ", i, j)
			word := wordsForTranslate[j]

			_, err := c.Translate(word, langs["english"], langs["russian"])
			if err != nil {
				log.Println(err)
			}

			time.Sleep(time.Second * 2)
		}
	}

	c.Close()
}
