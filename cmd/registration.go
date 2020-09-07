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
	"errors"
	"fmt"
	"os"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/YujiAzama/orionclient-go/orionclient"
)

var registrationFile string
var registration orionclient.Registration

var getRegistrationCmd = &cobra.Command{
	Use:   "registration",
	Short: "get registration",
	Long:  "get registration",
	Run: func(cmd *cobra.Command, args []string) {
		oc := orionclient.ClientConfig{Host: config.Host, Port: config.Port, TLS: config.TLS, Token: config.Token}
		client, err := orionclient.NewClient(oc)
		if err != nil {
			panic(err)
		}

                var registrations = []*orionclient.Registration{}
                if len(args) > 0 {
                        for _, id := range args {
                                registration, err := client.GetRegistration(context.Background(), id, fs, fsp)
                                if err != nil {
                                        panic(err)
                                }
                                registrations = append(registrations, registration)
                        }
                } else {
                        allRegistrations, err := client.GetRegistrations(context.Background(), fs, fsp)
                        if err != nil {
                                panic(err)
                        }
                        registrations = allRegistrations
                }

		table := uitable.New()
		table.MaxColWidth = 50
		table.AddRow("ID", "Provider URL", "Status")
		for _, registration := range registrations {
			table.AddRow(registration.Id, registration.Provider.HTTP.URL, registration.Status)
		}
		fmt.Println(table)
	},
}

var describeRegistrationCmd = &cobra.Command{
	Use:   "registration",
	Short: "describe registration",
	Long:  "describe registration",
	Run: func(cmd *cobra.Command, args []string) {
		oc := orionclient.ClientConfig{Host: config.Host, Port: config.Port, TLS: config.TLS, Token: config.Token}
		client, err := orionclient.NewClient(oc)
		if err != nil {
			panic(err)
		}

                var registrations = []*orionclient.Registration{}
                if len(args) > 0 {
                        for _, id := range args {
                                registration, err := client.GetRegistration(context.Background(), id, fs, fsp)
                                if err != nil {
                                        panic(err)
                                }
                                registrations = append(registrations, registration)
                        }
                } else {
                        allRegistrations, err := client.GetRegistrations(context.Background(), fs, fsp)
                        if err != nil {
                                panic(err)
                        }
                        registrations = allRegistrations
                }

		table := uitable.New()
		table.MaxColWidth = 80
		table.Wrap = true
		for _, registration := range registrations {
			table.AddRow("ID:", registration.Id)
			table.AddRow("DataProvided:")
			for i, entity := range registration.DataProvided.Entities {
				var value = ""
				if entity.IdPattern != "" {
					value = "IdPattern: " + entity.IdPattern
				}
				if entity.ID != "" {
					value = "Id: " + entity.ID
				}
				if i == 0 {
					table.AddRow("    Entities:", value + ", Type: " + entity.Type)
				} else {
					table.AddRow("             ", value + ", Type: " + entity.Type)
				}
			}
			for i, attr := range registration.DataProvided.Attrs {
				if i == 0 {
					table.AddRow("    Attrs:", attr)
				} else {
					table.AddRow("          ", attr)
				}
			}
			table.AddRow("Provider:")
			table.AddRow("    HTTP:")
			table.AddRow("        URL:", registration.Provider.HTTP.URL)
			table.AddRow("    LegacyForwarding:", registration.Provider.LegacyForwarding)
			table.AddRow("    SupportedForwardingMode:", registration.Provider.SupportedForwardingMode)
			table.AddRow("Status:", registration.Status)
			table.AddRow("")
		}
		fmt.Println(table)
	},
}

var createRegistrationCmd = &cobra.Command{
	Use:   "registration",
	Short: "create registration",
	Long:  "create registration resources by filename",
	Run:  func(cmd *cobra.Command, args []string) {
		viper.SetConfigName(registrationFile)
		viper.SetConfigType("yml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("yaml file read error")
			fmt.Println(err)
			os.Exit(1)
		}
		if err := viper.Unmarshal(&registration); err != nil {
			fmt.Println("registration file Unmarshal error")
			fmt.Println(err)
			os.Exit(1)
		}
		oc := orionclient.ClientConfig{Host: config.Host, Port: config.Port, TLS: config.TLS, Token: config.Token}
		client, err := orionclient.NewClient(oc)
		if err != nil {
			panic(err)
		}
		registrationId, err := client.CreateRegistration(context.Background(), registration, fs, fsp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("registration \"%s\" created\n", registrationId)
	},
}

var deleteRegistrationCmd = &cobra.Command{
	Use:   "registration",
	Short: "delete registration",
	Long:  "delete registration",
	Args:  func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a registration ID")
		}
		return nil
	},
	Run:  func(cmd *cobra.Command, args []string) {
		oc := orionclient.ClientConfig{Host: config.Host, Port: config.Port, TLS: config.TLS, Token: config.Token}
                client, err := orionclient.NewClient(oc)
		if err != nil {
                        panic(err)
                }
		for _, registrationId := range args {
			if err := client.DeleteRegistration(context.Background(), registrationId, fs, fsp); err != nil {
				fmt.Println(err)
			}
			fmt.Printf("registration \"%s\" deleted\n", registrationId)
		}
	},
}

func init() {
	getCmd.AddCommand(getRegistrationCmd)
	describeCmd.AddCommand(describeRegistrationCmd)
	createRegistrationCmd.Flags().StringVarP(&registrationFile, "registrationFile", "f", "", "Registration resource filename")
	createCmd.AddCommand(createRegistrationCmd)
	deleteCmd.AddCommand(deleteRegistrationCmd)
}
