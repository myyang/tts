package tts

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getResponse(req []byte) *bytes.Buffer {
	bodyReader := strings.NewReader(fmt.Sprintf(requestTemplate, req))
	response, err := http.Post(url, "text/xml", bodyReader)
	if err != nil {
		log.Fatalf("getResponse: post error: %v\n", err)
	}

	buf := make([]byte, response.ContentLength-1)
	_, err = response.Body.Read(buf)
	if err != nil {
		log.Fatalf("getResponse: read response error: %v\n", err)
	}

	return bytes.NewBuffer(buf)
}
