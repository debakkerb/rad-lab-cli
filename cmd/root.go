package cmd

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

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	billingAccount  string
	adminProjectID  string
	parentID        string
	adminBucketName string
)

var rootCmd = &cobra.Command{
	Use:     "radlab",
	Aliases: []string{"rl"},
	Short:   "Manage Rad Lab deployments, incl. UI",
	Long:    `Manage RAD Lab deployments on Google Cloud`,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&adminProjectID, "project-id", "p", "", "Project ID of the admin project for RAD Lab.  When specified, will overwrite local configuration.")
	rootCmd.PersistentFlags().StringVarP(&billingAccount, "billing-account-id", "b", "", "Billing account ID that will be added to all RAD Lab resources. When specified, will overwrite local configuration.")
	rootCmd.PersistentFlags().StringVarP(&parentID, "parent-id", "a", "", "ID of the parent for all RAD Lab resources.  Should be specified as 'organizations.123456789' or 'folders/123456'. When specified, will overwrite local configuration.")
	rootCmd.PersistentFlags().StringVarP(&adminBucketName, "bucket-name", "u", "", "Bucket name of the Admin project, where all Terraform state will be stored.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
