package tts

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestAdvancedTextStruct(t *testing.T) {
	text := &convertAdvancedText{
		Account: "Account", Password: "Password", TTStext: "Convert Text",
		TTSSpeaker: "A", Volume: "50", Speed: "5", OutType: "wav",
		PitchLevel: "-3", PitchSign: "0", PitchScale: "8",
	}
	mtext := text.Marshal()
	stext := []byte(
		`<ConvertAdvancedText><accountID>Account</accountID>` +
			`<password>Password</password><TTStext>Convert Text</TTStext>` +
			`<TTSSpeaker>A</TTSSpeaker><volume>50</volume><speed>5</speed><outType>wav</outType>` +
			`<PitchLevel>-3</PitchLevel><PitchSign>0</PitchSign><PitchScale>8</PitchScale></ConvertAdvancedText>`)
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
