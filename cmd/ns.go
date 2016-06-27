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

// nsCmd represents the ns command
var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Lookup nameserver(s) for a host",
	Long: `This command looks up the DNS nameservers for a host.

Example:

~/g/g/s/g/k/lookup ❯❯❯ lookup ns google.com
Name: 	google.com
Nameserver: 	ns3.google.com.
Nameserver: 	ns2.google.com.
Nameserver: 	ns1.google.com.
Nameserver: 	ns4.google.com.
`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, host := range args {
			nss, err := net.LookupNS(host)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Name: \t%s\n", host)
				for _, n := range nss {
					fmt.Printf("Nameserver: \t%s\n", n.Host)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(nsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
