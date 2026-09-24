package disasterkit

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type DisasterManifest struct {
	KitID           string   `json:"kit_id"`
	Hostname        string   `json:"hostname"`
	CreatedAt       string   `json:"created_at"`
	KenPanelVersion string   `json:"kenpanel_version"`
	RestoreSteps    []string `json:"restore_steps"`
}

func GenerateKit(hostname string) *DisasterManifest {
	return &DisasterManifest{
		KitID:           fmt.Sprintf("kit-%s-%d", hostname, time.Now().Unix()),
		Hostname:        hostname,
		CreatedAt:       time.Now().UTC().Format(time.RFC3339),
		KenPanelVersion: "2.3.0",
		RestoreSteps: []string{
			"1. Boot RecoveryOS or provision bare-metal target",
			"2. Install KenPanel agent",
			"3. Restore database engines and imports",
			"4. Reconstruct web server virtual hosts and SSL",
			"5. Synchronize document roots",
		},
	}
}

func SaveManifest(m *DisasterManifest, path string) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
