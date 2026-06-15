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

package cli

import (
	"fmt"
	"log"

	"github.com/Yanxinwu946/Ckyn/pkg/embedded"
	"github.com/Yanxinwu946/Ckyn/pkg/evaluate"
	"github.com/Yanxinwu946/Ckyn/pkg/plugin"
	"github.com/Yanxinwu946/Ckyn/pkg/tool/dockerd_api"
	"github.com/Yanxinwu946/Ckyn/pkg/tool/etcdctl"
	"github.com/Yanxinwu946/Ckyn/pkg/tool/kubectl"
	"github.com/Yanxinwu946/Ckyn/pkg/tool/probe"

	"os"
	"strconv"

	"github.com/docopt/docopt-go"
)

func PassInnerArgs() {
	os.Args = os.Args[1:]
}

func ParseCkynMain() bool {

	if len(os.Args) == 1 {
		docopt.PrintHelpAndExit(nil, BannerContainer)
	}

	// Check for release command before docopt parsing
	if len(os.Args) >= 2 && os.Args[1] == "release" {
		return handleRelease()
	}

	// docopt argparse start
	parseDocopt()

	// support for ckyn eva(Evangelion) and ckyn evaluate
	fok := Args["evaluate"]
	ok := Args["eva"]

	// docopt let fok = true, so we need to check it
	// fix #37
	if ok.(bool) || fok.(bool) {

		fmt.Printf(BannerHeader)
		profileID := evaluate.ProfileBasic
		if rawProfile, ok := Args["--profile"]; ok {
			if v, ok := rawProfile.(string); ok && v != "" {
				profileID = v
			}
		}
		if profileID == evaluate.ProfileBasic && Args["--full"].(bool) {
			profileID = evaluate.ProfileExtended
		}
		if err := evaluate.NewEvaluator().RunProfile(profileID, nil); err != nil {
			log.Printf("evaluate profile %q failed: %v", profileID, err)
		}
		return true
	}

	if Args["run"].(bool) {
		if Args["--list"].(bool) {
			plugin.ListAllExploit()
			os.Exit(0)
		}
		name := Args["<exploit>"].(string)
		if plugin.Exploits[name] == nil {
			fmt.Printf("\nInvalid script name: %s , available scripts:\n", name)
			plugin.ListAllExploit()
			return true
		}
		plugin.RunSingleExploit(name)
		return true
	}

	if Args["<tool>"] != nil {
		args := Args["<args>"].([]string)

		switch Args["<tool>"] {
		case "kcurl":
			kubectl.KubectlToolApi(args)
		case "ectl":
			etcdctl.EtcdctlToolApi(args)
		case "ucurl":
			dockerd_api.UcurlToolApi(args)
		case "dcurl":
			dockerd_api.DcurlToolApi(args)
		case "probe":
			if len(args) != 4 {
				log.Println("Invalid input args.")
				log.Println("usage: ckyn probe <ip> <port> <parallels> <timeout-ms>")
				log.Fatal("example: ckyn probe 192.168.1.0-255 22,80,100-110 50 1000")
			}
			parallel, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				log.Println("err found when parse input arg <parallel>")
				log.Fatal(err)
			}
			timeout, err := strconv.Atoi(args[3])
			if err != nil {
				log.Println("err found when parse input arg <timeout-ms>")
				log.Fatal(err)
			}
			probe.TCPScanToolAPI(args[0], args[1], parallel, timeout)
		default:
			docopt.PrintHelpAndExit(nil, BannerContainer)
		}
	}

	return false
}

func handleRelease() bool {
	fmt.Println("[*] Releasing embedded binaries...")

	// Release busybox
	busyboxPath, err := embedded.ReleaseBusybox()
	if err != nil {
		log.Printf("[-] Failed to release busybox: %v\n", err)
	} else {
		fmt.Printf("[+] Released busybox: %s\n", busyboxPath)
	}

	// Release exploit-passwd
	exploitPath, err := embedded.ReleaseExploitPasswd()
	if err != nil {
		log.Printf("[-] Failed to release exploit-passwd: %v\n", err)
	} else {
		fmt.Printf("[+] Released exploit-passwd: %s\n", exploitPath)
	}

	fmt.Println("[*] Done. You can now use busybox directly:")
	fmt.Printf("    %s/busybox --list\n", func() string {
		dir, _ := embedded.GetExecutableDir()
		return dir
	}())

	return true
}
