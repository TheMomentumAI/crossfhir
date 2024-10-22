# crossfhir

`crossfhir` is a CLI tool to easily migrate AWS HealthLake FHIR database to PostgreSQL

## Features overview

1. Connect to AWS HealthLake using credentials.
2. Export the FHIR repository to S3 and copy it from S3 to local storage.
5. PostgreSQL connector
4. Convert .ndjson files into a format that can be imported into Postgres and save them locally.
5. Optionally, import the data into Postgres.

https://github.com/fhirbase/fhirbase

v2
1. Integration with different fhir repos than heatlhlake

wybierasz rzeczy istotne z punktu widzenia wyszukiwania i dodajesz do kolumn a caly resource trzymasz na s3/w jsonie postgres

czy oznaczac transakcje chronologicznie jesli dodajemy nowe resource bo chcemy trzymac historie - (newest)

MUSIMY MIEC WSZYSTKIE WERSJE DANYCH I ZNAC AKTUALNE

TODO sprzatnac import export
TODO update resource -

---

1. run migracja by stworzyc tabele
2. iterowac sie po kolei po resourcach, parsowac, i wrzucac po kolei resourcey
3. handle update/new resources