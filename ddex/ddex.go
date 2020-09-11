package ddex

import "encoding/xml"

//DDEX ..
type DDEX struct {
	MessageHeader MessageHeader
}

//NewReleaseMessage ...
type NewReleaseMessage struct {
	XMLName                xml.Name      `xml:"NewReleaseMessage"`
	Text                   string        `xml:",chardata"`
	Ernm                   string        `xml:"ernm,attr"`
	Xs                     string        `xml:"xs,attr"`
	LanguageAndScriptCode  string        `xml:"LanguageAndScriptCode,attr"`
	MessageSchemaVersionId string        `xml:"MessageSchemaVersionId,attr"`
	SchemaLocation         string        `xml:"schemaLocation,attr"`
	MessageHeader          MessageHeader `xml:"MessageHeader"`
	UpdateIndicator        string        `xml:"UpdateIndicator"`
	ResourceList           ResourceList  `xml:"ResourceList"`
	ReleaseList            ReleaseList   `xml:"ReleaseList"`
	DealList               DealList      `xml:"DealList"`
}
