package internal

type Cmd = string

// Status represents the JSON used for status output in nibar
// type Status = map[string]interface{}
type Status struct {
	DateTime  *DateTime  `json:"datetime"`
	Battery   *Battery   `json:"battery"`
	Cpu       *Cpu       `json:"cpu"`
	Wifi      *Wifi      `json:"wifi"`
	Vpn       *Vpn       `json:"vpn"`
	Bluetooth *Bluetooth `json:"bluetooth"`
	Audio     *Audio     `json:"audio"`
	Dnd       Dnd        `json:"dnd"`
	Cgm       Cgm        `json:"cgm"`
}

func NewStatus() *Status {
	return &Status{
		DateTime:  &DateTime{},
		Battery:   &Battery{},
		Cpu:       &Cpu{},
		Wifi:      &Wifi{},
		Vpn:       &Vpn{},
		Bluetooth: &Bluetooth{},
		Audio:     &Audio{},
	}
}

type DateTime struct {
	Time string `json:"time"`
	Date string `json:"date"`
}

type Battery struct {
	Percentage string `json:"percentage"`
	Charging   string `json:"charging"`
	Remaining  string `json:"remaining"`
}

type Cpu struct {
	LoadAverage string `json:"loadAverage"`
}

type Wifi struct {
	Status          string `json:"status"`
	SSID            string `json:"ssid"`
	ActiveInterface string `json:"active_interface"`
	WifiInterface   string `json:"wifi_interface"`
}

type Vpn struct {
	Tunnelblick string `json:"tunnelblick"`
	PulseSecure string `json:"pulsesecure"`
}

type Bluetooth struct {
	On     string        `json:"on"`
	Paired []interface{} `json:"paired"`
}

type Audio struct {
	Input  string `json:"input"`
	Output string `json:"output"`
	Muted  string `json:"muted"`
}

type Dnd = string
type Cgm = map[string]interface{}
