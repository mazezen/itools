package logger

import (
	"testing"

	"go.uber.org/zap"
)

var tLog *LoggerOption

func init() {
	tLog = NewLogger(
		WithLoggerFilePath("./app.log"),
		WithLoggerConsole(false),
		WithLoggerFormatter("json"),
	)
}

func TestLogger(t *testing.T) {
	tt := map[string]struct{
		test func(t *testing.T)
	}{
		"debug": { testDebug },
		"info": { testInfo },
		"warn": { testWarn },
		"error": { testError },
		"DPanic": { testDPanic },
	}

	t.Parallel()
	for name, v := range tt {
		t.Run(name, v.test)
	}
}

func testDebug(t *testing.T) {
	type User struct {
		Name string
		Url string
	}

	u := []struct{
		Name string
		Url string
	}{
		{ Name: "go", Url: "go.dev" },
	}

	tLog.Debug("test debug ...", zap.Any("user", u))
}

func testInfo(t *testing.T) {
	type User struct {
		Name string
		Url string
	}

	u := []struct{
		Name string
		Url string
	}{
		{ Name: "go", Url: "go.dev" },
	}

	tLog.Info("test info ...", zap.Any("user", u))
}

func testWarn(t *testing.T) {
	type User struct {
		Name string
		Url string
	}

	u := []struct{
		Name string
		Url string
	}{
		{ Name: "go", Url: "go.dev" },
	}


	tLog.Warn("test Warn ...", zap.Any("user", u))
}

func testError(t *testing.T) {
	type User struct {
		Name string
		Url string
	}

	u := []struct{
		Name string
		Url string
	}{
		{ Name: "go", Url: "go.dev" },
	}


	tLog.Error("test Error ...", zap.Any("user", u))
}

func testDPanic(t *testing.T) {
	type User struct {
		Name string
		Url string
	}

	u := []struct{
		Name string
		Url string
	}{
		{ Name: "go", Url: "go.dev" },
	}


	tLog.DPanic("test DBpanic ...", zap.Any("user", u))
}

func TestPanic(t *testing.T) {
	type User struct {
		Name string
		Url string
	}

	u := []struct{
		Name string
		Url string
	}{
		{ Name: "go", Url: "go.dev" },
	}
	defer func() {
		if r := recover(); r != nil {
			t.Logf("成功捕获 Panic: %v", r)
		} else {
			t.Error("预期发生 Panic，但没有发生")
		}
	}()

	tLog.Panic("test panic ...", zap.Any("user", u))
}

func TestSugar(t *testing.T) {
	type User struct {Name string; Url string}
	u := []struct{Name string; Url string}{{ Name: "go", Url: "go.dev" }}

	tLog.SugarDebug("test sugar ..")
	tLog.SugarDebugf("test sugar u = %v", u)
	tLog.SugarDebugw("test sugar u", "user", u)

	tLog.SugarInfo("test sugar ..")
	tLog.SugarInfof("test sugar u = %v", u)
	tLog.SugarInfow("test sugar u", "user", u)

	tLog.SugarWarn("test sugar ..")
	tLog.SugarWarnf("test sugar u = %v", u)
	tLog.SugarWarnw("test sugar u", "user", u)

	tLog.SugarError("test sugar ..")
	tLog.SugarErrorf("test sugar u = %v", u)
	tLog.SugarErrorw("test sugar u", "user", u)

	tLog.SugarDPanic("test sugar ..")
	tLog.SugarDPanicf("test sugar u = %v", u)
	tLog.SugarDPanicw("test sugar u", "user", u)
}