package ddex

type DisplayArtist struct {
	Text           string    `xml:",chardata" json:"-"`
	SequenceNumber string    `xml:"SequenceNumber,attr"`
	PartyName      PartyName `xml:"PartyName"`
	ArtistRole     []string  `xml:"ArtistRole"`
}

type Title struct {
	Text                  string `xml:",chardata" json:"-"`
	LanguageAndScriptCode string `xml:"LanguageAndScriptCode,attr"`
	TitleType             string `xml:"TitleType,attr"`
	TitleText             string `xml:"TitleText"`
}

type ResourceContributorRole struct {
	Text             string `xml:",chardata" json:"-"`
	Namespace        string `xml:"Namespace,attr"`
	UserDefinedValue string `xml:"UserDefinedValue,attr"`
}

type LabelName struct {
	Text          string `xml:",chardata",json:"-"`
	LabelNameType string `xml:"LabelNameType,attr"`
}

type PLine struct {
	Text      string `xml:",chardata" json:"-"`
	Year      string `xml:"Year"`
	PLineText string `xml:"PLineText"`
}

type Genre struct {
	Text      string `xml:",chardata" json:"-"`
	GenreText string `xml:"GenreText"`
}

type ResourceContributor struct {
	Text                    string                    `xml:",chardata" json:"-"`
	SequenceNumber          string                    `xml:"SequenceNumber,attr"`
	PartyName               []PartyName               `xml:"PartyName"`
	ResourceContributorRole []ResourceContributorRole `xml:"ResourceContributorRole"`
}

type SoundRecordingDetailsByTerritory struct {
	Text                string                `xml:",chardata" json:"-"`
	TerritoryCode       string                `xml:"TerritoryCode"`
	Title               []Title               `xml:"Title"`
	DisplayArtist       DisplayArtist         `xml:"DisplayArtist"`
	ResourceContributor []ResourceContributor `xml:"ResourceContributor"`
	DisplayArtistName   string                `xml:"DisplayArtistName"`
	LabelName           LabelName             `xml:"LabelName"`
	PLine               PLine                 `xml:"PLine"`
	Genre               Genre                 `xml:"Genre"`
	ParentalWarningType string                `xml:"ParentalWarningType"`
}

type SoundRecordingId struct {
	Text string `xml:",chardata",json:"-"`
	ISRC string `xml:"ISRC"`
}

type ReferenceTitle struct {
	Text                  string `xml:",chardata" json:"-"`
	LanguageAndScriptCode string `xml:"LanguageAndScriptCode,attr"`
	TitleText             string `xml:"TitleText"`
}

type SoundRecording struct {
	Text                             string                           `xml:",chardata" json:"-"`
	SoundRecordingType               string                           `xml:"SoundRecordingType"`
	SoundRecordingId                 SoundRecordingId                 `xml:"SoundRecordingId"`
	ResourceReference                string                           `xml:"ResourceReference"`
	ReferenceTitle                   ReferenceTitle                   `xml:"ReferenceTitle"`
	LanguageOfPerformance            string                           `xml:"LanguageOfPerformance"`
	Duration                         string                           `xml:"Duration"`
	SoundRecordingDetailsByTerritory SoundRecordingDetailsByTerritory `xml:"SoundRecordingDetailsByTerritory"`
}
