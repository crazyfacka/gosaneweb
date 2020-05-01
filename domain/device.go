package domain

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	// NONE iota
	NONE = iota
	// MODE iota
	MODE
	// RESOLUTION iota
	RESOLUTION
	// SOURCE iota
	SOURCE
	// BRIGHTNESS iota
	BRIGHTNESS
	// CONTRAST iota
	CONTRAST
	// L iota
	L
	// T iota
	T
	// X iota
	X
	// Y iota
	Y
)

// Feature represents a capability of the scanner device
type Feature struct {
	Type    int
	Ranged  bool
	Values  []string
	Default string
}

// Features represent a set of Feature
type Features []Feature

// Device represents a scanner device
type Device struct {
	Name string
	Ft   Features
}

// Devices represent a set of Device
type Devices []Device

func parseValues(values string, def string) ([]string, bool) {
	var parsedValues []string
	var ranged bool

	if strings.Contains(values, "..") {
		ranged = true
		parsedValues = strings.Split(values, "..")
	} else {
		ranged = false
		parsedValues = strings.Split(values, "|")
	}
	_, err := strconv.ParseFloat(def, 32)
	if err == nil {
		re := regexp.MustCompile(`[0-9\.]+`)
		filtered := re.FindString(parsedValues[len(parsedValues)-1])
		parsedValues[len(parsedValues)-1] = filtered
	}

	return parsedValues, ranged
}

// ParseFeature parses a feature given three parameters
func (d *Device) ParseFeature(name string, values string, def string) Feature {
	var f Feature

	switch strings.TrimLeft(name, "-") {
	case "mode":
		f = Feature{Type: MODE}
	case "resolution":
		f = Feature{Type: RESOLUTION}
	case "source":
		f = Feature{Type: SOURCE}
	case "brightness":
		f = Feature{Type: BRIGHTNESS}
	case "contrast":
		f = Feature{Type: CONTRAST}
	case "l":
		f = Feature{Type: L}
	case "t":
		f = Feature{Type: T}
	case "x":
		f = Feature{Type: X}
	case "y":
		f = Feature{Type: Y}
	}

	if f.Type > NONE {
		vals, ranged := parseValues(values, def)

		f.Ranged = ranged
		f.Values = vals
		f.Default = def
	}

	return f
}
