package repository

import "github.com/crazyfacka/gosaneweb/domain"

// Scan interface that represent all that's possible with SANEd
type Scan interface {
	Devices() domain.Devices
	Features() domain.Features
}
