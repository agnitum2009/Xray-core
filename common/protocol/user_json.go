package protocol

import "encoding/json"

type jsonUser struct {
	Level uint32 `json:"level"`
	Email string `json:"email"`

	SpeedLimit          uint64 `json:"speed_limit"`
	SpeedLimitCamel     uint64 `json:"speedLimit"`
	DeviceLimit         uint32 `json:"device_limit"`
	DeviceLimitCamel    uint32 `json:"deviceLimit"`
	UpSpeedLimit        uint64 `json:"up_speed_limit"`
	UpSpeedLimitCamel   uint64 `json:"upSpeedLimit"`
	DownSpeedLimit      uint64 `json:"down_speed_limit"`
	DownSpeedLimitCamel uint64 `json:"downSpeedLimit"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	var raw jsonUser
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	u.Level = raw.Level
	u.Email = raw.Email
	u.SpeedLimit = raw.SpeedLimit
	if u.SpeedLimit == 0 {
		u.SpeedLimit = raw.SpeedLimitCamel
	}
	u.DeviceLimit = raw.DeviceLimit
	if u.DeviceLimit == 0 {
		u.DeviceLimit = raw.DeviceLimitCamel
	}
	u.UpSpeedLimit = raw.UpSpeedLimit
	if u.UpSpeedLimit == 0 {
		u.UpSpeedLimit = raw.UpSpeedLimitCamel
	}
	u.DownSpeedLimit = raw.DownSpeedLimit
	if u.DownSpeedLimit == 0 {
		u.DownSpeedLimit = raw.DownSpeedLimitCamel
	}
	return nil
}
