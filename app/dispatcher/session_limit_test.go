package dispatcher

import "testing"

func TestBeginSessionLimitsAndReleasesByEmail(t *testing.T) {
	sessionLimitUsers.Delete("session@example.com")

	release, ok := beginSession("session@example.com", 1)
	if !ok {
		t.Fatal("first session should be allowed")
	}
	if _, ok := beginSession("session@example.com", 1); ok {
		t.Fatal("second session should be rejected")
	}

	release()
	if releaseAgain, ok := beginSession("session@example.com", 1); !ok {
		t.Fatal("session slot was not released")
	} else {
		releaseAgain()
	}
}
