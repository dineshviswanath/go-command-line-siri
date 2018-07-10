package airplane

import (
	"log"
	"os/exec"
	"strings"
)

// Wireless Holds abstraction for all operating system airplane driver
type Wireless struct {
	DriverName      string
	OperatingSystem string
	DriverCommand   string
	Device          string
}

func IdentifyDriver() (Wireless, error) {
	// TODO: Add cases for Linux and Unix
	log.Printf("Running the command to identify driver...")
	_, err := exec.Command("which", "networksetup").Output()

	if err != nil {
		log.Printf("Finished with the value %s", err)
		return Wireless{}, err
	}

	return Wireless{
		DriverName:      "NetworkSetup",
		OperatingSystem: "MacOS",
		DriverCommand:   "networksetup",
	}, nil
}

func (w *Wireless) IdentifyDevice() (string, error) {

	log.Printf("Running the command to identify device...")
	out, err := exec.Command("networksetup", "-listallhardwareports").Output()

	if err != nil {
		log.Printf("Error in identifying device %s", err)
		return "", err
	}

	var device string
	isWifi := false
	for _, line := range strings.Split(string(out), "\n") {
		if isWifi {
			device = strings.Replace(line, "Device: ", "", -1)
			isWifi = false
			break
		}
		if !isWifi && strings.ContainsAny(line, "Wi-Fi") {
			isWifi = true
		}
	}
	w.Device = device
	log.Printf("Device is %s", device)
	return device, nil
}

func (w *Wireless) Toggle(state string) bool {
	log.Printf("About to Power %s: %s...", state, w.Device)
	_, err := exec.Command("networksetup", "-setairportpower", w.Device, state).Output()
	if err != nil {
		log.Printf("Error in toggling Power %s", err)
		return false
	}
	return true
}
