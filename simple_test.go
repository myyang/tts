package tts

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestSimpleStruct(t *testing.T) {
	simple := &convertSimple{Account: "Account", Password: "Password", TTStext: "Convert Text"}
	msimple := simple.Marshal()
	ssimple := []byte(
		`<ConvertSimple><accountID>Account</accountID>` +
			`<password>Password</password><TTStext>Convert Text</TTStext></ConvertSimple>`)
	if !reflect.DeepEqual(ssimple, msimple) {
		_, file, line, _ := runtime.Caller(0)
		t.Fatalf("%s:%d:\n\n\texp: %s\n\n\tgot: %s\n\n", filepath.Base(file), line, ssimple, msimple)
	}

	if "Account" != simple.getAccount() || "Password" != simple.getPassword() {
		_, file, line, _ := runtime.Caller(0)
		t.Fatalf(
			"%s:%d:\n\n\texp: Account:Password\n\tgot: %s:%s\n",
			filepath.Base(file), line, simple.getAccount(), simple.getPassword())
	}
}
