# User Guide

* [Configuration](#configuration)

## Configuration

| Parameter          | Description                                                                                                                                                            |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| billing-account-id | Billing account that will be attached to all RAD Lab resources.                                                                                                        |
| admin-project      | Project ID for the RAD Lab admin project.                                                                                                                              |
| admin-bucket       | Name of the storage bucket that contains all RAD Lab admin resources.                                                                                                  |
| rad-lab-dir        | Directory where RAD Lab sources are installed.                                                                                                                         |
| parent-id          | ID of the parent where all RAD Lab resources will be created. Should be configured as `organizations/12345678` or `folders/12345678`.                                  |
| region             | Google Cloud region where all resources will be created.  This should be an existing region name.  You can find region names by running `gcloud compute regions list`. |
| zone               | Google Cloud zone where all zonal resources will be created.  You can find all zone names by running `gcloud compute zones list`.                                      |