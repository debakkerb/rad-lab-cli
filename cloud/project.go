package cloud

import (
	billing "cloud.google.com/go/billing/apiv1"
	"cloud.google.com/go/billing/apiv1/billingpb"
	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	"cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"strings"
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

type GoogleProject struct {
	ProjectName      string
	ProjectID        string
	ProjectNumber    string
	ParentID         string
	BillingAccountID string
}

func (p *GoogleProject) Create() error {
	ctx := context.Background()

	projectClient, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		return err
	}
	defer projectClient.Close()

	err = projectExists(projectClient, ctx, p.ParentID, p.ProjectID)
	if err != nil {
		return err
	}

	err = p.createProject(ctx, projectClient)
	if err != nil {
		return err
	}

	err = p.addBilling(ctx)
	if err != nil {
		return err
	}

	err = p.enableServices(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *GoogleProject) enableServices(ctx context.Context) error {
	serviceClient, err := serviceusage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer serviceClient.Close()

	enableServiceRequest := &serviceusagepb.BatchEnableServicesRequest{
		Parent: fmt.Sprintf("projects/%s", p.ProjectID),
		ServiceIds: []string{
			"storage.googleapis.com",
			"cloudbilling.googleapis.com",
			"iam.googleapis.com",
		},
	}

	_, err = serviceClient.BatchEnableServices(ctx, enableServiceRequest)
	if err != nil {
		return err
	}

	return nil
}

func (p *GoogleProject) addBilling(ctx context.Context) error {
	billingClient, err := billing.NewCloudBillingClient(ctx)
	if err != nil {
		return err
	}
	defer billingClient.Close()

	updateBillingInfoRequest := &billingpb.UpdateProjectBillingInfoRequest{
		Name: fmt.Sprintf("projects/%s", p.ProjectID),
	}

	_, err = billingClient.UpdateProjectBillingInfo(ctx, updateBillingInfoRequest)
	if err != nil {
		return err
	}

	return nil
}
func (p *GoogleProject) createProject(ctx context.Context, projectClient *resourcemanager.ProjectsClient) error {
	projectDetails := &resourcemanagerpb.Project{
		Parent:      p.ParentID,
		DisplayName: p.ProjectID,
		ProjectId:   p.ProjectID,
		Labels: map[string]string{
			"created-by": "rad-lab-cli",
			"function":   "rad-lab",
		},
	}

	projectCreateRequest := &resourcemanagerpb.CreateProjectRequest{
		Project: projectDetails,
	}

	createProjectOperation, err := projectClient.CreateProject(ctx, projectCreateRequest)
	if err != nil {
		return err
	}

	project, err := createProjectOperation.Wait(ctx)
	if err != nil {
		return err
	}

	p.ProjectNumber = strings.SplitAfter(project.Name, "/")[0]
	p.ProjectName = project.Name

	return nil
}

func projectExists(projectClient *resourcemanager.ProjectsClient, context context.Context, parentID, projectID string) error {
	listProjectsRequest := &resourcemanagerpb.ListProjectsRequest{
		Parent:      parentID,
		ShowDeleted: true,
	}

	projectIterator := projectClient.ListProjects(context, listProjectsRequest)
	for {
		project, err := projectIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		if project.ProjectId == projectID {
			if project.State == resourcemanagerpb.Project_DELETE_REQUESTED {
				return errors.New("a project with this project id already exists, but is in the process of being deleted. please specify a different project id")
			}

			if project.State == resourcemanagerpb.Project_ACTIVE {
				return errors.New("a project with this name already exists.  please specify a different project id")
			}
		}
	}

	return nil
}
