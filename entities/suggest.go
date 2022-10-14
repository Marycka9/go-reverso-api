package entities

import (
	"encoding/json"
	"github.com/Marycka9/go-reverso-translate-api/languages"
	"net/url"
)

const urlContextSuggest = "https://context.reverso.net/bst-suggest-service"

type SuggestRequest struct {
	Search     string `json:"search"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
	MaxResults int64  `json:"max_results"`
	Mode       int64  `json:"mode"`
}

type SuggestResponse struct {
	Suggestions []FuzzySuggestion `json:"suggestions"`
	Fuzzy1      []FuzzySuggestion `json:"fuzzy1"`
	Fuzzy2      []FuzzySuggestion `json:"fuzzy2"`
	Request     SuggestRequest    `json:"request"`
	TimeMS      int64             `json:"time_ms"`
}

func NewSuggestRequest(text string, srcLang, dstLang *languages.Language) *SuggestRequest {
	return &SuggestRequest{
		Search:     text,
		SourceLang: srcLang.Code,
		TargetLang: dstLang.Code,
		MaxResults: 0,
		Mode:       0,
	}
}

func (s *SuggestRequest) GetUrl() string {
	base, err := url.Parse(urlContextSuggest)
	if err != nil {
		return ""
	}

	base.RawQuery = s.GetParams().Encode()

	return base.String()
}

func (s *SuggestRequest) GetParams() url.Values {
	return url.Values{
		"search":      []string{s.Search},
		"source_lang": []string{s.SourceLang},
		"target_lang": []string{s.TargetLang},
	}
}

func (s *SuggestRequest) MarshalJson() (string, error) {
	res, err := json.Marshal(&s)
	return string(res), err
}
