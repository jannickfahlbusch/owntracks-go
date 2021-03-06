// Code generated by "stringer -type=BatteryStatus -linecomment -trimprefix BatteryStatus"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BatteryStatusUnknown-0]
	_ = x[BatteryStatusUnplugged-1]
	_ = x[BatteryStatusCharging-2]
	_ = x[BatteryStatusFull-3]
}

const _BatteryStatus_name = "UnknownUnpluggedChargingFull"

var _BatteryStatus_index = [...]uint8{0, 7, 16, 24, 28}

func (i BatteryStatus) String() string {
	if i < 0 || i >= BatteryStatus(len(_BatteryStatus_index)-1) {
		return "BatteryStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BatteryStatus_name[_BatteryStatus_index[i]:_BatteryStatus_index[i+1]]
}
