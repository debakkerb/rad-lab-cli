package admin

import (
	billing "cloud.google.com/go/billing/apiv1"
	"cloud.google.com/go/billing/apiv1/billingpb"
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"context"
	"log"
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
	ctx := context.Background()
	client, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		log.Fatal("Error while trying to create a project", err)
	}
	defer client.Close()

	projectDetails := &resourcemanagerpb.Project{
		Name:      ap.ProjectName,
		Parent:    ap.ParentID,
		ProjectId: ap.ProjectID,

		Labels: map[string]string{
			"created-by": "rad-lab-cli",
			"function":   "rad-lab",
		},
	}

	request := &resourcemanagerpb.CreateProjectRequest{
		Project: projectDetails,
	}

	_, err = client.CreateProject(ctx, request)
	if err != nil {
		return err
	}

	billingClient, err := billing.NewCloudBillingClient(ctx)
	if err != nil {
		return err
	}
	defer billingClient.Close()

	req := &billingpb.UpdateProjectBillingInfoRequest{
		Name: "project-name",
	}

	_, err = billingClient.UpdateProjectBillingInfo(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
