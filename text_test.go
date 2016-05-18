package tts

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestTextStruct(t *testing.T) {
	text := &convertText{
		Account: "Account", Password: "Password", TTStext: "Convert Text",
		TTSSpeaker: "A", Volume: "50", Speed: "5", OutType: "wav",
	}
	mtext := text.Marshal()
	stext := []byte(
		`<ConvertText><accountID>Account</accountID>` +
			`<password>Password</password><TTStext>Convert Text</TTStext>` +
			`<TTSSpeaker>A</TTSSpeaker><volume>50</volume><speed>5</speed><outType>wav</outType></ConvertText>`)
	if !reflect.DeepEqual(stext, mtext) {
		_, file, line, _ := runtime.Caller(0)
		t.Fatalf("%s:%d:\n\n\texp: %s\n\n\tgot: %s\n\n", filepath.Base(file), line, stext, mtext)
	}

	if "Account" != text.getAccount() || "Password" != text.getPassword() {
		_, file, line, _ := runtime.Caller(0)
		t.Fatalf(
			"%s:%d:\n\n\texp: Account:Password\n\tgot: %s:%s\n",
			filepath.Base(file), line, text.getAccount(), text.getPassword())
	}
}
