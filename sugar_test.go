package sugar

import (
	"testing"
)

func TestNewSlog(t *testing.T) {
	lg := &Log{}
	log := NewSlog(lg)
	log.Debug("debug", "hello", "world")
	log.Info("info", "hello", "world")
	log.Error("error", "hello", "world")
}

func TestNewDefault(t *testing.T) {
	log := New()
	log.Debug("debug", "hello", "world")
	log.Info("info", "hello", "world")
	log.Error("error", "hello", "world")
}
