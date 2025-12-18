package global

import (
	"encoding/json"
	"os"
	"slg_sever/internal/config"
)

func InitMarchUnit() *config.MarchData {
	jsonData, err := os.ReadFile("data/match_data.json")
	if err != nil {
		panic(err)
	}

	var data config.MarchData
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		panic(err)
	}

	return &data
}
