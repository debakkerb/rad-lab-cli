package cmd

import (
	"github.com/debakkerb/rad-lab-cli/config"
	"github.com/spf13/cobra"
)

/**
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage local RAD Lab CLI configuration parameters",
	Long:  "Initialise and update RAD Lab configuration parameters through this command.\n\n\n" + config.Usage(),
	Run:   nil,
}

var initConfigCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise the local RAD Lab CLI configuration",
	Long:  "Local configuration, stored in ~/.config/radlab can be setup via this parameter.",
	Run:   nil,
}

var setConfigCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a certain parameter.",
	Long:  "Set a certain parameter in the local RAD Lab CLI configuration",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		config.SetConfigParameter(args[0], args[1])
	},
}

var showConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the local configuration.",
	Long:  "Display all local parameters of the RAD Lab CLI configuration",
	Run: func(cmd *cobra.Command, args []string) {
		config.Show()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(initConfigCmd)
	configCmd.AddCommand(setConfigCmd)
	configCmd.AddCommand(showConfigCmd)
}
