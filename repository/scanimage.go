package repository

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/crazyfacka/gosaneweb/domain"
	"github.com/rs/zerolog/log"
)

// ScanImage represents the struct of the scanimage binary handler
type ScanImage struct {
	binary   string
	devices  []string
	features []string
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
func (si *ScanImage) Devices() []string {
	var out bytes.Buffer

	if si.devices == nil {
		cmd := exec.Command(si.binary, "-A")
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			log.Error().Err(err).Msg("Error getting scan information")
		}

		fmt.Println(out.String())
	}

	return nil
}

// Features returns all available features for a given device
func (si *ScanImage) Features() []string {
	return nil
}
