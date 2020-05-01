package domain

import "fmt"

const (
	// MODE iota
	MODE = iota
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

// ParseFeature parses a feature given three parameters
func (d *Device) ParseFeature(name string, values string, def string) Feature {
	fmt.Println(name, values, def)
	return Feature{}
}
