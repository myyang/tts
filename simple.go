package tts

import (
	"bytes"
	"encoding/xml"
	"log"
)

type convertSimple struct {
	XMLName  xml.Name `xml:"ConvertSimple"`
	Account  string   `xml:"accountID"`
	Password string   `xml:"password"`
	TTStext  string   `xml:"TTStext"`
}

func (c *convertSimple) Marshal() []byte {
	m, err := xml.Marshal(c)
	if err != nil {
		log.Fatalf("Fail to marshal: %v\n", c)
	}
	return m
}

func (c *convertSimple) getAccount() string {
	return c.Account
}

func (c *convertSimple) getPassword() string {
	return c.Password
}

func (c *convertSimple) ParseResult(buf *bytes.Buffer) string {
	r := &simpleResult{}
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

type simpleResult struct {
	XMLName xml.Name `xml:"Envelope"`
	CResult string   `xml:"Body>ConvertSimpleResponse>Result"`
}

func (r *simpleResult) Result() string {
	return r.CResult
}
