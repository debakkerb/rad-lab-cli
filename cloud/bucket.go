package cloud

import (
	"cloud.google.com/go/storage"
	"context"
	"time"
)

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
