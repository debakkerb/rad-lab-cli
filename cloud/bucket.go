package cloud

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
	"cloud.google.com/go/storage"
	"context"
	"time"
)

// CreateBucket
//
// This function creates a Google Cloud storage bucket in the Admin project
// Its function is to store the Terraform state of modules that were created locally.
//
func CreateBucket(bucketName, projectID, region string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	bucketAttributes := &storage.BucketAttrs{
		Location:               region,
		PublicAccessPrevention: storage.PublicAccessPreventionEnforced,
		VersioningEnabled:      true,
		UniformBucketLevelAccess: storage.UniformBucketLevelAccess{
			Enabled: true,
		},
	}

	if err := bucket.Create(ctx, projectID, bucketAttributes); err != nil {
		return err
	}

	return nil
}
