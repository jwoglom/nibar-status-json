package battery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Charging(t *testing.T) {
	text := `Now drawing from 'AC Power'
 -InternalBattery-0 (id=9961571)	43%; charging; 2:08 remaining present: true`

	assert.Equal(t, BatteryDetails{
		BatteryPercent: "43%",
		State:          Charging,
		RemainingTime:  "2:08",
	}, Parse(text))
}

func TestParse_DrawingFromBatteryAndCharging(t *testing.T) {
	text := `Now drawing from 'Battery Power'
 -InternalBattery-0 (id=9961571)	43%; charging; 0:00 remaining present: true`

	assert.Equal(t, BatteryDetails{
		BatteryPercent: "43%",
		State:          Charging,
		RemainingTime:  "0:00",
	}, Parse(text))
}

func TestParse_DischargingWithNoEstimate(t *testing.T) {
	text := `Now drawing from 'Battery Power'
 -InternalBattery-0 (id=9961571)	44%; discharging; (no estimate) present: true`

	assert.Equal(t, BatteryDetails{
		BatteryPercent: "44%",
		State:          Discharging,
		RemainingTime:  "",
	}, Parse(text))
}

func TestParse_Discharging(t *testing.T) {
	text := `Now drawing from 'Battery Power'
 -InternalBattery-0 (id=9961571)	42%; discharging; 0:41 remaining present: true`

	assert.Equal(t, BatteryDetails{
		BatteryPercent: "42%",
		State:          Discharging,
		RemainingTime:  "0:41",
	}, Parse(text))
}

func TestParse_NotCharging(t *testing.T) {
	text := `Now drawing from 'AC Power'
 -InternalBattery-0 (id=9961571)	40%; AC attached; not charging present: true`

	assert.Equal(t, BatteryDetails{
		BatteryPercent: "40%",
		State:          NotCharging,
		RemainingTime:  "",
	}, Parse(text))
}
