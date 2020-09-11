package ddex

type ImageId struct {
	Text          string `xml:",chardata" json:"-"`
	ProprietaryId struct {
		Text      string `xml:",chardata"`
		Namespace string `xml:"Namespace,attr"`
	} `xml:"ProprietaryId"`
}

type Image struct {
	Text                    string                  `xml:",chardata" json:"-"`
	ImageType               string                  `xml:"ImageType"`
	ImageId                 ImageId                 `xml:"ImageId"`
	ResourceReference       string                  `xml:"ResourceReference"`
	ImageDetailsByTerritory ImageDetailsByTerritory `xml:"ImageDetailsByTerritory"`
}

type ImageDetailsByTerritory struct {
	Text                  string                `xml:",chardata" json:"-"`
	TerritoryCode         string                `xml:"TerritoryCode"`
	TechnicalImageDetails TechnicalImageDetails `xml:"TechnicalImageDetails"`
}

type TechnicalImageDetails struct {
	Text                              string `xml:",chardata" json:"-"`
	TechnicalResourceDetailsReference string `xml:"TechnicalResourceDetailsReference"`
	ImageCodecType                    string `xml:"ImageCodecType"`
	ImageHeight                       string `xml:"ImageHeight"`
	ImageWidth                        string `xml:"ImageWidth"`
	ImageResolution                   string `xml:"ImageResolution"`
	IsPreview                         string `xml:"IsPreview"`
	File                              File   `xml:"File"`
}

type File struct {
	Text    string `xml:",chardata" json:"-"`
	URL     string `xml:"URL"`
	HashSum struct {
		Text                 string `xml:",chardata" json:"-"`
		HashSum              string `xml:"HashSum"`
		HashSumAlgorithmType string `xml:"HashSumAlgorithmType"`
	} `xml:"HashSum"`
}
