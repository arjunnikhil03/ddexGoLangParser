package models

import (
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/moodleexpert/ddexGoLangParser/utils"
)

type db struct {
	DBConn *storm.DB
	err    error
}

//Connect ...
func (db *db) Connect() {
	if db.DBConn == nil {
		db.DBConn, db.err = storm.Open("my.db")
		if db.err != nil {
			utils.Log.Fatal("Error:", db.err.Error())
		}
	}
}

//Close
func (db *db) Close() {
	if db.DBConn != nil {
		db.DBConn.Close()
	}
	db.DBConn = nil
}

func (db *db) Save(data interface{}) (bool, error) {
	db.Connect()
	defer db.Close()
	err := db.DBConn.Save(data)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (db *db) Find(data interface{}, matchers q.Matcher) {
	db.Connect()
	defer db.Close()
	db.DBConn.Select(matchers).Find(data)
}

func (db *db) Count(data interface{}, matchers q.Matcher) (int, error) {
	db.Connect()
	defer db.Close()
	return db.DBConn.Select(matchers).Count(data)
}

func (db *db) First(data interface{}, matchers q.Matcher) {
	db.Connect()
	defer db.Close()
	db.DBConn.Select(matchers).First(data)
}

//NewDB ..
func NewDB() *db {
	return &db{}
}
