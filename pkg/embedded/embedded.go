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

// ExtractExploitPasswd extracts the embedded exploit-passwd binary to a temp file
// and returns its path. The caller is responsible for cleanup.
func ExtractExploitPasswd() (string, error) {
	return extractBinary("exploit-passwd", exploitPasswd)
}

// ExtractBusybox extracts the embedded busybox binary to a temp file
// and returns its path. The caller is responsible for cleanup.
func ExtractBusybox() (string, error) {
	return extractBinary("busybox", busyboxBinary)
}

// GetBusybox returns the embedded busybox binary bytes
func GetBusybox() []byte {
	return busyboxBinary
}

// GetExploitPasswd returns the embedded exploit-passwd binary bytes
func GetExploitPasswd() []byte {
	return exploitPasswd
}

func extractBinary(name string, data []byte) (string, error) {
	tmpDir := os.TempDir()
	binPath := filepath.Join(tmpDir, "ckyn-"+name)

	if err := os.WriteFile(binPath, data, 0755); err != nil {
		return "", fmt.Errorf("failed to extract %s: %v", name, err)
	}

	return binPath, nil
}
