package internal

import (
	"os"
	"time"
)

const defaultTimeout = time.Second

var wmScripts = os.Getenv("WMSCRIPTS")

var workFuncs = []workFunc{
	func(s *Status) {
		s.DateTime.Time = time.Now().Format("03:04 PM")
	},
	func(s *Status) {
		s.DateTime.Date = time.Now().Format("Mon")
	},
	func(s *Status) {
		lines := runlines(`
BATTERY_PERCENTAGE=$(pmset -g batt | egrep '([0-9]+\%).*' -o --colour=auto | cut -f1 -d'%')
BATTERY_STATUS=$(pmset -g batt | grep "'.*'" | sed "s/'//g" | cut -c 18-19)
BATTERY_REMAINING=$(pmset -g batt | egrep -o '([0-9]+%).*' | sed 's/attached; //' | cut -d\  -f3)

if [[ "$BATTERY_REMAINING" == "(no" ]]; then
    BATTERY_REMAINING=""
fi


if [[ "$BATTERY_REMAINING" == "charge;" ]]; then
    BATTERY_REMAINING=""
fi

if [[ "$BATTERY_REMAINING" == "not" ]]; then
    BATTERY_REMAINING="not charging"
fi

BATTERY_CHARGING=""
if [ "$BATTERY_STATUS" == "Ba" ]; then
  BATTERY_CHARGING="false"
elif [ "$BATTERY_STATUS" == "AC" ]; then
  BATTERY_CHARGING="true"
fi

echo $BATTERY_PERCENTAGE
echo $BATTERY_CHARGING
echo $BATTERY_REMAINING
`, defaultTimeout)
		if len(lines) == 3 {
			s.Battery.Percentage = lines[0]
			s.Battery.Charging = lines[1]
			s.Battery.Remaining = lines[2]
		}
	},
	func(s *Status) {
		s.Cpu.LoadAverage = "<unused>"
	},
	func(s *Status) {
		lines := runlines(`
WIFI_INTERFACE=en0
WIFI_ACTIVE_INTERFACE=$(route get 8.8.8.8 2>/dev/null | grep interface | cut -c 14-)
WIFI_STATUS=$(ifconfig $WIFI_INTERFACE | grep status | cut -c 10-)
WIFI_SSID=$(networksetup -getairportnetwork $WIFI_INTERFACE | cut -c 24-)

echo $WIFI_STATUS
echo $WIFI_SSID
echo $WIFI_ACTIVE_INTERFACE
echo $WIFI_INTERFACE
`, defaultTimeout)
		if len(lines) == 4 {
			s.Wifi.Status = lines[0]
			s.Wifi.SSID = lines[1]
			s.Wifi.ActiveInterface = lines[2]
			s.Wifi.WifiInterface = lines[3]
		}
	},
	func(s *Status) {
		s.Vpn.Tunnelblick = run(wmScripts+`/vpn_tunnelblick_status.sh`, defaultTimeout)
	},
	func(s *Status) {
		s.Vpn.PulseSecure = run(wmScripts+`/vpn_pulsesecure_status.sh`, defaultTimeout)
	},
	func(s *Status) {
		s.Bluetooth.On = run(`blueutil -p`, defaultTimeout)
	},
	func(s *Status) {
		out := run(`
blueutil --paired --format json 2> /dev/null | jq 'map(select(.connected == true))' 2> /dev/null || echo '[]'
`, defaultTimeout)
		s.Bluetooth.Paired = unmarshalJsonArray(out)
	},
	func(s *Status) {
		s.Audio.Input = run(`SwitchAudioSource -c -t input`, defaultTimeout)
	},
	func(s *Status) {
		s.Audio.Output = run(`SwitchAudioSource -c -t output`, defaultTimeout)
	},
	func(s *Status) {
		s.Audio.Muted = run(`osascript -e "output muted of (get volume settings)"`, defaultTimeout)
	},
	func(s *Status) {
		s.Dnd = run(`defaults -currentHost read com.apple.notificationcenterui doNotDisturb`, defaultTimeout)
	},
	func(s *Status) {
		run(`
if [ ! -f ~/.cache/cgm.json ]; then
    $WMSCRIPTS/update_cgm.sh
else
    $WMSCRIPTS/update_cgm_check.sh
fi
`, defaultTimeout)
		out := run(`cat ~/.cache/cgm.json 2> /dev/null || echo "{}"`, 100*time.Millisecond)
		s.Cgm = unmarshalJson(out)
	},
}
