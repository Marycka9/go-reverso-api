package entities

import (
	"fmt"
	"net/url"
)

const endpointSynonymsAutoComplete = "autocomplete/"

type AutoCompleteRequest struct{}

type AutoCompleteResponse []string

func NewAutoCompleteRequest() *AutoCompleteRequest {
	return &AutoCompleteRequest{}
}

func (s *AutoCompleteRequest) GetParams(text string) url.Values {
	return url.Values{
		"term":       []string{text},
		"rude":       []string{"true"},
		"colloquial": []string{"true"},
	}
}

func (s AutoCompleteRequest) GetUrl(code, text string) string {
	urlSynonym := fmt.Sprintf("%s%s%s/%s", urlSynonyms, endpointSynonymsAutoComplete, code, url.PathEscape(text))

	base, err := url.Parse(urlSynonym)
	if err != nil {
		return ""
	}

	base.RawQuery = s.GetParams(text).Encode()

	return base.String()
}
