package main

import (
	"github.com/marycka9/go-reverso-api/client"
	"github.com/marycka9/go-reverso-api/languages"
)

func simpleExamples() {
	c := client.NewClient()
	langs := languages.GetLanguages()

	resTranslate, err := c.Translate("indict", langs["english"], langs["russian"])
	res1, err := c.Synonyms("Hello", langs["english"])
	res2, err := c.AutoComplete("Hello", langs["english"])
	res3, err := c.Context("sky", langs["english"], langs["french"], 1)
	res4, err := c.Suggest("sky", langs["english"], langs["french"])
	err = c.Speak("example123", "data/user1", "sky", 128, 100)

	_ = resTranslate
	_ = res1
	_ = res2
	_ = res3
	_ = res4
	_ = err
}
