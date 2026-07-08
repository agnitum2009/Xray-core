package protocol

import (
	"encoding/json"
	"testing"
)

func TestUserUnmarshalJSONSpeedLimitAliases(t *testing.T) {
	var user User
	if err := json.Unmarshal([]byte(`{"email":"client@example.com","level":1,"speedLimit":3,"deviceLimit":4,"upSpeedLimit":5,"downSpeedLimit":6}`), &user); err != nil {
		t.Fatal(err)
	}
	if user.Email != "client@example.com" || user.Level != 1 || user.SpeedLimit != 3 || user.DeviceLimit != 4 || user.UpSpeedLimit != 5 || user.DownSpeedLimit != 6 {
		t.Fatalf("unexpected user: %+v", user)
	}
}

func TestUserUnmarshalJSONSnakeCaseSpeedLimits(t *testing.T) {
	var user User
	if err := json.Unmarshal([]byte(`{"email":"client@example.com","speed_limit":3,"device_limit":4,"up_speed_limit":5,"down_speed_limit":6}`), &user); err != nil {
		t.Fatal(err)
	}
	if user.SpeedLimit != 3 || user.DeviceLimit != 4 || user.UpSpeedLimit != 5 || user.DownSpeedLimit != 6 {
		t.Fatalf("unexpected limits: %+v", user)
	}
}
