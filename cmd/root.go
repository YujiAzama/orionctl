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
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var config Config

var host string
var port int
var tls  bool
var token string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "orionctl",
	Short: "This is a command line interface for control FIWARE Orion.",
	Long:  "This is a command line interface for control FIWARE Orion.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.orionctl.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "localhost", "Orion hostname or IP address")
	viper.BindPFlag("Host", rootCmd.PersistentFlags().Lookup("host"))
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 1026, "Orion port number")
	viper.BindPFlag("Port", rootCmd.PersistentFlags().Lookup("port"))
	rootCmd.PersistentFlags().BoolVarP(&tls, "tls", "k", false, "Enable TLS/SSL")
	viper.BindPFlag("TLS", rootCmd.PersistentFlags().Lookup("tls"))
	rootCmd.PersistentFlags().StringVarP(&token, "token", "T", "", "Access Token")
	viper.BindPFlag("Token", rootCmd.PersistentFlags().Lookup("token"))

	rootCmd.PersistentFlags().StringVarP(&fs, "fiware-service", "s", "", "FIWARE Service")
	rootCmd.PersistentFlags().StringVarP(&fsp, "fiware-servicepath", "P", "", "FIWARE Service Path")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".orionctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".orionctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
