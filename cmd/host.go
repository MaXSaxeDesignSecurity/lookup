// Copyright © 2016 Kevin Kirsche <kevin.kirsche@verizon.com> <kev.kirsche@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Find the IPv4 and IPv6 addrs for domain name",
	Long: `This command looks up a host using the local DNS resolver and returns
the IPv4 and IPv6 addresses for the host.

Example:

~/g/g/s/g/k/lookup ❯❯❯ lookup host google.com
Name: 	google.com
Address: 	172.217.0.46
Address: 	2607:f8b0:4006:807::200e
`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, host := range args {
			addrs, err := net.LookupHost(host)
			if err != nil {
				log.Panicln(err)
			}

			fmt.Printf("Name: \t%s\n", host)
			for _, addr := range addrs {
				fmt.Printf("Address: \t%s\n", addr)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(hostCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
