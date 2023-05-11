package admin

import (
	billing "cloud.google.com/go/billing/apiv1"
	"cloud.google.com/go/billing/apiv1/billingpb"
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"context"
	"errors"
	"github.com/debakkerb/rad-lab-cli/config"
	"google.golang.org/api/iterator"
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

type AdminProject struct {
	ProjectName      string
	ProjectID        string
	ProjectNumber    string
	ParentID         string
	BillingAccountID string
}

func (ap *AdminProject) Create() error {
	err := ap.checkPrerequisites()
	if err != nil {
		return err
	}

	ctx := context.Background()
	projectClient, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		return err
	}
	defer projectClient.Close()

	listProjectsRequest := &resourcemanagerpb.ListProjectsRequest{
		Parent:      ap.ParentID,
		ShowDeleted: true,
	}

	projectIterator := projectClient.ListProjects(ctx, listProjectsRequest)
	for {
		project, err := projectIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		if project.ProjectId == ap.ProjectID {
			if project.State == resourcemanagerpb.Project_DELETE_REQUESTED {
				return errors.New("a project with this project id already exists, but is in the process of being deleted. please specify a different project id")
			}

			if project.State == resourcemanagerpb.Project_ACTIVE {
				return errors.New("a project with this name already exists.  please specify a different project id")
			}
		}

	}

	projectDetails := &resourcemanagerpb.Project{
		Parent:      ap.ParentID,
		DisplayName: ap.ProjectID,
		ProjectId:   ap.ProjectID,
		Labels: map[string]string{
			"created-by": "rad-lab-cli",
			"function":   "rad-lab",
		},
	}

	projectCreateRequest := &resourcemanagerpb.CreateProjectRequest{
		Project: projectDetails,
	}

	_, err = projectClient.CreateProject(ctx, projectCreateRequest)
	if err != nil {
		return err
	}

	billingClient, err := billing.NewCloudBillingClient(ctx)
	if err != nil {
		return err
	}
	defer billingClient.Close()

	updateBillingInfoRequest := &billingpb.UpdateProjectBillingInfoRequest{
		Name: ap.ProjectID,
	}

	//req := &billingpb.UpdateProjectBillingInfoRequest{
	//	Name: "project-name",
	//}
	//
	//_, err = billingClient.UpdateProjectBillingInfo(ctx, req)
	//if err != nil {
	//	return err
	//}
	//
	return nil
}

func (ap *AdminProject) checkPrerequisites() error {
	if ap.ProjectID == "" {
		defaultProjectID := config.Get(config.ParameterAdminProject)
		if defaultProjectID != "" {
			ap.ProjectID = defaultProjectID
		} else {
			return errors.New("error while creating Admin project, Project ID has to be either passed in as command flag or configured as part of the RAD Lab configuration ('radlab config set')")
		}
	}

	if ap.ParentID == "" {
		parentID := config.Get(config.ParameterParentID)
		if parentID != "" {
			ap.ParentID = parentID
		} else {
			return errors.New("error while creating Admin project, Parent ID has to be either passed in as command flag or configured as part of the RAD Lab configuration ('radlab config set')")
		}
	}

	if ap.BillingAccountID == "" {
		billingAccountID := config.Get(config.ParameterBillingAccount)
		if billingAccountID != "" {
			ap.BillingAccountID = billingAccountID
		} else {
			return errors.New("error while creating Admin project, Billing Account ID has to be either passed in as command flag or configured as part of the RAD Lab configuration ('radlab config set')")
		}
	}

	return nil
}
