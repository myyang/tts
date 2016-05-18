package tts

// ConvertSimple return file URL with given text
func ConvertSimple(account, password, text string) string {
	q := &convertSimple{Account: account, Password: password, TTStext: text}
	return convertTemplate(q)
}

// ConvertText with optional speaker, volume, speed, outtype
func ConvertText(account, password, text, speaker, volume, speed, outtype string) string {
	q := &convertText{
		Account: account, Password: password, TTStext: text,
		TTSSpeaker: speaker, Volume: volume, Speed: speed, OutType: outtype,
	}
	return convertTemplate(q)
}

// ConvertAdvancedText with all available options
func ConvertAdvancedText(
	account, password, text, speaker, volume, speed, outtype,
	pitchSign, pitchLevel, pitchScale string) string {
	q := &convertAdvancedText{
		Account: account, Password: password, TTStext: text,
		TTSSpeaker: speaker, Volume: volume, Speed: speed, OutType: outtype,
		PitchLevel: pitchLevel, PitchSign: pitchSign, PitchScale: pitchScale,
	}
	return convertTemplate(q)
}

func convertTemplate(c convertType) string {
	m := c.Marshal()
	convertid := c.ParseResult(getResponse(m))
	statusQ := convertStatus{Account: c.getAccount(), Password: c.getPassword(), ConvertID: convertid}
	r := statusQ.WaitAndProcess()
	return <-r
}
