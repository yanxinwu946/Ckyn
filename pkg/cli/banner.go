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
var GitCommit string

var BannerTitle = `Ckyn`
var BannerVersion = fmt.Sprintf("%s %s", "Ckyn Version(GitCommit):", GitCommit)

var BannerHeader = fmt.Sprintf(`%s
%s
Lightweight container security assessment tool
`, util.GreenBold.Sprint(BannerTitle), BannerVersion)

var BannerContainerTpl = BannerHeader + `
%s
  ckyn eva
  ckyn eva --full
  ckyn evaluate [--full]
  ckyn run (--list | <exploit> [<args>...])
  ckyn <tool> [<args>...]

%s
  ckyn evaluate                              Gather information to find weakness inside container.
  ckyn eva                                   Alias of "ckyn evaluate".
  ckyn evaluate --full                       Enable file scan during information gathering.


%s
  ckyn run --list                            List all available exploits.
  ckyn run <exploit> [<args>...]             Run single exploit.

%s
  busybox [<args>...]                        Embedded busybox with common Unix utilities.
  busybox --list                             List available busybox applets.
  kcurl <path> (get|post) <uri> [<data>]    Make request to K8s api-server.
  ectl <endpoint> get <key>                 Unauthorized enumeration of etcd keys.
  ucurl (get|post) <socket> <uri> <data>    Make request to docker unix socket.
  probe <ip> <port> <parallel> <timeout-ms> TCP port scan, example: ckyn probe 10.0.1.0-255 80,8080-9443 50 1000

%s
  -h --help     Show this help msg.
  -v --version  Show version.
  --profile=<name> Select evaluation profile (basic, extended, additional).
`

// BannerContainer is the banner of Ckyn command line with colorful.
var BannerContainer = fmt.Sprintf(
	BannerContainerTpl,
	"Usage:",
	util.GreenBold.Sprint("Evaluate:"),
	util.GreenBold.Sprint("Exploit:"),
	util.GreenBold.Sprint("Tool:"),
	"Options:",
)

func parseDocopt() {
	args, err := docopt.ParseArgs(BannerContainer, os.Args[1:], BannerVersion)
	if err != nil {
		log.Fatalln("docopt err: ", err)
	}
	Args = args
}
