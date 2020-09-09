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

var subsFile string
var subscription orionclient.Subscription

var getSubscriptionCmd = &cobra.Command{
	Use:   "subscriptions",
	Aliases: []string{"subscription", "subs"},
	Short: "Get subscription. Aliases: [\"subscription\", \"subs\"]",
	Long:  "Get subscription",
	Run: func(cmd *cobra.Command, args []string) {
		oc := orionclient.ClientConfig{Host: config.Host, Port: config.Port, TLS: config.TLS, Token: config.Token}
		client, err := orionclient.NewClient(oc)
		if err != nil {
			panic(err)
		}

		var subscriptions = []*orionclient.Subscription{}
		if len(args) > 0 {
			for _, id := range args {
				subscription, err := client.GetSubscription(context.Background(), id, fs, fsp)
				if err != nil {
					panic(err)
				}
				subscriptions = append(subscriptions, subscription)
			}
		} else {
			allSubscriptions, err := client.GetSubscriptions(context.Background(), fs, fsp)
			if err != nil {
				panic(err)
			}
			subscriptions = allSubscriptions
		}

		table := uitable.New()
		table.MaxColWidth = 50
		table.AddRow("ID", "Description", "Notification URL", "LastSuccess")
		for _, subscription := range subscriptions {
			table.AddRow(subscription.Id, subscription.Description, subscription.Notification.HTTP.URL, subscription.Notification.LastSuccess)
		}
		fmt.Println(table)
	},
}

var describeSubscriptionCmd = &cobra.Command{
	Use:   "subscriptions",
	Aliases: []string{"subscription", "subs"},
	Short: "Describe subscription. Aliases: [\"subscription\", \"subs\"]",
	Long:  "Describe subscription",
	Run: func(cmd *cobra.Command, args []string) {
		oc := orionclient.ClientConfig{Host: config.Host, Port: config.Port, TLS: config.TLS, Token: config.Token}
		client, err := orionclient.NewClient(oc)
		if err != nil {
			panic(err)
		}

		var subscriptions = []*orionclient.Subscription{}
		if len(args) > 0 {
			for _, id := range args {
				subscription, err := client.GetSubscription(context.Background(), id, fs, fsp)
				if err != nil {
					panic(err)
				}
				subscriptions = append(subscriptions, subscription)
			}
		} else {
			allSubscriptions, err := client.GetSubscriptions(context.Background(), fs, fsp)
			if err != nil {
				panic(err)
			}
			subscriptions = allSubscriptions
		}

		table := uitable.New()
		table.MaxColWidth = 80
		table.Wrap = true
		for _, subscription := range subscriptions {
			table.AddRow("ID:", subscription.Id)
			table.AddRow("Description:", subscription.Description)
			table.AddRow("Subject:")
			for i, entity := range subscription.Subject.Entities {
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
			if len(subscription.Subject.Condition.Attrs) > 0 || subscription.Subject.Condition.Expression != nil {
				table.AddRow("    Condition:")
				for i, attr := range subscription.Subject.Condition.Attrs {
					if i == 0 {
						table.AddRow("        Attrs:", attr)
					} else {
						table.AddRow("              ", attr)
					}
				}
				if subscription.Subject.Condition.Expression != nil {
					table.AddRow("        Expression:")
					table.AddRow("            Q:", subscription.Subject.Condition.Expression.Q)
				}
			}
			table.AddRow("Notification:")
			table.AddRow("    HTTP:")
			table.AddRow("        URL:", subscription.Notification.HTTP.URL)
			for i, attr := range subscription.Notification.Attrs {
				if i == 0 {
					table.AddRow("    Attrs:", attr)
				} else {
					table.AddRow("          ", attr)
				}
			}
			table.AddRow("    AttrsFormat:", subscription.Notification.AttrsFormat)
			table.AddRow("    LastFailure:", subscription.Notification.LastFailure)
			table.AddRow("    LastFailureReason:", subscription.Notification.LastFailureReason)
			table.AddRow("    LastNotification:", subscription.Notification.LastNotification)
			table.AddRow("    LastSuccess:", subscription.Notification.LastSuccess)
			table.AddRow("    LastSuccessCode:", subscription.Notification.LastSuccessCode)
			table.AddRow("    OnlyChangedAttrs:", subscription.Notification.OnlyChangedAttrs)
			table.AddRow("    TimesSent:", subscription.Notification.TimesSent)
			table.AddRow("Expires:", subscription.Expires)
			table.AddRow("Throttling:", subscription.Throttling)
			table.AddRow("")
		}
		fmt.Println(table)
	},
}

var createSubscriptionCmd = &cobra.Command{
	Use:   "subscriptions",
	Aliases: []string{"subscription", "subs"},
	Short: "Create subscription. Aliases: [\"subscription\", \"subs\"]",
	Long:  "Create subscription resources by filename",
	Run:  func(cmd *cobra.Command, args []string) {
		viper.SetConfigName(subsFile)
		viper.SetConfigType("yml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("yaml file read error")
			fmt.Println(err)
			os.Exit(1)
		}
		if err := viper.Unmarshal(&subscription); err != nil {
			fmt.Println("subscription file Unmarshal error")
			fmt.Println(err)
			os.Exit(1)
		}
		oc := orionclient.ClientConfig{Host: config.Host, Port: config.Port, TLS: config.TLS, Token: config.Token}
		client, err := orionclient.NewClient(oc)
		if err != nil {
			panic(err)
		}
		subscriptionId, err := client.CreateSubscription(context.Background(), subscription, fs, fsp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("subscription \"%s\" created\n", subscriptionId)
	},
}

var deleteSubscriptionCmd = &cobra.Command{
	Use:   "subscriptions",
	Aliases: []string{"subscription", "subs"},
	Short: "Delete subscription. Aliases: [\"subscription\", \"subs\"]",
	Long:  "Delete subscription",
	Args:  func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a subscription ID")
		}
		return nil
	},
	Run:  func(cmd *cobra.Command, args []string) {
		oc := orionclient.ClientConfig{Host: config.Host, Port: config.Port, TLS: config.TLS, Token: config.Token}
                client, err := orionclient.NewClient(oc)
		if err != nil {
                        panic(err)
                }
		for _, subscriptionId := range args {
			if err := client.DeleteSubscription(context.Background(), subscriptionId, fs, fsp); err != nil {
				fmt.Println(err)
			}
			fmt.Printf("subscription \"%s\" deleted\n", subscriptionId)
		}
	},
}

func init() {
	getCmd.AddCommand(getSubscriptionCmd)
	describeCmd.AddCommand(describeSubscriptionCmd)
	createSubscriptionCmd.Flags().StringVarP(&subsFile, "subsFile", "f", "", "Subscription resource filename")
	createCmd.AddCommand(createSubscriptionCmd)
	deleteCmd.AddCommand(deleteSubscriptionCmd)
}
