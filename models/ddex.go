package models

import (
	"strings"
	"time"

	"github.com/asdine/storm/q"
	"github.com/moodleexpert/ddexGoLangParser/ddex"
	"github.com/moodleexpert/ddexGoLangParser/utils"
)

//XMLList ...
type XMLList struct {
	ID                     int    `storm:"id,increment"`
	Filename               string `storm:"index"`
	MessageThreadID        string
	MessageID              string
	MessageSenderID        string
	MessageSenderName      string
	MessageRecipientID     string
	MessageRecipientName   string
	MessageCreatedDateTime string
	XMLStatus              bool `storm:"index"`
	ProcessStatus          bool `storm:"index"`
	ProcessDate            time.Time
	TotalSongs             int
	DownloadSongs          int
	DownloadImages         int
	LogSend                bool
	UpdateIndicator        string
}

//ReleaseList ..
type ReleaseList struct {
	ID                        int    `storm:"id,increment"`
	Grid                      string `storm:"index"`
	ICPN                      string `storm:"index"`
	ISRC                      string `storm:"index"`
	DPID                      string
	XMLID                     int  `storm:"index"`
	ProcessStatus             bool `storm:"index"`
	ReleaseReference          string
	CatalogNumber             string
	Title                     string `storm:"index"`
	ReleaseResourceReference  []ddex.ReleaseResourceReference
	GlobalOriginalReleaseDate string
	Year                      string
	Label                     string
	Genre                     string
	ReleaseType               string `storm:"index"`
	IsMainRelease             string `storm:"index"`
}

//ResourceList ..
type ResourceList struct {
	ID                               int    `storm:"id,increment"`
	XMLID                            int    `storm:"index"`
	ISRC                             string `storm:"index"`
	Image                            ddex.Image
	ResourceReference                string
	ReferenceTitle                   string `storm:"index"`
	LanguageOfPerformance            string
	Duration                         string
	SoundRecordingDetailsByTerritory ddex.SoundRecordingDetailsByTerritory
}

//DealList ..
type DealList struct {
	ID                   int    `storm:"id,increment"`
	XMLID                int    `storm:"index"`
	ISRC                 string `storm:"index"`
	DealReleaseReference []string
	CommercialModelType  string `storm:"index"`
	Usage                string
	TerritoryCode        string
	ValidityPeriod       string `storm:"index"`
	TakenDown            string
	ValidityEndPeriod    string `storm:"index"`
	CatalogNumber        string `storm:"index"`
	Grid                 string `storm:"index"`
	ICPN                 string `storm:"index"`
	ReleaseType          string
}

//ReponseDDEX ...
type ReponseDDEX struct {
	XMLList      XMLList
	ReleaseList  []ReleaseList
	DealList     []DealList
	ResourceList []ResourceList
}

//ProcessDDEX ...
func ProcessDDEX(dx ddex.NewReleaseMessage, fileName string) ReponseDDEX {
	var rd ReponseDDEX
	rd.XMLList = messageHeader(dx, fileName)
	rd.ReleaseList = releaseList(dx, rd)
	rd.DealList = dealList(dx, rd)
	rd.ResourceList = resourceList(dx, rd)
	rd.XMLList.ProcessStatus = true
	save, err := NewDB().Save(&rd.XMLList)
	if !save && err != nil {
		utils.Log.Fatal("DB Error:", err)
	}

	utils.Log.Info("Update XML Data")
	return rd
}

func messageHeader(dx ddex.NewReleaseMessage, fileName string) XMLList {
	var xmlList XMLList

	xmlList.Filename = fileName
	xmlList.MessageThreadID = dx.MessageHeader.MessageThreadId
	xmlList.MessageID = dx.MessageHeader.MessageId
	xmlList.MessageSenderID = dx.MessageHeader.MessageSender.PartyId
	xmlList.MessageRecipientID = dx.MessageHeader.MessageRecipient.PartyId
	xmlList.MessageSenderName = dx.MessageHeader.MessageSender.PartyName.FullName
	xmlList.MessageRecipientName = dx.MessageHeader.MessageRecipient.PartyName.FullName
	xmlList.MessageCreatedDateTime = dx.MessageHeader.MessageCreatedDateTime
	xmlList.XMLStatus = true
	xmlList.ProcessStatus = false
	xmlList.ProcessDate = time.Now()
	xmlList.LogSend = false

	utils.Log.Info("Process the Message header")
	utils.Log.Printf("%+v\n", xmlList)

	save, err := NewDB().Save(&xmlList)
	if !save && err != nil {
		utils.Log.Fatal("DB Error:", err)
	}

	utils.Log.Info("Inserted Message Header")

	return xmlList
}

func resourceList(dx ddex.NewReleaseMessage, rd ReponseDDEX) []ResourceList {
	var rsArr []ResourceList
	image := dx.ResourceList.Image

	for _, sr := range dx.ResourceList.SoundRecording {
		var rs ResourceList
		rs.Image = image
		rs.ResourceReference = sr.ResourceReference
		rs.ISRC = sr.SoundRecordingId.ISRC
		rs.ReferenceTitle = sr.ReferenceTitle.TitleText
		rs.LanguageOfPerformance = sr.LanguageOfPerformance
		rs.Duration = sr.Duration
		rs.SoundRecordingDetailsByTerritory = sr.SoundRecordingDetailsByTerritory
		rs.XMLID = rd.XMLList.ID

		//utils.Log.Info("Process the SoundRecordingId")
		//utils.Log.Printf("%+v\n", rs)

		save, err := NewDB().Save(&rs)
		if !save && err != nil {
			utils.Log.Fatal("DB Error:", err)
		}

		utils.Log.Info("Inserted SoundRecordingId")
		rsArr = append(rsArr, rs)
	}
	return rsArr
}

func releaseList(dx ddex.NewReleaseMessage, rd ReponseDDEX) []ReleaseList {
	var relArr []ReleaseList
	for _, release := range dx.ReleaseList.Release {
		var rel ReleaseList
		rel.Grid = release.ReleaseId.GRid
		rel.IsMainRelease = release.IsMainRelease
		rel.ICPN = release.ReleaseId.ICPN
		rel.ISRC = release.ReleaseId.ISRC
		rel.CatalogNumber = release.ReleaseId.CatalogNumber.Text
		rel.Title = release.ReferenceTitle.TitleText
		rel.ReleaseReference = release.ReleaseReference
		rel.ReleaseResourceReference = release.ReleaseResourceReferenceList.ReleaseResourceReference
		rel.GlobalOriginalReleaseDate = release.GlobalOriginalReleaseDate
		rel.Year = release.PLine.Year
		rel.Label = release.PLine.PLineText
		rel.Genre = release.ReleaseDetailsByTerritory.Genre.GenreText
		rel.ReleaseType = release.ReleaseType
		rel.XMLID = rd.XMLList.ID
		utils.Log.Printf("%+v\n", rel)

		save, err := NewDB().Save(&rel)
		if !save && err != nil {
			utils.Log.Fatal("DB Error:", err)
		}

		utils.Log.Info("Inserted Release Header")

		relArr = append(relArr, rel)

	}
	return relArr
}

func dealList(dx ddex.NewReleaseMessage, rd ReponseDDEX) []DealList {
	var dealArr []DealList
	var dLArr []DealList
	for _, releasedeal := range dx.DealList.ReleaseDeal {
		var deal DealList
		deal.DealReleaseReference = releasedeal.DealReleaseReference
		for _, d := range releasedeal.Deal {
			deal.CommercialModelType = d.DealTerms.CommercialModelType
			deal.Usage = strings.Join(d.DealTerms.Usage.UseType, "/")
			deal.TerritoryCode = d.DealTerms.TerritoryCode.Text
			deal.ValidityPeriod = d.DealTerms.ValidityPeriod.StartDate
			deal.TakenDown = d.DealTerms.TakenDown
			if !utils.EmptyString(d.DealTerms.TakenDown) {
				deal.ValidityEndPeriod = d.DealTerms.ValidityPeriod.StartDate
			} else {
				deal.ValidityEndPeriod = ""
			}
		}

		dealArr = append(dealArr, deal)
	}

	for _, r := range rd.ReleaseList {
		for _, d := range dealArr {

			if utils.InArray(r.ReleaseReference, d.DealReleaseReference) {
				var dL DealList

				d.ISRC = r.ISRC
				d.ICPN = r.ICPN
				d.Grid = r.Grid
				d.CatalogNumber = r.CatalogNumber
				d.ReleaseType = r.ReleaseType

				if r.IsMainRelease == "true" {
					NewDB().First(&dL, q.And(q.Eq("ICPN", d.ICPN), q.Eq("CommercialModelType", d.CommercialModelType), q.Eq("CatalogNumber", r.CatalogNumber)))
				} else {
					NewDB().First(&dL, q.And(q.Eq("ISRC", d.ISRC), q.Eq("CommercialModelType", d.CommercialModelType)))
				}
				if dL.ID > 0 {
					d.ID = dL.ID
				}
				d.XMLID = rd.XMLList.ID
				dLArr = append(dLArr, d)
				NewDB().Save(&d)
				utils.Log.Info("Inserted Deal Header")
			}
		}
	}
	return dLArr
}
