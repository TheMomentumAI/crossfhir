<h1 align=center>crossfhir</h1>
<div align=center>
  <a href=mailto:hello@themomenum.ai?subject=crossfhir>
    <img src=https://img.shields.io/badge/Contact%20us-AFF476.svg alt="Contact us">
  </a>
    <a href="https://themomentum.ai">
    <img src=https://img.shields.io/badge/Check%20Momentum-1f6ff9.svg alt="Check">
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-636f5a.svg?longCache=true" alt="MIT License">
  </a>
</div>
<br>


`crossfhir` is a CLI tool to interact with AWS HealthLake FHIR datastore and easily migrate it to PostgreSQL.

## Features overview

### General features

- ✅ Connection to AWS HealthLake and client configuration.
- ✅ Command to trigger and monitor data export from HealthLake to S3.
- ✅ SMART on FHIR and AWS static credentials auth export trigger
- ✅ Command to pull exported data from S3 into local .ndjson files (possible to export and pull at once).

### PostgreSQL features

For more information, check the FAQ section.

- ✅ PostgreSQL connection setup.
- ✅ Migrating .ndjson files from FHIR v4 resources into specific tables.
- Indexing desired elements of resources and separating them from the jsonb column.
- Handling resource versioning and updates.

## Installation

Currently supports manual installation only. Homebrew support is coming soon.

The script automatically detects your operating system and architecture, downloads the appropriate binary, and installs it to /usr/local/bin.

```sh
curl -sSL https://raw.githubusercontent.com/TheMomentumAI/crossfhir/master/scripts/install.sh | bash
```

## Prerequisites

Each command requires a set of configuration parameters, which are validated during execution.

For example, several parameters are needed to export from HealthLake, but they aren't necessary when loading data into PostgreSQL.

For more details, you can check the AWS HealthLake module in the [HealthStack repository](https://github.com/TheMomentumAI/healthstack/tree/main/healthlake).
There, you will find Terraform modules to create everything you need. You can especially benefit from it when creating IAM permissions.

In the `example.config.toml` file, you can see a list of the required configuration parameters.

```toml
[aws]
# AWS credentials and configuration
aws_access_key = "AKIA..."
aws_secret_key = "Tdaz4e..."
aws_region = "us-east-1"
aws_s3_bucket = "s3://my-bucket"
aws_iam_export_role = "arn:aws:iam::123123123:role/IAMRole"
aws_datastore_id = "8699acc...c49744168"
aws_kms_key_id = "arn:aws:kms:us-east-1:123123123:key/749b1e97-85db-49af5"
aws_export_job_name = "my-export-job"
aws_datastore_fhir_url = "https://healthlake.us-east-1.amazonaws.com"

[db]
# PostgreSQL database configuration
db_host = "localhost"
db_port = "5432"
db_username = "postgres"
db_password = "password"
db_database = "postgres"

[smart]
# SMART on FHIR configuration
smart_client_id = "123123123123"
smart_client_secret = "12qwert123qwerty123"
smart_token_url = "https://fhir.auth.us-east-1.amazoncognito.com/oauth2/token"
smart_datastore_endpoint = "https://healthlake.us-east-1.amazonaws.com/datastore/123123/r4"
smart_scope = "system/*.*"
smart_grant_type = "client_credentials"
```

If you want to work directly with the source code, you need to have Go installed on your machine.

## Usage

There is a binary in this repository that you can use if you don't have Go installed.

```
$ crossfhir

crossfhir is a CLI tool for converting AWS HealthLake FHIR data to PostgreSQL
and interacting with the HealthLake FHIR REST API.

Usage:
  crossfhir [command]

Examples:

# Exporting data with pull into local directory
crossfhir export --pull --dir ./mydirectory

# Loading data from local directory to PostgreSQL with migration
crossfhir load -m --data ./mydirectory


Available Commands:
  completion  Generate the autocompletion script for the specified shell
  export      Export FHIR data from AWS Health Lake to S3 bucket
  help        Help about any command
  load        Load pulled FHIR data to PostgreSQL.
  pull        Pull FHIR data from S3 to local
  rest        Interact with FHIR REST API

Flags:
      --env-file string   environment file to load (default ".env")
  -h, --help              help for crossfhir
```

Exporting data with pull into local directory:

```sh
$ crossfhir export --pull --dir ./mydirectory
```

Pulling created export from S3 path:

```sh
$ crossfhir pull --url s3://fhir-bucket/fhir-export-123
```

Loading data from local directory to PostgreSQL with migration:

```sh
$ crossfhir load -m --data ./mydirectory
```

## FAQ

> Why not convert all the JSON data to relational data too, and only store everything in SQL after that?

Converting all in non-lossy way will produce **thousands** of tables, which will be hard to work with and
it is very complex task. The implementation to solve this issue is [SQL on FHIR](https://sql-on-fhir.org/ig/latest/StructureDefinition-ViewDefinition.html)

## Contribution

We are open to, and grateful for, any contributions made by the community.

The approach to PostgreSQL tables and jsonb is inspired by the [fhirbase](https://github.com/fhirbase/fhirbase) project by HealthSamurai. Many thanks to the creators and maintainers!


<a href="https://github.com/TheMomentumAI/crossfhir/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=TheMomentumAI/crossfhir" />
</a>


## License

This project is released under the MIT License.
