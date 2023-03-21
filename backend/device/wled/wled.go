package wled

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/AndreWongZH/iothome/models"
)

// this package contains structs related to wled interfaces

type Segment struct {
	Id    int     `json:"id"`
	Start int     `json:"start"`
	Stop  int     `json:"stop"`
	Len   int     `json:"len"`
	Grp   int     `json:"grp"`
	Spc   int     `json:"spc"`
	Fx    int     `json:"fx"`
	Sx    int     `json:"sx"`
	Ix    int     `json:"ix"`
	Pal   int     `json:"pal"`
	Col   [][]int `json:"col"`
}

type State struct {
	On         bool      `json:"on"`
	Bri        int       `json:"bri"`
	Transition int       `json:"transition"`
	Seg        []Segment `json:"seg"`
}

type WledConfig struct {
	State State `json:"state"`
	// Info     interface{}
	Effects  []string `json:"effects"`
	Palettes []string `json:"palettes"`
}

type WledSwitch struct {
	On bool `json:"on"`
}

// query status of wled devices, if no connection, no error is thrown
func QueryWledStatus(ipAddr string) models.DeviceStatus {
	devStatus := models.DeviceStatus{Connected: false, On_state: false}
	var wledConfig WledConfig

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := client.Get("http://" + ipAddr + "/json")
	if err != nil {
		return devStatus
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return devStatus
	}

	err = json.Unmarshal(body, &wledConfig)
	if err != nil {
		return devStatus
	}

	devStatus.Connected = true

	if wledConfig.State.On {

		devStatus.On_state = true
	} else {
		devStatus.On_state = false
	}

	return devStatus
}
