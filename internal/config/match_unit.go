package config

type MarchUnit struct {
	UnitID    string  `json:"unit_id"`
	PlayerID  int64   `json:"player_id"`
	TeamID    int64   `json:"team_id"`
	Path      []int64 `json:"path"`
	StartTick int64   `json:"start_tick"`
	Speed     int64   `json:"speed"` // tick per node
}

type MarchData struct {
	MarchUnits []MarchUnit `json:"march_units"`
}
