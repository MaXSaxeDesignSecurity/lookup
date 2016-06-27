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

// udpsrvCmd represents the udpsrv command
var udpsrvCmd = &cobra.Command{
	Use:   "udpservice",
	Short: "Lookup UDP Service Port",
	Long: `Looks up the port for a TCP service.

Example:

~/g/g/s/g/k/lookup ❯❯❯ lookup udpservice telnet
UDP Service: 	telnet
UDP Port: 	23
`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, service := range args {
			port, err := net.LookupPort("udp", service)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("UDP Service: \t%s\n", service)
				fmt.Printf("UDP Port: \t%d\n", port)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(udpsrvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// udpsrvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// udpsrvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
