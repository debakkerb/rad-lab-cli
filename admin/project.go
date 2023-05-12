package admin

import (
	"github.com/debakkerb/rad-lab-cli/cloud"
	"github.com/debakkerb/rad-lab-cli/config"
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

func CreateAdminProject(billingAccountID, adminProjectID, parentID string) error {
	billingAccount, err := checkConfigValue(billingAccountID, config.ParameterBillingAccount)
	if err != nil {
		return err
	}

	projectID, err := checkConfigValue(adminProjectID, config.ParameterAdminProject)
	if err != nil {
		return err
	}

	parent, err := checkConfigValue(parentID, config.ParameterParentID)
	if err != nil {
		return err
	}

	p := cloud.GoogleProject{
		ProjectID:        projectID,
		BillingAccountID: billingAccount,
		ParentID:         parent,
		ProjectName:      projectID,
	}

	return p.Create()
}

