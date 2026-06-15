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

package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/Yanxinwu946/Ckyn/pkg/util"
	"github.com/docopt/docopt-go"
)

var Args docopt.Opts

var BannerTitle = `Ckyn`

var BannerHeader = fmt.Sprintf(`%s
Lightweight container security assessment tool
`, util.GreenBold.Sprint(BannerTitle))

var BannerContainerTpl = BannerHeader + `
%s
  ckyn evaluate [--full|--profile=<name>]
  ckyn eva [--full|--profile=<name>]
  ckyn run (--list | <exploit> [<args>...])
  ckyn release
  ckyn <tool> [<args>...]
  ckyn -h | --help
  ckyn -v | --version

%s
  ckyn evaluate                              Run baseline evaluation (basic).
  ckyn eva                                   Alias of "ckyn evaluate".
    --full            Shortcut for --profile=extended.
    --profile=<name>  Select profile: basic (default), extended, additional.

%s
  ckyn run --list                            List all available exploits.
  ckyn run <exploit> [<args>...]             Run single exploit.

%s
  release                                    Release embedded busybox and so on to current directory.
  kcurl <path> (get|post) <uri> [<data>]    Make request to K8s api-server.
  ectl <endpoint> get <key>                 Unauthorized enumeration of etcd keys.
  ucurl (get|post) <socket> <uri> <data>    Make request to docker unix socket.
  probe <ip> <port> <parallel> <timeout-ms> TCP port scan
`

// BannerContainer is the banner of Ckyn command line with colorful.
var BannerContainer = fmt.Sprintf(
	BannerContainerTpl,
	"Usage:",
	util.GreenBold.Sprint("Evaluate:"),
	util.GreenBold.Sprint("Exploit:"),
	util.GreenBold.Sprint("Tool:"),
)

func parseDocopt() {
	args, err := docopt.ParseArgs(BannerContainer, os.Args[1:], "")
	if err != nil {
		log.Fatalln("docopt err: ", err)
	}
	Args = args
}
