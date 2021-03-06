package domain

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	// NONE const
	NONE string = "NONE"
	// MODE const
	MODE string = "MODE"
	// RESOLUTION const
	RESOLUTION string = "RESOLUTION"
	// SOURCE const
	SOURCE string = "SOURCE"
	// BRIGHTNESS const
	BRIGHTNESS string = "BRIGHTNESS"
	// CONTRAST const
	CONTRAST string = "CONTRAST"
	// L const
	L string = "L"
	// T const
	T string = "T"
	// X const
	X string = "X"
	// Y const
	Y string = "Y"
)

// Feature represents a capability of the scanner device
type Feature struct {
	Type    string   `json:"type"`
	Ranged  bool     `json:"ranged"`
	Values  []string `json:"values"`
	Default string   `json:"default"`
	ToUse   string   `json:"to_use"`
}

// Features represent a set of Feature
type Features map[string]*Feature

// Device represents a scanner device
type Device struct {
	Name string   `json:"name"`
	Ft   Features `json:"features"`
}

// Devices represent a set of Device
type Devices []Device

func parseValues(values string, def string) ([]string, string, bool) {
	var parsedValues []string
	var ranged bool

	if strings.Contains(values, "..") {
		ranged = true
		parsedValues = strings.Split(values, "..")
		for k, v := range parsedValues {
			if idx := strings.Index(v, "."); idx != -1 {
				parsedValues[k] = v[0:idx]
			}
		}
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

	if idx := strings.Index(def, "."); idx != -1 {
		def = def[0:idx]
	}

	return parsedValues, def, ranged
}

// ParseFeature parses a feature given three parameters
func (d *Device) ParseFeature(name string, values string, def string) *Feature {
	var f *Feature

	switch strings.TrimLeft(name, "-") {
	case "mode":
		f = &Feature{Type: MODE}
	case "resolution":
		f = &Feature{Type: RESOLUTION}
	case "source":
		f = &Feature{Type: SOURCE}
	case "brightness":
		f = &Feature{Type: BRIGHTNESS}
	case "contrast":
		f = &Feature{Type: CONTRAST}
	case "l":
		f = &Feature{Type: L}
	case "t":
		f = &Feature{Type: T}
	case "x":
		f = &Feature{Type: X}
	case "y":
		f = &Feature{Type: Y}
	}

	if f != nil {
		values, def, ranged := parseValues(values, def)

		if f.Type == MODE {
			var newValues []string
			for _, v := range values {
				if v != "Lineart" {
					newValues = append(newValues, v)
				}
			}

			if def == "Lineart" {
				def = newValues[0]
			}

			values = newValues
		}

		f.Ranged = ranged
		f.Values = values
		f.Default = def
		f.ToUse = def
	}

	return f
}
