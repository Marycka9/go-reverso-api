package languages

import (
	_ "embed"
	"encoding/json"
)

type Language struct {
	Code   string `json:"code"`
	Alpha3 string `json:"alpha3"`
}

//go:embed languages.json
var langData []byte

type Languages map[string]*Language

func GetLanguages() Languages {
	result := make(map[string]*Language, 0)
	_ = json.Unmarshal(langData, &result)
	return result
}
