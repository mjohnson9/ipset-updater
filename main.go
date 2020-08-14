package main

import (
	"log"
	"os"
	"sync"

	"github.com/mjohnson9/ipset-updater/internal/config"
	"github.com/mjohnson9/ipset-updater/internal/fetcher"
)

var logger = log.New(os.Stdout, "main: ", log.Flags())

func main() {
	flags := ParseFlags()

	config, err := config.Parse(flags.Config)
	if err != nil {
		panic(err)
	}

	fetcher.Prepare(config)

	var wg sync.WaitGroup

	for i := range config.Lists {
		wg.Add(1)
		fetcher.Fetch(&config.Lists[i], config, &wg)
	}

	wg.Wait()
}
