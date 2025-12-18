package config

type MapNode struct {
	ID    int64   `json:"id"`
	Links []int64 `json:"links"`
}

type MapConfig struct {
	Capital MapNode   `json:"capital"`
	Large   []MapNode `json:"large"`
	Medium  []MapNode `json:"medium"`
	HQ      []MapNode `json:"hq"`
	Small   []MapNode `json:"small"`
}
