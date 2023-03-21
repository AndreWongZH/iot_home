package tasmota

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/AndreWongZH/iothome/models"
)

// this package contains structs related to tasmota interfaces

// to be populated in the future
type innerStatus struct {
	DeviceName string `json:"DeviceName"`
	Power      int    `json:"Power"`
}

type Status struct {
	InnerStatus innerStatus `json:"Status"`
}

func QueryTasmotaStatus(ipAddr string) models.DeviceStatus {
	devStatus := models.DeviceStatus{Connected: false, On_state: false}

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	var tasmotaStatus Status

	resp, err := client.Get("http://" + ipAddr + "/cm?cmnd=Status")
	if err != nil {
		return devStatus
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return devStatus
	}

	err = json.Unmarshal(body, &tasmotaStatus)
	if err != nil {
		return devStatus
	}

	devStatus.Connected = true

	if tasmotaStatus.InnerStatus.Power == 1 {
		devStatus.On_state = true
	} else {
		devStatus.On_state = false
	}

	return devStatus
}
