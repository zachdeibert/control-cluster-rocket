package models

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// EngineType represents the type of an engine
type EngineType int

const (
	// SingleUse is an engine that can only be used once
	SingleUse EngineType = 0
	// Reloadable is a solid engine that can be reloaded (and the casing reused)
	Reloadable EngineType = 1
	// Hybrid is a hybrid engine
	Hybrid EngineType = 2
)

// DataPoint represents one data sample during the test
type DataPoint struct {
	XMLName xml.Name `xml:"eng-data"`
	Time    float32  `xml:"t,attr"`
	Force   float32  `xml:"f,attr"`
	Mass    float32  `xml:"m,attr"`
	CG      float32  `xml:"cg,attr"`
}

// Delays represents an array of different delay values
type Delays []float32

// EngineModel represents the thrust model for an engine
type EngineModel struct {
	XMLName         xml.Name    `xml:"engine"`
	Manufacturer    string      `xml:"mfg,attr"`
	Code            string      `xml:"code,attr"`
	Type            EngineType  `xml:"Type,attr"`
	Diameter        float32     `xml:"dia,attr"`
	Length          float32     `xml:"len,attr"`
	InitialMass     float32     `xml:"initWt,attr"`
	PropellantMass  float32     `xml:"propWt,attr"`
	Delays          Delays      `xml:"delays,attr"`
	AutoCalcMass    bool        `xml:"auto-calc-mass,attr"`
	AutoCalcCG      bool        `xml:"auto-calc-cg,attr"`
	AverageThrust   float32     `xml:"avgThrust,attr"`
	PeakThrust      float32     `xml:"peakThrust,attr"`
	ThroatDiameter  float32     `xml:"throatDia,attr"`
	ExitDiameter    float32     `xml:"exitDia,attr"`
	TotalImpulse    float32     `xml:"Itot,attr"`
	BurnTime        float32     `xml:"burn-time,attr"`
	MassFraction    float32     `xml:"massFrac,attr"`
	SpecificImpulse float32     `xml:"Isp,attr"`
	TimeDivision    float32     `xml:"tDiv,attr"`
	TimeStep        float32     `xml:"tStep,attr"`
	TimeFix         float32     `xml:"tFix,attr"`
	ForceDivision   float32     `xml:"FDiv,attr"`
	ForceStep       float32     `xml:"FStep,attr"`
	ForceFix        float32     `xml:"FFix,attr"`
	MassDivision    float32     `xml:"mDiv,attr"`
	MassStep        float32     `xml:"mStep,attr"`
	MassFix         float32     `xml:"mFix,attr"`
	CGDivision      float32     `xml:"cgDiv,attr"`
	CGStep          float32     `xml:"cgStep,attr"`
	CGFix           float32     `xml:"cgFix,attr"`
	Comment         string      `xml:"comments"`
	Data            []DataPoint `xml:"data>eng-data"`
}

// UnmarshalXMLAttr unmarshals an engine type from an XML attribute
func (t *EngineType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "single-use":
		*t = SingleUse
		break
	case "reloadable":
		*t = Reloadable
		break
	case "hybrid":
		*t = Hybrid
		break
	default:
		return errors.New("Invalid enumeration constant")
	}
	return nil
}

// UnmarshalXMLAttr unmarshals an engine type from an XML attribute
func (d *Delays) UnmarshalXMLAttr(attr xml.Attr) error {
	arr := strings.Split(attr.Value, ",")
	*d = make([]float32, len(arr))
	for i, x := range arr {
		f, err := strconv.ParseFloat(x, 32)
		if err != nil {
			return err
		}
		(*d)[i] = float32(f)
	}
	return nil
}

// String converts the EngineType to a string
func (t EngineType) String() string {
	switch t {
	case SingleUse:
		return "Single use"
	case Reloadable:
		return "Reloadable"
	case Hybrid:
		return "Hybrid"
	}
	panic(errors.New("Invalid enumeration constant"))
}

// String converts the Delays to a string
func (d Delays) String() string {
	str := "{"
	for i, x := range d {
		if i > 0 {
			str = str + ", "
		}
		str = str + strconv.FormatFloat(float64(x), 'f', -1, 32)
	}
	return str + "}"
}

// String converts the DataPoint to a string
func (data DataPoint) String() string {
	return fmt.Sprintf("(Time = %f, Force = %f, Mass = %f, CG = %f)", data.Time, data.Force, data.Mass, data.CG)
}

// String converts the EngineModel to a string
func (eng EngineModel) String() string {
	return fmt.Sprintf("Manufacturer = %s, Code = %s, Type = %s, Diameter = %f, Length = %f, Initial Mass = %f, "+
		"Propellant Mass = %f, Delays = %s, Auto Calc Mass = %t, Auto Calc CG = %t, Average Thrust = %f, "+
		"Peak Thrust = %f, Throat Diameter = %f, Exit Diameter = %f, Total Impulse = %f, Burn Time = %f, "+
		"Mass Fraction = %f, Specific Impulse = %f, Time Division = %f, Time Step = %f, Time Fix = %f, "+
		"Force Division = %f, Force Step = %f, Force Fix = %f, Mass Division = %f, Mass Step = %f, "+
		"Mass Fix = %f, CG Division = %f, CG Step = %f, CG Fix = %f, Comment = %s, Data = %s",
		eng.Manufacturer, eng.Code, eng.Type, eng.Diameter, eng.Length, eng.InitialMass, eng.PropellantMass,
		eng.Delays, eng.AutoCalcMass, eng.AutoCalcCG, eng.AverageThrust, eng.PeakThrust, eng.ThroatDiameter,
		eng.ExitDiameter, eng.TotalImpulse, eng.BurnTime, eng.MassFraction, eng.SpecificImpulse, eng.TimeDivision,
		eng.TimeStep, eng.TimeFix, eng.ForceDivision, eng.ForceStep, eng.ForceFix, eng.MassDivision, eng.MassStep,
		eng.MassFix, eng.CGDivision, eng.CGStep, eng.CGFix, eng.Comment, eng.Data)
}
