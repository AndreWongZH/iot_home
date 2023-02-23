package wled

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
	State interface{} `json:"state"`
	// Info     interface{}
	Effects  []string `json:"effects"`
	Palettes []string `json:"palettes"`
}

type WledSwitch struct {
	On bool `json:"on"`
}
