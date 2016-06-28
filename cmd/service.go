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

var (
	tcp bool
	udp bool
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Looks up the port for a named service",
	Long: `Looks up the port for a service.

Example:

~/g/g/s/g/k/lookup ❯❯❯ lookup service telnet
TCP Service: 	telnet
TCP Port: 	23
UDP Service: 	telnet
UDP Port: 	23

~/g/g/s/g/k/lookup ❯❯❯ lookup service -t telnet
TCP Service: 	telnet
TCP Port: 	23

~/g/g/s/g/k/lookup ❯❯❯ lookup service --tcp telnet
TCP Service: 	telnet
TCP Port: 	23

~/g/g/s/g/k/lookup ❯❯❯ lookup service -u telnet
UDP Service: 	telnet
UDP Port: 	23

~/g/g/s/g/k/lookup ❯❯❯ lookup service --udp telnet
UDP Service: 	telnet
UDP Port: 	23
	`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, service := range args {
			var (
				tport int
				uport int
				err   error
				terr  error
				uerr  error
			)
			if tcp {
				tport, err = net.LookupPort("tcp", service)
			} else if udp {
				uport, err = net.LookupPort("udp", service)
			} else {
				tport, terr = net.LookupPort("tcp", service)
				uport, uerr = net.LookupPort("udp", service)
			}
			if err != nil {
				fmt.Println(err)
			} else {
				if tcp {
					fmt.Printf("TCP Service: \t%s\n", service)
					fmt.Printf("TCP Port: \t%d\n", tport)
				} else if udp {
					fmt.Printf("UDP Service: \t%s\n", service)
					fmt.Printf("UDP Port: \t%d\n", uport)
				} else {
					if terr != nil {
						fmt.Println(terr)
					} else {
						fmt.Printf("TCP Service: \t%s\n", service)
						fmt.Printf("TCP Port: \t%d\n", tport)
					}

					if uerr != nil {
						fmt.Println(uerr)
					} else {
						fmt.Printf("UDP Service: \t%s\n", service)
						fmt.Printf("UDP Port: \t%d\n", uport)
					}
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(serviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serviceCmd.PersistentFlags().BoolVarP(&tcp, "tcp", "t", false, "TCP Service")
	serviceCmd.PersistentFlags().BoolVarP(&udp, "udp", "u", false, "UDP Service")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
