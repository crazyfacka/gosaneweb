package repository

// Scan interface that represent all that's possible with SANEd
type Scan interface {
	Devices() []string
	Features() []string
}
