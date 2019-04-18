package models

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	// Rocket denotes that a component is the airframe itself
	Rocket = 0
	// MainEngine denotes that a motor is the main engine that the flight
	// computer does not need to ignite ever (because an external system will)
	MainEngine = 1
	// BoosterEngine denotes that a motor can be lit by the flight computer
	BoosterEngine = 2
)

// IgnitionType describes what kind of ignition procedure a motor uses
type IgnitionType int

// IgnitionModel describes the ignition procedure a motor uses
type IgnitionModel struct {
	Type IgnitionType
	ID   int
}

// RocketComponent describes one piece of the rocket that needs to be modeled
// separately
type RocketComponent struct {
	XMLName  xml.Name `xml:"component"`
	ID       int
	Code     string `xml:"code,attr"`
	Model    EngineModel
	Position float32       `xml:"position,attr"`
	Ignition IgnitionModel `xml:"ignition,attr"`
}

// RocketModel describes the overall model for the rocket
type RocketModel struct {
	XMLName       xml.Name          `xml:"rocket"`
	CD            float32           `xml:"cd,attr"`
	ReferenceArea float32           `xml:"refA,attr"`
	Components    []RocketComponent `xml:"component"`
}

// UnmarshalXMLAttr unmarshals an ignition model from an XML attribute
func (ig *IgnitionModel) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "rocket":
		ig.Type = Rocket
		ig.ID = -1
		break
	case "main":
		ig.Type = MainEngine
		ig.ID = -1
		break
	default:
		ig.Type = BoosterEngine
		i, err := strconv.ParseInt(attr.Value, 10, 32)
		ig.ID = int(i)
		return err
	}
	return nil
}

// String converts the IgnitionModel to a string
func (ig IgnitionModel) String() string {
	switch ig.Type {
	case Rocket:
		return "Rocket"
	case MainEngine:
		return "Main Engine"
	case BoosterEngine:
		return fmt.Sprintf("Booster #%d", ig.ID)
	}
	panic(errors.New("Invalid enumeration constant"))
}

// String converts the RocketComponent to a string
func (rc RocketComponent) String() string {
	return fmt.Sprintf("(Code = %s, Position = %f, Ignition = %s, Model = {%s})", rc.Code, rc.Position, rc.Ignition, rc.Model)
}

// String converts the RocketModel to a string
func (rm RocketModel) String() string {
	return fmt.Sprintf("Rocket with CD=%f, ref. area=%f, components=%s", rm.CD, rm.ReferenceArea, rm.Components)
}

// Load loads a rocket xml file into memory
func (rm *RocketModel) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(bytes, rm)
	for i := range rm.Components {
		rm.Components[i].ID = i
	}
	return err
}

// LoadEngines loads the engine models from an engine database
func (rm *RocketModel) LoadEngines(db EngineDatabase) error {
	for i, comp := range rm.Components {
		found := false
		for _, eng := range db.Engines {
			if eng.Code == comp.Code {
				rm.Components[i].Model = eng
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("Unable to find engine with code %s", comp.Code)
		}
	}
	return nil
}
