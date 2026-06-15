//go:build linux && amd64
// +build linux,amd64

/*
Copyright 2026 Ckyn Authors .

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package busybox

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Yanxinwu946/Ckyn/pkg/embedded"
)

var busyboxPath string

// Init extracts the embedded busybox binary and returns its path.
// It caches the path for subsequent calls.
func Init() (string, error) {
	if busyboxPath != "" {
		return busyboxPath, nil
	}

	path, err := embedded.ExtractBusybox()
	if err != nil {
		return "", err
	}

	busyboxPath = path
	return busyboxPath, nil
}

// Cleanup removes the extracted busybox binary
func Cleanup() {
	if busyboxPath != "" {
		os.Remove(busyboxPath)
		busyboxPath = ""
	}
}

// Run executes a busybox command with the given arguments
func Run(args ...string) error {
	path, err := Init()
	if err != nil {
		return err
	}

	cmd := exec.Command(path, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// RunApplet runs a specific busybox applet
func RunApplet(applet string, args ...string) error {
	path, err := Init()
	if err != nil {
		return err
	}

	// busybox uses argv[0] to determine the applet
	cmdArgs := append([]string{applet}, args...)
	cmd := exec.Command(path, cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// ListApplets returns a list of available busybox applets
func ListApplets() error {
	path, err := Init()
	if err != nil {
		return err
	}

	cmd := exec.Command(path, "--list")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	applets := strings.Split(string(output), "\n")
	fmt.Printf("Available busybox applets (%d):\n", len(applets))
	for _, applet := range applets {
		if applet != "" {
			fmt.Printf("  %s\n", applet)
		}
	}

	return nil
}

// GetPath returns the path to the extracted busybox binary
func GetPath() string {
	return busyboxPath
}

// Symlink creates a symlink for a specific applet
func Symlink(applet string) error {
	path, err := Init()
	if err != nil {
		return err
	}

	linkPath := filepath.Join(os.TempDir(), "ckyn-"+applet)
	if err := os.Symlink(path, linkPath); err != nil {
		return fmt.Errorf("failed to create symlink: %v", err)
	}

	log.Printf("[+] Created symlink: %s -> %s\n", linkPath, path)
	return nil
}
