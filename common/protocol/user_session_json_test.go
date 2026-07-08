package protocol

import (
	"encoding/json"
	"testing"
)

func TestUserUnmarshalJSONSessionLimitAliases(t *testing.T) {
	var user User
	if err := json.Unmarshal([]byte(`{"email":"client@example.com","sessionLimit":12}`), &user); err != nil {
		t.Fatal(err)
	}
	if user.SessionLimit != 12 {
		t.Fatalf("expected camelCase session limit, got %d", user.SessionLimit)
	}

	if err := json.Unmarshal([]byte(`{"email":"client@example.com","session_limit":13}`), &user); err != nil {
		t.Fatal(err)
	}
	if user.SessionLimit != 13 {
		t.Fatalf("expected snake_case session limit, got %d", user.SessionLimit)
	}
}
