# cmdb

_This generated README.md file loosely follows a [popular template](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)._

One paragraph of project description goes here.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

```bash
$ atlas init-app -db -debug -gateway -registry soheileizadi/cmdb -name cmdb
```

### Prerequisites

What things you need to install the software and how to install them.

```
Give examples
```

### Database Migration

For migrating the database schema, [golang-migrate](https://github.com/golang-migrate/migrate) framework is used.

***Steps for local development***
1. The migration files must be placed `/db/migrations/` directory.
2. Make sure you have [migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cli)
3. Run migration
```bash
$ make migrate-up
```

#### Schema
The database schema will have following tables:
```bash
postgres-# \dt
                 List of relations
 Schema |          Name           | Type  |  Owner
--------+-------------------------+-------+---------
 public | application_container   | table | seizadi
 public | applications            | table | seizadi
 public | artifacts               | table | seizadi
 public | kubernetes              | table | seizadi
 public | aws_rds                 | table | seizadi
 public | aws_services            | table | seizadi
 public | aws_to_rds              | table | seizadi
 public | containers              | table | seizadi
 public | deployments             | table | seizadi
 public | environment_application | table | seizadi
 public | environments            | table | seizadi
 public | manifests               | table | seizadi
 public | region_environment      | table | seizadi
 public | regions                 | table | seizadi
 public | secrets                 | table | seizadi
 public | vault_secret            | table | seizadi
 public | vaults                  | table | seizadi
 public | version_tags            | table | seizadi
(18 rows)
```
Here is the general format for one of the tables:
```bash
postgres-# \d+ applications
                                                             Table "public.applications"
     Column     |           Type           | Collation | Nullable |                 Default                  | Storage  | Stats target | Description
----------------+--------------------------+-----------+----------+------------------------------------------+----------+--------------+-------------
 id             | integer                  |           | not null | nextval('applications_id_seq'::regclass) | plain    |              |
 created_at     | timestamp with time zone |           |          |                                          | plain    |              |
 updated_at     | timestamp with time zone |           |          |                                          | plain    |              |
 name           | text                     |           |          |                                          | extended |              |
 description    | text                     |           |          |                                          | extended |              |
 code           | text                     |           |          |                                          | extended |              |
 version_tag_id | integer                  |           |          |                                          | plain    |              |
 manifest_id    | integer                  |           |          |                                          | plain    |              |
Indexes:
    "applications_pkey" PRIMARY KEY, btree (id)
```
You should reference the migration files (./db/migrations) for more detail on the Tables.
### Development

Run App server:

```bash
$ go run ./cmd/server/*.go -db "host=localhost port=5432 user=seizadi password= sslmode=disable dbname=cmdb"
```

### Installing

A step-by-step series of examples that tell you have to get a development environment running.

Say what the step will be.

```
Give the example
```

And repeat.

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo.

## Deployment

Add additional notes about how to deploy this application. Maybe list some common pitfalls or debugging strategies.

## Running the tests

Explain how to run the automated tests for this system.

```
Give an example
```

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/seizadi/cmdb/tags).
