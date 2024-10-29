<h1 align=center>crossfhir</h1>
<div align=center>
  <a href=mailto:hello@themomenum.ai?subject=crossfhir>
    <img src=https://img.shields.io/badge/Contact%20us-AFF476.svg alt="Contact us">
  </a>
    <a href="https://themomentum.ai">
    <img src=https://img.shields.io/badge/Check%20Momentum-1f6ff9.svg alt="Check">
  </a>
  <a href="LICENSE.md">
    <img src="https://img.shields.io/badge/License-MIT-636f5a.svg?longCache=true" alt="MIT License">
  </a>
</div>
<br>


`crossfhir` is a CLI tool to interact with AWS HealthLake FHIR datastore and easily migrate it to PostgreSQL.

## Features overview

### General features

- ✅ Connection to AWS HealthLake and client configuration.
- ✅ Command to trigger and monitor data export from HealthLake to S3.
- ✅ Command to pull exported data from S3 into local .ndjson files (possible to export and pull at once).
- Interaction with FHIR resources on AWS HealthLake via REST API.
  - ✅ GET
  - ✅ PUT
  - POST
  - DELETE
- SMART on FHIR integration

### PostgreSQL features

For more information, check the PostgreSQL section.

- PostgreSQL connection setup.
- Migrating .ndjson files from FHIR v4 resources into specific tables.
- Indexing desired elements of resources and separating them from the jsonb column.
- Handling resource versioning and updates.

## Prerequisites

To use this tool, several environment variables are required to conduct an export from HealthLake.

For more details, you can check the AWS HealthLake module in the [HealthStack repository](https://github.com/TheMomentumAI/healthstack/tree/main/healthlake).
There, you will find Terraform modules to create everything you need. You can especially benefit from it when creating IAM permissions.

In the .env.example file, you can see a list of the required environment variables.

```
# AWS credentials to access your account
export AWS_ACCESS_KEY="AKIA123123123..."
export AWS_SECRET_KEY="XWSWAD123123123..."
export AWS_REGION="us-east-1"

# The name of the S3 bucket
export AWS_S3_BUCKET="s3://fhir-bucket"

# The ARN of an IAM Role to conduct the export, which has access to S3, KMS, and HealthLake
export AWS_IAM_EXPORT_ROLE="arn:aws:iam::123:role/Role"

# The ID of your HealthLake FHIR datastore
export AWS_DATASTORE_ID="123123123123"

# The ARN of a KMS key that encrypts the S3 bucket (not that encrypts the HealthLake datastore)
export AWS_KMS_KEY_ID_ARN="arn:aws:kms:region:123123123:key/123123123"

# The export job name for tracking purposes
export AWS_EXPORT_JOB_NAME="my-export-job"
```

If you want to work directly with the source code, you need to have Go installed on your machine.

## Usage

There is a binary in this repository that you can use if you don't have Go installed.

```
./bin/crossfhir

crossfhir is a CLI for converting AWS Health Lake FHIR data to PostgreSQL

Usage:
  crossfhir [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  convert     Convert FHIR data to Postgres data
  export      Export FHIR data from AWS Health Lake
  help        Help about any command
  pull        Pull FHIR data from S3 to local

Flags:
      --env-file string   environment file to load (default ".env")
  -h, --help              help for crossfhir

Use "crossfhir [command] --help" for more information about a command.
```

Exporting data with pull into local directory:

```sh
./bin/crossfhir export --pull --dir ./mydirectory
```

Pulling created export from S3 path:

```sh
./bin/crossfhir pull --url s3://fhir-bucket/fhir-export-123
```

## Contribution

We are open to, and grateful for, any contributions made by the community.


<a href="https://github.com/TheMomentumAI/crossfhir/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=TheMomentumAI/crossfhir" />
</a>

## License

This project is released under the MIT License.
