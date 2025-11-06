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

// Start begins the routine that fetches the latest version every hour
func (vm *VersionManager) Start() {
	// Fetch immediately on start
	if err := vm.fetchLatestVersion(); err != nil {
		log.Printf("Error fetching initial version: %v", err)
	}

	// Start the ticker for hourly updates
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			if err := vm.fetchLatestVersion(); err != nil {
				log.Printf("Error fetching version: %v", err)
			}
		}
	}()
}

// fetchLatestVersion retrieves the latest version from the DDragon API
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

// GetLatestVersion returns the currently stored latest version
func (vm *VersionManager) GetLatestVersion() string {
	vm.mu.RLock()
	defer vm.mu.RUnlock()
	return vm.latestVersion
}
