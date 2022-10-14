package main

import (
	"github.com/Marycka9/go-reverso-translate-api/client"
	"github.com/Marycka9/go-reverso-translate-api/languages"
)

func main() {
	client := client.NewClient()
	langs := languages.GetLanguages()

	res, err := client.Translate("Hello", langs["english"], langs["russian"])
	res1, err := client.Synonyms("Hello", langs["english"])
	res2, err := client.AutoComplete("Hello", langs["english"])
	res3, err := client.Context("sky", langs["english"], langs["french"], 1)
	res4, err := client.Suggest("sky", langs["english"], langs["french"])
	err = client.Speak("example123", "data/user1", "sky", 128, 100)

	_ = res
	_ = res1
	_ = res2
	_ = res3
	_ = res4
	_ = err
}
