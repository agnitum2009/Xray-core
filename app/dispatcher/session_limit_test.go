package dispatcher

import (
	"context"
	"testing"
	"time"
)

func TestAllowSessionLimitsAndReleasesByEmail(t *testing.T) {
	sessionLimitUsers.Delete("session@example.com")
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()

	if !allowSession(ctx1, "session@example.com", 1) {
		t.Fatal("first session should be allowed")
	}
	if allowSession(ctx2, "session@example.com", 1) {
		t.Fatal("second session should be rejected")
	}

	cancel1()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		ctx3, cancel3 := context.WithCancel(context.Background())
		if allowSession(ctx3, "session@example.com", 1) {
			cancel3()
			return
		}
		cancel3()
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("session slot was not released")
}
