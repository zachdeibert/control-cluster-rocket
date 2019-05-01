package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const cutset = " \r\n\t"

// FlightProfile represents the sensor data from a flight
type FlightProfile struct {
	Fields []string
	Data   [][]float64
}

// LoadFlightProfile opens a file with the given file name and reads it into memory
func LoadFlightProfile(filename string) (FlightProfile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return FlightProfile{}, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return FlightProfile{}, err
	}
	lines := strings.Split(string(data), "\n")
	header := strings.Split(lines[0], ",")
	profile := FlightProfile{
		Fields: make([]string, len(header)),
		Data:   make([][]float64, len(lines)-1),
	}
	for i, field := range header {
		profile.Fields[i] = strings.Trim(field, cutset)
	}
	for i := range profile.Data {
		if len(strings.Trim(lines[i+1], cutset)) == 0 {
			continue
		}
		d := strings.Split(lines[i+1], ",")
		if len(d) != len(profile.Fields) {
			return FlightProfile{}, fmt.Errorf("Invalid number of cells in row %d", i+1)
		}
		floats := make([]float64, len(d))
		for ix, s := range d {
			if floats[ix], err = strconv.ParseFloat(strings.Trim(s, cutset), 64); err != nil {
				return FlightProfile{}, err
			}
		}
		profile.Data[i] = floats
	}
	return profile, nil
}

// NumDataPoints returns the number of data points contained in the profile
func (profile FlightProfile) NumDataPoints() int {
	return len(profile.Data)
}

// ContainsField determines if a given field is contained within the profile data
func (profile FlightProfile) ContainsField(field string) bool {
	for _, f := range profile.Fields {
		if f == field {
			return true
		}
	}
	return false
}

// GetData gets the data at a specific data point
func (profile FlightProfile) GetData(field string, dataPoint int) float64 {
	var idx = -1
	for i, f := range profile.Fields {
		if f == field {
			idx = i
			break
		}
	}
	if idx < 0 {
		return -1
	}
	return profile.Data[dataPoint][idx]
}

// String converts a flight profile to a string representation
func (profile FlightProfile) String() string {
	return fmt.Sprintf("Flight Profile of {%s} with %d data points", strings.Join(profile.Fields, ", "), len(profile.Data))
}
