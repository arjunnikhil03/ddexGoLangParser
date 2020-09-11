package ddex

type ResourceList struct {
	Text           string           `xml:",chardata" json:"-"`
	SoundRecording []SoundRecording `xml:"SoundRecording"`
	Image          Image            `xml:"Image"`
}
