package entities

import (
	"github.com/Marycka9/go-reverso-translate-api/languages"
	"net/url"
	"strconv"
)

const urlContextQuery = "https://context.reverso.net/bst-query-service"
const UserAgentContextApp = "Dalvik/2.1.0 (Linux; U; Android 9; ONEPLUS A5000 Build/PKQ1.180716.001) ReversoContext"

type ContextRequest struct {
	SourceText     string `json:"source_text"`
	TargetText     string `json:"target_text"`
	SourceLang     string `json:"source_lang"`
	TargetLang     string `json:"target_lang"`
	Corpus         string `json:"corpus"`
	NonadaptedText string `json:"nonadapted_text"`
	SourcePos      string `json:"source_pos"`
	Nrows          int    `json:"nrows"`
	Npage          int    `json:"npage"`
	Mode           int    `json:"mode"`
	ExprSug        int    `json:"expr_sug"`
	PosReorder     int    `json:"pos_reorder"`
	Device         int    `json:"device"`
	Adapted        bool   `json:"adapted"`
	RudeWords      bool   `json:"rude_words"`
	Colloquialisms bool   `json:"colloquialisms"`
	RiskyWords     bool   `json:"risky_words"`
	DymApply       bool   `json:"dym_apply"`
	SplitLong      bool   `json:"split_long"`
	HasLocd        bool   `json:"has_locd"`
}

func NewContextRequest(text string, srcLang, dstLang *languages.Language, page int) *ContextRequest {
	return &ContextRequest{
		SourceText: text,
		SourceLang: srcLang.Code,
		TargetLang: dstLang.Code,
		Npage:      page,
		Nrows:      4,
		ExprSug:    1,
		DymApply:   true,
		PosReorder: 5,
	}
}

type SearchResult struct {
	SText string `json:"s_text"`
	TText string `json:"t_text"`
	Ref   string `json:"ref"`
	Cname string `json:"cname"`
	URL   string `json:"url"`
	Ctags string `json:"ctags"`
	Pba   bool   `json:"pba"`
}

type DictionaryEntryList struct {
	Frequency        int64                 `json:"frequency"`
	Term             string                `json:"term"`
	IsFromDict       bool                  `json:"isFromDict"`
	IsPrecomputed    bool                  `json:"isPrecomputed"`
	Stags            []string              `json:"stags"`
	Pos              *string               `json:"pos"`
	Sourcepos        []string              `json:"sourcepos"`
	Variant          interface{}           `json:"variant"`
	Domain           interface{}           `json:"domain"`
	Definition       *string               `json:"definition"`
	Vowels2          interface{}           `json:"vowels2"`
	Transliteration2 interface{}           `json:"transliteration2"`
	AlignFreq        int64                 `json:"alignFreq"`
	ReverseValidated bool                  `json:"reverseValidated"`
	PosGroup         int64                 `json:"pos_group"`
	IsTranslation    bool                  `json:"isTranslation"`
	IsFromLOCD       bool                  `json:"isFromLOCD"`
	InflectedForms   []DictionaryEntryList `json:"inflectedForms"`
}

type FuzzySuggestion struct {
	Lang       string `json:"lang"`
	Suggestion string `json:"suggestion"`
	Weight     int64  `json:"weight"`
	IsFromDict bool   `json:"isFromDict"`
}

type ContextResponse struct {
	List                     []SearchResult        `json:"list"`
	Nrows                    int64                 `json:"nrows"`
	NrowsExact               int64                 `json:"nrows_exact"`
	Pagesize                 int64                 `json:"pagesize"`
	Npages                   int64                 `json:"npages"`
	Page                     int64                 `json:"page"`
	RemovedSuperstrings      []string              `json:"removed_superstrings"`
	DictionaryEntryList      []DictionaryEntryList `json:"dictionary_entry_list"`
	DictionaryOtherFrequency int64                 `json:"dictionary_other_frequency"`
	DictionaryNrows          int64                 `json:"dictionary_nrows"`
	TimeMS                   int64                 `json:"time_ms"`
	Request                  ContextRequest        `json:"request"`
	Suggestions              []FuzzySuggestion     `json:"suggestions"`
	DymCase                  int64                 `json:"dym_case"`
	DymList                  []interface{}         `json:"dym_list"`
	DymApplied               interface{}           `json:"dym_applied"`
	DymNonadaptedSearch      interface{}           `json:"dym_nonadapted_search"`
	DymPairApplied           interface{}           `json:"dym_pair_applied"`
	DymNonadaptedSearchPair  interface{}           `json:"dym_nonadapted_search_pair"`
	DymPair                  interface{}           `json:"dym_pair"`
	ExtractedPhrases         []interface{}         `json:"extracted_phrases"`
}

func (s *ContextRequest) GetParams() url.Values {
	return url.Values{
		"source_text": []string{s.SourceText},
		"source_lang": []string{s.SourceLang},
		"target_lang": []string{s.TargetLang},
		"npage":       []string{strconv.Itoa(s.Npage)},
		"nrows":       []string{strconv.Itoa(s.Nrows)},
		"expr_sug":    []string{strconv.Itoa(s.ExprSug)},
		"json":        []string{"1"},
		"dym_apply":   []string{"true"},
		"pos_reorder": []string{strconv.Itoa(s.PosReorder)},
	}
}

func (s ContextRequest) GetUrl() string {
	base, err := url.Parse(urlContextQuery)
	if err != nil {
		return ""
	}

	base.RawQuery = s.GetParams().Encode()

	return base.String()
}
