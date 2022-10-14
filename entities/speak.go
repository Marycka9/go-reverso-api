package entities

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type SpeakRequest struct {
	FileName   string
	FilePath   string
	Text       string
	Voice      string
	Mp3BitRate int
	VoiceSpeed int
}

type SpeakResponse struct {
}

const UrlSpeak = "https://voice.reverso.net/RestPronunciation.svc/v1/output=json/GetVoiceStream/voiceName=%s"

func NewSpeakRequest(fileName, filePath, text, voice string, mp3BitRate, voiceSpeed int) (*SpeakRequest, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	splitter := "/"
	if filePath[:1] == "/" {
		splitter = ""
	}

	filePath = fmt.Sprintf("%s%s%s", path, splitter, filePath)

	return &SpeakRequest{
		FileName:   fileName,
		FilePath:   filePath,
		Text:       text,
		Voice:      voice,
		Mp3BitRate: mp3BitRate,
		VoiceSpeed: voiceSpeed,
	}, nil
}

func (s *SpeakRequest) GetPath() string {
	path := s.FilePath
	if len(path) > 0 {
		b := path[len(path)-1:] == "/"
		if !b {
			path += "/"
		}
	}

	path += fmt.Sprintf("%s.mp3", s.FileName)

	return path
}

func (s SpeakRequest) GetParams() url.Values {
	return url.Values{
		"inputText":  []string{base64.URLEncoding.EncodeToString([]byte(s.Text))},
		"mp3BitRate": []string{strconv.Itoa(s.Mp3BitRate)},
		"voiceSpeed": []string{strconv.Itoa(s.VoiceSpeed)},
	}
}

func (s *SpeakRequest) GetUrl(voice string) string {
	base, err := url.Parse(fmt.Sprintf(UrlSpeak, voice))
	if err != nil {
		return ""
	}

	base.RawQuery = s.GetParams().Encode()

	return base.String()
}
