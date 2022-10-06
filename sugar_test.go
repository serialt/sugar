package sugar

import "testing"

func TestLog(t *testing.T) {
	SetLog("error", "")
	Debug("debug logSugar")
	Info("info logSugar")
	Error("error logSugar")
}
