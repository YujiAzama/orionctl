/*
Copyright © 2020 Yuji Azama <yuji.azama@gmail.com>

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
package cmd

import (
	"context"
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/YujiAzama/orionclient-go/orionclient"
)

var version orionclient.Version

var getVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get version",
	Long:  "Get version",
	Run: func(cmd *cobra.Command, args []string) {
		oc := orionclient.ClientConfig{Host: config.Host, Port: config.Port, TLS: config.TLS, Token: config.Token}
		client, err := orionclient.NewClient(oc)
		if err != nil {
			panic(err)
		}
		version, err := client.GetVersion(context.Background())
		if err != nil {
			panic(err)
		}

		table := uitable.New()
		table.MaxColWidth = 80
		table.Wrap = true
		table.AddRow("Version", version.Orion.Version)
		table.AddRow("Uptime", version.Orion.Uptime)
		table.AddRow("GitHash", version.Orion.GitHash)
		table.AddRow("CompileTime", version.Orion.CompileTime)
		table.AddRow("CompiledBy", version.Orion.CompiledBy)
		table.AddRow("CompiledIn", version.Orion.CompiledIn)
		table.AddRow("ReleaseDate", version.Orion.ReleaseDate)
		table.AddRow("Doc", version.Orion.Doc)
		fmt.Println(table)
	},
}

func init() {
	rootCmd.AddCommand(getVersionCmd)
}
