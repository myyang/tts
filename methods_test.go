package tts

import (
	"fmt"
	"testing"
)

var text = `
	我現在 schedule 上總 total 有10個 case 在 run, 等等還要跟我的 team 再 confirm 一下 format, 可能要再 review 一下新版的 checklist, 看 data 現在處理的 process 到哪邊,都 check 完、confirm了都 OK 的話,就只要給他們去 maintain 就好了,Anyway, 明天跟 RD 部門的 leader meeting還是 focus 在 interface 和 menu 上面, 反正他們都有 for 新平台的 know how 了，照 S O P 做,我 concern 的是 lab 裡面有什麼 special 的 idea. 感覺新來的比較沒 sense, present 的時候感覺進度一直 delay, 搞不好 boss 過幾天就會找他 talk 一下了`

func TestConvertSimple(t *testing.T) {
	url := ConvertSimple("ttsaccount", "ttspassword", text)
	fmt.Printf("\n\nURL: %#s\n\n", url)
	if url == "" {
		t.Fatalf("Fail to get file")
	}
}

func TestConvertText(t *testing.T) {
	url := ConvertText("ttsaccount", "ttspassword", text, "MCHEN_Joddess", "80", "3", "wav")
	fmt.Printf("\n\nURL: %#s\n\n", url)
	if url == "" {
		t.Fatalf("Fail to get file")
	}
}

func TestConvertAdvancedText(t *testing.T) {
	url := ConvertAdvancedText("ttsaccount", "ttspassword", text, "Angela", "95", "2", "wav", "1", "1", "10")
	fmt.Printf("\n\nURL: %#s\n\n", url)
	if url == "" {
		t.Fatalf("Fail to get file")
	}
}
