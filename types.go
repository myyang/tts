package tts

import "bytes"

const (
	url             string = "http://tts.itri.org.tw/TTSService/Soap_1_3.php?wsdl"
	requestTemplate string = "" +
		"<soap12:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"" +
		" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\"" +
		" xmlns:soap12=\"http://www.w3.org/2003/05/soap-envelope\">" +
		"<soap12:Body>%s</soap12:Body></soap12:Envelope>"
)

type convertType interface {
	Marshal() []byte
	ParseResult(*bytes.Buffer) string
	getAccount() string
	getPassword() string
}

type convertResult interface {
	Result() string
}
