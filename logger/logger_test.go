package logger

import (
	"testing"

	"go.uber.org/zap"
)

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
		NewLogger(
			WithLoggerFilePath("./app.log"),
			WithLoggerConsole(false),
			WithLoggerFormatter("json"),
		)
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

	Logger.Debug("test debug ...", zap.Any("user", u))
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

	Logger.Info("test info ...", zap.Any("user", u))
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


	Logger.Warn("test Warn ...", zap.Any("user", u))
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


	Logger.Error("test Error ...", zap.Any("user", u))
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


	Logger.DPanic("test DBpanic ...", zap.Any("user", u))
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

	Logger.Panic("test panic ...", zap.Any("user", u))
}

func TestSugar(t *testing.T) {
	type User struct {Name string; Url string}
	u := []struct{Name string; Url string}{{ Name: "go", Url: "go.dev" }}

	Logger.SugarDebug("test sugar ..")
	Logger.SugarDebugf("test sugar u = %v", u)
	Logger.SugarDebugw("test sugar u", "user", u)

	Logger.SugarInfo("test sugar ..")
	Logger.SugarInfof("test sugar u = %v", u)
	Logger.SugarInfow("test sugar u", "user", u)

	Logger.SugarWarn("test sugar ..")
	Logger.SugarWarnf("test sugar u = %v", u)
	Logger.SugarWarnw("test sugar u", "user", u)

	Logger.SugarError("test sugar ..")
	Logger.SugarErrorf("test sugar u = %v", u)
	Logger.SugarErrorw("test sugar u", "user", u)

	Logger.SugarDPanic("test sugar ..")
	Logger.SugarDPanicf("test sugar u = %v", u)
	Logger.SugarDPanicw("test sugar u", "user", u)
}