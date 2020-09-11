package controllers

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/moodleexpert/ddexGoLangParser/ddex"
	"github.com/moodleexpert/ddexGoLangParser/models"
	"github.com/moodleexpert/ddexGoLangParser/utils"
)

//Process ...
func Process(res http.ResponseWriter, req *http.Request) {
	fileName := "xmls/A10301A0001716010S.xml"
	xmlFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		utils.Log.Error(err)
	}

	utils.Log.Info("Successfully Opened users.xml")

	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var dx ddex.NewReleaseMessage

	xml.Unmarshal(byteValue, &dx)

	message := utils.Message(true, "Proccessed XML", 200)
	message["data"] = models.ProcessDDEX(dx, fileName)
	utils.Respond(res, message)
}
