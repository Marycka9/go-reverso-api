package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/marycka9/go-reverso-api/entities"
	"github.com/marycka9/go-reverso-api/languages"
	"github.com/marycka9/go-reverso-api/voices"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Client struct {
	Client *http.Client
}

func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

func (c *Client) Close() {
	c.Close()
}

func (c *Client) Translate(text string, srcLang, dstLang *languages.Language) (*entities.TranslateResponse, error) {
	translateReq := entities.NewTranslateRequest(text, srcLang, dstLang)
	requestBody, err := translateReq.MarshalJson()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		translateReq.GetUrl(),
		strings.NewReader(requestBody),
	)

	userAgent, userAgentIndex := entities.GetUserAgentContextBrowser()

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-Reverso-Origin", "context.web")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("sec-ch-ua", "\"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"138\", \"Google Chrome\";v=\"138\"")
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("host", "api.reverso.net")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Reverso returned status code: %d\nFor the next text \"%s\"\nUser-Agent-Index: %d\nUser-Agent: \"%s\"", resp.StatusCode, text, userAgentIndex, userAgent))
	}

	var translate *entities.TranslateResponse
	if err := json.NewDecoder(resp.Body).Decode(&translate); err != nil {
		return nil, err
	}

	_ = resp.Body.Close()

	return translate, nil
}

func (c *Client) Synonyms(text string, language *languages.Language) (*entities.SynonymsResponse, error) {
	synonymRequest := entities.NewSynonymRequest(text, language)

	req, err := http.NewRequest(
		http.MethodGet,
		synonymRequest.GetUrl(language.Code, text),
		nil,
	)

	userAgent, _ := entities.GetUserAgentContextBrowser()

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-Reverso-Origin", "context.web")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("sec-ch-ua", "\"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"138\", \"Google Chrome\";v=\"138\"")
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("host", "api.reverso.net")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var synonym *entities.SynonymsResponse
	if err := json.NewDecoder(resp.Body).Decode(&synonym); err != nil {
		return nil, err
	}

	_ = resp.Body.Close()

	return synonym, nil
}

func (c *Client) AutoComplete(text string, language *languages.Language) (*entities.AutoCompleteResponse, error) {
	autoCompleteRequest := entities.NewAutoCompleteRequest()

	req, err := http.NewRequest(
		http.MethodGet,
		autoCompleteRequest.GetUrl(language.Code, text),
		nil,
	)

	userAgent, _ := entities.GetUserAgentContextBrowser()

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("x-reverso-origin", "synonymapp")
	req.Header.Add("x-reverso-ui-lang", "en")
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", entities.BearerSynonyms))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	autocomplete := make(entities.AutoCompleteResponse, 0)
	if err := json.NewDecoder(resp.Body).Decode(&autocomplete); err != nil {
		return nil, err
	}

	_ = resp.Body.Close()

	return &autocomplete, nil
}

func (c *Client) Context(text string, srcLang, dstLang *languages.Language, page int) (*entities.ContextResponse, error) {
	queryReq := entities.NewContextRequest(text, srcLang, dstLang, page)

	req, err := http.NewRequest(
		http.MethodPost,
		queryReq.GetUrl(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	userAgent, _ := entities.GetUserAgentContextBrowser()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("User-Agent", userAgent)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var query *entities.ContextResponse
	if err := json.NewDecoder(resp.Body).Decode(&query); err != nil {
		return nil, err
	}

	_ = resp.Body.Close()

	return query, nil
}

func (c *Client) Suggest(text string, srcLang, dstLang *languages.Language) (*entities.SuggestResponse, error) {
	suggestReq := entities.NewSuggestRequest(text, srcLang, dstLang)

	req, err := http.NewRequest(
		http.MethodGet,
		suggestReq.GetUrl(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	userAgent, _ := entities.GetUserAgentContextBrowser()

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("User-Agent", userAgent)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var query *entities.SuggestResponse
	if err := json.NewDecoder(resp.Body).Decode(&query); err != nil {
		return nil, err
	}

	_ = resp.Body.Close()

	return query, nil
}

func (c *Client) Speak(fileName, filePath, text string, mp3BitRate, voiceSpeed int) error {
	speakRequest, err := entities.NewSpeakRequest(fileName, filePath, text, voices.VoiceEnglishFemale, mp3BitRate, voiceSpeed)
	if err != nil {
		return err
	}

	if _, err := os.Stat(speakRequest.FilePath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(speakRequest.FilePath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	fileOut, err := os.OpenFile(speakRequest.GetPath(), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodGet,
		speakRequest.GetUrl(voices.VoiceEnglishFemale),
		nil,
	)
	if err != nil {
		return err
	}

	userAgent, _ := entities.GetUserAgentContextBrowser()

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("User-Agent", userAgent)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	var buffer *bytes.Buffer
	var body []byte

	buffer = bytes.NewBuffer(body)

	if resp.ContentLength == -1 {
		_, err = buffer.ReadFrom(resp.Body)
	} else {
		body = make([]byte, resp.ContentLength)
		_, err = io.Copy(buffer, resp.Body)
	}

	body = buffer.Bytes()
	_ = resp.Body.Close()

	if _, err = fileOut.Write(body); err != nil {
		return err
	}

	_ = fileOut.Close()

	return nil
}
