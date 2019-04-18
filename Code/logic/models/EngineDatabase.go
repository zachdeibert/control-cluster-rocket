package models

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// EngineDatabase represents a database of engine data
type EngineDatabase struct {
	XMLName xml.Name      `xml:"engine-database"`
	Engines []EngineModel `xml:"engine-list>engine"`
}

// Load loads a database xml file in RSE format into memory
func (db *EngineDatabase) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(bytes, db)
	return err
}

// String converts the EngineDatabase to a string
func (db EngineDatabase) String() string {
	str := ""
	for i, engine := range db.Engines {
		if i > 0 {
			str = str + "\n"
		}
		str = str + engine.String()
	}
	return str
}
