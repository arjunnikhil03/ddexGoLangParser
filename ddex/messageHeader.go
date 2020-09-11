package ddex

type PartyName struct {
	Text                  string `xml:",chardata" json:"-"`
	LanguageAndScriptCode string `xml:"LanguageAndScriptCode,attr"`
	FullName              string `xml:"FullName"`
}

type MessageSender struct {
	Text      string    `xml:",chardata" json:"-"`
	PartyId   string    `xml:"PartyId"`
	PartyName PartyName `xml:"PartyName"`
}

type MessageRecipient struct {
	Text      string    `xml:",chardata" json:"-"`
	PartyId   string    `xml:"PartyId"`
	PartyName PartyName `xml:"PartyName"`
}

type MessageHeader struct {
	Text                   string           `xml:",chardata" json:"-"`
	MessageThreadId        string           `xml:"MessageThreadId"`
	MessageId              string           `xml:"MessageId"`
	MessageSender          MessageSender    `xml:"MessageSender"`
	MessageRecipient       MessageRecipient `xml:"MessageRecipient"`
	MessageCreatedDateTime string           `xml:"MessageCreatedDateTime"`
	MessageControlType     string           `xml:"MessageControlType"`
}
