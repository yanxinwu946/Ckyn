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

package cli_test

import (
	"bytes"
	// "io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Yanxinwu946/Ckyn/pkg/cli"
	_ "github.com/Yanxinwu946/Ckyn/pkg/exploit" // register all exploits
)

type testArgsCase struct {
	name       string
	args       []string
	successStr string
}

const parseTimeout = 5 * time.Second

func doParseCkynMainWithTimeout() {

	result := make(chan bool, 1)

	go func() {
		result <- cli.ParseCkynMain()
	}()

	select {
	case <-time.After(parseTimeout):
		log.Printf("check run ok, timeout reached in %s, and return.", parseTimeout)
		return
	case <-result:
		return
	}

}

func TestParseCkynMain(t *testing.T) {

	// ./ckyn eva 2>&1 | head
	// ./ckyn run test-poc | head
	// ./ckyn ifconfig | head

	tests := []testArgsCase{
		{
			name:       "./ckyn eva",
			args:       []string{"./ckyn_cli_path", "eva"},
			successStr: "current user",
		},
		{
			name:       "./ckyn run test-poc",
			args:       []string{"./ckyn_cli_path", "run", "test-poc"},
			successStr: "run success",
		},
	}

	for _, tt := range tests {

		// fmt.Print and log.Print to buffer, and check output
		var buf bytes.Buffer
		log.SetOutput(&buf)

		// hook fmt.X to buffer, hook os.Stdout
		// oldStdout := os.Stdout
		// r, w, _ := os.Pipe()
		// os.Stdout = w

		// hook os.Args
		args := tt.args
		os.Args = args

		t.Run(tt.name, func(t *testing.T) {
			doParseCkynMainWithTimeout()
			// out, _ := ioutil.ReadAll(r)

			// check success string in buf and out
			// if !bytes.Contains(buf.Bytes(), []byte(tt.successStr)) && !bytes.Contains(out, []byte(tt.successStr)) {
			// 	t.Errorf(("parse ckyn main failed, name: %s, args: %v, buf: %s, out: %s"), tt.name, tt.args, buf.String()[:1000], string(out)[:1000])
			// }

			if !bytes.Contains(buf.Bytes(), []byte(tt.successStr)) {

				// get sub string from buf, lenght is 1000
				str := buf.String()
				if len(str) > 1000 {
					str = str[:1000]
				}

				t.Errorf(("parse ckyn main failed, name: %s, args: %v, buf: %s"), tt.name, tt.args, str)
			}

		})

		// return to os.Stdout default
		// os.Stdout = oldStdout
		// w.Close()
	}
}
