package entities

import (
	"github.com/marycka9/go-reverso-api/languages"
	"math/rand"
	"net/url"
	"strconv"
)

const urlContextQuery = "https://context.reverso.net/bst-query-service"
const UserAgentContextApp = "Dalvik/2.1.0 (Linux; U; Android 9; ONEPLUS A5000 Build/PKQ1.180716.001) ReversoContext"

//const UserAgentContextBrowser = "Mobile User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 17_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1.2 Mobile/15E148 Safari/604.1"

type userAgentContextBrowser []string

// TODO: from time to time update user agents

var userAgent = userAgentContextBrowser{
	"Mobile User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 17_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1.2 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:139.0) Gecko/20100101 Firefox/139.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 14.7; rv:139.0) Gecko/20100101 Firefox/139.0",
	"Mozilla/5.0 (X11; Linux i686; rv:139.0) Gecko/20100101 Firefox/139.0",
	"Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0",
	"Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:139.0) Gecko/20100101 Firefox/139.0",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0",
	"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 14_7_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.4 Safari/605.1.15",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)",
	"Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 6.0; Trident/5.0)",
	"Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
	"Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.2; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Edg/137.0.3296.52",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Edg/137.0.3296.52",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 OPR/119.0.0.0",
	"Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 OPR/119.0.0.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 14_7_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 OPR/119.0.0.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36 OPR/119.0.0.0",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Vivaldi/7.4.3684.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Vivaldi/7.4.3684.38",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 14_7_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Vivaldi/7.4.3684.38",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36 Vivaldi/7.4.3684.38",
	"Mozilla/5.0 (X11; Linux i686) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36 Vivaldi/7.4.3684.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 YaBrowser/25.4.1.1052 Yowser/2.5 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 YaBrowser/25.4.1.1052 Yowser/2.5 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/137.0.7151.51 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/137.0.7151.51 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (iPod; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/137.0.7151.51 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.127 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 14_7_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/139.0 Mobile/15E148 Safari/605.1.15",
	"Mozilla/5.0 (iPad; CPU OS 14_7_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/139.0 Mobile/15E148 Safari/605.1.15",
	"Mozilla/5.0 (iPod touch; CPU iPhone OS 14_7_6 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/139.0 Mobile/15E148 Safari/605.1.15",
	"Mozilla/5.0 (Android 15; Mobile; rv:139.0) Gecko/139.0 Firefox/139.0",
	"Mozilla/5.0 (Android 15; Mobile; LG-M255; rv:139.0) Gecko/139.0 Firefox/139.0",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 17_7_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.4 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (iPad; CPU OS 17_7_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.4 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (iPod touch; CPU iPhone 17_7_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.4 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.127 Mobile Safari/537.36 EdgA/136.0.3240.91",
	"Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.127 Mobile Safari/537.36 EdgA/136.0.3240.91",
	"Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.127 Mobile Safari/537.36 EdgA/136.0.3240.91",
	"Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.127 Mobile Safari/537.36 EdgA/136.0.3240.91",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 17_7_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 EdgiOS/136.3240.91 Mobile/15E148 Safari/605.1.15",
	"Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.127 Mobile Safari/537.36 OPR/76.2.4027.73374",
	"Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.127 Mobile Safari/537.36 OPR/76.2.4027.73374",
	"Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.127 Mobile Safari/537.36 OPR/76.2.4027.73374",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 17_7_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.4 YaBrowser/25.4.6.729 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (iPad; CPU OS 17_7_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.4 YaBrowser/25.4.6.729 Mobile/15E148 Safari/605.1",
	"Mozilla/5.0 (iPod touch; CPU iPhone 17_7_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.4 YaBrowser/25.4.6.729 Mobile/15E148 Safari/605.1",
	"Mozilla/5.0 (Linux; arm_64; Android 15; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.7103.127 YaBrowser/25.4.4.105 Mobile Safari/537.36",
}

func GetUserAgentContextBrowser() string {
	return userAgent[rand.Intn(len(userAgent))]
}

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
