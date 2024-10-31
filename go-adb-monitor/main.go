package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func RunADBCommand(cmdArgs ...string) (string, error) {
	cmd := exec.Command("adb", cmdArgs...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func GetConnectedDevices() ([]string, error) {
	out, err := RunADBCommand("devices")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(out, "\n")
	var devices []string
	for _, line := range lines {
		if strings.Contains(line, "device") && !strings.Contains(line, "List of devices") {
			parts := strings.Fields(line)
			if len(parts) > 0 {
				devices = append(devices, parts[0])
			}
		}
	}
	return devices, nil
}

func MonitorDevices() {
	fmt.Println("Starting to monitor connected devices...")
	previousDevices := make(map[string]bool)
	for {
		currentDevices, err := GetConnectedDevices()
		if err != nil {
			fmt.Printf("Error fetching devices: %v\n", err)
			time.Sleep(2 * time.Second)
			continue
		}
		currentMap := make(map[string]bool)
		for _, device := range currentDevices {
			currentMap[device] = true
			if !previousDevices[device] {
				fmt.Printf("Device connected: %s\n", device)
			}
		}
		for device := range previousDevices {
			if !currentMap[device] {
				fmt.Printf("Device disconnected: %s\n", device)
			}
		}
		previousDevices = currentMap
		time.Sleep(2 * time.Second)
	}
}

// replace "cut", "-c", "52-66", "|", "tr", "-d", "'.[:space:]'"
func ParseIMEI(input string) string {
	re := regexp.MustCompile(`'([^']*)'`)
	matches := re.FindAllStringSubmatch(input, -1)
	var imei string
	for _, match := range matches {
		if len(match) > 1 {
			cleaned := strings.ReplaceAll(match[1], ".", "")
			imei += strings.TrimSpace(cleaned)
		}
	}
	return imei
}

func GetDeviceInfo() (map[string]string, error) {
	fmt.Println("Get device information...")
	commands := map[string][]string{
		"SerialNumber":    {"shell", "getprop", "ro.serialno"},
		"Model":           {"shell", "getprop", "ro.product.model"},
		"Manufacturer":    {"shell", "getprop", "ro.product.manufacturer"},
		"AndroidVersion":  {"shell", "getprop", "ro.build.version.release"},
		"NetworkOperator": {"shell", "getprop", "gsm.operator.alpha"},
		"NetworkType":     {"shell", "getprop", "gsm.network.type"},
		"PhoneNumber":     {"shell", "getprop", "line1.number"},
		"IMEI1":           {"shell", "service", "call", "iphonesubinfo", "1"},
		"IMEI2":           {"shell", "service", "call", "iphonesubinfo", "2"},
	}
	info := make(map[string]string)
	for key, cmd := range commands {
		output, err := RunADBCommand(cmd...)
		if err != nil {
			info[key] = "Unavailable"
			continue
		}
		if key == "IMEI1" || key == "IMEI2" {
			info[key] = ParseIMEI(output)
			continue
		}
		info[key] = output
	}
	return info, nil
}

func main() {
	// MonitorDevices()
	deviceInfo, err := GetDeviceInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for key, value := range deviceInfo {
		fmt.Printf("%s: %s\n", key, value)
	}
}
