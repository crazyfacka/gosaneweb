package repository

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

		if len(deviceMatches) >= 2 {
			device := domain.Device{
				Name: deviceMatches[1],
			}

			device.Ft = make(map[string]*domain.Feature)

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
	}

	if si.devices == nil {
		log.Info().Msg("Loading device info from cache")
		dat, err := ioutil.ReadFile("devices.json")
		if err != nil {
			log.Error().Err(err).Msg("Error reading devices from file")
		}

		err = json.Unmarshal(dat, &si.devices)
		if err != nil {
			log.Error().Err(err).Msg("Error unmarshalling data")
		}
	} else {
		jsonDevices, err := json.MarshalIndent(si.devices, "", "  ")
		if err != nil {
			log.Error().Err(err).Msg("Error marshalling devices data to JSON")
		}

		err = ioutil.WriteFile("devices.json", jsonDevices, 0644)
		if err != nil {
			log.Error().Err(err).Msg("Error writing devices data to output file")
		}
	}

	return si.devices
}

// Scan scans an image
func (si *ScanImage) Scan(device domain.Device) ([]byte, error) {
	var out bytes.Buffer
	var args []string

	log.Info().Str("driver", "ScanImage").Msg("Scanning image")

	if v, ok := device.Ft[domain.MODE]; ok {
		args = append(args, "--mode", v.ToUse)
	}

	if v, ok := device.Ft[domain.SOURCE]; ok {
		args = append(args, "--source", v.ToUse)
	}

	if v, ok := device.Ft[domain.RESOLUTION]; ok {
		args = append(args, "--resolution", v.ToUse)
	}

	if v, ok := device.Ft[domain.BRIGHTNESS]; ok {
		args = append(args, "--brightness", v.ToUse)
	}

	if v, ok := device.Ft[domain.CONTRAST]; ok {
		args = append(args, "--contrast", v.ToUse)
	}

	if v, ok := device.Ft[domain.L]; ok {
		args = append(args, "-l", v.ToUse)
	}

	if v, ok := device.Ft[domain.T]; ok {
		args = append(args, "-t", v.ToUse)
	}

	if v, ok := device.Ft[domain.X]; ok {
		args = append(args, "-x", v.ToUse)
	}

	if v, ok := device.Ft[domain.Y]; ok {
		args = append(args, "-y", v.ToUse)
	}

	args = append(args, "--format", "png")

	log.Debug().Str("cmd", si.binary).Strs("args", args).Msg("Executing")

	cmd := exec.Command(si.binary, args...)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Error().Strs("params", args).Err(err).Msg("Error doing scan")
		return nil, err
	}

	log.Debug().Int("bytes", out.Len()).Msg("Scan complete")
	return out.Bytes(), nil
}
