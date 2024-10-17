# crossfhir

`crossfhir` is a CLI tool to easily migrate AWS HealthLake FHIR database to PostgreSQL

## Features overview

1. Connect to AWS HealthLake using credentials.
2. Export the FHIR repository to S3 and copy it from S3 to local storage.
5. PostgreSQL connector 
4. Convert .ndjson files into a format that can be imported into Postgres and save them locally.
5. Optionally, import the data into Postgres.

https://github.com/fhirbase/fhirbase