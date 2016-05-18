package tts

import (
	"bytes"
	"encoding/xml"
	"log"
	"time"
)

type convertStatus struct {
	XMLName   xml.Name `xml:"GetConvertStatus"`
	Account   string   `xml:"accountID"`
	Password  string   `xml:"password"`
	ConvertID string   `xml:"convertID"`
}

func (c *convertStatus) Marshal() []byte {
	m, err := xml.Marshal(c)
	if err != nil {
		log.Fatalf("Fail to marshal: %v\n", c)
	}
	return m
}

func (c *convertStatus) getAccount() string {
	return c.Account
}

func (c *convertStatus) getPassword() string {
	return c.Password
}

func (c *convertStatus) ParseResult(buf *bytes.Buffer) string {
	r := statusResult{}
	err := xml.Unmarshal(buf.Bytes(), &r)
	if err != nil {
		log.Fatalf("Unmarshal error: %v\n", err)
	}
	return matchStatusResult(r)
}

func (c *convertStatus) WaitAndProcess() chan string {
	result := make(chan string)
	go func() {
		fail, r := 3, ""
		for fail > 0 {
			r = c.ParseResult(getResponse(c.Marshal()))
			if r != "" {
				log.Printf("Convertion done. File: %s\n", r)
				result <- r
				break
			}
			log.Printf("Wait 5 seconds and retry to fetch file...\n")
			time.Sleep(5 * time.Second)
			fail--
		}

		if fail < 0 {
			log.Printf("Exceed maximum failure times\n")
		}

		result <- ""
	}()
	return result
}

type statusResult struct {
	XMLName xml.Name `xml:"Envelope"`
	CResult string   `xml:"Body>GetConvertStatusResponse>Result"`
}

func (r *statusResult) Result() string {
	return r.CResult
}
