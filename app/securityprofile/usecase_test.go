package securityprofile

import (
	"ai-dekawaf/app/f5"
	"errors"
	"fmt"
	"testing"
)

func TestDeleteF5Resource(t *testing.T) {
	called := false
	if err := deleteF5Resource("x", "", func() error { called = true; return nil }); err != nil || called {
		t.Fatalf("empty path must be skipped: called=%v err=%v", called, err)
	}
	if err := deleteF5Resource("x", "/Common/a", func() error { return fmt.Errorf("wrap: %w", f5.ErrNotFound) }); err != nil {
		t.Fatalf("not found must be treated as deleted: %v", err)
	}
	boom := errors.New("boom")
	if err := deleteF5Resource("x", "/Common/a", func() error { return boom }); !errors.Is(err, boom) {
		t.Fatalf("expected boom, got %v", err)
	}
}
