package main

import (
	"flag"
	"io"
	"os"
)

type Flags struct {
	ConfigPath string
	Config     io.Reader
}

func (f *Flags) Validate() {
	var err error
	f.Config, err = os.OpenFile(f.ConfigPath, os.O_RDONLY, 0644)
	if err != nil {
		logger.Fatal(err)
	}
}

func ParseFlags() *Flags {
	flags := new(Flags)

	flag.StringVar(&flags.ConfigPath, "config", "ipset-updater.toml", "the configuration `file`")

	flag.Parse()

	flags.Validate()

	return flags
}
