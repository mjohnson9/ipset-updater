package fetcher

import (
	"log"
	"os"
	"path"
	"sync"

	"github.com/adrg/xdg"

	"github.com/mjohnson9/ipset-updater/internal/config"
)

func Fetch(list *config.List, config *config.Config, wg *sync.WaitGroup) {
	defer wg.Done()

	logger := log.New(os.Stdout, "fetcher: "+list.Name+": ", log.Flags())

	cachedFile, err := checkCachedFile(list, config)
	if err != nil {
		logger.Printf("failed to get cached file: %s", err)
		return
	}
	defer cachedFile.Close()
}

func Prepare(config *config.Config) error {
	return nil
}

func checkCachedFile(list *config.List, config *config.Config) (*os.File, error) {
	cachedPath, err := xdg.SearchCacheFile(path.Join("ipset-updater", list.Name+".txt"))
	if err != nil {
		// No cached file could be found
		return nil, nil
	}

	f, err := os.OpenFile(cachedPath, os.O_RDWR, 0640)
	if err != nil {
		return nil, err
	}

	return f, nil
}
