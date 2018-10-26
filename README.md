# atlas-contacts-app

_This generated README.md file loosely follows a [popular template](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)._

One paragraph of project description goes here.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

**Install go dep**

``` sh
go get -u github.com/golang/dep/cmd/dep
```

### Database Migration
[Architecture](https://docs.google.com/document/d/1wM26w4Fdkvw0mWKKupkiLwHQSBLeN5H0euieizKue-Q/edit?usp=sharing)

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
***Steps***
1. The service container is modified to include migration files in `/db/migrations/` directory.
2. Then `initContainers` are leveraged to copy the migration files from service image to migration image as outlined in `deploy/migrations.yaml`.
3. To create migrations update `migrations.yaml` with **database details** and run:

```
  # Modify database details in migrations.yaml (database name, user, password, address)
  kubectl apply -f deploy/migration.yaml
```

***Docker Usage***

In case you don't want to follow the kubernetes deployment; build service image following Step 1 and run
```
docker create --name init-container1 -v migrations:/db/migrations infobloxcto/contacts-server:latest
docker run --name contacts-app-migration --volumes-from init-container1 --network host infoblox/migrate:latest --verbose --path=/db/migrations/ --database.address=postgres.contacts:5432 --database.user=postgres --database.password=postgres --database.name=atlas_contacts_app up
```

**NOTE**: migrate support database connection string (passed as --database.dsn) as well as individual parameters (passed as --database.driver, --database.address, --database.name, --database.user, --database.password and --database.ssl)

### Local development setup

Please note that you should have the following ports opened on you local workstation: `:8080 :8081 :9090 :5432`.
If they are busy - please change them via corresponding parameters of `gateway` and `server` binaries or postgres container run.

Run PostgresDB:

```sh
docker run --name contacts-db -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=contacts -p 5432:5432 -d postgres:9.4
```

Table creation should be done manually by running the migrations scripts or following the steps defined in database migration section. Scripts can be found at `./db/migrations/`

Create vendor directory with required golang packages
``` sh
make vendor
```

Run App server:

``` sh
go run ./cmd/server/*.go
```

#### Try atlas-contacts-app

For Multi-Account environment, Authorization token (Bearer) is required. You can generate it using https://jwt.io/ with following Payload:
```
{
  "AccountID": YourAccountID
}
```

Example:
```
{
  "AccountID": 1
}
```
Bearer
``` sh
export JWT="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEal3O7t3vLI3edA7dGU"
```

Request examples:
``` sh
curl -H "Authorization: Bearer $JWT" \
http://localhost:8080/v1/contacts -d '{"first_name": "Mike", "primary_email": "mike@example.com"}'
```

``` sh
curl -H "Authorization: Bearer $JWT" \
http://localhost:8080/v1/contacts -d \
'{"first_name": "Robert", "primary_email": "robert@example.com", "nicknames": ["bob", "robbie"]}'
```

``` sh
curl -H "Authorization: Bearer $JWT" \
http://localhost:8080/v1/contacts?_filter='first_name=="Mike"'
```
Note, that `JWT` should contain AccountID field.

#### Build docker images

``` sh
make
```
Will be created docker images 'infoblox/contacts-gateway' and 'infoblox/contacts-server'.

If this process finished with errors it's likely that docker doesn't allow to mount host directory in its container.
Therefore you are proposed to run `su -c "setenforce 0"` command to fix this issue.

### Local Kubernetes setup

##### Prerequisites

Make sure nginx is deployed in your K8s. Otherwise you can deploy it using

``` sh
make nginx-up
```

##### Deployment
To deploy atlas-contacts-app use

``` sh
make up
```
Will be used latest Docker Hub images: 'infoblox/contacts-gateway:latest', 'infoblox/contacts-server:latest'.

To deploy authN stub, clone atlas-stubs repo (https://github.com/infobloxopen/atlas-stubs.git) and then execute deployment script inside authn-stub package or:

``` sh
curl https://raw.githubusercontent.com/infobloxopen/atlas-stubs/master/authn-stub/deploy/authn-stub.yaml | kubectl apply -f -
```

This will start AuthN stub that maps `User-And-Pass` header on JWT tokens, with following meaning:

```
admin1:admin -> AccountID=1
admin2:admin -> AccountID=2
```

##### Usage

Try it out by executing following curl commands:

``` sh
# Create some profiles
curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/profiles -d '{"name": "personal", "notes": "Used for personal aims"}' | jq

curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/profiles -d '{"name": "work", "notes": "Used for work aims"}' | jq

# Create some groups assigned to profiles
curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/groups -d '{"name": "schoolmates", "profile_id": "atlas-contacts-app/profiles/1"}' | jq

curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/groups -d '{"name": "family", "profile_id": "atlas-contacts-app/profiles/1"}' | jq

curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/groups -d '{"name": "accountants", "profile_id": "atlas-contacts-app/profiles/2"}' | jq

# Add some contacts assigned to profiles and groups
curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/contacts -d '{"first_name": "Mike", "primary_email": "mike@gmail.com", "profile_id": "atlas-contacts-app/profiles/1", "groups": ["atlas-contacts-app/groups/1", "atlas-contacts-app/groups/2"], "home_address": {"city": "Minneapolis", "state": "Minnesota", "country": "US"}}' | jq

curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/contacts -d '{"first_name": "John", "primary_email": "john@gmail.com", "profile_id": "atlas-contacts-app/profiles/2", "work_address": {"city": "St.Paul", "state": "Minnesota", "country": "US"}}' | jq

# Patch zip value for a particular contact
curl -k -H "User-And-Pass: admin1:admin" -X PATCH \
https://$(minikube ip)/atlas-contacts-app/v1/contacts/{contact_id} -d '{"work_address": {"zip": "161"}}' | jq

# Read created resources
curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/profiles  | jq

curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/groups | jq

curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/contacts | jq

# Read groups which belong to a profile
curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/profiles/1/groups  | jq

# Read contacts which belong to a profile
curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/profiles/1/contacts  | jq

# Read contacts which belong to a group
curl -k -H "User-And-Pass: admin1:admin" \
https://$(minikube ip)/atlas-contacts-app/v1/groups/1/contacts  | jq
```

##### API documentation

API documentation in k8s deployment could be found on following link, note that no credentials needed to access it:

```
https://<minikube address>/atlas-contacts-app/apidoc/index
```

NOTE: This documentation page is only for demo purposes, do not copy and use it in your own production due to risk of injections.

##### Pagination (page token)

**DISCLAIMER**: it is intended only for demonstration purposes and should not be emulated.

Contacts App implements pagination in by adding application **specific** page token implementation.

Actually the service supports "composite" pagination in a specific way:

- limit and offset are still supported but without page token

- if an user requests page token and provides limit then limit value will be used as a step for all further requests
		`page_token = null & limit = 2 -> page_token=base64(offset=2:limit=2)`

- if an user requests page token and provides offset then only first time the provided offset is applied
		`page_token = null & offset = 2 & limit = 2 -> page_token=base64(offset=4:limit=2)`

Get all contacts: `GET http://localhost:8080/v1/contacts`
```json
{
  "results": [
    {
      "emails": [
        {
          "address": "one@mail.com",
          "id": "1"
        }
      ],
      "first_name": "Mike",
      "id": "1",
      "primary_email": "one@mail.com"
    },
    {
      "emails": [
        {
          "address": "two@mail.com",
          "id": "2"
        }
      ],
      "first_name": "Mike",
      "id": "2",
      "primary_email": "two@mail.com"
    },
    {
      "emails": [
        {
          "address": "three@mail.com",
          "id": "3"
        }
      ],
      "first_name": "Mike",
      "id": "3",
      "primary_email": "three@mail.com"
    }
  ],
  "success": {
    "status": 200,
    "code": "OK"
  }
}
```

Default pagination (supported by atlas-app-toolkit): `GET http://localhost:8080/v1/contacts?_limit=1&_offset=1`
```json
{
  "results": [
    {
      "emails": [
        {
          "address": "two@mail.com",
          "id": "2"
        }
      ],
      "first_name": "Mike",
      "id": "2",
      "primary_email": "two@mail.com"
    }
  ],
  "success": {
    "status": 200,
    "code": "OK"
  }
}
```

Request **specific** page token: `GET http://localhost:8080/v1/contacts?_page_token=null&_limit=2`
```json
{
  "results": [
    {
      "emails": [
        {
          "address": "one@mail.com",
          "id": "1"
        }
      ],
      "first_name": "Mike",
      "id": "1",
      "primary_email": "one@mail.com"
    },
    {
      "emails": [
        {
          "address": "two@mail.com",
          "id": "2"
        }
      ],
      "first_name": "Mike",
      "id": "2",
      "primary_email": "two@mail.com"
    }
  ],
  "success": {
    "status": 200,
    "code": "OK",
    "_page_token": "NDo0"
  }
}
```

Get next page via page token: `GET http://localhost:8080/v1/contacts?_page_token=NDo0`
```json
{
  "results": [
    {
      "emails": [
        {
          "address": "three@mail.com",
          "id": "3"
        }
      ],
      "first_name": "Mike",
      "id": "3",
      "primary_email": "three@mail.com"
    }
  ],
  "success": {
    "status": 200,
    "code": "OK",
    "_page_token": "NTo0"
  }
}
```

Get next page: `GET http://localhost:8080/v1/contacts?_page_token=NTo0`
The `"_page_token": "null"` means there are no more pages
```json
{
  "success": {
    "status": 200,
    "code": "OK",
    "_page_token": "null"
  }
}
```

## Deployment

Add additional notes about how to deploy this application. Maybe list some common pitfalls or debugging strategies.

## Running the tests

Explain how to run the automated tests for this system.

```
Give an example
```

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Debugging
Here is a real life case, I downloaded the application and did migration:
```bash
$ createdb atlas_contacts_app
$ make migrate-up
```
Then built and ran the server:
```bash
$ go run ./cmd/server/*.go
```
Then you do the REST call:
```bash
$ export JWT="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEal3O7t3vLI3edA7dGU"
$ curl -H "Authorization: Bearer $JWT" \
http://localhost:8080/v1/contacts -d '{"first_name": "Mike", "primary_email": "mike@example.com"}'
```
Here is the error I get:
```bash
(pq: relation "contacts" does not exist)
[2018-10-23 14:44:24]
ERRO[6186] finished unary call with code Internal        Request-Id=73690003-d318-424e-b752-a30f452deab0 account_id=1 error="Internal error occured. For more details see log for request 73690003-d318-424e-b752-a30f452deab0" grpc.code=Internal grpc.method=Create grpc.service=api.contacts.Contacts grpc.start_time="2018-10-23T14:44:24-07:00" grpc.time_ms=2.06 internal-error="pq: relation \"contacts\" does not exist" span.kind=server system=grpc
```
The service oriented pattern using gRPC means that it is difficult to
track issues unless you have instrumented your system and the error
events are clear.

At the top level you would look at ./cmd/server/grpc.go
```bash
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
}
```
This is the function that gets called to handle a request. After this
point most of the other functions are proto generated code, so you
need to find these stubs to instrument. In case of creating a
contact for example, this is the server function that is invoked:
```go
// Create ...
func (m *ContactsDefaultServer) Create(ctx context.Context, in *CreateContactRequest) (*CreateContactResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(ContactsContactWithBeforeCreate); ok {
		var err error
		if db, err = custom.BeforeCreate(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateContact(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	...
}
```
This server will call function to interact with Gorm for persistence:
```go
// DefaultCreateContact executes a basic gorm create call
func DefaultCreateContact(ctx context.Context, in *Contact, db *gorm1.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateContact")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	...
}
```
This ends up calling another generated stub function 'ToORM()'into
the Gorm interface:
```bash
func (m *Contact) ToORM(ctx context.Context) (ContactORM, error) {
	to := ContactORM{}
	var err error
	if prehook, ok := interface{}(m).(ContactWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
```
A good reference would be the object definition 'ContactORM{}' for the
resource which again a generated code:
```bash
type ContactORM struct {
	AccountID   string
	Emails      []*EmailORM `gorm:"foreignkey:ContactId;association_foreignkey:Id"`
	FirstName   string
	Groups      []*GroupORM `gorm:"many2many:group_contacts;jointable_foreignkey:contact_id;association_jointable_foreignkey:group_id;association_autoupdate:false;association_autocreate:false"`
	HomeAddress *AddressORM `gorm:"foreignkey:HomeAddressContactId;association_foreignkey:Id"`
	Id          int64       `gorm:"type:serial;primary_key"`
	LastName    string
	MiddleName  string
	Nicknames   *postgres1.Jsonb `gorm:"type:jsonb"`
	Notes       string
	ProfileId   *int64      `gorm:"type:integer"`
	WorkAddress *AddressORM `gorm:"foreignkey:WorkAddressContactId;association_foreignkey:Id"`
}
```
Then we can check this against the database to make sure that the migration
is in sync with the resource definition.
```bash
$ psql atlas_contacts_app
atlas_contacts_app=# \dt
              List of relations
 Schema |       Name        | Type  |  Owner
--------+-------------------+-------+---------
 public | addresses         | table | seizadi
 public | contacts          | table | seizadi
 public | emails            | table | seizadi
 public | group_contacts    | table | seizadi
 public | groups            | table | seizadi
 public | profiles          | table | seizadi
 public | schema_migrations | table | seizadi
(7 rows)

atlas_contacts_app=# \d+ contacts
                                                           Table "public.contacts"
   Column    |           Type           | Collation | Nullable |               Default                | Storage  | Stats target | Description
-------------+--------------------------+-----------+----------+--------------------------------------+----------+--------------+-------------
 id          | integer                  |           | not null | nextval('contacts_id_seq'::regclass) | plain    |              |
 account_id  | character varying(255)   |           |          |                                      | extended |              |
 created_at  | timestamp with time zone |           |          | CURRENT_TIMESTAMP                    | plain    |              |
 updated_at  | timestamp with time zone |           |          |                                      | plain    |              |
 first_name  | character varying(255)   |           |          | NULL::character varying              | extended |              |
 middle_name | character varying(255)   |           |          | NULL::character varying              | extended |              |
 last_name   | character varying(255)   |           |          | NULL::character varying              | extended |              |
 nicknames   | jsonb                    |           |          |                                      | extended |              |
 notes       | text                     |           |          |                                      | extended |              |
 profile_id  | integer                  |           |          |                                      | plain    |              |
Indexes:
    "contacts_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "contacts_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE SET NULL
Referenced by:
    TABLE "addresses" CONSTRAINT "addresses_home_address_contact_id_fkey" FOREIGN KEY (home_address_contact_id) REFERENCES contacts(id) ON DELETE CASCADE
    TABLE "addresses" CONSTRAINT "addresses_work_address_contact_id_fkey" FOREIGN KEY (work_address_contact_id) REFERENCES contacts(id) ON DELETE CASCADE
    TABLE "emails" CONSTRAINT "emails_contact_id_fkey" FOREIGN KEY (contact_id) REFERENCES contacts(id) ON DELETE CASCADE
    TABLE "group_contacts" CONSTRAINT "group_contacts_contact_id_fkey" FOREIGN KEY (contact_id) REFERENCES contacts(id) ON DELETE CASCADE
Triggers:
    contacts_updated_at BEFORE INSERT OR UPDATE ON contacts FOR EACH ROW EXECUTE PROCEDURE set_updated_at()
```
For the first FKey constraints lets check the emails table:
```bash
atlas_contacts_app=# \d+ emails
                                                           Table "public.emails"
   Column   |           Type           | Collation | Nullable |              Default               | Storage  | Stats target | Description
------------+--------------------------+-----------+----------+------------------------------------+----------+--------------+-------------
 id         | integer                  |           | not null | nextval('emails_id_seq'::regclass) | plain    |              |
 created_at | timestamp with time zone |           |          | CURRENT_TIMESTAMP                  | plain    |              |
 updated_at | timestamp with time zone |           |          |                                    | plain    |              |
 is_primary | boolean                  |           |          | false                              | plain    |              |
 address    | character varying(255)   |           |          | NULL::character varying            | extended |              |
 account_id | character varying(255)   |           |          |                                    | extended |              |
 contact_id | integer                  |           |          |                                    | plain    |              |
Indexes:
    "emails_pkey" PRIMARY KEY, btree (id)
    "emails_address_key" UNIQUE CONSTRAINT, btree (address)
Foreign-key constraints:
    "emails_contact_id_fkey" FOREIGN KEY (contact_id) REFERENCES contacts(id) ON DELETE CASCADE
Triggers:
    emails_updated_at BEFORE INSERT OR UPDATE ON emails FOR EACH ROW EXECUTE PROCEDURE set_updated_at()
```

There is a lot of indirection in the Gorm interface starting with this
stub:
```go
	if hook, ok := interface{}(&ormObj).(ContactORMWithBeforeCreate); ok {
```
which calls gorm [Create()](https://github.com/jinzhu/gorm/blob/master/main.go#L440)
calls [callCallbacks](https://github.com/jinzhu/gorm/blob/master/scope.go#L857)
which calls [beforeCreateCallback()](https://github.com/jinzhu/gorm/blob/master/callback_create.go#L22)
At the end of this chain Gorm builds the SQL Excution that gets sent here to database.
[You can find it here](https://github.com/jinzhu/gorm/blob/master/callback_create.go#L119)
You can look at scope.SQL to find the command. We execute following insert:
```sh
INSERT INTO "contacts" ("account_id","first_name","last_name","middle_name","nicknames","notes","profile_id") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "contacts"."id"
$1 = "1"
$2 = "Mike"
$3 = ""
$4 = ""
$5 = nil
$6 = ""
$7 = nil
```

The best place to catch Gorm Errors is where we check for operation errors
and set flag in scope [here](https://github.com/jinzhu/gorm/blob/master/scope.go#L85)
The call fails in
[saveAfterAssociationsCallback()](https://github.com/jinzhu/gorm/blob/master/callback_save.go#L98)
We are trying to satisfy the 'field.Field' Email which 'has_many' relationship to contacts.

### AuthN and AuthZ
You should start your search by looking at the JWT Header and its claims.
You will find the code for processing the token in
[jwt.go](https://github.com/infobloxopen/atlas-app-toolkit/blob/master/auth/jwt.go)
For example follow
[GetAccountID](https://github.com/infobloxopen/atlas-app-toolkit/blob/master/auth/jwt.go#L59)
to see how we extract the account_id (aka AccountID) value from the JWT Token, this uses
[GetJWTField()](https://github.com/infobloxopen/atlas-app-toolkit/blob/master/auth/jwt.go#L52)
which calls
[getToken()](https://github.com/infobloxopen/atlas-app-toolkit/blob/master/auth/jwt.go#L72)
which leverage the
[grpc middleware toolkit](https://github.com/grpc-ecosystem/go-grpc-middleware),
for the case of reading the header we use
[AuthFromMD()](https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/auth/metadata.go#L24)
which look for the header pattern 'Authorization: Bearer $JWT' in header
and returns the $JWT object. The toolkit uses
[jwt-go package](https://github.com/dgrijalva/jwt-go) to enumerate the JWT claims using
[Parser()](https://github.com/dgrijalva/jwt-go/blob/master/parser.go#L19).

I figure out I can manually put enteries into the database! I can cause
a problem just query not just create. See
```bash
$ curl -H "Authorization: Bearer $JWT" http://localhost:8080/v1/contacts
{"error":{"status":500,"code":"INTERNAL","message":"Internal error occured. For more details see log for request 1f4b0fe4-175d-45a0-9377-5025e8998f6c"}}
```
More useful is the Posgres log:
```bash
2018-10-25 17:28:27.248 PDT [19556] ERROR:  relation "contacts" does not exist at character 15
2018-10-25 17:28:27.248 PDT [19556] STATEMENT:  SELECT * FROM "contacts"  WHERE ("contacts"."account_id" = $1) ORDER BY "id" LIMIT 1000 OFFSET 0
```

Let's look at the query since simpler chain, still trigger on the error in Err()
```bash
UnaryServerInterceptor() -> ChainUnaryServer() -> _Contacts_List_Handler() -> ...
... (s *contactsServer) List() -> (m *ContactsDefaultServer) List() -> ...
... DefaultListContact() -> ...
Gorm Layer [(s *DB) Find() -> (scope *Scope) callCallbacks() -> ...
... -> queryCallback(scope *Scope) -> (scope *Scope) Err(err error)
```

I am trying to start a new project, atlas cli is not in good state so I
tried to use contacts-app as a starting point, before I made any changes
I tried to run it and had issue with Gorm interface getting this obscure
error from postgres: “(pq: relation “contacts” does not exist)”
Here are the steps I followed:
```sh
$ git clone git@github.com:Infoblox-CTO/atlas.contacts.app.git
$ git status
..
	modified:   Makefile
	modified:   config/kubernetes.yaml

$ git diff config/kubernetes.yaml
diff --git a/config/kubernetes.yaml b/config/kubernetes.yaml
index 70d3992..36806c6 100644
--- a/config/kubernetes.yaml
+++ b/config/kubernetes.yaml
@@ -8,10 +8,10 @@ gateway:
   swaggerFile: ./www/contacts.swagger.json
   swaggerUI: ./www/swagger-ui-dist/
 database:
-  address: postgres.contacts:5432
+  address: localhost:5432
   name: atlas_contacts_app
-  user: postgres
-  password: postgres
+  user: seizadi
+  password:
   ssl: disable
 atlas.authz:
   enable: false
@@ -21,4 +21,4 @@ internal:
   health: /healthz
   readiness: /ready
 logging:
-  level: debug
\ No newline at end of file
+  level: debug
```
Changed Makefile so I can run migrations:

```sh
$ git diff Makefile
diff --git a/Makefile b/Makefile
index 934b4fc..1d095dd 100644
--- a/Makefile
+++ b/Makefile
@@ -101,3 +101,10 @@ nginx-up:
 nginx-down:
        kubectl delete -f deploy/nginx.yaml

+.PHONY: migrate-up
+migrate-up:
+        @migrate -database 'postgres://$(DATABASE_ADDRESS)/atlas_contacts_app?sslmode=disable' -path ./db/migrations up
+
+.PHONY: migrate-down
+migrate-down:
+        @migrate -database 'postgres://$(DATABASE_ADDRESS):5432/atlas_contacts_app?sslmode=disable' -path ./db/migrations down

$ createdb atlas_contacts_app
$ make migrate-up
1/u contacts (22.234579ms)
2/u emails (34.376535ms)
3/u groupprofileaddress (46.925954ms)

$ go run ./cmd/server/*.go
DEBU[0000] serving internal http at "0.0.0.0:8081"
DEBU[0000] serving gRPC at "0.0.0.0:9090"
DEBU[0000] serving http at "0.0.0.0:8080"

$ curl -H "Authorization: Bearer $JWT" http://localhost:8080/v1/contacts
{"error":{"status":500,"code":"INTERNAL","message":"Internal error occured. For more details see log for request ff527439-7a70-4222-b374-44c22b6483f8"}}
```
Log from contact-app server:
```sh
(pq: relation "contacts" does not exist)
[2018-10-25 22:11:31]
ERRO[0046] finished unary call with code Internal        Request-Id=ff527439-7a70-4222-b374-44c22b6483f8 account_id=1 error="Internal error occured. For more details see log for request ff527439-7a70-4222-b374-44c22b6483f8" grpc.code=Internal grpc.method=List grpc.service=api.contacts.Contacts grpc.start_time="2018-10-25T22:11:31-07:00" grpc.time_ms=2.239 internal-error="pq: relation \"contacts\" does not exist" span.kind=server system=grpc

```
Log from Postgres Server:
```sh
2018-10-25 22:11:31.767 PDT [24996] ERROR:  relation "contacts" does not exist at character 15
2018-10-25 22:11:31.767 PDT [24996] STATEMENT:  SELECT * FROM "contacts"  WHERE ("contacts"."account_id" = $1) ORDER BY "id" LIMIT 1000 OFFSET 0
```
