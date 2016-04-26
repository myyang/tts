// Package tts provides ...
package tts

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	url             string = "http://tts.itri.org.tw/TTSService/Soap_1_3.php?wsdl"
	requestTemplate string = "" +
		"<soap12:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"" +
		" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap12=\"http://www.w3.org/2003/05/soap-envelope\">" +
		"<soap12:Body>%s</soap12:Body></soap12:Envelope>"
)

type convertInfo struct {
	AccountID  string `xml:"accountID"`
	Password   string `xml:"password"`
	TTStext    string `xml:"TTStext"`
	TTSSpeaker string `xml:"TTSSpeaker,omitempty"`
	Volume     string `xml:"volume,omitempty"`
	Speed      string `xml:"speed,omitempty"`
	OutType    string `xml:"outType,omitempty"`
	PitchLevel string `xml:"PitchLevel,omitempty"`
	PitchSign  string `xml:"PitchSign,omitempty"`
	PitchScale string `xml:"PitchScale,omitempty"`
}

type convertSimpleXML struct {
	XMLName xml.Name `xml:"ConvertSimple"`
	convertInfo
}

type convertTextXML struct {
	XMLName xml.Name `xml:"ConvertText"`
	convertInfo
}

type convertAdvancedTextXML struct {
	XMLName xml.Name `xml:"ConvertAdvancedText"`
	convertInfo
}

type convertResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    body     `xml:"Body"`
}

type body struct {
	Simple        result `xml:"ConvertSimpleResponse,omitempty"`
	Text          result `xml:"ConvertTextResponse,omitempty"`
	AdvancedText  result `xml:"ConvertAdvancedTextResponse,omitempty"`
	ConvertStatus result `xml:"GetConvertStatusResponse,omitempty"`
}

type result struct {
	RPCResult string `xml:"result"`
	Result    string `xml:"Result"`
}

type convertStatusXML struct {
	XMLName   xml.Name `xml:"GetConvertStatus"`
	AccountID string   `xml:"accountID"`
	Password  string   `xml:"password"`
	ConvertID string   `xml:"convertID"`
}

// ConvertSimple provides shortcut to get converted sound file url
func ConvertSimple(account, password, text string) string {
	x := &convertSimpleXML{}
	x.convertInfo = convertInfo{AccountID: account, Password: password, TTStext: text}

	output, err := xml.Marshal(x)
	if err != nil {
		log.Fatalf("ConvertSimple: marshal error: %v\n", err)
	}

	buf := getResponse(output)
	convertID := parseConvertResult(buf)

	return getConvertStatus(account, password, convertID)
}

func getResponse(output []byte) []byte {
	bodyReader := strings.NewReader(fmt.Sprintf(requestTemplate, output))
	response, err := http.Post(url, "text/xml", bodyReader)
	if err != nil {
		log.Fatalf("getResponse: post error: %v\n", err)
	}

	buf := make([]byte, response.ContentLength-1)
	_, err = response.Body.Read(buf)
	if err != nil {
		log.Fatalf("getResponse: read response error: %v\n", err)
	}

	return buf
}

func parseConvertResult(buf []byte) string {
	r := convertResponse{}
	err := xml.Unmarshal(buf, &r)
	if err != nil {
		log.Fatalf("parseConvertResult: unmarshal error: %v\n", err)
	}

	s := ""
	switch {
	case r.Body.Simple != result{}:
		s = r.Body.Simple.Result
	case r.Body.Text != result{}:
		s = r.Body.Text.Result
	case r.Body.AdvancedText != result{}:
		s = r.Body.AdvancedText.Result
	}

	re, err := regexp.Compile(`(?P<resultCode>[[:digit:]]+)&(?P<resultMsg>[\s[:alnum:]]+)&?(?P<covertID>[[:digit:]]+)?`)
	if err != nil {
		log.Fatalf("parseConvertResult: regular exp fail: %v\n", re)
	}

	if !re.MatchString(s) {
		log.Fatalf("parseConvertResult: %s, not match %v\n", s, re.FindStringSubmatch(s))
	}

	log.Printf("parseConvertResult match: %v\n", re.FindStringSubmatch(s))

	if re.FindStringSubmatch(s)[1] != "0" {
		log.Fatal("parseConvertResult: Remote error: %v\n", re.FindStringSubmatch(s))
	}

	return re.FindStringSubmatch(s)[3]

}

func getConvertStatus(accountID, password, convertID string) string {
	x := &convertStatusXML{AccountID: accountID, Password: password, ConvertID: convertID}

	output, err := xml.Marshal(x)
	if err != nil {
		log.Fatalf("ConvertSimple: marshal error: %v\n", err)
	}

	fail, r := 3, ""
	for fail > 0 {
		buf := getResponse(output)
		r = parseConvertStatus(buf)
		if r != "" {
			log.Printf("Convertion done. File: %s\n", r)
			break
		}
		log.Printf("Wait and retry to fetch file...\n")
		time.Sleep(2 * time.Second)
		fail--
	}
	return r
}

func parseConvertStatus(buf []byte) string {
	r := convertResponse{}
	err := xml.Unmarshal(buf, &r)
	if err != nil {
		log.Fatalf("parseConvertStatus: unmarshal error: %v\n", err)
	}
	log.Printf("parseConvertStatus: response: %v\n", r)
	s := r.Body.ConvertStatus.Result
	re, err := regexp.Compile(`(?P<resultCode>[[:digit:]]+)&(?P<resultMsg>[[:alnum:]]+)&(?P<statusCode>[[:digit:]]+)&(?P<statusMsg>[[:alnum:]]+)&?(?P<url>.*)?`)
	if err != nil {
		log.Fatalf("parseConvertStatus: regular exp fail: %v\n", re)
	}

	if !re.MatchString(s) {
		log.Printf("parseConvertStatus: not match %v\n", re.FindStringSubmatch(s))
		return ""
	}
	return re.FindStringSubmatch(s)[5]
}

// ConvertText provides shortcut to get converted sound file url
func ConvertText(account, password, text, speaker, volume, speed, outtype string) string {
	x := &convertTextXML{}
	x.convertInfo = convertInfo{
		AccountID: account, Password: password, TTStext: text,
		TTSSpeaker: speaker, Volume: volume, Speed: speed, OutType: outtype,
	}

	output, err := xml.Marshal(x)
	if err != nil {
		log.Fatalf("ConvertText: marshal error: %v\n", err)
	}

	buf := getResponse(output)
	convertID := parseConvertResult(buf)

	return getConvertStatus(account, password, convertID)
}

// ConvertAdvancedText provides shortcut to get converted sound file url
func ConvertAdvancedText(
	account, password, text, speaker, volume, speed, outtype,
	pitchSign, pitchLevel, pitchScale string) string {
	x := &convertAdvancedTextXML{}
	x.convertInfo = convertInfo{
		AccountID: account, Password: password, TTStext: text,
		TTSSpeaker: speaker, Volume: volume, Speed: speed, OutType: outtype,
		PitchLevel: pitchLevel, PitchSign: pitchSign, PitchScale: pitchScale,
	}

	output, err := xml.Marshal(x)
	if err != nil {
		log.Fatalf("ConvertAdvancedText: marshal error: %v\n", err)
	}

	buf := getResponse(output)
	convertID := parseConvertResult(buf)

	return getConvertStatus(account, password, convertID)
}
