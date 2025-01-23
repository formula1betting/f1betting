package race_info

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Lap struct {
	DriverNumber    int     `json:"driver_number"`
	LapNumber       int     `json:"lap_number"`
	LapDuration     float64 `json:"lap_duration"`
	DurationSector1 float64 `json:"duration_sector_1"`
	DurationSector2 float64 `json:"duration_sector_2"`
	DurationSector3 float64 `json:"duration_sector_3"`
}

func GetFastestDriverEachLap(sessionKey int, lapNumber int) (*Lap, error) {
	url := fmt.Sprintf("https://api.openf1.org/v1/laps?session_key=%v", sessionKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var laps []Lap
	if err := json.NewDecoder(resp.Body).Decode(&laps); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	if len(laps) == 0 {
		return nil, fmt.Errorf("no laps found")
	}

	var fastestLap *Lap
	for _, lap := range laps {
		if lap.LapNumber == lapNumber && (fastestLap == nil || lap.LapDuration < fastestLap.LapDuration) && lap.LapDuration > 0 {
			fastestLap = &lap
		}
	}

	if fastestLap == nil {
		return nil, fmt.Errorf("no laps found for lap number %d", lapNumber)
	}

	return fastestLap, nil
}

func GetFastestDriverWholeRace(sessionKey int) (*Lap, error) {
	url := fmt.Sprintf("https://api.openf1.org/v1/laps?session_key=%v", sessionKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var laps []Lap
	if err := json.NewDecoder(resp.Body).Decode(&laps); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	if len(laps) == 0 {
		return nil, fmt.Errorf("no laps found")
	}

	fastestLap := Lap{LapDuration: 1e9} // Set initial fastest lap duration to a very high value (infinity)
	for _, lap := range laps {
		if lap.LapDuration < fastestLap.LapDuration && lap.LapDuration > 0 {
			fastestLap = lap
		}
	}

	return &fastestLap, nil
}
