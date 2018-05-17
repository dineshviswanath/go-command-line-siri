package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWireless_IdentifyDriver(t *testing.T) {
	w, err := IdentifyDriver()
	assert.NotNil(t, w, "Driver should be identified")
	assert.Nil(t, err, "Err should be nil")
	assert.Equal(t, "networksetup", w.DriverCommand)
	assert.Equal(t, "MacOS", w.OperatingSystem)
}


func TestWireless_IdentifyPort(t *testing.T) {
	w := Wireless{
		DriverName: "NetworkSetup",
		OperatingSystem: "MacOS",
		DriverCommand: "networksetup",
	}
	device, err := w.IdentifyDevice()
	assert.NotEmpty(t, device, "Device should be identified")
	assert.Equal(t, "en0", device)
	assert.Nil(t, err, "Error should be Nil")
}