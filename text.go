package tts

import (
	"bytes"
	"encoding/xml"
	"log"
)

type convertText struct {
	XMLName    xml.Name `xml:"ConvertText"`
	Account    string   `xml:"accountID"`
	Password   string   `xml:"password"`
	TTStext    string   `xml:"TTStext"`
	TTSSpeaker string   `xml:"TTSSpeaker"`
	Volume     string   `xml:"volume"`
	Speed      string   `xml:"speed"`
	OutType    string   `xml:"outType"`
}

func (c *convertText) Marshal() []byte {
	m, err := xml.Marshal(c)
	if err != nil {
		log.Fatalf("Fail to marshal: %v\n", c)
	}
	return m
}

func (c *convertText) getAccount() string {
	return c.Account
}

func (c *convertText) getPassword() string {
	return c.Password
}

func (c *convertText) ParseResult(buf *bytes.Buffer) string {
	r := &textResult{}
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

type textResult struct {
	XMLName xml.Name `xml:"Envelope"`
	CResult string   `xml:"Body>ConvertTextResponse>Result"`
}

func (r *textResult) Result() string {
	return r.CResult
}
