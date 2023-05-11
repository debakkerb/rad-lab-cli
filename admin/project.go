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
	err := checkValues(billingAccountID, adminProjectID, parentID)
	if err != nil {
		return err
	}

	p := cloud.GoogleProject{
		ProjectID:        adminProjectID,
		BillingAccountID: billingAccountID,
		ParentID:         parentID,
		ProjectName:      adminProjectID,
	}

	return p.Create()
}

func checkValues(billingAccountID, adminProjectID, parentID string) error {
	if billingAccountID == "" {
		billingAccountID = config.Get(config.ParameterBillingAccount)
	}

	if adminProjectID == "" {
		adminProjectID = config.Get(config.ParameterAdminProject)
	}

	if parentID == "" {
		parentId
	}
}
