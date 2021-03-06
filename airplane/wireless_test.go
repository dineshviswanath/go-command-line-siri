package airplane

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWireless_IdentifyDriver(t *testing.T) {
	w, err := IdentifyDriver()
	// TODO: Add mocks so it works on all OS
	assert.NotNil(t, w, "Driver should be identified")
	assert.Nil(t, err, "Err should be nil")
	assert.Equal(t, "networksetup", w.DriverCommand)
	assert.Equal(t, "MacOS", w.OperatingSystem)
}

func TestWireless_IdentifyPort(t *testing.T) {
	w := Wireless{
		DriverName:      "NetworkSetup",
		OperatingSystem: "MacOS",
		DriverCommand:   "networksetup",
	}
	// TODO: Add mocks so it works on all OS
	device, err := w.IdentifyDevice()
	assert.NotEmpty(t, device, "Device should be identified")
	assert.Equal(t, "en0", device)
	assert.Nil(t, err, "Error should be Nil")
}

func TestWireless_Toggle(t *testing.T) {
	w := Wireless{
		DriverName:      "NetworkSetup",
		OperatingSystem: "MacOS",
		DriverCommand:   "networksetup",
	}
	// TODO: Add mocks so it works on all OS
	w.IdentifyDevice()
	ok := w.Toggle("ON")
	assert.True(t, ok)
}
