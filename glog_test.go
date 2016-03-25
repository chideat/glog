package glog

import (
	"testing"
)

func TestInfo(test *testing.T) {
	Info("test info")
}

func TestInfof(test *testing.T) {
	Infof("%s", "test infof")
}

func TestWarn(test *testing.T) {
	Warn("test warn")
}

func TestWarnf(test *testing.T) {
	Warnf("%s", "test warnf")
}

func TestError(test *testing.T) {
	Error("test error")
}

func TestErrorf(test *testing.T) {
	Errorf("%s", "test errorf")
}
