package repository

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"

	"github.com/crazyfacka/gosaneweb/domain"
	"github.com/rs/zerolog/log"
)

// ScanImage represents the struct of the scanimage binary handler
type ScanImage struct {
	binary  string
	devices domain.Devices
}

// InitScanImage initializes the ScanImage handler
func InitScanImage() (*ScanImage, error) {
	si := &ScanImage{
		binary: domain.Confs.ScanImageBinary,
	}

	if _, err := os.Stat(si.binary); err != nil {
		return nil, err
	}

	return si, nil
}

// Devices returns all available devices
func (si *ScanImage) Devices() domain.Devices {
	var out bytes.Buffer

	if si.devices == nil {
		cmd := exec.Command(si.binary, "-A")
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			log.Error().Err(err).Msg("Error getting scan information")
		}

		output := out.String()

		devicesRe := regexp.MustCompile("All options specific to device `(.*)'")
		deviceMatches := devicesRe.FindStringSubmatch(output)

		featuresRe := regexp.MustCompile(`\s+([-]{1,2}[-a-zA-Z0-9]+) ?(.*) \[(.*)\]\n`)
		featureMatches := featuresRe.FindAllStringSubmatch(output, -1)

		device := domain.Device{
			Name: deviceMatches[1],
		}

		for _, m := range featureMatches {
			feature := device.ParseFeature(m[1], m[2], m[3])
			if feature.Type > domain.NONE {
				device.Ft = append(device.Ft, feature)
			}
		}

		si.devices = append(si.devices, device)
	}

	return si.devices
}

// Features returns all available features for a given device
func (si *ScanImage) Features() domain.Features {
	return nil
}
