package routine

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const versionsURL = "https://ddragon.leagueoflegends.com/api/versions.json"

type VersionManager struct {
	latestVersion string
	mu            sync.RWMutex
	httpClient    *http.Client
}

func NewVersionManager(httpClient *http.Client) *VersionManager {
	return &VersionManager{
		httpClient:    httpClient,
		latestVersion: "",
	}
}

func (vm *VersionManager) Start() {
	if err := vm.fetchLatestVersion(); err != nil {
		log.Printf("Error fetching initial version: %v", err)
	}

	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			if err := vm.fetchLatestVersion(); err != nil {
				log.Printf("Error fetching version: %v", err)
			}
		}
	}()
}

func (vm *VersionManager) fetchLatestVersion() error {
	resp, err := vm.httpClient.Get(versionsURL)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var versions []string
	if err := json.NewDecoder(resp.Body).Decode(&versions); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if len(versions) == 0 {
		return fmt.Errorf("no versions returned from API")
	}

	vm.mu.Lock()
	vm.latestVersion = versions[0]
	vm.mu.Unlock()

	log.Printf("Updated API version to: %s", vm.latestVersion)
	return nil
}

func (vm *VersionManager) GetLatestVersion() string {
	vm.mu.RLock()
	defer vm.mu.RUnlock()
	return vm.latestVersion
}
