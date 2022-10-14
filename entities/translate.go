package entities

import (
	"encoding/json"
	"github.com/Marycka9/go-reverso-translate-api/languages"
)

type TranslateOptions struct {
	Origin            string `json:"origin"`
	SentenceSplitter  bool   `json:"sentenceSplitter"`
	ContextResults    bool   `json:"contextResults,omitempty"`
	LanguageDetection bool   `json:"languageDetection,omitempty"`
}

type TranslateRequest struct {
	Input   string           `json:"input"`
	From    string           `json:"from"`
	To      string           `json:"to"`
	Format  string           `json:"format"`
	Options TranslateOptions `json:"options"`
}

type LanguageDetection struct {
	DetectedLanguage                string `json:"detectedLanguage"`
	IsDirectionChanged              bool   `json:"isDirectionChanged"`
	OriginalDirection               string `json:"originalDirection"`
	OriginalDirectionContextMatches int64  `json:"originalDirectionContextMatches"`
	ChangedDirectionContextMatches  int64  `json:"changedDirectionContextMatches"`
}

type TranslateResult struct {
	Translation    string   `json:"translation"`
	SourceExamples []string `json:"sourceExamples"`
	TargetExamples []string `json:"targetExamples"`
	Rude           bool     `json:"rude"`
	Colloquial     bool     `json:"colloquial"`
	PartOfSpeech   string   `json:"partOfSpeech"`
}

type ContextResults struct {
	RudeWords             bool              `json:"rudeWords"`
	Colloquialisms        bool              `json:"colloquialisms"`
	RiskyWords            bool              `json:"riskyWords"`
	Results               []TranslateResult `json:"results"`
	TotalContextCallsMade int64             `json:"totalContextCallsMade"`
	TimeTakenContext      int64             `json:"timeTakenContext"`
}

type TranslateResponse struct {
	ID                string            `json:"id"`
	From              string            `json:"from"`
	To                string            `json:"to"`
	Input             []string          `json:"input"`
	CorrectedText     interface{}       `json:"correctedText"`
	Translation       []string          `json:"translation"`
	Engines           []string          `json:"engines"`
	LanguageDetection LanguageDetection `json:"languageDetection"`
	ContextResults    ContextResults    `json:"contextResults"`
	Truncated         bool              `json:"truncated"`
	TimeTaken         int64             `json:"timeTaken"`
}

const urlTranslate = "https://api.reverso.net/translate/v1/translation"

func NewTranslateRequest(text string, fromLang, toLang *languages.Language) *TranslateRequest {
	return &TranslateRequest{
		Input:  text,
		From:   fromLang.Alpha3,
		To:     toLang.Alpha3,
		Format: "text",
		Options: TranslateOptions{
			Origin:            "contextappandroid", // contextappandroid vs reversomobile
			SentenceSplitter:  false,
			ContextResults:    true,
			LanguageDetection: false,
		},
	}
}

func (t TranslateRequest) GetUrl() string {
	return urlTranslate
}

func (r *TranslateRequest) MarshalJson() (string, error) {
	res, err := json.Marshal(&r)
	return string(res), err
}
