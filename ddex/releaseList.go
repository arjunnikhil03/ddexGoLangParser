package ddex

type ReleaseId struct {
	Text          string `xml:",chardata" json:"-"`
	GRid          string `xml:"GRid"`
	ICPN          string `xml:"ICPN"`
	CatalogNumber struct {
		Text      string `xml:",chardata"`
		Namespace string `xml:"Namespace,attr"`
	} `xml:"CatalogNumber"`
	ISRC string `xml:"ISRC"`
}

type ReleaseResourceReference struct {
	Text                string `xml:",chardata" json:"-"`
	ReleaseResourceType string `xml:"ReleaseResourceType,attr"`
}

type ReleaseResourceReferenceList struct {
	Text                     string                     `xml:",chardata" json:"-"`
	ReleaseResourceReference []ReleaseResourceReference `xml:"ReleaseResourceReference"`
}

type Release struct {
	Text                         string                       `xml:",chardata" json:"-"`
	IsMainRelease                string                       `xml:"IsMainRelease,attr"`
	ReleaseId                    ReleaseId                    `xml:"ReleaseId"`
	ReleaseReference             string                       `xml:"ReleaseReference"`
	ReferenceTitle               ReferenceTitle               `xml:"ReferenceTitle"`
	ReleaseResourceReferenceList ReleaseResourceReferenceList `xml:"ReleaseResourceReferenceList"`
	ReleaseType                  string                       `xml:"ReleaseType"`
	ReleaseDetailsByTerritory    ReleaseDetailsByTerritory    `xml:"ReleaseDetailsByTerritory"`
	Duration                     string                       `xml:"Duration"`
	PLine                        PLine                        `xml:"PLine"`
	GlobalOriginalReleaseDate    string                       `xml:"GlobalOriginalReleaseDate"`
}

type ReleaseList struct {
	Text    string    `xml:",chardata",json:"-"`
	Release []Release `xml:"Release"`
}

type ResourceGroupContentItem struct {
	Text                     string                   `xml:",chardata" json:"-"`
	SequenceNumber           string                   `xml:"SequenceNumber"`
	ResourceType             string                   `xml:"ResourceType"`
	ReleaseResourceReference ReleaseResourceReference `xml:"ReleaseResourceReference"`
}

type ResourceGroup struct {
	Text                     string                     `xml:",chardata" json:"-"`
	Title                    Title                      `xml:"Title"`
	SequenceNumber           string                     `xml:"SequenceNumber"`
	ResourceGroupContentItem []ResourceGroupContentItem `xml:"ResourceGroupContentItem"`
}

type ReleaseDetailsByTerritory struct {
	Text                string        `xml:",chardata" json:"-"`
	TerritoryCode       string        `xml:"TerritoryCode"`
	DisplayArtistName   string        `xml:"DisplayArtistName"`
	LabelName           LabelName     `xml:"LabelName"`
	Title               []Title       `xml:"Title"`
	DisplayArtist       DisplayArtist `xml:"DisplayArtist"`
	ParentalWarningType string        `xml:"ParentalWarningType"`
	ResourceGroup       struct {
		Text                     string                   `xml:",chardata" json:"-"`
		ResourceGroup            ResourceGroup            `xml:"ResourceGroup"`
		ResourceGroupContentItem ResourceGroupContentItem `xml:"ResourceGroupContentItem"`
	} `xml:"ResourceGroup"`
	Genre Genre `xml:"Genre"`
}
