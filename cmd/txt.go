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

// txtCmd represents the txt command
var txtCmd = &cobra.Command{
	Use:   "txt",
	Short: "Lookup TXT records for a given domain",
	Long: `Looks up the DNS TXT records for the given domain name.

Example:

~/g/g/s/g/k/lookup ❯❯❯ lookup txt google.com
Name: 	google.com
TXT: 	v=spf1 include:_spf.google.com ~all
`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, host := range args {
			txts, err := net.LookupTXT(host)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Name: \t%s\n", host)
				for _, txt := range txts {
					fmt.Printf("TXT: \t%s\n", txt)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(txtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// txtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// txtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
