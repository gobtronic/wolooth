package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"tinygo.org/x/bluetooth"
)

type MACAddr = string

type Config struct {
	// Network interface's MAC address of the device to wake
	TargetWOLMacAddress MACAddr `env:"WOL_TARGET,required"`
	// List of bluetooth devices MAC address to monitor for activity
	MonitoredDevices []MACAddr `env:"BT_DEVICES,required"`
}

// Timestamp of when a monitored device has been last seen
var devices_last_seen = map[MACAddr]time.Time{}

func main() {
	godotenv.Load()
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		panic("An error occured while parsing config: " + err.Error())
	}

	adapter := bluetooth.DefaultAdapter
	err = adapter.Enable()
	if err != nil {
		panic("Failed to enable bluetooth adapter: " + err.Error())
	}

	err = adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		device_addr := device.Address.String()
		// If the device is a monitored device
		if slices.Contains(cfg.MonitoredDevices, device_addr) {
			curr := time.Now()
			last_seen := devices_last_seen[device_addr]
			devices_last_seen[device_addr] = curr
			// If the device hasn't been seen for more than 5 seconds
			if curr.Unix()-last_seen.Unix() > 5 {
				err = send_wol_signal(cfg.TargetWOLMacAddress)
				if err != nil {
					fmt.Fprintf(os.Stderr, "An error occured while sending WOL signal: %s\n", err.Error())
				}
			}
		}
	})
	if err != nil {
		panic("Failed to start scan: " + err.Error())
	}
}

func send_wol_signal(address MACAddr) error {
	cmd := exec.Command("etherwake", address)
	_, err := cmd.Output()
	return err
}
