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
		log.Info().Str("driver", "ScanImage").Msg("Searching for devices")

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

		device.Ft = make(map[int]*domain.Feature)

		for _, m := range featureMatches {
			feature := device.ParseFeature(m[1], m[2], m[3])
			if feature != nil {
				device.Ft[feature.Type] = feature
			}
		}

		si.devices = append(si.devices, device)

		log.Info().Int("count", len(deviceMatches)-1).Msg("Found devices")
		log.Debug().Interface("data", si.devices).Msg("Device data")
	}

	return si.devices
}

// Scan scans an image
func (si *ScanImage) Scan(device domain.Device) ([]byte, error) {
	var out bytes.Buffer

	log.Info().Str("driver", "ScanImage").Msg("Scanning image")

	args := []string{
		"--mode", device.Ft[domain.MODE].ToUse,
		"--resolution", device.Ft[domain.RESOLUTION].ToUse,
		"--brightness", device.Ft[domain.BRIGHTNESS].ToUse,
		"--contrast", device.Ft[domain.CONTRAST].ToUse,
		"-l", device.Ft[domain.L].ToUse,
		"-t", device.Ft[domain.T].ToUse,
		"-x", device.Ft[domain.X].ToUse,
		"-y", device.Ft[domain.Y].ToUse,
		"--format", "png",
	}

	log.Debug().Str("cmd", si.binary).Strs("args", args).Msg("Executing")

	cmd := exec.Command(si.binary, args...)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Error().Strs("params", args).Err(err).Msg("Error doing scan")
		return nil, err
	}

	return out.Bytes(), nil
}
