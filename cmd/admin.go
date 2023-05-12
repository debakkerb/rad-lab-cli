package cmd

import (
	"github.com/debakkerb/rad-lab-cli/admin"
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

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Command to manage all RAD Lab admin resources.",
}

var adminProjectCmd = &cobra.Command{
	Use: "project",
}

var adminProjectCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create admin project.",
	Long:  "Creates the RAD Lab admin project, where all Terraform state is stored.",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(admin.CreateAdminProject(billingAccount, adminProjectID, parentID))
	},
}

var adminProjectShowCmd = &cobra.Command{
	Use: "show",
}

var adminBucketCmd = &cobra.Command{
	Use: "bucket",
}

var adminBucketCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create admin bucket",
	Long:  "Create the storage bucket for all Admin resources",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(admin.CreateAdminBucket())
	},
}

var adminBucketShowCmd = &cobra.Command{
	Use: "show",
}

func init() {
	rootCmd.AddCommand(adminCmd)

	adminCmd.AddCommand(adminProjectCmd)
	adminCmd.AddCommand(adminBucketCmd)

	adminProjectCmd.AddCommand(adminProjectCreateCmd)
	adminProjectCmd.AddCommand(adminProjectShowCmd)

	adminBucketCmd.AddCommand(adminBucketCreateCmd)
	adminBucketCmd.AddCommand(adminBucketShowCmd)

}
