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

package embedded

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed assets/exploit-passwd
var exploitPasswd []byte

//go:embed assets/busybox
var busyboxBinary []byte

// GetExecutableDir returns the directory of the current executable
func GetExecutableDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(exe), nil
}

// ReleaseBusybox releases the embedded busybox binary to the same directory as ckyn
func ReleaseBusybox() (string, error) {
	dir, err := GetExecutableDir()
	if err != nil {
		return "", err
	}

	busyboxPath := filepath.Join(dir, "busybox")

	// Check if already exists
	if _, err := os.Stat(busyboxPath); err == nil {
		return busyboxPath, nil
	}

	if err := os.WriteFile(busyboxPath, busyboxBinary, 0755); err != nil {
		return "", fmt.Errorf("failed to release busybox: %v", err)
	}

	return busyboxPath, nil
}

// ReleaseExploitPasswd releases the embedded exploit-passwd binary to the same directory as ckyn
func ReleaseExploitPasswd() (string, error) {
	dir, err := GetExecutableDir()
	if err != nil {
		return "", err
	}

	exploitPath := filepath.Join(dir, "exploit-passwd")

	// Check if already exists
	if _, err := os.Stat(exploitPath); err == nil {
		return exploitPath, nil
	}

	if err := os.WriteFile(exploitPath, exploitPasswd, 0755); err != nil {
		return "", fmt.Errorf("failed to release exploit-passwd: %v", err)
	}

	return exploitPath, nil
}

// ExtractExploitPasswd extracts the embedded exploit-passwd binary to a temp file
// and returns its path. The caller is responsible for cleanup.
func ExtractExploitPasswd() (string, error) {
	return extractBinary("exploit-passwd", exploitPasswd)
}

func extractBinary(name string, data []byte) (string, error) {
	tmpDir := os.TempDir()
	binPath := filepath.Join(tmpDir, "ckyn-"+name)

	if err := os.WriteFile(binPath, data, 0755); err != nil {
		return "", fmt.Errorf("failed to extract %s: %v", name, err)
	}

	return binPath, nil
}
