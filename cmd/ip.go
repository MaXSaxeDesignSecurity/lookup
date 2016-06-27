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
	"net"

	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Reverse lookup for the given IP Address",
	Long: `This command performs a reverse lookup to find the DNS name(s) that an
IPv4 or IPv6 address resolves to.

Example:

~/g/g/s/g/k/lookup ❯❯❯ lookup ip 172.217.0.46
Address: 	172.217.0.46
Name: 	google.com
`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, addr := range args {
			names, err := net.LookupAddr(addr)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Address: \t%s\n", addr)
				for _, name := range names {
					fmt.Printf("Name: \t%s\n", name)
					addrs, _ := net.LookupHost(name)
					if len(addrs) > 1 {
						fmt.Println()
						for _, adaddr := range addrs {
							if adaddr == addr {
								continue
							}
							fmt.Printf("Additional address: \t%s\n", adaddr)
						}
					}
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(ipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
