package domain

import (
	"flag"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Confs variable holding all the loaded configurations
var Confs *c

type c struct {
	Debug           bool
	Port            int
	ScanImageBinary string
}

// LoadConfiguration loads all configs (flags and file) to memory
func LoadConfiguration() error {
	debug := flag.Bool("debug", false, "sets log level to debug and enables /test route")
	flag.Parse()

	Confs = &c{
		Debug: *debug,
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Info().Msg("Loading gosaneweb")

	viper.SetConfigName(".gosaneweb")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	log.Debug().Interface(".gosaneweb", viper.AllSettings()).Msg("Loaded configuration")

	Confs.Port = viper.GetInt("port")
	Confs.ScanImageBinary = viper.GetString("scanimage")

	return nil
}
