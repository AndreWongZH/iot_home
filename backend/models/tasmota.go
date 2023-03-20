package models

// to be populated in the future
type innerStatus struct {
	DeviceName string `json:"DeviceName"`
	Power      int    `json:"Power"`
}

type TasmotaStatus struct {
	Status innerStatus `json:"Status"`
}
