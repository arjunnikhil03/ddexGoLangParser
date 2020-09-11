package ddex

type DealList struct {
	Text        string        `xml:",chardata" json:"-"`
	ReleaseDeal []ReleaseDeal `xml:"ReleaseDeal"`
}

type ReleaseDeal struct {
	Text                 string   `xml:",chardata" json:"-"`
	DealReleaseReference []string `xml:"DealReleaseReference"`
	Deal                 []Deal   `xml:"Deal"`
}

type Deal struct {
	Text      string    `xml:",chardata" json:"-"`
	DealTerms DealTerms `xml:"DealTerms"`
}

type Usage struct {
	Text    string   `xml:",chardata" json:"-"`
	UseType []string `xml:"UseType"`
}

type DealTerms struct {
	Text                string `xml:",chardata" json:"-"`
	CommercialModelType string `xml:"CommercialModelType"`
	Usage               Usage  `xml:"Usage"`
	TerritoryCode       struct {
		Text           string `xml:",chardata" json:"-"`
		IdentifierType string `xml:"IdentifierType,attr"`
	} `xml:"TerritoryCode"`
	ValidityPeriod struct {
		Text      string `xml:",chardata" json:"-"`
		StartDate string `xml:"StartDate"`
	} `xml:"ValidityPeriod"`
	TakenDown string `xml:"TakenDown"`
}
