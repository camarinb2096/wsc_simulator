package server

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	t.Run("Should return a new Server instance", func(t *testing.T) {
		got := NewServer()

		if got.router == nil {
			t.Errorf("Expected router to be initialized")
		}

	})
}
