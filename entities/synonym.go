package entities

import (
	"encoding/json"
	"fmt"
	"github.com/Marycka9/go-reverso-translate-api/languages"
	"net/url"
)

const urlSynonyms = "https://synonyms.reverso.net/api/v2/"

const endpointSynonymsSearch = "search/"

const BearerSynonyms = "c3lub255bXM6REtiZTUyRjNZRExZdVFFOHlk" // synonyms:DKbe52F3YDLYuQE8yd  // probably just an identifier

type SynonymRequest struct {
	Input string `json:"input"`
	Lang  string `json:"lang"`
}

type Pos struct {
	Mask int64    `json:"mask"`
	Desc []string `json:"desc"`
}

type Cluster struct {
	ID           int64   `json:"id"`
	Word         string  `json:"word"`
	Language     string  `json:"language"`
	Cluster      int64   `json:"cluster"`
	Weight       int64   `json:"weight"`
	Nrows        int64   `json:"nrows"`
	Isentry      bool    `json:"isentry"`
	Pos          Pos     `json:"pos"`
	Rude         bool    `json:"rude"`
	Colloquial   bool    `json:"colloquial"`
	Relevance    float64 `json:"relevance"`
	MostRelevant bool    `json:"mostRelevant"`
}

type Example struct {
	ID      int64  `json:"id"`
	Cluster int64  `json:"cluster"`
	Example string `json:"example"`
	Pos     Pos    `json:"pos"`
}

type Stopwatch struct {
	Start   float64 `json:"start"`
	Ended   float64 `json:"ended"`
	Elapsed float64 `json:"elapsed"`
}

type SynonymResult struct {
	Pos               Pos           `json:"pos"`
	Weight            int64         `json:"weight"`
	Nrows             int64         `json:"nrows"`
	Relevance         int64         `json:"relevance"`
	RudeResults       int64         `json:"rudeResults"`
	ColloquialResults int64         `json:"colloquialResults"`
	Merged            interface{}   `json:"merged"`
	Relevantitems     int64         `json:"relevantitems"`
	Cluster           []Cluster     `json:"cluster"`
	Examples          []Example     `json:"examples"`
	Antonyms          []interface{} `json:"antonyms"`
}

type Related struct {
	ID         int64  `json:"id"`
	Word       string `json:"word"`
	Pos        Pos    `json:"pos"`
	Language   string `json:"language"`
	Rude       bool   `json:"rude"`
	Colloquial bool   `json:"colloquial"`
}

type SynonymsResponse struct {
	ID                  int64           `json:"id"`
	Search              string          `json:"search"`
	Language            string          `json:"language"`
	Fuzzyhash           string          `json:"fuzzyhash"`
	Input               string          `json:"input"`
	Pos                 Pos             `json:"pos"`
	SearchType          string          `json:"searchType"`
	Allowredirect       bool            `json:"allowredirect"`
	Rude                bool            `json:"rude"`
	Colloquial          bool            `json:"colloquial"`
	PotentialRude       int64           `json:"potentialRude"`
	PotentialColloquial int64           `json:"potentialColloquial"`
	Groupable           bool            `json:"groupable"`
	ResultsCount        int64           `json:"resultsCount"`
	Results             []SynonymResult `json:"results"`
	Suggestions         []interface{}   `json:"suggestions"`
	Antonyms            []interface{}   `json:"antonyms"`
	Related             []Related       `json:"related"`
	Stopwatch           Stopwatch       `json:"stopwatch"`
}

func NewSynonymRequest(text string, language *languages.Language) *SynonymRequest {
	return &SynonymRequest{
		Input: text,
		Lang:  language.Alpha3,
	}
}

func (s *SynonymRequest) GetParams() url.Values {
	return url.Values{
		"limit": []string{"50"},
		"rude":  []string{"true"},
		"abc":   []string{"false"},
		"merge": []string{"true"},
	}
}

func (s SynonymRequest) GetUrl(code, text string) string {
	urlSynonym := fmt.Sprintf("%s%s%s/%s", urlSynonyms, endpointSynonymsSearch, code, url.PathEscape(text))

	base, err := url.Parse(urlSynonym)
	if err != nil {
		return ""
	}

	base.RawQuery = s.GetParams().Encode()

	return base.String()
}

func (s *SynonymRequest) MarshalJson() (string, error) {
	res, err := json.Marshal(&s)
	return string(res), err
}
