package cmd

import (
	"github.com/debakkerb/rad-lab-cli/modules"
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

var (
	variablesFile string
)

var modulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "Manage RAD Lab modules",
	Long:  "Manage RAD Lab modules in Google Cloud",
	Run:   nil,
}

var listModulesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all RAD Lab modules",
	Long:  "List all RAD Lab modules, incl. their status.",
}

var createModulesCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a RAD Lab module",
	Long:  "This command creates all the necessary RAD Lab resources for a specific module",
	//Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(modules.StartWizard(variablesFile))
	},
}

var deleteModulesCmd = &cobra.Command{
	Use: "delete",
}

func init() {
	rootCmd.AddCommand(modulesCmd)

	modulesCmd.PersistentFlags().StringVarP(&variablesFile, "variables", "v", "", "specify a key value file for the variables")

	modulesCmd.AddCommand(listModulesCmd)
	modulesCmd.AddCommand(createModulesCmd)
	modulesCmd.AddCommand(deleteModulesCmd)
}
