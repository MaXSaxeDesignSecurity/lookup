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

// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var serviceport = map[string]map[int]string{
	"tcp": {80: "http"},
}

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:   "port",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open("/etc/services")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for _, servicename := range args {
			fmt.Printf("Service: %s\n", servicename)
			found := false
			for scanner.Scan() {
				// "http 80/tcp www www-http # World Wide Web HTTP"
				line := scanner.Text()
				i := strings.IndexByte(line, '#')
				if i >= 0 {
					line = line[0:i]
				}

				f := getFields(line)
				if len(f) < 2 {
					continue
				}

				if servicename == f[0] {

					portnet := f[1] // "80/tcp"
					port, offset, ok := dtoi(portnet, 0)
					if !ok ||
						port <= 0 ||
						offset >= len(portnet) ||
						portnet[offset] != '/' {
						continue
					}
					network := portnet[offset+1 : offset+4] // "tcp"
					found = true
					fmt.Printf("%s Port: %d\n", strings.ToUpper(string(network)), port)
				}
			}
			if !found {
				fmt.Println("No service port found")
			}
		}
	},
}

func getFields(s string) []string {
	return splitAtBytes(s, " \r\t\n")
}

func splitAtBytes(s string, t string) []string {
	a := make([]string, 1+countAnyByte(s, t))
	n := 0
	last := 0
	for i := 0; i < len(s); i++ {
		ind := strings.IndexByte(t, s[i])
		if ind >= 0 {
			if last < i {
				a[n] = string(s[last:i])
				n++
			}
			last = i + 1
		}
	}
	if last < len(s) {
		a[n] = string(s[last:])
		n++
	}
	return a[0:n]
}

func countAnyByte(s string, t string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		ind := strings.IndexByte(t, s[i])
		if ind >= 0 {
			n++
		}
	}

	return n
}

// Decimal to integer starting at &s[i0].
// Returns number, new offset, success.
func dtoi(s string, i0 int) (n int, i int, ok bool) {
	n = 0
	big := 0xFFFFFF
	neg := false
	if len(s) > 0 && s[0] == '-' {
		neg = true
		s = s[1:]
	}
	for i = i0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		n = n*10 + int(s[i]-'0')
		if n >= big {
			if neg {
				return -big, i + 1, false
			}
			return big, i, false
		}
	}
	if i == i0 {
		return 0, i, false
	}
	if neg {
		n = -n
		i++
	}
	return n, i, true
}

func init() {
	RootCmd.AddCommand(portCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
