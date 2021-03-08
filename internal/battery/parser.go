package battery

import (
	"strings"
)

type BatteryDetails struct {
	BatteryPercent string
	State          batteryState
	RemainingTime  string
}

type batteryState int

const (
	Charging batteryState = iota
	Discharging
	NotCharging
)

func Parse(pmsetOutput string) BatteryDetails {
	lines := strings.Split(pmsetOutput, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, " -") {
			parts := strings.Split(line, "\t")
			return parse(parts[1])
		}
	}
	return BatteryDetails{}
}

func parse(detailLine string) BatteryDetails {
	parts := strings.Split(detailLine, "; ")

	percent := ""
	if strings.Contains(parts[0], "%") {
		percent = parts[0]
	}

	var state batteryState
	if s := parts[1]; s == "charging" {
		state = Charging
	} else if s == "discharging" {
		state = Discharging
	} else if s == "AC attached" {
		state = NotCharging
	}

	remaining := ""
	if strings.Contains(parts[2], " remaining") {
		remaining = strings.Split(parts[2], " remaining")[0]
	}

	return BatteryDetails{
		BatteryPercent: percent,
		State:          state,
		RemainingTime:  remaining,
	}
}
