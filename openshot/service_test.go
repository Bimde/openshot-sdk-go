package openshot

import (
	"config"
	"os"
	"testing"
)

const (
	// testUsername = "demo-cloud"
	// testPassword = "demo-password"

	testUsername = "cloud-admin"
	testPassword = "m5lKdOMATC"
)

var (
	openShot *OpenShot
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	config.Set(config.Username, testUsername)
	config.Set(config.Password, testPassword)
	openShot = New()
}

func shutdown() {
	openShot = nil
}
