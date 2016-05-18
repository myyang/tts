package tts

import (
	"bytes"
	"encoding/xml"
	"log"
)

type convertAdvancedText struct {
	XMLName    xml.Name `xml:"ConvertAdvancedText"`
	Account    string   `xml:"accountID"`
	Password   string   `xml:"password"`
	TTStext    string   `xml:"TTStext"`
	TTSSpeaker string   `xml:"TTSSpeaker"`
	Volume     string   `xml:"volume"`
	Speed      string   `xml:"speed"`
	OutType    string   `xml:"outType"`
	PitchLevel string   `xml:"PitchLevel"`
	PitchSign  string   `xml:"PitchSign"`
	PitchScale string   `xml:"PitchScale"`
}

func (c *convertAdvancedText) Marshal() []byte {
	m, err := xml.Marshal(c)
	if err != nil {
		log.Fatalf("Fail to marshal: %v\n", c)
	}
	return m
}

func (c *convertAdvancedText) getAccount() string {
	return c.Account
}

func (c *convertAdvancedText) getPassword() string {
	return c.Password
}

func (c *convertAdvancedText) ParseResult(buf *bytes.Buffer) string {
	r := &advancedTextResult{}
	err := xml.Unmarshal(buf.Bytes(), r)
	if err != nil {
		log.Fatalf("Unmarshal error: %v\n", err)
	}

	s, err := matchConvertResult(r)
	if err != nil {
		log.Fatalf("Can't match result: %v\n", err)
	}
	return s
}

type advancedTextResult struct {
	XMLName xml.Name `xml:"Envelope"`
	CResult string   `xml:"Body>ConvertAdvancedTextResponse>Result"`
}

func (r *advancedTextResult) Result() string {
	return r.CResult
}
