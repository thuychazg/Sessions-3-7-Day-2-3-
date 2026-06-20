package model

import "time"

type Asset struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateAssetRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type BatchCreateRequest struct {
	Assets []CreateAssetRequest `json:"assets"`
}

type ScanJob struct {
	ID        string      `json:"id"`
	AssetID   string      `json:"asset_id"`
	ScanType  string      `json:"scan_type"`
	Status    string      `json:"status"`
	Error     string      `json:"error"`
	Results   interface{} `json:"results"`
	CreatedAt time.Time   `json:"created_at"`
}
