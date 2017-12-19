package system

import (
	"bytes"
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

func GetAddress(intf string) (string, error) {
	var stdoutBuff, stderrBuff bytes.Buffer

	args := []string{"-4", "addr", "show", "dev", intf}
	cmd := exec.Command("ip", args...)
	cmd.Stdout = &stdoutBuff
	cmd.Stderr = &stderrBuff

	err := cmd.Start()
	if err != nil {
		return "", fmt.Errorf("GetAddress cmd.Start: %v", err)
	}

	err = cmd.Wait()

	stdout := ""
	if stdoutBuff.Len() > 0 {
		stdout = stdoutBuff.String()
	}

	if err != nil {
		return "", fmt.Errorf("GetAddress cmd.Wait: %v", err)
	}

	address := ""
	for _, line := range strings.Split(stdout, "\n") {
		result := reAddress.FindAllStringSubmatch(line, -1)
		if len(result) > 0 {
			address = result[0][1]
			break
		}
	}

	if address == "" {
		return "", fmt.Errorf("GetAddress: No ip address found for %s", intf)
	}

	return address, nil
}

func GetFirstInterface() (string, error) {
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command("ifconfig", "-a")
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Start()
	if err != nil {
		return "", fmt.Errorf("GetFirstInterface cmd.Start: %v", err)
	}

	err = cmd.Wait()
	if err != nil {
		return "", fmt.Errorf("GetFirstInterface cmd.Wait: %v", err)
	}

	allInterfaces := []string{}
	for _, line := range strings.Split(stdoutBuf.String(), "\n") {
		result := reInterface.FindAllStringSubmatch(line, -1)
		if len(result) > 0 {
			intf := result[0][1]
			if intf == "lo" {
				continue
			}
			allInterfaces = append(allInterfaces, intf)
		}
	}

	if len(allInterfaces) == 0 {
		return "", fmt.Errorf("GetFirstInterface: no interfaces found")
	}

	sort.Strings(allInterfaces)

	return allInterfaces[0], nil
}
