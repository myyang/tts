package tts

import (
	"log"
	"regexp"
)

func matchConvertResult(r convertResult) (string, error) {
	s := r.Result()
	re, _ := regexp.Compile(`(?P<resultCode>[[:digit:]]+)&(?P<resultMsg>[\s[:alnum:]]+)&?(?P<covertID>[[:digit:]]+)?`)

	if !re.MatchString(s) || re.FindStringSubmatch(s)[1] != "0" {
		log.Fatalf("matchConvertResult: %s, not match or fail: %v\n", s, re.FindStringSubmatch(s))
	}

	log.Printf("matchConvertResult: %v\n", re.FindStringSubmatch(s))
	return re.FindStringSubmatch(s)[3], nil
}

func matchStatusResult(r statusResult) string {
	s := r.Result()
	re, _ := regexp.Compile(
		`(?P<resultCode>[[:digit:]]+)&(?P<resultMsg>[[:alnum:]]+)` +
			`&(?P<statusCode>[[:digit:]]+)&(?P<statusMsg>[[:alnum:]]+)&?(?P<url>.*)?`)

	if !re.MatchString(s) {
		log.Printf("matchStatusResult: not match %v\n", re.FindStringSubmatch(s))
		return ""
	} else if re.FindStringSubmatch(s)[3] != "2" {
		log.Printf("matchStatusResult: not completed yet %v\n", re.FindStringSubmatch(s))
		return ""
	}

	return re.FindStringSubmatch(s)[5]
}
